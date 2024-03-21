import (
	c "context"

	language "golang.org/x/text/language"
)

mdData := param("content")
transData, err := community.TranslateMarkdownText(c.TODO(), mdData, "auto", language.English)
if err != nil {
	json {
		"code": 500,
		"err":  err.Error(),
	}
}
json {
	"code": 200,
	"data": transData, // translation markdown content
}
