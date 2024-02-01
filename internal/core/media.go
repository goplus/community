package core

import (
	"context"
	"github.com/goplus/yap"
	"github.com/qiniu/x/xlog"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	_ "github.com/qiniu/go-cdk-driver/kodoblob"
)

type File struct {
	Id       int
	FileKey  string
	Format   string
	UserId   int
	Size     int64
	CreateAt time.Time
	UpdateAt time.Time
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

	uId, err := strconv.Atoi(userId)
	if err != nil {
		return err
	}
	mId, err := strconv.Atoi(mediaId)
	if err != nil {
		return err
	}
	// del db media
	res, err := c.db.ExecContext(ctx, "delete from file where user_id = ? and id = ?", uId, mId)
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

	// del cloud oss media
	fileKey, err := c.GetMediaUrl(ctx, mediaId)
	if err != nil {
		return err
	}
	if err := c.bucket.Delete(context.Background(), fileKey); err != nil {
		return err
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
	row.Scan(&fileKey)

	return fileKey, nil
}

func (c *Community) SaveMedia(ctx context.Context, userId string, data []byte) (int64, error) {

	// upload cloud oss
	fileKey := uuid.New().String()
	err := c.uploadMedia(fileKey, data)
	if err != nil {
		return 0, err
	}
	// get fileInfo
	fileInfo, err := c.getMediaInfo(fileKey)
	if err != nil {
		return 0, err
	}
	// save
	stem, err := c.db.Prepare(`insert into file (file_key,format,size,user_id,create_at,update_at) VALUES (?,?,?,?,?,?)`)
	if err != nil {
		return 0, err
	}
	res, err := stem.Exec(fileKey, fileInfo.Format, fileInfo.Size, userId, time.Now(), time.Now())

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// for internal use,no need to add ctx
func (c *Community) getMediaInfo(fileKey string) (*File, error) {

	bucket := c.bucket
	r, err := bucket.NewReader(context.Background(), fileKey, nil)
	defer r.Close()

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

func (c *Community) uploadMedia(fileKey string, data []byte) error {

	w, err := c.bucket.NewWriter(context.Background(), fileKey, nil)
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

func UploadFile(ctx *yap.Context, community *Community) {
	xLog := xlog.New("")
	file, header, err := ctx.FormFile("file")
	filename := header.Filename
	ctx.ParseMultipartForm(10 << 20)
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
	uid, err := community.ParseJwtToken(token.Value)
	if err != nil {
		ctx.JSON(500, err.Error())
	}
	id, err := community.SaveMedia(context.Background(), uid, bytes)
	if err != nil {
		xLog.Error("save file", err.Error())
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, id)
}
