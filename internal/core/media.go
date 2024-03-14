package core

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gocloud.dev/blob"

	"github.com/google/uuid"
	_ "github.com/qiniu/go-cdk-driver/kodoblob"
)

// todo
type VideoSubtitle struct {
	VideoId    int
	SubtitleId int
	UserId     int
	Language   string
}

type File struct {
	Id       int
	FileKey  string
	Format   string
	UserId   string
	Size     int64
	CreateAt time.Time
	UpdateAt time.Time
	Duration *string
	Vtt      string
}

func (c *Community) DelMedias(ctx context.Context, userId string, ids []string) error {

	uId, err := strconv.Atoi(userId)
	if err != nil {
		return err
	}
	idStr := strings.Join(ids, ",")
	_, err = c.db.ExecContext(ctx, "delete from file where user_id = ? and id in( ? )", uId, idStr)

	return err
}

func (c *Community) DelMedia(ctx context.Context, userId, mediaId string) error {
	// del cloud oss media
	fileKey, err := c.GetMediaUrl(ctx, mediaId)
	if err != nil {
		return err
	}
	if err := c.S3Service.Delete(context.Background(), fileKey); err != nil {
		return err
	}
	// del db media
	res, err := c.db.ExecContext(ctx, "delete from file where user_id = ? and id = ?", userId, mediaId)
	if err != nil {
		return err
	}
	aff, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if aff == 0 {
		c.xLog.Warn("no need del data")
		return nil
	}
	return nil
}

// get file key
func (c *Community) GetMediaUrl(ctx context.Context, mediaId string) (string, error) {
	ID, err := strconv.Atoi(mediaId)
	if err != nil {
		return "", err
	}
	row := c.db.QueryRow(`select file_key from file where id = ?`, ID)
	var fileKey string
	err = row.Scan(&fileKey)
	if err != nil {
		return "", err
	}

	return fileKey, nil
}

func (c *Community) GetMediaType(ctx context.Context, mediaId string) (string, error) {
	ID, err := strconv.Atoi(mediaId)
	if err != nil {
		return "", err
	}
	row := c.db.QueryRow(`select format from file where id = ?`, ID)
	var Format string
	err = row.Scan(&Format)
	if err != nil {
		return "", err
	}
	return Format, nil

}
func (c *Community) GetVideoSubtitle(ctx context.Context, mediaId string) (string, string, error) {
	ID, err := strconv.Atoi(mediaId)
	if err != nil {
		return "", "-1", err
	}
	row := c.db.QueryRow(`select output, status from video_task where resource_id = ?`, ID)
	var fileKey string
	var status string
	err = row.Scan(&fileKey, &status)
	if err != nil {
		return "", "-1", err
	}

	return fileKey, status, nil
}

