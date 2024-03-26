import (
	c "context"
	"fmt"
	"regexp"
)

id := param("id")
fileKey, err := community.GetMediaUrl(c.TODO(), id)
m := make(map[string]string, 4)
format, err := community.GetMediaType(c.TODO(), id)
if err != nil {
	json {
		"code": 500,
		"err":  err.Error(),
	}
}
htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
if err != nil {
	json {
		"code": 500,
		"err":  "have no html media",
	}
}
m["fileKey"] = htmlUrl
m["type"] = format
m["subtitle"] = domain + id
m["status"] = "0"
match, _ := regexp.MatchString("^video", format)
if match {
	subtitle, status, err := community.GetVideoSubtitle(c.TODO(), id)
	if err != nil {
		json {
			"code": 200,
			"url":  m,
		}
	}
	if status == "1" {
		m["subtitle"] = domain + subtitle
	}
	m["status"] = status
}
json {
	"code": 200,
	"url":  m,
}
