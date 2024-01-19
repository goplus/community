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
	"testing"
	"time"
)

func TestPutUser(t *testing.T) {
	conf := &Config{}
	todo := context.TODO()
	community, err := New(todo, conf)

	if err != nil {
		t.Skip(err)
	}

	// test data
	user := &User{
		Name:     "Test User",
		Avatar:   "avatar1",
		Password: "password1",
		Ctime:    time.Now(),
		Mtime:    time.Now(),
	}
	userUpdate := &User{
		ID:       "3",
		Name:     "Update User",
		Avatar:   "avatar2",
		Password: "password2",
		Ctime:    time.Now(),
		Mtime:    time.Now(),
	}

	tests := []struct {
		user       *User
		expectedID string
	}{
		{user, "4"},       // insert
		{userUpdate, "3"}, // update
	}

	for _, tt := range tests {
		id, _ := community.PutUser(todo, tt.user)

		if id != tt.expectedID {
			t.Errorf("PutUser( %+v) returned ID %s, expected: %s", tt.user, id, tt.expectedID)
		}
	}
}

func TestGetUser(t *testing.T) {
	conf := &Config{}
	todo := context.TODO()
	community, err := New(todo, conf)

	if err != nil {
		t.Skip(err)
	}

	// test data
	tests := []struct {
		id            string
		expectedID    string
		expectedError error
	}{
		{"1", "1", nil},
		{"10", "", ErrNotExist},
	}

	for _, tt := range tests {
		article, err := community.GetUser(todo, tt.id)

		if article.ID != tt.expectedID {
			t.Errorf("GetUser(%s) returned id is %s, expected: %s", tt.id, article.ID, tt.expectedID)
		}
		if err != tt.expectedError {
			t.Errorf("GetUser(%s) returned err is %s, expected: %s", tt.id, err, tt.expectedError)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	conf := &Config{}
	todo := context.TODO()
	community, err := New(todo, conf)

	if err != nil {
		t.Skip(err)
	}

	// test data
	tests := []struct {
		id          string
		expectedErr error
	}{
		{"1", nil},
		{"6", nil},
	}

	for _, tt := range tests {
		err := community.DeleteUser(todo, tt.id)
		if err != tt.expectedErr {
			t.Errorf("DeleteUser(%s) returned error: %v, expected: %v", tt.id, err, tt.expectedErr)
		}
		fmt.Println(err)
	}
}
