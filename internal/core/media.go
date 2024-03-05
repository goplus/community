package core

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"

	"gocloud.dev/blob"

	"github.com/google/uuid"
	"github.com/goplus/yap"
	_ "github.com/qiniu/go-cdk-driver/kodoblob"
	"github.com/qiniu/x/xlog"
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
	UserId   int
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
	if err := c.bucket.Delete(context.Background(), fileKey); err != nil {
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
	c.xLog.Info(fileKey)
	if strings.Contains(fileExt, "video") {
		fileExt = ".mp4"
		encoding := base64.URLEncoding.EncodeToString([]byte(c.bucketName + ":" + fileKey + fileExt))
		opts.Metadata = make(map[string]string)
		opts.Metadata["persistent-ops"] = "avthumb/mp4/vcodec/libx264|saveas/" + encoding
	}
	fileKey = fileKey + fileExt
	c.xLog.Info(fileKey)
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
		duration, err = GetVideoDuration(c.domain + fileKey + "?avinfo")
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

// for internal use,no need to add ctx
func (c *Community) getMediaInfo(fileKey string) (*File, error) {

	bucket := c.bucket
	r, err := bucket.NewReader(context.Background(), fileKey, nil)
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

func GetVideoDuration(url string) (duration string, err error) {
	resp, err := http.Get(url)
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
	if c.bucket == nil {
		return nil
	}

	w, err := c.bucket.NewWriter(context.Background(), fileKey, opts)
	if err != nil {
		return err
	}
	defer w.Close()
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Community) UploadFile(ctx *yap.Context) {
	xLog := xlog.New("")
	file, header, err := ctx.FormFile("file")
	if err != nil {
		xLog.Error("upload file error:", header)
		ctx.JSON(500, err.Error())
		return
	}
	filename := header.Filename
	err = ctx.ParseMultipartForm(10 << 20)
	if err != nil {
		xLog.Error("upload file error:", filename)
		ctx.JSON(500, err.Error())
		return
	}

	dst, err := os.Create(filename)
	if err != nil {
		xLog.Error("create file error:", file)
		ctx.JSON(500, err.Error())
		return
	}
	defer func() {
		file.Close()
		dst.Close()
		err = os.Remove(filename)
		if err != nil {
			xLog.Error("delete file error:", filename)
			return
		}
	}()

	_, err = io.Copy(dst, file)
	if err != nil {
		xLog.Error("copy file errer:", filename)
		ctx.JSON(500, err.Error())
		return
	}

	bytes, err := os.ReadFile(filename)
	if err != nil {
		xLog.Error("read file errer:", filename)
		ctx.JSON(500, err.Error())
		return
	}

	token, err := GetToken(ctx)
	if err != nil {
		ctx.JSON(500, err.Error())
	}

	uid, err := c.ParseJwtToken(token.Value)
	if err != nil {
		ctx.JSON(500, err.Error())
	}
	ext := mimetype.Detect(bytes).String()
	id, err := c.SaveMedia(context.Background(), uid, bytes, ext)
	if err != nil {
		xLog.Error("save file", err.Error())
		ctx.JSON(500, err.Error())
		return
	}

	// Judge the file type and start the corresponding task
	if strings.Contains(ext, "video") {
		// Start captioning task
		err := c.NewVideoTask(context.Background(), uid, strconv.FormatInt(id, 10))
		if err != nil {
			xLog.Error("start video task error:", err)
		}
	}

	ctx.JSON(200, id)
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

func (c *Community) ListSubtitleByVideoId(ctx context.Context, videoId int) ([]VideoSubtitle, error) {
	sqlStr := "select * from video_subtitle where video_id =? and "
	rows, err := c.db.Query(sqlStr, videoId)
	var videoSubtitles []VideoSubtitle
	if err != nil {
		return videoSubtitles, err
	}
	defer rows.Close()

	for rows.Next() {
		var videoSubtitle VideoSubtitle
		err = rows.Scan(&videoSubtitle.VideoId, &videoSubtitle.SubtitleId, &videoSubtitle.UserId, &videoSubtitle.Language)
		if err != nil {
			return videoSubtitles, err
		}
		videoSubtitles = append(videoSubtitles, videoSubtitle)
	}
	return videoSubtitles, nil
}

func (c *Community) AddSubtitle(ctx context.Context, videoId, subtitleId, userId int, language string) error {
	sqlStr := "INSERT INTO video_subtitle (video_id,user_id,subtitle_id,language) values (?,?,?,?)"
	_, err := c.db.Exec(sqlStr, videoId, userId, subtitleId, language)
	if err != nil {
		c.xLog.Fatalln(err.Error())
		return err
	}
	return nil
}

func (c *Community) DelSubtitle(ctx context.Context, videoId, subtitleId, userId int) error {
	sqlStr := "DELETE FROM video_subtitle where video_id = ? and subtitle_id = ? and user_id = ?"
	_, err := c.db.Exec(sqlStr, videoId, subtitleId, userId)
	if err != nil {
		c.xLog.Fatalln(err.Error())
		return err
	}
	return nil
}
