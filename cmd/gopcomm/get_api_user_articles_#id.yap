import (
	c "context"
	"github.com/qiniu/x/xlog"
	"strconv"
)


xLog := xlog.New("")
todo := c.TODO()
id := param("id")
page := param("page")
limit := param("limit")

limitInt, err := strconv.Atoi(limit)
if err != nil {
	limitInt = mediaLimitConst
}
items, total, err := community.GetArticlesByUid(todo, id, page, limitInt)
if err != nil {
	xLog.Error("get article list error:", err)
	json {
		"code":  0,
		"err":   err.Error(),
		"total": 0,
	}
}
json {
	"code":  200,
	"items": items,
	"total": total,
}
