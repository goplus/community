# Account module

## Module purpose

This module is responsible for user account management, including login, logout, and user information query.

## Module scope

This module get user information from the database, and return the information to the caller. It gets the OAuth login authentication information from github twitter facebook and stores the user information temporarily in a cookie, which is used to verify the user login status and parse the user information.

- Input: Twitter/Facebook/Github
- Output: User Token
- Dependency: Casdoor OAuth endpoint

## Module structure

None.

## Module Interface

None.

## Functions

### Login

- Function: Login with OAuth
- Input: OAuth provider
- Return: User Token
- Error: None

Example:

```go
// Login
redirectURL := fmt.Sprintf("%s/%s", ctx.Request.Referer() + "callback")
loginURL := community.RedirectToCasdoor(redirectURL)
ctx.Redirect loginURL, http.StatusFound
```

### Logout

- Function: Logout
- Input: None
- Return: None
- Error: None

Example:

```go
tokenCookie, err := ctx.Request.Cookie("token")
if err != nil {
    zlog.Error("get token error:", err)
}

// Delete token
tokenCookie.MaxAge = -1
http.SetCookie(ctx.ResponseWriter, tokenCookie)

// Redirect to home page
http.Redirect(ctx.ResponseWriter, ctx.Request, fmt.Sprintf("http://localhost:8080"), http.StatusFound)
```

### GetUserInfo

- Function: Get user information
- Input: User Id
- Return: User information
- Error: None

Example:

```go
claim, err := casdoorsdk.GetUserByUserId(uid)
if err != nil {
    p.zlog.Error(err)
    return &User{}, ErrNotExist
}
user = &User{
    Name:   claim.Name,
    Avatar: claim.Avatar,
    Id:     claim.Id,
}
```