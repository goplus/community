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

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestParseJwtToken(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	tests := []struct {
		token  string
		userId string
	}{
		{
			token:  "123",
			userId: "1",
		},
	}

	for _, test := range tests {
		userId, err := community.ParseJwtToken(test.token)
		assert.Nil(t, err)
		assert.Equal(t, test.userId, userId)
	}
}

func TestCasdoorConfigInit(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	CasdoorConfigInit()
}

func TestGetUser(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	tests := []struct {
		token string
		name  string
	}{
		{
			token: "test",
			name:  "test",
		},
	}

	for _, test := range tests {
		// Replace CasdoorSDKService with mock
		community.CasdoorSDKService = &mockCasdoorSDKService{
			DoGetUser: func(name string) (user *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoParseJwtToken: func(token string) (claims *casdoorsdk.Claims, err error) {
				return &casdoorsdk.Claims{
					User: casdoorsdk.User{
						Id:          "1",
						DisplayName: "test",
					},
				}, nil
			},
			DoGetUserClaim: func(uid string) (claim *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoGetUserByUserId: func(uid string) (user *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoGetUserById: func(uid string) (user *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoUpdateUser: func(user *UserInfo) (res bool, err error) {
				return true, nil
			},
			DoUpdateUserById: func(uid string, user *UserInfo) (res bool, err error) {
				return true, nil
			},
			DoGetOAuthToken: func(code string, state string) (token *oauth2.Token, err error) {
				return &oauth2.Token{
					AccessToken: "access_token",
				}, nil
			},
			DoGetApplication: func(name string) (*casdoorsdk.Application, error) {
				return &casdoorsdk.Application{
					Name: "test",
				}, nil
			},
		}

		user, err := community.GetUser(test.token)
		assert.Nil(t, err)

		assert.Equal(t, test.name, user.Name)
	}
}

func TestGetUserClaim(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	tests := []struct {
		userId string
		name   string
	}{
		{
			userId: "1",
			name:   "test",
		},
	}

	for _, test := range tests {
		// Replace CasdoorSDKService with mock
		community.CasdoorSDKService = &mockCasdoorSDKService{
			DoGetUser: func(name string) (user *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoParseJwtToken: func(token string) (claims *casdoorsdk.Claims, err error) {
				return &casdoorsdk.Claims{
					User: casdoorsdk.User{
						Id:          "1",
						DisplayName: "test",
					},
				}, nil
			},
			DoGetUserClaim: func(uid string) (claim *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoGetUserByUserId: func(uid string) (user *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoGetUserById: func(uid string) (user *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoUpdateUser: func(user *UserInfo) (res bool, err error) {
				return true, nil
			},
			DoUpdateUserById: func(uid string, user *UserInfo) (res bool, err error) {
				return true, nil
			},
			DoGetOAuthToken: func(code string, state string) (token *oauth2.Token, err error) {
				return &oauth2.Token{
					AccessToken: "access_token",
				}, nil
			},
			DoGetApplication: func(name string) (*casdoorsdk.Application, error) {
				return &casdoorsdk.Application{
					Name: "test",
				}, nil
			},
		}

		user, err := community.GetUserClaim(test.userId)
		assert.Nil(t, err)
		assert.Equal(t, test.name, user.DisplayName)
	}
}

func TestUpdateUserById(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	tests := []struct {
		userId string
		user   *UserInfo
	}{
		{
			userId: "1",
			user: &UserInfo{
				Id:          "1",
				DisplayName: "test",
			},
		},
	}

	for _, test := range tests {
		// Replace CasdoorSDKService with mock
		community.CasdoorSDKService = &mockCasdoorSDKService{
			DoGetUser: func(name string) (user *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoParseJwtToken: func(token string) (claims *casdoorsdk.Claims, err error) {
				return &casdoorsdk.Claims{
					User: casdoorsdk.User{
						Id:          "1",
						DisplayName: "test",
					},
				}, nil
			},
			DoGetUserClaim: func(uid string) (claim *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoGetUserByUserId: func(uid string) (user *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoGetUserById: func(uid string) (user *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoUpdateUser: func(user *UserInfo) (res bool, err error) {
				return true, nil
			},
			DoUpdateUserById: func(uid string, user *UserInfo) (res bool, err error) {
				return true, nil
			},
			DoGetOAuthToken: func(code string, state string) (token *oauth2.Token, err error) {
				return &oauth2.Token{
					AccessToken: "access_token",
				}, nil
			},
			DoGetApplication: func(name string) (*casdoorsdk.Application, error) {
				return &casdoorsdk.Application{
					Name: "test",
				}, nil
			},
		}

		res, err := community.UpdateUserById(test.userId, test.user)
		assert.Nil(t, err)
		assert.True(t, res)
	}
}

func TestUpdateUser(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	tests := []struct {
		userId string
		user   *UserInfo
	}{
		{
			userId: "1",
			user: &UserInfo{
				Id:          "1",
				DisplayName: "test",
			},
		},
	}

	for _, test := range tests {
		// Replace CasdoorSDKService with mock
		community.CasdoorSDKService = &mockCasdoorSDKService{
			DoGetUser: func(name string) (user *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoParseJwtToken: func(token string) (claims *casdoorsdk.Claims, err error) {
				return &casdoorsdk.Claims{
					User: casdoorsdk.User{
						Id:          "1",
						DisplayName: "test",
					},
				}, nil
			},
			DoGetUserClaim: func(uid string) (claim *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoGetUserByUserId: func(uid string) (user *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoGetUserById: func(uid string) (user *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoUpdateUser: func(user *UserInfo) (res bool, err error) {
				return true, nil
			},
			DoUpdateUserById: func(uid string, user *UserInfo) (res bool, err error) {
				return true, nil
			},
			DoGetOAuthToken: func(code string, state string) (token *oauth2.Token, err error) {
				return &oauth2.Token{
					AccessToken: "access_token",
				}, nil
			},
			DoGetApplication: func(name string) (*casdoorsdk.Application, error) {
				return &casdoorsdk.Application{
					Name: "test",
				}, nil
			},
		}

		res, err := community.UpdateUser(test.user)
		assert.Nil(t, err)
		assert.True(t, res)
	}
}

func TestGetOAuthToken(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	tests := []struct {
		code  string
		state string
	}{
		{
			code:  "123",
			state: "123",
		},
	}

	for _, test := range tests {
		// Replace CasdoorSDKService with mock
		community.CasdoorSDKService = &mockCasdoorSDKService{
			DoGetUser: func(name string) (user *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoParseJwtToken: func(token string) (claims *casdoorsdk.Claims, err error) {
				return &casdoorsdk.Claims{
					User: casdoorsdk.User{
						Id:          "1",
						DisplayName: "test",
					},
				}, nil
			},
			DoGetUserClaim: func(uid string) (claim *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoGetUserByUserId: func(uid string) (user *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoGetUserById: func(uid string) (user *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoUpdateUser: func(user *UserInfo) (res bool, err error) {
				return true, nil
			},
			DoUpdateUserById: func(uid string, user *UserInfo) (res bool, err error) {
				return true, nil
			},
			DoGetOAuthToken: func(code string, state string) (token *oauth2.Token, err error) {
				return &oauth2.Token{
					AccessToken: "access_token",
				}, nil
			},
			DoGetApplication: func(name string) (*casdoorsdk.Application, error) {
				return &casdoorsdk.Application{
					Name: "test",
				}, nil
			},
		}

		token, err := community.GetOAuthToken(test.code, test.state)
		assert.Nil(t, err)
		assert.NotNil(t, token)
	}
}

func TestGetUserById(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	tests := []struct {
		userId string
	}{
		{
			userId: "1",
		},
	}

	for _, test := range tests {
		// Replace CasdoorSDKService with mock
		community.CasdoorSDKService = &mockCasdoorSDKService{
			DoGetUser: func(name string) (user *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoParseJwtToken: func(token string) (claims *casdoorsdk.Claims, err error) {
				return &casdoorsdk.Claims{
					User: casdoorsdk.User{
						Id:          "1",
						DisplayName: "test",
					},
				}, nil
			},
			DoGetUserClaim: func(uid string) (claim *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoGetUserByUserId: func(uid string) (user *casdoorsdk.User, err error) {
				return nil, nil
			},
			DoGetUserById: func(uid string) (user *casdoorsdk.User, err error) {
				return &casdoorsdk.User{
					Id:          "1",
					DisplayName: "test",
				}, nil
			},
			DoUpdateUser: func(user *UserInfo) (res bool, err error) {
				return true, nil
			},
			DoUpdateUserById: func(uid string, user *UserInfo) (res bool, err error) {
				return true, nil
			},
			DoGetOAuthToken: func(code string, state string) (token *oauth2.Token, err error) {
				return &oauth2.Token{
					AccessToken: "access_token",
				}, nil
			},
			DoGetApplication: func(name string) (*casdoorsdk.Application, error) {
				return &casdoorsdk.Application{
					Name: "test",
				}, nil
			},
		}

		user, err := community.GetUserById(test.userId)
		assert.Nil(t, err)
		assert.NotNil(t, user)
	}
}
