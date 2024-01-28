# Article module

## Module purpose

This module is used to add、delete、query and update the article.

## Module scope

This module get article information from the database, and return the information to the caller. In addition to this, the article can also be modified. This module is the core of the gop community, and it not only connects the user module, but also integrates the article translation and resource upload module.

## Module structure

None.

## Module Interface

None.

## Struct

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

type Article struct {
	ArticleEntry
	Content string // in markdown
	// Status  int    // published or draft
	HtmlUrl  string // parsed html file url
	HtmlData string
}
```

## Functions

### Article

> Query article information by the article ID

- Function: Article
- Input: id string
- Return: article *Article
- Error: None


示例：

```go
// test data
tests := []struct {
    id            string
    expectedID    string
    expectedError error
}{
    {"1", "1", nil},
    {"10", "", ErrNotExist},
}

for _, tt := range tests {
    article, err := community.Article(todo, tt.id)

    if article.ID != tt.expectedID {
        t.Errorf("Article(%s) returned id is %s, expected: %s", tt.id, article.ID, tt.expectedID)
    }
    if err != tt.expectedError {
        t.Errorf("Article(%s) returned err is %s, expected: %s", tt.id, err, tt.expectedError)
    }
}
```

### CanEditable

> Check whether the user has the operation permission

- Function: CanEditable
- Input: uid, id string
- Return: editable bool
- Error: None

示例：

```go
// test data
tests := []struct {
    uid           string
    articleID     string
    expectedEdit  bool
    expectedError error
}{
    {"1", "1", true, nil},
    {"2", "1", false, ErrPermission},
}

for _, tt := range tests {
    canEdit, _ := community.CanEditable(todo, tt.uid, tt.articleID)

    if canEdit != tt.expectedEdit {
        t.Errorf("CanEditable(%s, %s) returned %t, expected: %t", tt.uid, tt.articleID, canEdit, tt.expectedEdit)
    }
}
```

### SaveHtml

> When you click the translation button on the front-end, you need to store the original MD and HTML files in the article table of the database, and return the ID of the article so that the translated article information can be inserted

- Function: SaveHtml
- Input: uid, htmlStr, mdData, id string (if id == "", add; else update)
- Return: articleId string
- Error: None

示例：

```go
tests := []struct {
    uid            string
    htmlStr        string
    mdData         string
    id            string
    expectedArticleID    string
    expectedError error
}{
    {"1", "<html><body><p>Hello, World!<p></body></html>","##Hello, World!", "", "15", nil},
    {"1", "<html><body><p>Hello, World!<p></body></html>","##Hello, World!", "15", "15", nil},
}

for _, tt := range tests {
    articleId, err := community.SaveHtml(todo, tt.uid, tt.htmlStr, tt.mdData, tt.id)

    if articleId != tt.expectedArticleID {
        t.Errorf("SaveHtml(%s,%s,%s,%s) returned id is %s, expected: %s", tt.uid, tt.htmlStr, tt.mdData, tt.id, articleId, tt.expectedArticleID)
    }
    if err != tt.expectedError {
        t.Errorf("SaveHtml(%s,%s,%s,%s) returned err is %s, expected: %s", tt.uid, tt.htmlStr, tt.mdData, tt.id, err, tt.expectedError)
    }
}
```

###  PutArticle

> add/update article

- Function: PutArticle
- Input: uid string, trans string, article *Article (if trans == "", origin md and html; else translate md and html)
- Return: id string
- Error: None

```go
// test data
article := &Article{
    ArticleEntry: ArticleEntry{
        Title: "Test Article",
        UId:   "1",
        Cover: "cover1",
        Tags:  "tag1",
        Ctime: time.Now(),
        Mtime: time.Now(),
    },
    Content: "This is a test article.",
}
articleUpdate := &Article{
    ArticleEntry: ArticleEntry{
        ID:    "1",
        Title: "Test Article",
        UId:   "1",
        Cover: "cover1",
        Tags:  "tag1",
        Ctime: time.Now(),
        Mtime: time.Now(),
    },
    Content: "This is a test article.",
}

tests := []struct {
    uid        string
    trans		string
    article    *Article
    expectedID string
}{
    {"1", "", article, "1"},       // insert
    {"1", "trans", article, "1"},       // insert trans
    {"1", "", articleUpdate, "1"}, // update
}

for _, tt := range tests {
    id, _ := community.PutArticle(todo, tt.uid, tt.trans, tt.article)

    if id != tt.expectedID {
        t.Errorf("PutArticle(%s, %s, %+v) returned ID %s, expected: %s", tt.uid,tt.trans, tt.article, id, tt.expectedID)
    }
}
```

###  DeleteArticle

> Delete article by user id and article id

- Function: DeleteArticle
- Input: uid, id string
- Return: None
- Error: error

```go
tests := []struct {
    uid         string
    articleID   string
    expectedErr error
}{
    {"22", "2", ErrPermission}, // no permission
    {"1", "1", nil},
}

for _, tt := range tests {
    err := community.DeleteArticle(todo, tt.uid, tt.articleID)

    if err != tt.expectedErr {
        t.Errorf("DeleteArticle(%s, %s) returned error: %v, expected: %v", tt.uid, tt.articleID, err, tt.expectedErr)
    }
}
```

### DeleteArticles

> It is necessary to delete user-written articles at the same time as deleting the user

- Function: DeleteArticles
- Input: uid string
- Return: None
- Error: error


### Articles

> Paginate the list of articles (the same interface is used for search)
> To be modified (change to "Bottom-out loading")

- Function: Articles
- Input: page int, limit int, searchValue string (if searchValue == "", home article list; else search)
- Return: items []*ArticleEntry, total int ("total" indicates the number of articles that meet the requirements)
- Error: None

```go
tests := []struct {
    page         int
    limit   int
    searchValue   string
    expectedTotal int
}{
    {1, 10, "", 10}, // home
    {1, 10, "test", 10},  // search
}

for _, tt := range tests {
    _,total,_ := community.Articles(todo, tt.page, tt.limit, tt.searchValue)

    if total != tt.expectedTotal {
        t.Errorf("Articles(%d, %d, %s) returned total: %v, expected: %v", ttt.page, tt.limit, tt.searchValue, total, tt.expectedTotal)
    }
}
```

### GetArticlesByUid

> Get user-written article list by user id
> To be modified (change to "Bottom-out loading" or "Paginate")

- Function: GetArticlesByUid
- Input: uid string
- Return: items []*ArticleEntry
- Error: None

```go
// test data
tests := []struct {
    uid         string
    expectedError error
}{
    {"1", nil},
}

for _, tt := range tests {
    _, err := community.GetArticlesByUid(todo, tt.uid)

    if err != tt.expectedError {
        t.Errorf("GetArticlesByUid(%s) returned err: %v, expected: %v", ttt.uid, err, tt.expectedTotal)
    }
}
```

## Collaboration

### Web

At present, the collaboration with the front-end is mainly that the front-end sends me the written html file, and then I make adjustments and bug fixes. (In general, this method is slow)

### Markdown && Media

Invoke functions