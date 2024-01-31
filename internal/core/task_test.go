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
	"database/sql"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/qiniu/x/xlog"
	"gocloud.dev/blob"
)

func TestCommunity_createVideoTask(t *testing.T) {
	type fields struct {
		bucket        *blob.Bucket
		db            *sql.DB
		domain        string
		casdoorConfig *CasdoorConfig
		xLog          *xlog.Logger
		taskMap       *VideoTaskCache
	}
	type args struct {
		ctx        context.Context
		resourceId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Community{
				bucket:        tt.fields.bucket,
				db:            tt.fields.db,
				domain:        tt.fields.domain,
				casdoorConfig: tt.fields.casdoorConfig,
				xLog:          tt.fields.xLog,
				taskMap:       tt.fields.taskMap,
			}
			if err := c.createVideoTask(tt.args.ctx, tt.args.resourceId); (err != nil) != tt.wantErr {
				t.Errorf("Community.createVideoTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommunity_updateVideoTaskOutput(t *testing.T) {
	type fields struct {
		bucket        *blob.Bucket
		db            *sql.DB
		domain        string
		casdoorConfig *CasdoorConfig
		xLog          *xlog.Logger
		taskMap       *VideoTaskCache
	}
	type args struct {
		ctx        context.Context
		resourceId string
		output     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Community{
				bucket:        tt.fields.bucket,
				db:            tt.fields.db,
				domain:        tt.fields.domain,
				casdoorConfig: tt.fields.casdoorConfig,
				xLog:          tt.fields.xLog,
				taskMap:       tt.fields.taskMap,
			}
			if err := c.updateVideoTaskOutput(tt.args.ctx, tt.args.resourceId, tt.args.output); (err != nil) != tt.wantErr {
				t.Errorf("Community.updateVideoTaskOutput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommunity_updateVideoTaskStatus(t *testing.T) {
	type fields struct {
		bucket        *blob.Bucket
		db            *sql.DB
		domain        string
		casdoorConfig *CasdoorConfig
		xLog          *xlog.Logger
		taskMap       *VideoTaskCache
	}
	type args struct {
		ctx        context.Context
		resourceId string
		status     int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Community{
				bucket:        tt.fields.bucket,
				db:            tt.fields.db,
				domain:        tt.fields.domain,
				casdoorConfig: tt.fields.casdoorConfig,
				xLog:          tt.fields.xLog,
				taskMap:       tt.fields.taskMap,
			}
			if err := c.updateVideoTaskStatus(tt.args.ctx, tt.args.resourceId, tt.args.status); (err != nil) != tt.wantErr {
				t.Errorf("Community.updateVideoTaskStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommunity_deleteVideoTask(t *testing.T) {
	type fields struct {
		bucket        *blob.Bucket
		db            *sql.DB
		domain        string
		casdoorConfig *CasdoorConfig
		xLog          *xlog.Logger
		taskMap       *VideoTaskCache
	}
	type args struct {
		ctx        context.Context
		resourceId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Community{
				bucket:        tt.fields.bucket,
				db:            tt.fields.db,
				domain:        tt.fields.domain,
				casdoorConfig: tt.fields.casdoorConfig,
				xLog:          tt.fields.xLog,
				taskMap:       tt.fields.taskMap,
			}
			if err := c.deleteVideoTask(tt.args.ctx, tt.args.resourceId); (err != nil) != tt.wantErr {
				t.Errorf("Community.deleteVideoTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommunity_getVideoTask(t *testing.T) {
	type fields struct {
		bucket        *blob.Bucket
		db            *sql.DB
		domain        string
		casdoorConfig *CasdoorConfig
		xLog          *xlog.Logger
		taskMap       *VideoTaskCache
	}
	type args struct {
		ctx        context.Context
		resourceId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *VideoTask
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Community{
				bucket:        tt.fields.bucket,
				db:            tt.fields.db,
				domain:        tt.fields.domain,
				casdoorConfig: tt.fields.casdoorConfig,
				xLog:          tt.fields.xLog,
				taskMap:       tt.fields.taskMap,
			}
			got, err := c.getVideoTask(tt.args.ctx, tt.args.resourceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Community.getVideoTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Community.getVideoTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommunity_NewVideoTask(t *testing.T) {
	type fields struct {
		bucket        *blob.Bucket
		db            *sql.DB
		domain        string
		casdoorConfig *CasdoorConfig
		xLog          *xlog.Logger
		taskMap       *VideoTaskCache
	}
	type args struct {
		ctx        context.Context
		resourceId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Community{
				bucket:        tt.fields.bucket,
				db:            tt.fields.db,
				domain:        tt.fields.domain,
				casdoorConfig: tt.fields.casdoorConfig,
				xLog:          tt.fields.xLog,
				taskMap:       tt.fields.taskMap,
			}
			if err := c.NewVideoTask(tt.args.ctx, tt.args.resourceId); (err != nil) != tt.wantErr {
				t.Errorf("Community.NewVideoTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommunity_SetVideoTaskSuccess(t *testing.T) {
	type fields struct {
		bucket        *blob.Bucket
		db            *sql.DB
		domain        string
		casdoorConfig *CasdoorConfig
		xLog          *xlog.Logger
		taskMap       *VideoTaskCache
	}
	type args struct {
		ctx        context.Context
		resourceId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Community{
				bucket:        tt.fields.bucket,
				db:            tt.fields.db,
				domain:        tt.fields.domain,
				casdoorConfig: tt.fields.casdoorConfig,
				xLog:          tt.fields.xLog,
				taskMap:       tt.fields.taskMap,
			}
			if err := c.SetVideoTaskSuccess(tt.args.ctx, tt.args.resourceId); (err != nil) != tt.wantErr {
				t.Errorf("Community.SetVideoTaskSuccess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommunity_SetVideoTaskFailed(t *testing.T) {
	type fields struct {
		bucket        *blob.Bucket
		db            *sql.DB
		domain        string
		casdoorConfig *CasdoorConfig
		xLog          *xlog.Logger
		taskMap       *VideoTaskCache
	}
	type args struct {
		ctx        context.Context
		resourceId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Community{
				bucket:        tt.fields.bucket,
				db:            tt.fields.db,
				domain:        tt.fields.domain,
				casdoorConfig: tt.fields.casdoorConfig,
				xLog:          tt.fields.xLog,
				taskMap:       tt.fields.taskMap,
			}
			if err := c.SetVideoTaskFailed(tt.args.ctx, tt.args.resourceId); (err != nil) != tt.wantErr {
				t.Errorf("Community.SetVideoTaskFailed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommunity_SetVideoTaskOutput(t *testing.T) {
	type fields struct {
		bucket        *blob.Bucket
		db            *sql.DB
		domain        string
		casdoorConfig *CasdoorConfig
		xLog          *xlog.Logger
		taskMap       *VideoTaskCache
	}
	type args struct {
		ctx        context.Context
		resourceId string
		output     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Community{
				bucket:        tt.fields.bucket,
				db:            tt.fields.db,
				domain:        tt.fields.domain,
				casdoorConfig: tt.fields.casdoorConfig,
				xLog:          tt.fields.xLog,
				taskMap:       tt.fields.taskMap,
			}
			if err := c.SetVideoTaskOutput(tt.args.ctx, tt.args.resourceId, tt.args.output); (err != nil) != tt.wantErr {
				t.Errorf("Community.SetVideoTaskOutput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewVideoTaskCache(t *testing.T) {
	tests := []struct {
		name string
		want *VideoTaskCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVideoTaskCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVideoTaskCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVideoTaskCache_Get(t *testing.T) {
	type fields struct {
		RWMutex      sync.RWMutex
		videoTaskMap VideoTaskMap
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   VideoTaskTimestamp
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &VideoTaskCache{
				RWMutex:      tt.fields.RWMutex,
				videoTaskMap: tt.fields.videoTaskMap,
			}
			got, got1 := c.Get(tt.args.key)
			if got != tt.want {
				t.Errorf("VideoTaskCache.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("VideoTaskCache.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestVideoTaskCache_Set(t *testing.T) {
	type fields struct {
		RWMutex      sync.RWMutex
		videoTaskMap VideoTaskMap
	}
	type args struct {
		key   string
		value VideoTaskTimestamp
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &VideoTaskCache{
				RWMutex:      tt.fields.RWMutex,
				videoTaskMap: tt.fields.videoTaskMap,
			}
			c.Set(tt.args.key, tt.args.value)
		})
	}
}

func TestVideoTaskCache_Delete(t *testing.T) {
	type fields struct {
		RWMutex      sync.RWMutex
		videoTaskMap VideoTaskMap
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &VideoTaskCache{
				RWMutex:      tt.fields.RWMutex,
				videoTaskMap: tt.fields.videoTaskMap,
			}
			c.Delete(tt.args.key)
		})
	}
}

func TestVideoTaskCache_Clear(t *testing.T) {
	type fields struct {
		RWMutex      sync.RWMutex
		videoTaskMap VideoTaskMap
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "test",
			fields: fields{
				RWMutex:      sync.RWMutex{},
				videoTaskMap: VideoTaskMap{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &VideoTaskCache{
				RWMutex:      tt.fields.RWMutex,
				videoTaskMap: tt.fields.videoTaskMap,
			}
			c.Clear()
		})
	}
}

func TestCommunity_TimedCheckVideoTask(t *testing.T) {
	type fields struct {
		bucket        *blob.Bucket
		db            *sql.DB
		domain        string
		casdoorConfig *CasdoorConfig
		xLog          *xlog.Logger
		taskMap       *VideoTaskCache
	}
	type args struct {
		ctx        context.Context
		resourceId string
		timeout    time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Community{
				bucket:        tt.fields.bucket,
				db:            tt.fields.db,
				domain:        tt.fields.domain,
				casdoorConfig: tt.fields.casdoorConfig,
				xLog:          tt.fields.xLog,
				taskMap:       tt.fields.taskMap,
			}
			got, err := c.TimedCheckVideoTask(tt.args.ctx, tt.args.resourceId, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("Community.TimedCheckVideoTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Community.TimedCheckVideoTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
