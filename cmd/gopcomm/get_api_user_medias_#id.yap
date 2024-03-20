import (
	c "context"
	"strconv"
)

todo := c.TODO()
format := param("format")
uid := param("id")
page := param("page")
limit := param("limit")

limitInt, err := strconv.Atoi(limit)
if err != nil {
	limitInt = mediaLimitConst
}
pageInt, err := strconv.Atoi(page)
if err != nil {
	json {
		"code":  400,
		"total": 0,
		"err":   err.Error(),
	}
}
files, total, err := community.ListMediaByUserId(todo, uid, format, pageInt, limitInt)
if err != nil {
	json {
		"code":  0,
		"total": total,
		"err":   err.Error(),
	}
} else {
	json {
		"code":  200,
		"total": total,
		"items": files,
	}
}
