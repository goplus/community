import (
	c "context"
	"strconv"

	"github.com/qiniu/x/log"
)


from := param("from")
limit := param("limit")
searchValue := param("value")
label := param("label")

limitInt, err := strconv.Atoi(limit)
if err != nil {
	limitInt = limitConst
}

// Get Article Info
articles, next, err := community.ListArticle(c.TODO(), from, limitInt, searchValue, label)
if err != nil {
	log.Error("get article error:", err)
	json {
		"code": 0,
		"err":  "get article failed",
	}
}
json {
	"code":  200,
	"items": articles,
	"next":  next,
	"value": searchValue,
}
