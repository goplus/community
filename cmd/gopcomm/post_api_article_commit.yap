import (
	c "context"
	"github.com/goplus/community/internal/core"
	"github.com/goplus/community/translation"
	"github.com/qiniu/x/xlog"
	"strconv"
)

var (
	trans *translation.Engine
)

xLog := xlog.New("")
todo := c.TODO()
id := param("id")
content := param("content")
title := param("title")
tags := param("tags")
abstract := param("abstract")
label := param("label")
trans, err := strconv.ParseBool(param("trans"))
if err != nil {
	xLog.Error("parse bool error:", err)
}

// get user id
token, err := Request.Cookie("token")
if err != nil {
	xLog.Info("token", err)
	json {
		"code": 0,
		"err":  "no token",
	}
}
uid, err := community.ParseJwtToken(token.Value)
if err != nil {
	xLog.Info("uid", err)
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

// if trans == true translate the article
if trans {
	article, _ = community.TranslateArticle(todo, article)
}

id, err = community.PutArticle(todo, uid, article)
if err != nil {
	xLog.Info(err)
	json {
		"code": 0,
		"err":  "add failed",
	}
}
json {
	"code": 200,
	"data": id,
}
