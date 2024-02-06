/*
 * Copyright (c) 2023 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package core

import (
	"context"
	"errors"
	"sync"
	"time"

	language "golang.org/x/text/internal/language"
)

var (
	ErrTimeout = errors.New("timeout")
)

// VideoTask video task define
type VideoTask struct {
	Id         int
	ResourceId string
	TaskId     string // ASR task id
	Output     string // Recognition result link
	Status     int    // 0: progress, -1: failed, 1: ok
	CreateAt   time.Time
	UpdateAt   time.Time
}

func (c *Community) createVideoTask(ctx context.Context, resourceId string) error {
	_, err := c.db.ExecContext(ctx, "insert into video_task(resource_id, status, create_at, update_at) values(?, ?, ?, ?)", resourceId, 0, time.Now(), time.Now())
	return err
}

func (c *Community) updateVideoTaskOutput(ctx context.Context, resourceId string, output string) error {
	_, err := c.db.ExecContext(ctx, "update video_task set output = ? where resource_id = ?", output, resourceId)
	return err
}

func (c *Community) updateVideoTaskStatus(ctx context.Context, resourceId string, status int) error {
	_, err := c.db.ExecContext(ctx, "update video_task set status = ? where resource_id = ?", status, resourceId)
	return err
}

func (c *Community) deleteVideoTask(ctx context.Context, resourceId string) error {
	_, err := c.db.ExecContext(ctx, "delete from video_task where resource_id = ?", resourceId)
	return err
}

func (c *Community) getVideoTask(ctx context.Context, resourceId string) (*VideoTask, error) {
	var task VideoTask
	err := c.db.QueryRow("select id, resource_id, task_id, output, status, create_at, update_at from video_task where resource_id = ?", resourceId).Scan(&task.Id, &task.ResourceId, &task.TaskId, &task.Output, &task.Status, &task.CreateAt, &task.UpdateAt)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// NewVideoTask create new video task
func (c *Community) NewVideoTask(ctx context.Context, resourceId string) error {
	err := c.createVideoTask(ctx, resourceId)
	if err != nil {
		return err
	}

	// Set video task cache
	c.SetVideoTaskCache(resourceId, VideoTaskTimestamp(time.Now().Unix()))

	return nil
}

// SetVideoTaskSuccess set video task success
func (c *Community) SetVideoTaskSuccess(ctx context.Context, resourceId string) error {
	return c.updateVideoTaskStatus(ctx, resourceId, 1)
}

// SetVideoTaskFailed set video task failed
func (c *Community) SetVideoTaskFailed(ctx context.Context, resourceId string) error {
	return c.updateVideoTaskStatus(ctx, resourceId, -1)
}

// SetVideoTaskOutput set video task output link
func (c *Community) SetVideoTaskOutput(ctx context.Context, resourceId string, output string) error {
	return c.updateVideoTaskOutput(ctx, resourceId, output)
}

type VideoTaskTimestamp int64
type VideoTaskMap map[string]VideoTaskTimestamp

// VideoTaskCache video task cache
// Simple cache, no expiration
type VideoTaskCache struct {
	sync.RWMutex
	videoTaskMap VideoTaskMap

	// Check task status
	isActive bool
}

// NewVideoTaskCache create new video task cache
func NewVideoTaskCache() *VideoTaskCache {
	return &VideoTaskCache{
		videoTaskMap: make(VideoTaskMap),
		isActive:     false,
	}
}

// Get get video task timestamp
func (c *VideoTaskCache) Get(key string) (VideoTaskTimestamp, bool) {
	c.RLock()
	defer c.RUnlock()
	value, ok := c.videoTaskMap[key]
	return value, ok
}

// Set set video task timestamp
func (c *VideoTaskCache) Set(key string, value VideoTaskTimestamp) {
	c.Lock()
	defer c.Unlock()
	c.videoTaskMap[key] = value
}

// Delete delete video task
func (c *VideoTaskCache) Delete(key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.videoTaskMap, key)
}

// Clear clear video task cache
func (c *VideoTaskCache) Clear() {
	c.Lock()
	defer c.Unlock()
	c.videoTaskMap = make(VideoTaskMap)
}

// SetVideoTaskCache set video task cache
func (c *Community) SetVideoTaskCache(key string, value VideoTaskTimestamp) {
	c.translation.VideoTaskCache.Lock()
	defer c.translation.VideoTaskCache.Unlock()

	c.translation.VideoTaskCache.Set(key, value)

	// Start check task status
	if !c.translation.VideoTaskCache.isActive {
		c.translation.VideoTaskCache.isActive = true
		go c.TimedCheckVideoTask(context.Background(), 30*time.Second)
	}
}

// Timed check status of video task with timeout
func (c *Community) TimedCheckVideoTask(ctx context.Context, timeout time.Duration) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	for {
		select {
		case <-ticker.C:
			// Iterator video task cache
			c.translation.VideoTaskCache.RLock()
			defer c.translation.VideoTaskCache.RUnlock()
			for resourceId, timestamp := range c.translation.VideoTaskCache.videoTaskMap {
				if time.Now().Unix()-int64(timestamp) > int64(timeout) {
					// Delete expired video task
					c.translation.VideoTaskCache.Delete(resourceId)
				} else {
					// Check status of video task
					task, err := c.getVideoTask(ctx, resourceId)
					if err != nil {
						c.xLog.Errorf("TimedCheckVideoTask getVideoTask failed, resourceId: %s, err: %v", resourceId, err)
						continue
					}
					if task.Status == 1 {
						// Set video task success
						c.translation.VideoTaskCache.Delete(resourceId)
					} else if task.Status == -1 {
						// Request for ASR task
						asrTaskData, err := c.translation.Engine.QueryVideo2TextTask(task.TaskId)
						if err != nil {
							c.xLog.Errorf("TimedCheckVideoTask QueryVideo2TextTask failed, resourceId: %s, err: %v", resourceId, err)
							continue
						}

						if asrTaskData.Rtn == 0 {
							// Upload ASR result
							c.translation.Engine.GenerateWebVTTFileWithTranslation(asrTaskData, language.Chinese, language.English)

							// Set video task success
							c.translation.VideoTaskCache.Delete(resourceId)
						}
					}
				}
			}
		case <-timer.C:
			c.xLog.Errorf("TimedCheckVideoTask timeout, timeout: %v", timeout)

			// Update status of video task
			// c.videoTaskCache.Clear()
			c.translation.VideoTaskCache.isActive = false

			return
		}
	}
}
