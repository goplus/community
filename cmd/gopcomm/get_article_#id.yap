import (
	c "context"

	"github.com/qiniu/x/log"
	"github.com/goplus/community/internal/core"
)

uid := ""
var user *core.User
token, err := Request.Cookie("token")
if token!=nil{
    user, err = community.GetUser(token.Value)
    if err != nil {
    	log.Error("get user error")
    	json {
    		"code": 0,
    		"err":  err.Error(),
    	}
    }
}
uid = user.Id
id := param("id")
platform := Param("platform")
ip := community.GetClientIP(Request)
community.ArticleLView(c.TODO(), id, ip, uid, platform)
article, err := community.Article(c.TODO(), id)
if err != nil {
	log.Error("get article error:", err)
	return
}
likeState, err := community.ArticleLikeState(c.TODO(), uid, id)
if err != nil {
	log.Error("article state err:", err)
	return
}

yap "article", {
	"User":      user,
	"Article":   article,
	"LikeState": likeState,
}
