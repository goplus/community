package main

import (
	"embed"
	"io/fs"
	"log"
	"github.com/goplus/yap"
)
//line cmd/gopcomm/community.gop:9:1
func check(err error) {
//line cmd/gopcomm/community.gop:10:1
	if err != nil {
//line cmd/gopcomm/community.gop:11:1
		log.Panicln(err)
	}
}

var yapFS embed.FS
//line cmd/gopcomm/community.gop:18
func main() {
//line cmd/gopcomm/community.gop:18:1
	fsYap, err := fs.Sub(yapFS, "yap")
//line cmd/gopcomm/community.gop:19:1
	check(err)
//line cmd/gopcomm/community.gop:21:1
	y := yap.New(fsYap)
//line cmd/gopcomm/community.gop:23:1
	y.Handle("/p/", func(ctx *yap.Context) {
	})
}
