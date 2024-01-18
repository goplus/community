package core_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/goplus/community/internal/core"
)

var c *core.Community

func TestMain(t *testing.M) {
	config := &core.Config{
		Driver: "mysql",
		DSN:    "",
		BlobUS: "",
	}
	ret, err := core.New(context.Background(), config)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	c = ret
	t.Run()
}

func TestGetMediaUrl(t *testing.T) {
	url, err := c.GetMediaUrl(context.Background(), "10")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	log.Println(url)
}

func TestSaveMedia(t *testing.T) {
	data, err := readFileByte("xxx")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	id, err := c.SaveMedia(context.Background(), "1", data)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	log.Println("id:", id)

}

func TestDelMeida(t *testing.T) {
	err := c.DelMedia(context.Background(), "1", "10")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	log.Println(err.Error())
}

func TestDelMedias(t *testing.T) {
	err := c.DelMedias(context.Background(), "1", []string{"1"})
	if err != nil {

		log.Fatalln(err.Error())
		return
	}

}

func readFileByte(filePath string) ([]byte, error) {
	data, err := os.ReadFile("sss")
	return data, err
}
