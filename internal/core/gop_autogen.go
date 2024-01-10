package core
// Article Config
//
// ArticleConfig is the config for article.
type ArticleConfig struct {
}
// articleHandler
//
// articleHandler is the interface for article handler,
// which defines the methods that must be implemented by the article handler.
// We need to implement this interface in the community/internal/core package.
type articleHandler interface {
	GetArticle(id string) (Article, error)
}
// ArticleHandler
//
// ArticleHandler is the handler for article.
type articleHandlerImpl struct {
	config *ArticleConfig
}
// Config
//
// Top level config for community handlers.
type Config struct {
	ArticleConfig
}
// communityHandler
//
// communityHandler is the interface for community handler,
// which defines the methods that must be implemented by the community handler.
// It returns the handler for each service.
type communityHandler interface {
	GetArticleHandler() articleHandler
}
// CommunityHandler
//
// CommunityHandler is the top handler for community.
// It contains all the handlers for each service.
type communityHandlerImpl struct {
	articleHandler
}
//Article Article struct
type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
// GetArticle gets article by id.
//
//line internal/core/article.gop:52:1
func (h *articleHandlerImpl) GetArticle(id string) (Article, error) {
//line internal/core/article.gop:55:1
	return Article{Id: id, Title: "Goplus Community Article Content: " + string(id), Content: "Goplus Community Article Content: " + string(id)}, nil
}
// NewArticleHandler creates a new article handler.
//
//line internal/core/article.gop:45:1
func newArticleHandler(conf *ArticleConfig) *articleHandlerImpl {
//line internal/core/article.gop:46:1
	return &articleHandlerImpl{config: conf}
}
// GetArticleHandler gets article handler.
//
//line internal/core/community.gop:54:1
func (p *communityHandlerImpl) GetArticleHandler() articleHandler {
//line internal/core/community.gop:55:1
	return p.articleHandler
}
// New creates a new community handler.
//
//line internal/core/community.gop:47:1
func New(conf *Config) *communityHandlerImpl {
//line internal/core/community.gop:48:1
	return &communityHandlerImpl{articleHandler: newArticleHandler(&conf.ArticleConfig)}
}
// Check interface implementation.
var _ articleHandler = (*articleHandlerImpl)(nil)
// Check interface implementation.
var _ communityHandler = (*communityHandlerImpl)(nil)
