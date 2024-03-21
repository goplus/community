import (
	c "context"
	"strconv"

	"github.com/goplus/community/internal/core"
	"github.com/qiniu/x/log"
)


id := param("id")
content := param("content")
title := param("title")
tags := param("tags")
abstract := param("abstract")
label := param("label")
trans, err := strconv.ParseBool(param("trans"))
if err != nil {
	log.Error("parse bool error:", err)
}

// get user id
token, err := Request.Cookie("token")
if err != nil {
	log.Info("token", err)
	json {
		"code": 0,
		"err":  "no token",
	}
}
uid, err := community.ParseJwtToken(token.Value)
if err != nil {
	log.Info("uid", err)
	json {
		"code": 0,
		"err":  err.Error(),
	}
}
article := &core.Article{
	ArticleEntry: core.ArticleEntry{
		ID:       id,
		Title:    title,
		UId:      uid,
		Cover:    param("cover"),
		Tags:     tags,
		Abstract: abstract,
		Label:    label,
	},
	Content: content,
	Trans:   trans,
	VttId:   param("vtt_id"),
}

if trans {
	article, _ = community.TranslateArticle(c.TODO(), article)
}

id, err = community.PutArticle(c.TODO(), uid, article)
if err != nil {
	log.Info(err)
	json {
		"code": 0,
		"err":  "add failed",
	}
}
json {
	"code": 200,
	"data": id,
}
