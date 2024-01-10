package core

type ArticleService struct {
}
type Config struct {
}
type Community struct {
}

//Article Article struct
type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

//line internal/core/article.gop:28:1
func (s *ArticleService) GetArticle(id string) (Article, error) {
//line internal/core/article.gop:31:1
	return Article{Id: id, Title: "Goplus Community Article Content: " + string(id), Content: "Goplus Community Article Content: " + string(id)}, nil
}

//line internal/core/article.gop:24:1
func NewArticleService() *ArticleService {
//line internal/core/article.gop:25:1
	return &ArticleService{}
}

//line internal/core/community.gop:25:1
func New(conf *Config) *Community {
//line internal/core/community.gop:26:1
	return &Community{}
}
