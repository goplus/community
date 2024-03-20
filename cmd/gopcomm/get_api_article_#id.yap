import (
	c "context"
	"github.com/qiniu/x/xlog"
)

xLog := xlog.New("")
todo := c.TODO()
id := param("id")
article, err := community.Article(todo, id)
if err != nil {
	xLog.Error("get article error:", err)
	json {
		"code": 0,
		"err":  "get article failed",
	}
}
json {
	"code": 200,
	"data": article,
}
