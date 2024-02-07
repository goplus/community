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
	"testing"
)

func Test_VideoTaskCache_Set(t *testing.T) {
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
			if got, ok := c.Get(tt.args.id); !ok || got != tt.want {
				t.Errorf("VideoTaskCache.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}
