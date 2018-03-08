package files

import (
	"net/http"
	"gopkg.in/h2non/filetype.v1"
	"mime/multipart"
	"io/ioutil"
)


type UploadConfer struct {
	name   string
	path   string
	sha1   string
	module string
	cate   string
	bucket string
	key    string
	sid    string
	suffix string
}

func GetNewConfer(r *http.Request, file multipart.File, sid string) *UploadConfer {
	buf, _ := ioutil.ReadAll(file)
	kind, _ := filetype.Match(buf)
	//size := file.(Size).Size()
	conf := &UploadConfer{
		bucket: "",
		sha1:   r.Form.Get("sha1"),
		module: r.Form.Get("module"),
		cate:   r.Form.Get("type"),
		sid:    sid,
		suffix: kind.Extension,
	}
	generate(conf)

	return conf
}

func (uploadConfer *UploadConfer) SetSuffix(suffix string) {
	uploadConfer.suffix = suffix
}
func (uploadConfer *UploadConfer) SaveName() string {
	return uploadConfer.name
}

func (uploadConfer *UploadConfer) SavePath() string {
	return uploadConfer.path
}

func (uploadConfer *UploadConfer) Sha1() string {
	return uploadConfer.sha1
}

func (uploadConfer *UploadConfer) GetKey() string {
	return uploadConfer.key
}

func (uploadConfer *UploadConfer) Bucket() string {
	return uploadConfer.bucket
}

func generate(uploadConfer *UploadConfer) {
	sha1 := uploadConfer.sha1

	shaSplit := string([]byte(sha1)[:1]) + "/" + string([]byte(sha1)[1:2]) + "/" + string([]byte(sha1)[2:3]) + "/" + string([]byte(sha1)[3:4])
	uploadConfer.name = sha1 + ".png"
	uploadConfer.path = uploadConfer.sid + "/" +
		uploadConfer.module + "/" +
		uploadConfer.cate + "/" +
		shaSplit + "/"

}
