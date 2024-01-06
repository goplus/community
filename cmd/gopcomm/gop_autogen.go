package main

import (
	"embed"
	"io/fs"
	"github.com/goplus/yap"
	"github.com/qiniu/x/errors"
)

type article struct {
	ID string
}
//go:embed yap
var yapFS embed.FS
//line cmd/gopcomm/community.gop:15
func main() {
//line cmd/gopcomm/community.gop:15:1
	fsYap := func() (_gop_ret fs.FS) {
//line cmd/gopcomm/community.gop:15:1
		var _gop_err error
//line cmd/gopcomm/community.gop:15:1
		_gop_ret, _gop_err = fs.Sub(yapFS, "yap")
//line cmd/gopcomm/community.gop:15:1
		if _gop_err != nil {
//line cmd/gopcomm/community.gop:15:1
			_gop_err = errors.NewFrame(_gop_err, "fs.sub(yapFS, \"yap\")", "cmd/gopcomm/community.gop", 15, "main.main")
//line cmd/gopcomm/community.gop:15:1
			panic(_gop_err)
		}
//line cmd/gopcomm/community.gop:15:1
		return
	}()
//line cmd/gopcomm/community.gop:16:1
	y := yap.New(fsYap)
//line cmd/gopcomm/community.gop:18:1
	y.GET("/p/:id", func(ctx *yap.Context) {
//line cmd/gopcomm/community.gop:19:1
		ctx.YAP(200, "article", article{ID: ctx.Param("id")})
	})
//line cmd/gopcomm/community.gop:24:1
	y.Run(":8080")
}
