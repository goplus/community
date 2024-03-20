import (
	c "context"
	"github.com/qiniu/x/xlog"
)


xLog := xlog.New("")
todo := c.TODO()
id := param("id")
article, err := community.GetTranslateArticle(todo, id)

if err != nil {
	xLog.Info(err)
	json {
		"code": 0,
	}
}
json {
	"code":    200,
	"content": article.Content,
	"tags":    article.Tags,
	"title":   article.Title,
}
