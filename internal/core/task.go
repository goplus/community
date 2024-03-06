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
	"fmt"
	"sync"
	"time"

	language "golang.org/x/text/language"
)

var (
	ErrTimeout = errors.New("timeout")
)

// VideoTask video task define
type VideoTask struct {
	Id         int
	ResourceId string
	UserId     string // User id
	TaskId     string // ASR task id
	Output     string // Recognition result link
	Status     int    // 0: progress, -1: failed, 1: ok
	CreateAt   time.Time
	UpdateAt   time.Time
}

func (c *Community) createVideoTask(ctx context.Context, userId, resourceId, taskId string) error {
	_, err := c.db.ExecContext(ctx, "insert into video_task (user_id, resource_id, task_id, output, status, create_at, update_at) values (?, ?, ?, ?, ?, ?, ?)", userId, resourceId, taskId, "", 0, time.Now(), time.Now())
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

//nolint:unused
func (c *Community) deleteVideoTask(ctx context.Context, resourceId string) error {
	_, err := c.db.ExecContext(ctx, "delete from video_task where resource_id = ?", resourceId)
	return err
}

func (c *Community) getVideoTask(ctx context.Context, resourceId string) (*VideoTask, error) {
	var task VideoTask
	err := c.db.QueryRow("select id, user_id, resource_id, task_id, output, status, create_at, update_at from video_task where resource_id = ?", resourceId).Scan(&task.Id, &task.UserId, &task.ResourceId, &task.TaskId, &task.Output, &task.Status, &task.CreateAt, &task.UpdateAt)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// NewVideoTask create new video task
func (c *Community) NewVideoTask(ctx context.Context, userId, resourceId string) error {
	urlKey, err := c.GetMediaUrl(ctx, resourceId)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s%s", c.domain, urlKey)

	resp, err := c.translation.Engine.Video2Text(url, "")
	if err != nil {
		return err
	}

	err = c.createVideoTask(ctx, userId, resourceId, resp.TaskId)
	if err != nil {
		return err
	}

	// Set video task cache
	c.SetVideoTaskCache(resourceId, VideoTaskTimestamp(time.Now().Unix()))

	c.xLog.Infof("NewVideoTask success, resourceId: %s, taskId: %s", resourceId, resp.TaskId)

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
			c.xLog.Infof("TimedCheckVideoTask start, timeout: %v", timeout)
			// Iterator video task cache
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
					if task.Status == 1 || task.Status == -1 {
						// Set video task success or failed
						c.translation.VideoTaskCache.Delete(resourceId)
					} else if task.Status == 0 {
						// Update ASR result
						err = c.updateASRResult(ctx, resourceId, task)
						if err != nil {
							c.xLog.Errorf("TimedCheckVideoTask updateASRResult failed, resourceId: %s, err: %v", resourceId, err)

							err = c.SetVideoTaskFailed(ctx, resourceId)
							if err != nil {
								c.xLog.Errorf("TimedCheckVideoTask SetVideoTaskStatus failed, resourceId: %s, err: %v", resourceId, err)
							}

							continue
						}

						// Set video task success
						c.translation.VideoTaskCache.Delete(resourceId)
					}
				}
			}

			if len(c.translation.VideoTaskCache.videoTaskMap) == 0 {
				c.translation.VideoTaskCache.isActive = false
				return
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

func (c *Community) updateASRResult(ctx context.Context, resourceId string, task *VideoTask) error {
	// Request for ASR task
	asrTaskData, err := c.translation.Engine.QueryVideo2TextTask(task.TaskId)
	if err != nil {
		c.xLog.Errorf("TimedCheckVideoTask QueryVideo2TextTask failed, resourceId: %s, err: %v", resourceId, err)

		return err
	}

	if asrTaskData.Rtn == 0 && asrTaskData.Data.StatusCode == 3 {
		// Upload ASR result
		buffer, err := c.translation.Engine.GenerateWebVTTBytesWithTranslation(*asrTaskData, language.Chinese, language.English)
		if err != nil {
			c.xLog.Errorf("TimedCheckVideoTask GenerateWebVTTBytesWithTranslation failed, resourceId: %s, err: %v", resourceId, err)
			// Can not parse ASR result
			// continue

			return err
		}

		// Upload ASR result
		captionId, err := c.SaveMedia(ctx, task.UserId, buffer.Bytes(), ".vtt")
		if err != nil {
			c.xLog.Errorf("TimedCheckVideoTask SaveMedia failed, resourceId: %s, err: %v", resourceId, err)
			// Can not save ASR result
			return err
		}

		// Get ASR result id
		output, err := c.GetMediaUrl(ctx, fmt.Sprintf("%d", captionId))
		if err != nil {
			c.xLog.Errorf("TimedCheckVideoTask GetMediaURL failed, resourceId: %s, err: %v", resourceId, err)
			// Can not get ASR result link
			return err
		}

		err = c.SetVideoTaskSuccess(ctx, resourceId)
		if err != nil {
			c.xLog.Errorf("TimedCheckVideoTask SetVideoTaskStatus failed, resourceId: %s, err: %v", resourceId, err)

			return err
		}
		c.xLog.Infof("TimedCheckVideoTask GetMediaURL success, resourceId: %s, output: %s", resourceId, output)

		// Update status of video task
		err = c.SetVideoTaskOutput(ctx, resourceId, output)
		if err != nil {
			c.xLog.Errorf("TimedCheckVideoTask SetVideoTaskOutput failed, resourceId: %s, err: %v", resourceId, err)

			return err
		}
	} else {
		// Update status of video task
		err = c.SetVideoTaskFailed(ctx, resourceId)
		if err != nil {
			c.xLog.Errorf("TimedCheckVideoTask SetVideoTaskFailed failed, resourceId: %s, err: %v", resourceId, err)

			return err
		}
	}

	return nil
}
