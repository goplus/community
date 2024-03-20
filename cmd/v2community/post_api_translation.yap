import (
	c "context"
	language "golang.org/x/text/language"
)



todo := c.TODO()
mdData := param("content")
transData, err := community.TranslateMarkdownText(todo, mdData, "auto", language.English)
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
