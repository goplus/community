import (
	c "context"
	"github.com/qiniu/x/xlog"
	"strconv"
)

xLog := xlog.New("")
todo := c.TODO()
from := param("from")
limit := param("limit")
searchValue := param("value")
label := param("label")

limitInt, err := strconv.Atoi(limit)
if err != nil {
	limitInt = limitConst
}

// Get Article Info
articles, next, err := community.ListArticle(todo, from, limitInt, searchValue, label)
if err != nil {
	xLog.Error("get article error:", err)
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
