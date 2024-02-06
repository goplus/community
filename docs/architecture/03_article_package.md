# Article module

## Module purpose

This module is used to add、delete、query and update the article.

## Module scope

This module get article information from the database, and return the information to the caller. In addition to this, the article can also be modified. This module is the core of the goplus community, and it not only connects the user module, but also integrates the article translation and resource upload module.

## Module structure

| Field         | Type         | Null | Key | Default | Extra                       |
|----|----|----|----|----|----|
| id            | int unsigned | NO   | PRI | NULL    | auto_increment              |
| title         | varchar(255) | YES  |     |         |                             |
| user_id       | varchar(255) | NO   |     |         |                             |
| tags          | varchar(255) | YES  |     |         |                             |
| cover         | varchar(255) | YES  |     |         |                             |
| content       | longtext     | YES  |     | NULL    |                             |
| trans_content | longtext     | YES  |     | NULL    |                             |
| html_id       | int          | YES  |     | 0    |                             |
| trans_html_id | int          | YES  |     | 0    |                             |
| ctime         | datetime     | YES  |     | NULL    |                             |
| mtime         | datetime     | YES  |     | NULL    | on update CURRENT_TIMESTAMP |
| abstract      | varchar(255) | YES  |     |         |                             |
## Module Interface

None.

## Functions

### Overview

> Before invoke the article functions, we must connect mysql database to get article information. Since the article operation involves the storage of HTML files, it also needs to be connected to Qiniu cloud storage.

Example:

```go
conf := &core.Config{}
community, _ = core.New(todo, conf)  // Instantiate the community
```

**community-defined Types and Functions:**

 [type ArticleEntry](#type-articleentry)

[type Article](#type-article)

[func Article(ctx context.Context, id string) (article *Article, err error)](#func-article)

[func PutArticle(ctx context.Context, uid string, trans string, article *Article) (id string, err error)](#func-putarticle)

[func TransHtmlUrl(ctx context.Context, id string) (htmlUrl string, err error)](https://github.com/goplus/community/blob/dev/internal/core/community.go#L163)

[func SaveHtml(ctx context.Context, uid, htmlStr, mdData, id string) (articleId string, err error)](https://github.com/goplus/community/blob/dev/internal/core/community.go#L193)

[func DeleteArticle(ctx context.Context, uid, id string) (err error)](https://github.com/goplus/community/blob/dev/internal/core/community.go#L292)

[func DeleteArticles(ctx context.Context, uid string) (err error)](https://github.com/goplus/community/blob/dev/internal/core/community.go#L327)

[func ListArticle(ctx context.Context, from string, limit int, searchValue string) (items []*ArticleEntry, next string, err error)](https://github.com/goplus/community/blob/dev/internal/core/community.go#L390)

[func GetArticlesByUid(ctx context.Context, uid string) (items []*ArticleEntry, err error)](https://github.com/goplus/community/blob/dev/internal/core/community.go#L438)



### Types

#### type [ArticleEntry](https://github.com/goplus/community/blob/dev/internal/core/community.go#L49)

> Article Brief information.

```go
type ArticleEntry struct {
	ID       string
	Title    string
	UId      string
	Cover    string
	Tags     string
	User     User
	Abstract string
	Ctime    time.Time
	Mtime    time.Time
}
```

#### type [Article](https://github.com/goplus/community/blob/dev/internal/core/community.go#L61)

> All information of article.

```go
type Article struct {
	ArticleEntry
	Content string // in markdown
	HtmlUrl  string // parsed html file url
	HtmlData string
}
```

### Functions

#### func [Article](https://github.com/goplus/community/blob/dev/internal/core/community.go#L135)

```go
func Article(ctx context.Context, id string) (article *Article, err error)
```

> All users can view articles in the community, so they don't need user-related permissions, and they can view the corresponding article information only by the article ID. 
>
> However, it also needs to support the modification of the article, and the user needs to have the operation permission to make the modification.
>
> We use Article() function to get article by article ID. However, for the above two scenarios, need to use the CanEditable() function to determine the permission before getting the article information.

**Example：**

```go
// if edit the article, we must check the permission
if editable, _ := community.canEditable(todo, uid, id); !editable {
    zlog.Error("no permissions")
    http.Redirect(ctx.ResponseWriter, ctx.Request, "/error", http.StatusTemporaryRedirect)
}
article, _ := community.article(todo, id)
```

#### func [PutArticle](https://github.com/goplus/community/blob/dev/internal/core/community.go#L226)

```go
func PutArticle(ctx context.Context, uid string, trans string, article *Article) (id string, err error)
```

> Add or update article. When the ID in the article is not empty, it means the add operation, otherwise it is the update operation.
>
> Since both updates and additions operation require specific user actions, the user ID is required in the input parameters.
>
> When trans is empty, it means add or update origin article content, otherwise translate content.

**Example:**

```go

article := &Article{
    ArticleEntry: ArticleEntry{
        ID:"",  // add
        Title: "Test Article",
        UId:   "1",
        Cover: "cover1",
        Tags:  "tag1",
        Abstract: "abstract1"
        Ctime: time.Now(),
        Mtime: time.Now(),
    },
    Content: "This is a test article.",
}

uid = "testuid"

id, _ := community.PutArticle(todo, uid, "", article)
```



## Collaboration

### Web

At present, the collaboration with the front-end is mainly that the front-end sends me the written html file, and then I make adjustments and bug fixes. 

### Markdown && Media

Invoke functions

