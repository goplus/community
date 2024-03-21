import (
	c "context"

	"github.com/qiniu/x/log"
)

id := param("id")
token, err := Request.Cookie("token")
uid, err := community.ParseJwtToken(token.Value)
if err != nil {
	log.Error("token parse error")
	json {
		"code": 0,
		"err":  err.Error(),
	}
}
err = community.DelMedia(c.TODO(), uid, id)
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
