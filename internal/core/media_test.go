package core

import (
	"context"
	"fmt"
	"testing"

	"github.com/gabriel-vasile/mimetype"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/assert"
)

func TestSaveMedia(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	tests := []struct {
		data   []byte
		userId string
	}{
		{
			data:   []byte("test"),
			userId: "1",
		},
	}
	for _, test := range tests {
		id, err := community.SaveMedia(context.Background(), test.userId, test.data, mimetype.Detect(test.data).String())
		assert.Nil(t, err)
		assert.NotEqual(t, 0, id)
	}
}

func TestGetMediaUrl(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	tests := []struct {
		data   []byte
		userId string
	}{
		{
			data:   []byte("test"),
			userId: "1",
		},
	}

	for _, test := range tests {
		id, err := community.SaveMedia(context.Background(), test.userId, test.data, mimetype.Detect(test.data).String())
		assert.Nil(t, err)

		url, err := community.GetMediaUrl(context.Background(), fmt.Sprintf("%d", id))
		assert.Nil(t, err)
		assert.NotEqual(t, "", url)
	}
}

func TestDelMeida(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	tests := []struct {
		data   []byte
		userId string
	}{
		{
			data:   []byte("test"),
			userId: "1",
		},
	}

	for _, test := range tests {
		id, err := community.SaveMedia(context.Background(), test.userId, test.data, mimetype.Detect(test.data).String())
		assert.Nil(t, err)

		err = community.DelMedia(context.Background(), test.userId, fmt.Sprintf("%d", id))
		assert.Nil(t, err)
	}
}

func TestDelMedias(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	tests := []struct {
		data   []byte
		userId string
	}{
		{
			data:   []byte("test"),
			userId: "1",
		},
	}

	for _, test := range tests {
		id, err := community.SaveMedia(context.Background(), test.userId, test.data, mimetype.Detect(test.data).String())
		assert.Nil(t, err)

		err = community.DelMedias(context.Background(), test.userId, []string{fmt.Sprintf("%d", id)})
		assert.Nil(t, err)
	}

}
