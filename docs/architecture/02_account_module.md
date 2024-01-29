# Account module

## Module purpose

This module is responsible for user account management, including login, logout, and user information query.

## Module scope

This module get user information from the database, and return the information to the caller. It gets the OAuth login authentication information from github twitter facebook and stores the user information temporarily in a cookie, which is used to verify the user login status and parse the user information.

- Input: Twitter/Facebook/Github
- Output: User Token
- Dependency: Casdoor OAuth endpoint

## Module structure

```go
    type User struct {
        Id                string   `xorm:"varchar(100) index" json:"id"`
        Type              string   `xorm:"varchar(100)" json:"type"`
        Name              string `xorm:"varchar(100) notnull pk" json:"name"`
        Password          string   `xorm:"varchar(100)" json:"password"`
        FirstName         string   `xorm:"varchar(100)" json:"firstName"`
        LastName          string   `xorm:"varchar(100)" json:"lastName"`
        Avatar            string   `xorm:"varchar(500)" json:"avatar"`
        Email             string   `xorm:"varchar(100) index" json:"email"`
        Phone             string   `xorm:"varchar(20) index" json:"phone"`
        GitHub            string `xorm:"github varchar(100)" json:"github"`
        WeChat            string `xorm:"wechat varchar(100)" json:"wechat"`
        Facebook          string `xorm:"facebook varchar(100)" json:"facebook"`
        Twitter           string `xorm:"twitter varchar(100)" json:"twitter"`
        CreatedTime       string `xorm:"varchar(100) index" json:"createdTime"`
        UpdatedTime       string `xorm:"varchar(100)" json:"updatedTime"`
    }
```

## Module Interface

None.

## Functions

### Login
```go
// Get login url and user login 
...
// Get token and save token
...
// Redirect to home page
...
```

### Logout
```go
//Get token
...
func DeleteToken(token *Token) (bool, error)
// Redirect to home page
...
```

### GetUserInfo

```go
//Get token 
...
//Parse token
func ParseJwtToken(token string) (*Claims, error)
//Get UserId or ...
...

func (c *Client) GetUserByUserId(userId string) (*User, error)

func (c *Client) GetUserByPhone(phone string) (*User, error)

func (c *Client) GetUserByEmail(email string) (*User, error)

```