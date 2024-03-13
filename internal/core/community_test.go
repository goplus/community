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
	"net/http"
	"testing"
	"time"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/stretchr/testify/assert"
	"gocloud.dev/blob"
	"golang.org/x/oauth2"
	language "golang.org/x/text/language"
)

var (
	community *Community
	todo      context.Context

	_ CasdoorSDKService = &mockCasdoorSDKService{}
	_ S3Service         = &mockS3Service{}
	_ S3Reader          = &mockS3Reader{}
	_ S3Writer          = &mockS3Writer{}
)

// Mock S3Service
type mockS3Service struct {
	DoNewReader func(ctx context.Context, key string, opts *blob.ReaderOptions) (_ S3Reader, err error)
	DoNewWriter func(ctx context.Context, key string, opts *blob.WriterOptions) (_ S3Writer, err error)
	DoDelete    func(ctx context.Context, key string) (err error)
}

func (m *mockS3Service) NewReader(ctx context.Context, key string, opts *blob.ReaderOptions) (_ S3Reader, err error) {
	return m.DoNewReader(ctx, key, opts)
}

func (m *mockS3Service) NewWriter(ctx context.Context, key string, opts *blob.WriterOptions) (_ S3Writer, err error) {
	return m.DoNewWriter(ctx, key, opts)
}

func (m *mockS3Service) Delete(ctx context.Context, key string) (err error) {
	return m.DoDelete(ctx, key)
}

// Mock S3Reader
type mockS3Reader struct {
	DoRead        func(p []byte) (int, error)
	DoClose       func() error
	DoContentType func() string
	DoSize        func() int64
}

func (m *mockS3Reader) Read(p []byte) (int, error) {
	return m.DoRead(p)
}

func (m *mockS3Reader) Close() error {
	return m.DoClose()
}

func (m *mockS3Reader) ContentType() string {
	// fmt.Sprintf("mock S3Reader ContentType:%#v", m.DoContentType())
	return m.DoContentType()
	// return "video/mp4"
}

func (m *mockS3Reader) Size() int64 {
	return m.DoSize()
}

// Mock S3Writer
//
//lint:ignore U1000 This is a mock
type mockS3Writer struct {
	DoWrite func(p []byte) (n int, err error)
	DoClose func() (err error)
}

//lint:ignore U1000 This is a mock
func (m *mockS3Writer) Write(p []byte) (int, error) {
	return m.DoWrite(p)
}

//lint:ignore U1000 This is a mock
func (m *mockS3Writer) Close() error {
	return m.DoClose()
}

// Mock CasdoorSDKService
type mockCasdoorSDKService struct {
	DoGetUser               func(name string) (user *casdoorsdk.User, err error)
	DoGetUsers              func() ([]*casdoorsdk.User, error)
	DoParseJwtToken         func(token string) (claims *casdoorsdk.Claims, err error)
	DoGetUserClaim          func(uid string) (claim *casdoorsdk.User, err error)
	DoGetUserById           func(uid string) (user *casdoorsdk.User, err error)
	DoUpdateUserById        func(uid string, user *UserInfo) (res bool, err error)
	DoUpdateCasdoorUserById func(uid string, user *casdoorsdk.User) (res bool, err error)
	DoUpdateUser            func(user *UserInfo) (res bool, err error)
	DoGetUserByUserId       func(uid string) (user *casdoorsdk.User, err error)
	DoGetOAuthToken         func(code string, state string) (token *oauth2.Token, err error)
	DoGetApplication        func(name string) (*casdoorsdk.Application, error)
}

func (m *mockCasdoorSDKService) GetUser(name string) (user *casdoorsdk.User, err error) {
	return m.DoGetUser(name)
}

func (m *mockCasdoorSDKService) GetUsers() ([]*casdoorsdk.User, error) {
	return m.DoGetUsers()
}

func (m *mockCasdoorSDKService) ParseJwtToken(token string) (claims *casdoorsdk.Claims, err error) {
	return m.DoParseJwtToken(token)
}

func (m *mockCasdoorSDKService) GetUserClaim(uid string) (claim *casdoorsdk.User, err error) {
	return m.DoGetUserClaim(uid)
}

func (m *mockCasdoorSDKService) GetUserById(uid string) (user *casdoorsdk.User, err error) {
	return m.DoGetUserById(uid)
}

