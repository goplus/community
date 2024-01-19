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
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gocloud.dev/blob"
)

var (
	ErrNotExist   = os.ErrNotExist
	ErrPermission = os.ErrPermission
)

type Config struct {
	Driver string // database driver. default is `mysql`.
	DSN    string // database data source name
	BlobUS string // blob URL scheme
}

type ArticleEntry struct {
	ID    string
	Title string
	UId   string
	Cover string
	Tags  string
	User  User
	Ctime time.Time
	Mtime time.Time
}

type Article struct {
	ArticleEntry
	Content string // in markdown
	// Status  int    // published or draft
	// HtmlUrl string // parsed html file url
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
		log.Println(err)
		return
	}
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Println(err)
		return
	}
	return &Community{bucket, db}, nil
}

// Article returns an article.
func (p *Community) Article(ctx context.Context, id string) (article *Article, err error) {
	article = &Article{}
	sqlStr := "select id,title,user_id,cover,tags,content,ctime,mtime from article where id=?"
	err = p.db.QueryRow(sqlStr, id).Scan(&article.ID, &article.Title, &article.UId, &article.Cover, &article.Tags, &article.Content, &article.Ctime, &article.Mtime)
	if err == sql.ErrNoRows {
		log.Println("not found the article")
		return article, ErrNotExist
	} else if err != nil {
		return article, err
	}
	// TODO add author info
	user, err := p.GetUser(ctx, article.UId)
	if err != nil {
		return
	}
	article.User = *user
	return
}

// CanEditable determine whether the user has the permission to operate.
func (p *Community) CanEditable(ctx context.Context, uid, id string) (editable bool, err error) {
	sqlStr := "select id from article where id=? and user_id = ?"
	err = p.db.QueryRow(sqlStr, id, uid).Scan(&id)
	if err != nil {
		return false, ErrPermission
	}
	return true, nil
}

// PutArticle adds new article (ID == "") or edits an existing article (ID != "").
func (p *Community) PutArticle(ctx context.Context, uid string, article *Article) (id string, err error) {
	// new article
	if article.ID == "" {
		sqlStr := "insert into article (title, ctime, mtime, user_id, tags, cover, content) values (?, ?, ?, ?, ?, ?, ?)"
		res, err := p.db.Exec(sqlStr, &article.Title, time.Now(), time.Now(), uid, &article.Tags, &article.Cover, &article.Content)
		if err != nil {
			return "", err
		}
		idInt, err := res.LastInsertId()
		return strconv.FormatInt(idInt, 10), nil
	}
	// edit article
	sqlStr := "update article set title=?, mtime=?, tags=?, cover=?, content=? where id=?"
	_, err = p.db.Exec(sqlStr, &article.Title, time.Now(), &article.Tags, &article.Cover, &article.Content, &article.ID)
	return article.ID, err
}

// DeleteArticle delete the article.
func (p *Community) DeleteArticle(ctx context.Context, uid, id string) (err error) {
	sqlStr := "delete from article where id=? and user_id=?"
	_, err = p.db.Exec(sqlStr, id, uid)
	// TODO delete the media
	return
}

// DeleteArticles delete the articles by uid.
func (p *Community) DeleteArticles(ctx context.Context, uid string) (err error) {
	sqlStr := "delete from article where user_id=?"
	_, err = p.db.Exec(sqlStr, uid)
	// TODO delete the media
	return
}

const (
	MarkBegin = ""
	MarkEnd   = "eof"
)

// ListArticle lists articles from an position.
func (p *Community) ListArticle(ctx context.Context, from string, limit int) (items []*ArticleEntry, next string, err error) {
	if from == MarkBegin {
		from = "0"
	} else if from == MarkEnd {
		return []*ArticleEntry{}, from, nil
	}
	fromInt, err := strconv.Atoi(from)
	if err != nil {
		return []*ArticleEntry{}, from, err
	}
	sqlStr := "select id, title, ctime, user_id, tags, cover from article limit ? offset ?"
	rows, err := p.db.Query(sqlStr, limit, fromInt)
	if err != nil {
		return []*ArticleEntry{}, from, err
	}
	defer rows.Close()

	var rowLen int
	for rows.Next() {
		article := &ArticleEntry{}
		err := rows.Scan(&article.ID, &article.Title, &article.Ctime, &article.UId, &article.Tags, &article.Cover)
		if err != nil {
			return []*ArticleEntry{}, from, err
		}
		// TODO add author info
		user, err := p.GetUser(ctx, article.UId)
		if err != nil {
			return []*ArticleEntry{}, from, err
		}
		article.User = *user

		items = append(items, article)
		rowLen++
	}
	// have no article
	if rowLen == 0 {
		return []*ArticleEntry{}, MarkEnd, io.EOF
	}
	next = strconv.Itoa(fromInt + rowLen)
	return items, next, nil
}
