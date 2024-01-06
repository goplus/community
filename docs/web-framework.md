Web Framework
=====

## Choose Web Framework

TODO

#### Yap

* https://github.com/goplus/yap (from Dec 2023)
* demo ([blog.go](https://github.com/goplus/yap/blob/main/demo/blog/blog.go), [article.yap](https://github.com/goplus/yap/blob/main/demo/blog/yap/article.yap))

```go
type article struct {
	ID string
}

//go:embed yap
var yapFS embed.FS

fsYap, _ := fs.Sub(yapFS, "yap")
y := yap.New(fsYap)

y.GET("/p/:id", func(ctx *yap.Context) {
	ctx.YAP(200, "article", article{
		ID: ctx.Param("id"),
	})
})

y.Run(":8080")
```

#### Gin

* https://github.com/gin-gonic/gin (73.5k+ stars, from Jun 2014)

```go
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
```

#### Fiber

* https://github.com/gofiber/fiber (29.8k+ stars, from Dec 2019)

```go
    app := fiber.New()

    // GET /api/register
    app.Get("/api/*", func(c *fiber.Ctx) error {
        msg := fmt.Sprintf("✋ %s", c.Params("*"))
        return c.SendString(msg) // => ✋ register
    })

    // GET /flights/LAX-SFO
    app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
        msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
        return c.SendString(msg) // => 💸 From: LAX, To: SFO
    })

    // GET /dictionary.txt
    app.Get("/:file.:ext", func(c *fiber.Ctx) error {
        msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
        return c.SendString(msg) // => 📃 dictionary.txt
    })

    // GET /john/75
    app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
        msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
        return c.SendString(msg) // => 👴 john is 75 years old
    })

    // GET /john
    app.Get("/:name", func(c *fiber.Ctx) error {
        msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
        return c.SendString(msg) // => Hello john 👋!
    })

    app.Listen(":3000")
```

#### Iris

* https://github.com/kataras/iris (24.6k+ stars, from May 2016)

```go
  app := iris.New()
  app.Use(iris.Compression)

  app.Get("/", func(ctx iris.Context) {
    ctx.HTML("Hello <strong>%s</strong>!", "World")
  })

  app.Listen(":8080")
```