func (m *mockCasdoorSDKService) UpdateUserById(uid string, user *UserInfo) (res bool, err error) {
	return m.DoUpdateUserById(uid, user)
}

func (m *mockCasdoorSDKService) UpdateCasdoorUserById(uid string, user *casdoorsdk.User) (res bool, err error) {
	return m.DoUpdateCasdoorUserById(uid, user)
}

func (m *mockCasdoorSDKService) UpdateUser(user *UserInfo) (res bool, err error) {
	return m.DoUpdateUser(user)
}

func (m *mockCasdoorSDKService) GetUserByUserId(uid string) (user *casdoorsdk.User, err error) {
	return m.DoGetUserByUserId(uid)
}

func (m *mockCasdoorSDKService) GetOAuthToken(code string, state string) (token *oauth2.Token, err error) {
	return m.DoGetOAuthToken(code, state)
}

func (m *mockCasdoorSDKService) GetApplication(name string) (*casdoorsdk.Application, error) {
	return m.DoGetApplication(name)
}

// MockWriter for blob.Writer
type mockWriter struct {
	DoWrite func(p []byte) (n int, err error)
	DoClose func() (err error)
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	return m.DoWrite(p)
}

func (m *mockWriter) Close() (err error) {
	return m.DoClose()
}

func TestMain(m *testing.M) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	m.Run()
}

func initClient() {
	conf := &Config{
		DBConfig: DBConfig{
			DSN:    ":memory:",
			Driver: "sqlite3",
		},
		QiNiuConfig: QiNiuConfig{
			BlobUS: "kodo://123:456@gop-community?useHttps",
		},
	}
	todo = context.TODO()
	comm, err := New(todo, conf)
	if err != nil {
		panic(err)
	}
	community = comm

	// Replace S3Service with mock
	community.S3Service = &mockS3Service{
		DoNewReader: func(ctx context.Context, key string, opts *blob.ReaderOptions) (_ S3Reader, err error) {
			return &mockS3Reader{
				DoRead: func(p []byte) (n int, err error) {
					return 0, nil
				},
				DoClose: func() (err error) {
					return nil
				},
				DoContentType: func() (contentType string) {
					return "video/mp4"
				},
				DoSize: func() (size int64) {
					return 0
				},
			}, nil
		},
		DoNewWriter: func(ctx context.Context, key string, opts *blob.WriterOptions) (_ S3Writer, err error) {
			return &mockWriter{
				DoWrite: func(p []byte) (n int, err error) {
					return 0, nil
				},
				DoClose: func() (err error) {
					return nil
				},
			}, nil
		},
		DoDelete: func(ctx context.Context, key string) (err error) {
			return nil
		},
	}
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

	// Replace S3Service with mock
	community.S3Service = &mockS3Service{
		DoNewReader: func(ctx context.Context, key string, opts *blob.ReaderOptions) (_ S3Reader, err error) {
			return &mockS3Reader{
				DoRead: func(p []byte) (n int, err error) {
					return 0, nil
				},
				DoClose: func() (err error) {
					return nil
				},
				DoContentType: func() (contentType string) {
					return ""
				},
				DoSize: func() (size int64) {
					return 0
				},
			}, nil
		},
		DoNewWriter: func(ctx context.Context, key string, opts *blob.WriterOptions) (_ S3Writer, err error) {
			return &mockWriter{
				DoWrite: func(p []byte) (n int, err error) {
					return 0, nil
				},
				DoClose: func() (err error) {
					return nil
				},
			}, nil
		},
		DoDelete: func(ctx context.Context, key string) (err error) {
			return nil
		},
	}
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
}

