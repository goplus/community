import (
	c "context"

	"github.com/qiniu/x/log"
)

id := param("id")
article, err := community.GetTranslateArticle(c.TODO(), id)

if err != nil {
	log.Info(err)
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
