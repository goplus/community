import (
	c "context"
	"strconv"

	"github.com/qiniu/x/log"
)

id := param("id")
page := param("page")
limit := param("limit")

limitInt, err := strconv.Atoi(limit)
if err != nil {
	limitInt = mediaLimitConst
}
items, total, err := community.GetArticlesByUid(c.TODO(), id, page, limitInt)
if err != nil {
	log.Error("get article list error:", err)
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