func initDB() {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS article (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title VARCHAR(255),
			user_id VARCHAR(255) NOT NULL,
			tags VARCHAR(255),
			cover VARCHAR(255),
			content TEXT,
			trans_content TEXT,
			html_id INTEGER,
			trans_html_id INTEGER,
			ctime DATETIME,
			mtime DATETIME DEFAULT CURRENT_TIMESTAMP,
			abstract VARCHAR(511),
			trans TINYINT(1) DEFAULT 0,
			trans_tags VARCHAR(255),
			trans_abstract VARCHAR(511),
			trans_title VARCHAR(255),
			like_count INTEGER NOT NULL DEFAULT 0,
			label VARCHAR(255) DEFAULT 'article',
			view_count INTEGER NOT NULL DEFAULT 0,
			vtt_id VARCHAR(255)
		);`,
		`CREATE TABLE IF NOT EXISTS video_task (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			resource_id TEXT,
			output TEXT,
			status INTEGER,
			create_at DATETIME,
			update_at DATETIME,
			user_id TEXT,
			task_id TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS video_subtitle (
			video_id INTEGER NOT NULL,
			subtitle_id INTEGER,
			user_id INTEGER NOT NULL,
			language TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS share (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			platform TEXT NOT NULL,
			user_id INTEGER,
			ip TEXT,
			index_key TEXT UNIQUE,
			create_at TEXT DEFAULT (datetime('now','localtime'))
		);`,
		`CREATE TABLE IF NOT EXISTS file (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			create_at DATETIME,
			update_at DATETIME,
			file_key TEXT,
			format TEXT,
			user_id TEXT,
			size BIGINT,
			duration TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS article_like (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			article_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL
		);
		CREATE UNIQUE INDEX idx_article_id_user_id ON article_like(article_id, user_id);
		`,
		`CREATE TABLE article_view (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			ip TEXT NOT NULL,
			user_id INTEGER NOT NULL DEFAULT 0,
			article_id INTEGER NOT NULL,
			created_at TEXT NOT NULL, 
			"index" TEXT,
			platform TEXT NOT NULL DEFAULT ''
		);`,
	}

	for _, table := range tables {
		_, err := community.db.Exec(table)
		if err != nil {
			panic(err)
		}
	}
}

func TestPutArticle(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// test data
	article := &Article{
		ArticleEntry: ArticleEntry{
			Title: "Test Article",
			UId:   "1",
			Cover: "cover1",
			Tags:  "tag1",
			Ctime: time.Now(),
			Mtime: time.Now(),
		},
		Content: "This is a test article.",
	}
	articleUpdate := &Article{
		ArticleEntry: ArticleEntry{
			ID:    "1",
			Title: "Test Article",
			UId:   "1",
			Cover: "cover1",
			Tags:  "tag1",
			Ctime: time.Now(),
			Mtime: time.Now(),
		},
		Content: "This is a test article.",
	}

	tests := []struct {
		uid        string
		article    *Article
		expectedID string
	}{
		{"1", article, "1"},       // insert
		{"1", article, "2"},       // insert trans
		{"1", articleUpdate, "1"}, // update
	}

	for _, tt := range tests {
		id, err := community.PutArticle(todo, tt.uid, tt.article)
		assert.Nil(t, err)
		assert.Equal(t, tt.expectedID, id)
	}
}

func TestCanEditable(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	article := &Article{
		ArticleEntry: ArticleEntry{
			Title: "Test Article",
			UId:   "1",
			Cover: "cover1",
			Tags:  "tag1",
			Ctime: time.Now(),
			Mtime: time.Now(),
		},
		Content: "This is a test article.",
	}
	_, err := community.PutArticle(todo, "1", article)
	assert.Nil(t, err)

	// test data
	tests := []struct {
		uid           string
		articleID     string
		expectedEdit  bool
		expectedError error
	}{
		{"1", "1", true, nil},
		{"2", "1", false, ErrPermission},
	}

	for _, tt := range tests {
		canEdit, err := community.CanEditable(todo, tt.uid, tt.articleID)
		assert.Equal(t, err, tt.expectedError)
		assert.Equal(t, tt.expectedEdit, canEdit)
	}
}

func TestArticle(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	article := &Article{
		ArticleEntry: ArticleEntry{
			Title: "Test Article",
			UId:   "1",
			Cover: "cover1",
			Tags:  "tag1",
			Ctime: time.Now(),
			Mtime: time.Now(),
		},
		Content: "This is a test article.",
		VttId:   "1",
	}
	_, err := community.PutArticle(todo, "1", article)
	assert.Nil(t, err)
	// Create video task
	_, err = community.db.ExecContext(
		todo,
		"insert into video_task (user_id, resource_id, task_id, output, status, create_at, update_at) values (?, ?, ?, ?, ?, ?, ?)",
		"1",
		"1",
		"1",
		"1",
		1,
		time.Now(),
		time.Now(),
	)
	assert.Nil(t, err)

	// Mock CasdoorSDKService data
	community.CasdoorSDKService = &mockCasdoorSDKService{
		DoGetUserByUserId: func(uid string) (user *casdoorsdk.User, err error) {
			return &casdoorsdk.User{
				Id:          "1",
				DisplayName: "test",
			}, nil
		},
	}

	// test data
	tests := []struct {
		id            string
		expectedID    string
		expectedError error
	}{
		{"1", "1", nil},
		// {"10", "", ErrNotExist},
	}

	for _, tt := range tests {
		article, err := community.Article(todo, tt.id)

		assert.Equal(t, tt.expectedError, err)
		assert.Equal(t, tt.expectedID, article.ID)
	}
}

func TestListArticle(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	for i := 0; i < 10; i++ {
		article := &Article{
			ArticleEntry: ArticleEntry{
				Title: "Test Article",
				UId:   fmt.Sprintf("%d", i),
				Cover: "cover1",
				Tags:  "tag1",
				Ctime: time.Now(),
				Mtime: time.Now(),
			},
			Content: "This is a test article.",
		}
		_, err := community.PutArticle(todo, fmt.Sprintf("%d", i), article)
		assert.Nil(t, err)
	}

	tests := []struct {
		from         string
		limit        int
		expectedLen  int
		expectedNext string
	}{
		{MarkBegin, 5, 5, "5"},
	}

	for _, tt := range tests {
		items, next, err := community.ListArticle(todo, tt.from, tt.limit, "", "")

		assert.Nil(t, err)
		assert.EqualValues(t, tt.expectedLen, len(items))
		assert.EqualValues(t, tt.expectedNext, next)
	}
}

func TestDeleteArticle(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	article := &Article{
		ArticleEntry: ArticleEntry{
			Title: "Test Article",
			UId:   "1",
			Cover: "cover1",
			Tags:  "tag1",
			Ctime: time.Now(),
			Mtime: time.Now(),
		},
		Content: "This is a test article.",
	}
	_, err := community.PutArticle(todo, "1", article)
	assert.Nil(t, err)

	// test data
	tests := []struct {
		uid         string
		articleID   string
		expectedErr error
	}{
		{"1", "1", nil},
	}

	for _, tt := range tests {
		err := community.DeleteArticle(todo, tt.uid, tt.articleID)
		assert.EqualValues(t, tt.expectedErr, err)
	}
}

func TestGetArticlesByUid(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// test data
	tests := []struct {
		from          string
		limit         int
		uid           string
		expectedError error
	}{
		{"", 10, "1", nil},
	}

	for _, tt := range tests {
		_, _, err := community.GetArticlesByUid(todo, tt.uid, tt.from, tt.limit)

		assert.EqualValues(t, tt.expectedError, err)
	}
}
func TestArticleLView(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// TODO: Update test data
	community.ArticleLView(todo, "14", "127.0.0.1", "73917397", "")
}

func TestArticleLike(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	article := &Article{
		ArticleEntry: ArticleEntry{
			Title: "Test Article",
			UId:   "1",
			Cover: "cover1",
			Tags:  "tag1",
			Ctime: time.Now(),
			Mtime: time.Now(),
		},
		Content: "This is a test article.",
	}
	_, err := community.PutArticle(todo, "1", article)
	assert.Nil(t, err)

	tests := []struct {
		uid           string
		articleID     int
		result        bool
		expectedError error
	}{
		{"1", 14, true, nil},
		{"1", 14, false, nil},
	}

	for _, tt := range tests {
		ok, err := community.ArticleLike(todo, tt.articleID, tt.uid)
		assert.Nil(t, err)
		assert.Equal(t, tt.result, ok)
	}
}

func TestRedirectToCasdoor(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// test data
	tests := []struct {
		redirectURL   string
		expectedError error
	}{
		{"http://123.com", nil},
	}

	for _, tt := range tests {
		loginURL := community.RedirectToCasdoor(tt.redirectURL)

		assert.NotEqualValues(t, "", loginURL)
	}
}

func TestGetAccessToken(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// test data
	tests := []struct {
		code          string
		state         string
		expectedError error
	}{
		{"123", "123", nil},
	}

	for _, tt := range tests {
		token, err := community.GetAccessToken(tt.code, tt.state)

		assert.EqualValues(t, tt.expectedError, err)
		assert.NotEqualValues(t, "", token)
	}
}

func TestShare(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// test data
	tests := []struct {
		platform      string
		uid           string
		ip            string
		articleId     string
		expectedError error
	}{
		{"facebook", "1", "1", "1", nil},
	}

	for _, tt := range tests {
		community.Share(tt.platform, tt.uid, tt.ip, tt.articleId)
	}
}

func TestGetApplicationInfo(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// test data
	tests := []struct {
		uid           string
		expectedError error
	}{
		{"1", nil},
	}

	for _, tt := range tests {
		app, err := community.GetApplicationInfo()

		assert.EqualValues(t, tt.expectedError, err)
		assert.NotEqualValues(t, "", app)
	}
}

func TestGetClientIP(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// test data
	tests := []struct {
		req           *http.Request
		expectedError error
	}{
		{&http.Request{}, nil},
	}

	for _, tt := range tests {
		ip := community.GetClientIP(tt.req)

		assert.EqualValues(t, tt.expectedError, nil)
		assert.EqualValues(t, "", ip)
	}
}

func TestTranslateMarkdownText(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// test data
	tests := []struct {
		src           string
		from          string
		to            language.Tag
		expectedError error
	}{
		{"##Hello, World!", "en", language.English, nil},
	}

	for _, tt := range tests {
		html, err := community.TranslateMarkdownText(context.Background(), tt.src, tt.from, tt.to)

		assert.EqualValues(t, tt.expectedError, err)
		assert.NotEqualValues(t, "", html)
	}
}

func TestTranslateArticle(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// test data
	tests := []struct {
		article           *Article
		translatedArticle *Article
		expectedError     error
	}{
		{
			&Article{
				ArticleEntry: ArticleEntry{
					Title: "Test Article",
					UId:   "1",
					Cover: "cover1",
					Tags:  "tag1",
					Ctime: time.Now(),
					Mtime: time.Now(),
				},
				Content: "This is a test article.",
			},
			&Article{
				ArticleEntry: ArticleEntry{
					Title: "Test Article",
					UId:   "1",
					Cover: "cover1",
					Tags:  "tag1",
					Ctime: time.Now(),
					Mtime: time.Now(),
				},
				Content: "This is a test article.",
			},
			nil,
		},
	}

	for _, tt := range tests {
		translatedArticle, err := community.TranslateArticle(context.Background(), tt.article)
		assert.EqualValues(t, tt.translatedArticle.Content, translatedArticle.Content)
		assert.EqualValues(t, tt.expectedError, err)
	}
}

func TestGetTranslateArticle(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	article := &Article{
		ArticleEntry: ArticleEntry{
			Title: "Test Article",
			UId:   "1",
			Cover: "cover1",
			Tags:  "tag1",
			Ctime: time.Now(),
			Mtime: time.Now(),
		},
		Content: "This is a test article.",
	}
	_, err := community.PutArticle(todo, "1", article)
	assert.Nil(t, err)

	// test data
	tests := []struct {
		uid           string
		id            string
		expectedError error
	}{
		{"1", "1", nil},
	}

	for _, tt := range tests {
		article, err := community.GetTranslateArticle(todo, tt.id)

		assert.EqualValues(t, tt.expectedError, err)
		assert.NotNil(t, article)
	}
}

func TestArticleLikeState(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	article := &Article{
		ArticleEntry: ArticleEntry{
			Title: "Test Article",
			UId:   "1",
			Cover: "cover1",
			Tags:  "tag1",
			Ctime: time.Now(),
			Mtime: time.Now(),
		},
		Content: "This is a test article.",
	}
	_, err := community.PutArticle(todo, "1", article)
	assert.Nil(t, err)

	// test data
	tests := []struct {
		uid           string
		articleID     string
		expectedError error
	}{
		{"1", "1", nil},
	}

	for _, tt := range tests {
		likeState, err := community.ArticleLikeState(todo, tt.uid, tt.articleID)
		assert.Nil(t, err)
		assert.Equal(t, false, likeState)
	}
}
