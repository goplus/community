import (
	c "context"
	"fmt"
)


todo := c.TODO()
id := param("id")
fileKey, err := community.GetMediaUrl(todo, id)
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
