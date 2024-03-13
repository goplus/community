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
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/goplus/community/translation"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/stretchr/testify/assert"
)

// Mock API Client
type mockQiNiuSDKClient struct {
	DoFunc func(ctx context.Context, cred *auth.Credentials, tokenType auth.TokenType, ret interface{},
		method, reqUrl string, headers http.Header, param interface{}) error
}

// Implement NiuTransClient interface
func (m *mockQiNiuSDKClient) CredentialedCallWithJson(ctx context.Context, cred *auth.Credentials, tokenType auth.TokenType, ret interface{},
	method, reqUrl string, headers http.Header, param interface{}) error {
	return m.DoFunc(ctx, cred, tokenType, ret, method, reqUrl, headers, param)
}

// For single benchmark
func Test_VideoTaskCache_Set(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	type fields struct {
		key string
		val VideoTaskTimestamp
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   VideoTaskTimestamp
	}{
		{
			name: "Test_VideoTaskCache_Set",
			fields: fields{
				key: "1",
				val: 111,
			},
			args: args{
				id: "1",
			},
			want: 111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewVideoTaskCache()
			c.Set(tt.args.id, tt.fields.val)
			if got, ok := c.Get(tt.args.id); !ok || got != tt.want {
				t.Errorf("VideoTaskCache.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_VideoTaskCache_Set(b *testing.B) {
	c := NewVideoTaskCache()
	for i := 0; i < b.N; i++ {
		c.Set(fmt.Sprint(i), 111)
	}
}

func Test_VideoTaskCache_Get(t *testing.T) {
	type fields struct {
		key string
		val VideoTaskTimestamp
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   VideoTaskTimestamp
	}{
		{
			name: "Test_VideoTaskCache_Get",
			fields: fields{
				key: "1",
				val: 111,
			},
			args: args{
				id: "1",
			},
			want: 111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewVideoTaskCache()
			c.Set(tt.args.id, tt.fields.val)
			if got, ok := c.Get(tt.args.id); !ok || got != tt.want {
				t.Errorf("VideoTaskCache.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_VideoTaskCache_Get(b *testing.B) {
	c := NewVideoTaskCache()
	// Add b.N items to the cache
	for i := 0; i < b.N; i++ {
		c.Set(fmt.Sprint(i), 111)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.Get(fmt.Sprint(i))
	}
}

func Test_VideoTaskCache_Delete(t *testing.T) {
	type fields struct {
		key string
		val VideoTaskTimestamp
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   VideoTaskTimestamp
	}{
		{
			name: "Test_VideoTaskCache_Delete",
			fields: fields{
				key: "1",
				val: 111,
			},
			args: args{
				id: "1",
			},
			want: 111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewVideoTaskCache()
			c.Set(tt.args.id, tt.fields.val)
			c.Delete(tt.args.id)
			if got, ok := c.Get(tt.args.id); ok {
				t.Errorf("VideoTaskCache.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_VideoTaskCache_Delete(b *testing.B) {
	c := NewVideoTaskCache()
	// Add b.N items to the cache
	for i := 0; i < b.N; i++ {
		c.Set(fmt.Sprint(i), 111)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.Delete(fmt.Sprint(i))
	}
}

func Test_VideoTaskCache_Clear(t *testing.T) {
	type fields struct {
		key string
		val VideoTaskTimestamp
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   VideoTaskTimestamp
	}{
		{
			name: "Test_VideoTaskCache_Clear",
			fields: fields{
				key: "1",
				val: 111,
			},
			args: args{
				id: "1",
			},
			want: 111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewVideoTaskCache()
			c.Set(tt.args.id, tt.fields.val)
			c.Clear()
			if got, ok := c.Get(tt.args.id); ok {
				t.Errorf("VideoTaskCache.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_VideoTaskCache_Clear(b *testing.B) {
	c := NewVideoTaskCache()
	// Add b.N items to the cache
	for i := 0; i < b.N; i++ {
		c.Clear()
	}
}

// For logic benchmark
func Benchmark_VideoTaskCache_CheckVideoTask(b *testing.B) {
	c := NewVideoTaskCache()
	// Add b.N items to the cache
	for i := 0; i < b.N; i++ {
		c.Set(fmt.Sprint(i), VideoTaskTimestamp(rand.Intn(9999)))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Iterate over the cache
		for k, v := range c.videoTaskMap {
			// Mock some work
			time.Sleep(100 * time.Microsecond)

			// Delete the item from the cache with a 20% probability
			if v%5 == 0 {
				c.Delete(k)
			}
		}
	}
}

func TestNewVideoTask(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare test data
	// Prepare data
	community.db.Exec(
		"insert into file (file_key,format,size,user_id,create_at,update_at,duration) values (?,?,?,?,?,?,?)",
		"test",
		"test",
		1,
		"1",
		"2021-01-01",
		"2021-01-01",
		1,
	)

	tests := []struct {
		userId     string
		resourceId string
	}{
		{
			userId:     "1",
			resourceId: "1",
		},
	}

	for _, test := range tests {
		community.translation.Engine.QiNiuSDKClient = &mockQiNiuSDKClient{
			DoFunc: func(ctx context.Context, cred *auth.Credentials, tokenType auth.TokenType, ret interface{},
				method, reqUrl string, headers http.Header, param interface{}) error {
				return nil
			},
		}

		err := community.NewVideoTask(context.Background(), test.userId, test.resourceId)
		assert.Nil(t, err)
	}
}

func TestTimedCheckVideoTask(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare test data
	// Prepare data
	community.db.Exec(
		"insert into file (file_key,format,size,user_id,create_at,update_at,duration) values (?,?,?,?,?,?,?)",
		"test",
		"test",
		1,
		"1",
		"2021-01-01",
		"2021-01-01",
		1,
	)

	// Set video task cache
	community.SetVideoTaskCache("1", VideoTaskTimestamp(time.Now().Unix()))

	// Mock API Client
	community.translation.Engine.QiNiuSDKClient = &mockQiNiuSDKClient{
		DoFunc: func(ctx context.Context, cred *auth.Credentials, tokenType auth.TokenType, ret interface{},
			method, reqUrl string, headers http.Header, param interface{}) error {
			return nil
		},
	}

	// Timed check video task
	community.TimedCheckVideoTask(context.Background(), time.Duration(1000))
}

func Test_getVideoTask(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	community.createVideoTask(context.Background(), "1", "1", "1")

	tests := []struct {
		resourceId string
	}{
		{
			resourceId: "1",
		},
	}

	for _, test := range tests {
		task, err := community.getVideoTask(context.Background(), test.resourceId)
		assert.Nil(t, err)
		assert.NotNil(t, task)
	}
}

func Test_updateVideoTaskOutput(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	community.createVideoTask(context.Background(), "1", "1", "1")

	tests := []struct {
		resourceId string
		output     string
	}{
		{
			resourceId: "1",
			output:     "test",
		},
	}

	for _, test := range tests {
		err := community.updateVideoTaskOutput(context.Background(), test.resourceId, test.output)
		assert.Nil(t, err)
	}
}

func Test_updateVideoTaskStatus(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	community.createVideoTask(context.Background(), "1", "1", "1")

	tests := []struct {
		resourceId string
		status     int
	}{
		{
			resourceId: "1",
			status:     1,
		},
	}

	for _, test := range tests {
		err := community.updateVideoTaskStatus(context.Background(), test.resourceId, test.status)
		assert.Nil(t, err)
	}
}

func TestSetVideoTaskSuccess(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	community.createVideoTask(context.Background(), "1", "1", "1")

	tests := []struct {
		resourceId string
	}{
		{
			resourceId: "1",
		},
	}

	for _, test := range tests {
		err := community.SetVideoTaskSuccess(context.Background(), test.resourceId)
		assert.Nil(t, err)
	}
}

func TestSetVideoTaskOutput(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	community.createVideoTask(context.Background(), "1", "1", "1")

	tests := []struct {
		resourceId string
		output     string
	}{
		{
			resourceId: "1",
			output:     "test",
		},
	}

	for _, test := range tests {
		err := community.SetVideoTaskOutput(context.Background(), test.resourceId, test.output)
		assert.Nil(t, err)
	}
}

func Test_deleteVideoTask(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	community.createVideoTask(context.Background(), "1", "1", "1")

	tests := []struct {
		resourceId string
	}{
		{
			resourceId: "1",
		},
	}

	for _, test := range tests {
		err := community.deleteVideoTask(context.Background(), test.resourceId)
		assert.Nil(t, err)
	}
}

func Test_updateASRResult(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	community.createVideoTask(context.Background(), "1", "1", "1")

	tests := []struct {
		resourceId string
		task       *VideoTask
	}{
		{
			resourceId: "1",
			task: &VideoTask{
				Id:         1,
				UserId:     "1",
				ResourceId: "1",
				TaskId:     "1",
				Output:     "test",
				Status:     1,
			},
		},
	}

	for _, test := range tests {
		community.translation.Engine.QiNiuSDKClient = &mockQiNiuSDKClient{
			DoFunc: func(ctx context.Context, cred *auth.Credentials, tokenType auth.TokenType, ret interface{},
				method, reqUrl string, headers http.Header, param interface{}) error {
				ret = &translation.ASRTaskData{
					Rtn: 0,
					Data: translation.Data{
						StatusCode: 3,
					},
				}

				return nil
			},
		}

		err := community.updateASRResult(context.Background(), test.resourceId, test.task)
		assert.Nil(t, err)
	}
}
