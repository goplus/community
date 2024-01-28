# Media module

## Module purpose

The media module is mainly responsible for unified management of all media resources in the community.

## Module scope

This module is used to uniformly manage resources in cloud storage in projects.

- Input: resources
- Output: ID after adding the resource record
- Dependencies: Cloud Storage

## Module structure

None

## Module Interface

None

## Functions

### Upload resources

- Function: Upload resources
- Input parameters: resources
- Return: file table id
- Error: None

Example:

```go
file, header, err := ctx.FormFile("file")
filename := header.Filename

ctx.ParseMultipartForm(10 << 20)

if err != nil {
zlog.Error("upload file error:", filename)
ctx.JSON(500, err.Error())
return
}


dst, err := os.Create(filename)
if err != nil {
zlog.Error("create file error:", file)
ctx.JSON(500, err.Error())
return
}
defer func() {
file.Close()
dst.Close()
err = os.Remove(filename)
if err != nil {
zlog.Error("delete file error:", filename)
return
}
}()


_, err = io.Copy(dst, file)
if err != nil {
zlog.Error("copy file errer:", filename)
ctx.JSON(500, err.Error())
return
}
bytes, err := os.ReadFile(filename)
if err != nil {
zlog.Error("read file errer:", filename)
ctx.JSON(500, err.Error())
return
}
token, err := ctx.Request.Cookie("token")
if err != nil {
ctx.json {
"code": 500,
"err": "no token",
}
}
uid, err := community.ParseJwtToken(token.Value)
if err != nil {
ctx.json {
"code": 500,
"err": err.Error(),
}
}
id,err:=community.SaveMedia(context.Background(), uid, bytes)
if err!=nil {
zlog.Error("save file",err.Error())
ctx.JSON(500, err.Error())
return
}
// todo append current project ip and getMedia
// sample: 127.0.0.1:8080/getMedia/ + id
ctx.JSON(200,id)
```

### Access to resources

- Function: Get the URL of the resource
- Input parameter: resource id
- Return: resource table fileKey
- Error: The resource may not exist, returning 404

Example:

```go
id := ctx.param("id")
fileKey, err := community.GetMediaUrl(todo, id)
htmlUrl := fmt.Sprintf("%s%s", domain, fileKey)
if err != nil {
ctx.json {
"code": 500,
"err": "have no html media",
}
}
ctx.json {
"code": 200,
"url": htmlUrl,
}
```