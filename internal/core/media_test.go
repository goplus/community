package core_test

import (
	"context"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"log"
	"os"
	"testing"

	"github.com/goplus/community/internal/core"
)

var c *core.Community

func TestGetMediaUrl(t *testing.T) {
	url, err := c.GetMediaUrl(context.Background(), "10")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	log.Println(url)
}

func TestSaveMedia(t *testing.T) {

	conf := &core.Config{}
	todo := context.TODO()
	c, err := core.New(todo, conf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data, err := readFileByte("")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	id, err := c.SaveMedia(context.Background(), "", data, mimetype.Detect(data).String())
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
	data, err := os.ReadFile(filePath)
	return data, err
}
