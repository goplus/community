import (
	c "context"
	"github.com/qiniu/x/xlog"
)


xLog := xlog.New("")
todo := c.TODO()
id := param("id")
token, err := Request.Cookie("token")
uid, err := community.ParseJwtToken(token.Value)
if err != nil {
	xLog.Error("token parse error")
	json {
		"code": 0,
		"err":  err.Error(),
	}
}
err = community.DeleteArticle(todo, uid, id)
if err != nil {
	json {
		"code": 0,
		"err":  "delete failed",
	}
} else {
	json {
		"code": 200,
		"msg":  "delete success",
	}
}
