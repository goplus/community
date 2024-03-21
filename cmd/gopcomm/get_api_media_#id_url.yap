import (
	c "context"
	"fmt"
)

id := param("id")
fileKey, err := community.GetMediaUrl(c.TODO(), id)
htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
if err != nil {
	json {
		"code": 500,
		"err":  "have no html media",
	}
}
json {
	"code": 200,
	"url":  htmlUrl,
}
