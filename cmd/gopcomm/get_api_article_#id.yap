import (
	c "context"

	"github.com/qiniu/x/log"
)

id := param("id")
article, err := community.Article(c.TODO(), id)
if err != nil {
	log.Error("get article error:", err)
	json {
		"code": 0,
		"err":  "get article failed",
	}
}
json {
	"code": 200,
	"data": article,
}
