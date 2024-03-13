package core

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gocloud.dev/blob"
)

// Mock API Client
type mockVideoDurationClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Implement NiuTransClient interface
func (m *mockVideoDurationClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestSaveMedia(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	tests := []struct {
		data   []byte
		ext    string
		userId string
	}{
		{
			data:   []byte("test"),
			ext:    "video/mp4",
			userId: "1",
		},
	}
	for _, test := range tests {
		id, err := community.SaveMedia(context.Background(), test.userId, test.data, test.ext)
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
		ext    string
		userId string
	}{
		{
			data:   []byte("test"),
			ext:    "video/mp4",
			userId: "1",
		},
	}

	for _, test := range tests {
		id, err := community.SaveMedia(context.Background(), test.userId, test.data, test.ext)
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
		ext    string
		userId string
	}{
		{
			data:   []byte("test"),
			ext:    "video/mp4",
			userId: "1",
		},
	}

	for _, test := range tests {
		id, err := community.SaveMedia(context.Background(), test.userId, test.data, test.ext)
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
		ext    string
		userId string
	}{
		{
			data:   []byte("test"),
			ext:    "video/mp4",
			userId: "1",
		},
	}

	for _, test := range tests {
		id, err := community.SaveMedia(context.Background(), test.userId, test.data, test.ext)
		assert.Nil(t, err)

		err = community.DelMedias(context.Background(), test.userId, []string{fmt.Sprintf("%d", id)})
		assert.Nil(t, err)
	}
}

func TestGetMediaType(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	_, err := community.db.Exec(
		"insert into file (file_key,format,size,user_id,create_at,update_at,duration) values (?,?,?,?,?,?,?)",
		"test",
		"test",
		1,
		"1",
		"2021-01-01",
		"2021-01-01",
		1,
	)
	assert.Nil(t, err)

	tests := []struct {
		data   []byte
		ext    string
		userId string
	}{
		{
			data:   []byte("test"),
			ext:    "video/mp4",
			userId: "1",
		},
	}

	for _, test := range tests {
		community.S3Service = &mockS3Service{
			DoNewReader: func(ctx context.Context, key string, opts *blob.ReaderOptions) (_ S3Reader, err error) {
				return &mockS3Reader{
					DoRead: func(p []byte) (n int, err error) {
						return 0, nil
					},
					DoClose: func() (err error) {
						return nil
					},
					DoContentType: func() (contentType string) {
						return "video/mp4"
					},
					DoSize: func() (size int64) {
						return 0
					},
				}, nil
			},
			DoNewWriter: func(ctx context.Context, key string, opts *blob.WriterOptions) (_ S3Writer, err error) {
				return &mockWriter{
					DoWrite: func(p []byte) (n int, err error) {
						return 0, nil
					},
					DoClose: func() (err error) {
						return nil
					},
				}, nil
			},
			DoDelete: func(ctx context.Context, key string) (err error) {
				return nil
			},
		}

		mockFormat := `{"format":{"duration":"1"}}`
		community.translation.Engine.HTTPClient = &mockVideoDurationClient{
			DoFunc: func(req *http.Request) (resp *http.Response, err error) {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewReader([]byte(mockFormat))),
				}, nil
			},
		}

		id, err := community.SaveMedia(context.Background(), test.userId, test.data, test.ext)
		assert.Nil(t, err)

		format, err := community.GetMediaType(context.Background(), fmt.Sprintf("%d", id))
		assert.Nil(t, err)
		assert.Equal(t, test.ext, format)
	}
}

func TestGetVideoSubtitle(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	_, err := community.db.Exec(
		"insert into file (file_key,format,size,user_id,create_at,update_at,duration) values (?,?,?,?,?,?,?)",
		"test",
		"test",
		1,
		"1",
		"2021-01-01",
		"2021-01-01",
		1,
	)
	assert.Nil(t, err)
	_, err = community.db.Exec(
		"insert into video_task (user_id, resource_id, task_id, output, status, create_at, update_at) values (?,?,?,?,?,?,?)",
		"1",
		"1",
		"1",
		"test",
		1,
		"2021-01-01",
		"2021-01-01",
	)
	assert.Nil(t, err)

	tests := []struct {
		mediaId string
		fileKey string
		status  string
	}{
		{
			mediaId: "1",
			fileKey: "test",
			status:  "1",
		},
	}

	for _, test := range tests {
		resFileKey, resStatus, err := community.GetVideoSubtitle(context.Background(), test.mediaId)

		assert.Nil(t, err)
		assert.Equal(t, test.fileKey, resFileKey)
		assert.Equal(t, test.status, resStatus)
	}
}

func TestRetryCaptionGenerate(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	_, err := community.db.Exec(
		"insert into file (file_key,format,size,user_id,create_at,update_at,duration) values (?,?,?,?,?,?,?)",
		"test",
		"test",
		1,
		"1",
		"2021-01-01",
		"2021-01-01",
		1,
	)
	assert.Nil(t, err)
	_, err = community.db.Exec(
		"insert into video_task (user_id, resource_id, task_id, output, status, create_at, update_at) values (?,?,?,?,?,?,?)",
		"1",
		"1",
		"1",
		"test",
		1,
		"2021-01-01",
		"2021-01-01",
	)
	assert.Nil(t, err)

	tests := []struct {
		userId  string
		videoId string
	}{
		{
			userId:  "1",
			videoId: "1",
		},
	}

	for _, test := range tests {
		err := community.RetryCaptionGenerate(context.Background(), test.userId, test.videoId)
		assert.NotNil(t, err)
	}
}

func TestListMediaByUserId(t *testing.T) {
	// In memory db
	initClient()
	// Create article table
	initDB()

	// Prepare data
	_, err := community.db.Exec(
		"insert into file (file_key,format,size,user_id,create_at,update_at,duration) values (?,?,?,?,?,?,?)",
		"test",
		"test",
		1,
		"1",
		"2021-01-01",
		"2021-01-01",
		1,
	)
	assert.Nil(t, err)

	tests := []struct {
		userId   string
		format   string
		page     int
		limitInt int
	}{
		{
			userId:   "1",
			format:   "test",
			page:     1,
			limitInt: 10,
		},
	}

	for _, test := range tests {
		medias, _, err := community.ListMediaByUserId(context.Background(), test.userId, test.format, test.page, test.limitInt)
		assert.Nil(t, err)
		assert.NotEqual(t, 0, len(medias))
	}
}
