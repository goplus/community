package core

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
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
		log.Println("no need del data")
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
	uId, err := strconv.Atoi(userId)
	if err != nil {
		return 0, err
	}
	// save
	stem, err := c.db.Prepare(`insert into file (file_key,format,size,user_id,create_at,update_at) VALUES (?,?,?,?,?,?)`)
	if err != nil {
		return 0, err
	}
	defer c.db.Close()
	res, err := stem.Exec(fileKey, fileInfo.Format, fileInfo.Size, uId, time.Now(), time.Now())

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// for internal use,no need to add ctx
func (c *Community) getMediaInfo(fileKey string) (*File, error) {

	bucket := c.bucket
	r, err := bucket.NewReader(context.Background(), fileKey, nil)

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

	w, err := c.bucket.NewWriter(context.Background(), "", nil)
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
