import (
	c "context"
	"strconv"

	"github.com/qiniu/x/log"
)

articleId := param("articleId")
articleIdInt, err := strconv.Atoi(articleId)
if err != nil {
	json {
		"code": 500,
		"err":  err.Error(),
	}
}
token, err := Request.Cookie("token")
if err != nil {
	log.Info("token", err)
	json {
		"code": 0,
		"err":  "no token",
	}
}
uid, err := community.ParseJwtToken(token.Value)
if err != nil {
	json {
		"code": 0,
		"err":  err.Error(),
	}
}
res, err := community.ArticleLike(c.TODO(), articleIdInt, uid)
if err != nil {
	json {
		"code": 500,
		"err":  err.Error(),
	}
}
json {
	"code": 200,
	"data": res,
}
