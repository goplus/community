import (
	c "context"
)

todo := c.TODO()
uid := param("uid")
vid := param("vid")

if uid == "" || vid == "" {
	json {
		"code": 200,
		"msg":  "Invalid param",
	}
}

if err := community.RetryCaptionGenerate(todo, uid, vid); err != nil {
	json {
		"code": 200,
		"msg":  "Request task error",
	}
}

json {
	"code": 200,
	"msg":  "Ok",
}