func (c *Community) SaveMedia(ctx context.Context, userId string, data []byte, fileExt string) (int64, error) {
	opts := &blob.WriterOptions{}
	// upload cloud oss
	fileKey := uuid.New().String()
	if strings.Contains(fileExt, "video") {
		fileExt = ".mp4"
		encoding := base64.URLEncoding.EncodeToString([]byte(c.bucketName + ":" + fileKey + fileExt))
		opts.Metadata = make(map[string]string)
		opts.Metadata["persistent-ops"] = "avthumb/mp4/vcodec/libx264|saveas/" + encoding
	}
	fileKey = fileKey + fileExt
	err := c.uploadMedia(fileKey, data, opts)
	if err != nil {
		c.xLog.Warn(err.Error())
		return 0, err
	}
	// get fileInfo
	fileInfo, err := c.getMediaInfo(fileKey)
	if err != nil {
		c.xLog.Warn(err.Error())
		return 0, err
	}

	var duration string = ""
	if fileInfo.Format == "video/mp4" {
		duration, err = c.GetVideoDuration(c.domain + fileKey + "?avinfo")
		if err != nil {
			c.xLog.Warn(err.Error())
			return 0, err
		}
		duration, err = formatVideoDuration(duration)
		c.xLog.Info(duration)
		if err != nil {
			c.xLog.Warn(err.Error())
			return 0, err
		}
	}
	// save
	stem, err := c.db.Prepare(`insert into file (file_key,format,size,user_id,create_at,update_at,duration) VALUES (?,?,?,?,?,?,?)`)
	if err != nil {
		c.xLog.Warn(err.Error())
		return 0, err
	}
	res, err := stem.Exec(fileKey, fileInfo.Format, fileInfo.Size, userId, time.Now(), time.Now(), duration)

	if err != nil {
		c.xLog.Warn(err.Error())
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}

func formatVideoDuration(duration string) (string, error) {
	seconds, err := strconv.ParseFloat(duration, 64)
	if err != nil {
		return "", err
	}

	minutes := math.Floor(seconds / 60)
	secondsRemainder := int(seconds) % 60

	formattedTime := fmt.Sprintf("%d:%02d", int(minutes), secondsRemainder)
	return formattedTime, nil
}

// for internal use,no need to add ctx
func (c *Community) getMediaInfo(fileKey string) (*File, error) {
	r, err := c.S3Service.NewReader(context.Background(), fileKey, nil)
	defer func() {
		err = r.Close()
		if err != nil {
			c.xLog.Error("close file error:", err)
		}
	}()

	if err != nil {
		return nil, err
	}
	fileInfo := &File{
		Format:  r.ContentType(),
		Size:    r.Size(),
		FileKey: fileKey,
	}

	return fileInfo, nil
}

func (c *Community) GetVideoDuration(url string) (duration string, err error) {
	// resp, err := http.Get(url)
	var req *http.Request
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Use mock client instead of real client
	resp, err := c.translation.Engine.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	type Format struct {
		Duration string `json:"duration"`
	}
	var Data struct {
		Format Format `json:"format"`
	}

	err = json.Unmarshal(body, &Data)

	if err != nil {
		return "", err
	}
	duration = Data.Format.Duration

	decimalIndex := -1
	for i := 0; i < len(duration); i++ {
		if duration[i] == '.' {
			decimalIndex = i
			break
		}
	}
	if decimalIndex != -1 {
		duration = duration[:decimalIndex+3]
	}
	return duration, nil
}

func (c *Community) uploadMedia(fileKey string, data []byte, opts *blob.WriterOptions) error {
	if c.S3Service == nil {
		return nil
	}

	w, err := c.S3Service.NewWriter(context.Background(), fileKey, nil)
	if err != nil {
		c.xLog.Warn(err.Error())
		return err
	}
	defer w.Close()
	_, err = w.Write(data)
	if err != nil {
		c.xLog.Warn(err.Error())
		return err
	}
	return nil
}

func (c *Community) RetryCaptionGenerate(ctx context.Context, userId, videoId string) error {
	err := c.NewVideoTask(ctx, userId, videoId)
	if err != nil {
		return err
	}
	return nil
}

// func (c *Community) ListMediaByUserId(ctx context.Context, userId string, format string) ([]File, error) {
func (c *Community) ListMediaByUserId(ctx context.Context, userId string, format string, page, limitInt int) ([]File, int, error) {
	sqlStr := "select * from file where user_id = ? "
	var args []any
	args = append(args, userId)
	var rows *sql.Rows
	var err error
	if format != "" {
		sqlStr += " and format like ?"
		args = append(args, "%"+format+"%")
	}
	sqlStr = sqlStr + " order by create_at desc limit ? offset ?"
	args = append(args, limitInt)
	args = append(args, (page-1)*limitInt)

	rows, err = c.db.Query(sqlStr, args...)

	var files []File
	if err != nil {
		return files, 0, err
	}
	for rows.Next() {
		var file File
		if err := rows.Scan(&file.Id, &file.CreateAt, &file.UpdateAt, &file.FileKey, &file.Format, &file.UserId, &file.Size, &file.Duration); err != nil {
			return files, 0, err
		}
		file.FileKey = c.domain + file.FileKey
		if format == "video" {
			vtt, status, _ := c.GetVideoSubtitle(ctx, strconv.Itoa(file.Id))
			if status == "1" {
				file.Vtt = c.domain + vtt
			}
		}

		files = append(files, file)

	}
	var count int
	err = c.db.QueryRow("select count(*) from file where user_id = ? and format like ?", userId, "%"+format+"%").Scan(&count)
	if err != nil {
		return files, 0, err
	}
	return files, count, nil
}
