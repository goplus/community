import (
    "os"
	"io"
	c "context"
	"strings"
	"strconv"

	"github.com/gabriel-vasile/mimetype"
	"github.com/qiniu/x/log"
)

file, header, err := FormFile("file")
if err != nil {
	log.Error("upload file error:", header)
	JSON(500, err.Error())
	return
}
filename := header.Filename
err = ParseMultipartForm(10 << 20)
if err != nil {
	log.Error("upload file error:", filename)
	JSON(500, err.Error())
	return
}

dst, err := os.Create(filename)
if err != nil {
	log.Error("create file error:", file)
	JSON(500, err.Error())
	return
}
defer func() {
	file.Close()
	dst.Close()
	err = os.Remove(filename)
	if err != nil {
		log.Error("delete file error:", filename)
		return
	}
}()

_, err = io.Copy(dst, file)
if err != nil {
	log.Error("copy file errer:", filename)
	JSON(500, err.Error())
	return
}

bytes, err := os.ReadFile(filename)
if err != nil {
	log.Error("read file errer:", filename)
	JSON(500, err.Error())
	return
}

token, err := Request.Cookie("token")
if err != nil {
	JSON(500, err.Error())
}

uid, err := community.ParseJwtToken(token.Value)
if err != nil {
	JSON(500, err.Error())
}
ext := mimetype.Detect(bytes).String()
id, err := community.SaveMedia(c.TODO(), uid, bytes, ext)
if err != nil {
	log.Error("save file", err.Error())
	JSON(500, err.Error())
	return
}

// Judge the file type and start the corresponding task
if strings.Contains(ext, "video") {
	// Start captioning task
	err := community.NewVideoTask(c.TODO(), uid, strconv.FormatInt(id, 10))
	if err != nil {
		log.Error("start video task error:", err)
	}
}

JSON(200, id)
