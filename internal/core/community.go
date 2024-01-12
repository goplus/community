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
	"io"
	"os"
	"time"

	"gocloud.dev/blob"
)

var (
	ErrNotExist = os.ErrNotExist
)

type Config struct {
	Driver string // database driver. default is `mysql`.
	DSN    string // database data source name
	BlobUS string // blob URL scheme
}

type ArticleEntry struct {
	ID    string
	Title string
	Ctime time.Time
	Mtime time.Time
}

type Article struct {
	ArticleEntry
	Content string // in markdown
}

type Community struct {
	bucket *blob.Bucket
	db     *sql.DB
}

func New(ctx context.Context, conf *Config) (ret *Community, err error) {
	if conf == nil {
		conf = new(Config)
	}
	driver := conf.Driver
	dsn := conf.DSN
	bus := conf.BlobUS
	if driver == "" {
		driver = "mysql"
	}
	if dsn == "" {
		dsn = os.Getenv("GOP_COMMUNITY_DSN")
	}
	if bus == "" {
		bus = os.Getenv("GOP_COMMUNITY_BLOBUS")
	}
	bucket, err := blob.OpenBucket(ctx, bus)
	if err != nil {
		return
	}
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return
	}
	return &Community{bucket, db}, nil
}

const contentSummary = `
Content
====

Text body
`

// Article returns an article.
func (p *Community) Article(ctx context.Context, id string) (article *Article, err error) {
	if id == "123" {
		article = &Article{
			ArticleEntry{
				ID:    id,
				Title: "Title",
			},
			contentSummary,
		}
		return
	}
	return nil, ErrNotExist
}

// CanEditable
func (p *Community) CanEditable(ctx context.Context, uid, id string) (editable bool, err error) {
	return true, nil
}

// PutArticle adds new article (ID == "") or edits an existing article (ID != "").
func (p *Community) PutArticle(ctx context.Context, uid string, article *Article) (id string, err error) {
	return
}

func (p *Community) DeleteArticle(ctx context.Context, uid, id string) (err error) {
	return
}

const (
	MarkBegin = ""
	MarkEnd   = "eof"
)

// ListArticle lists articles from an position.
func (p *Community) ListArticle(ctx context.Context, from string, limit int) (items []*ArticleEntry, next string, err error) {
	if from == MarkBegin {
		item := &ArticleEntry{
			ID:    "123",
			Title: "Title",
		}
		return []*ArticleEntry{item}, MarkEnd, nil
	}
	return nil, MarkEnd, io.EOF
}
