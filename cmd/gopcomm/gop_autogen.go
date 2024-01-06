package main

import (
	"embed"
	"io/fs"
	"github.com/goplus/yap"
	"github.com/qiniu/x/errors"
)
//go:embed yap
var yapFS embed.FS
//line community.gop:11
func main() {
//line community.gop:11:1
	fsYap := func() (_gop_ret fs.FS) {
//line community.gop:11:1
		var _gop_err error
//line community.gop:11:1
		_gop_ret, _gop_err = fs.Sub(yapFS, "yap")
//line community.gop:11:1
		if _gop_err != nil {
//line community.gop:11:1
			_gop_err = errors.NewFrame(_gop_err, "fs.sub(yapFS, \"yap\")", "community.gop", 11, "main.main")
//line community.gop:11:1
			panic(_gop_err)
		}
//line community.gop:11:1
		return
	}()
//line community.gop:12:1
	y := yap.New(fsYap)
//line community.gop:14:1
	y.GET("/p/:id", func(ctx *yap.Context) {
//line community.gop:15:1
		ctx.YAP(200, "article", yap.H{"id": ctx.Param("id")})
	})
//line community.gop:20:1
	y.Run(":8080")
}
