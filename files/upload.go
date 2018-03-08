package files

import (
	"net/http"
	"mime/multipart"
	"fmt"
)

type UploadFile struct {
	file multipart.File
	header *multipart.FileHeader
	uploadConfer *UploadConfer
	error error
}

type UploadInterface interface {
	Judge()
	GetMimeType() string
}

// 获取文件大小的接口
type Size interface {
	Size() int64
}

func  Init(r *http.Request) (uploadFile *UploadFile, err error)  {
	r.ParseMultipartForm(90<<20)  //90M + 10M
	file, header, err := r.FormFile("file")
	sid := r.Header.Get("x-seller-long")
	uploader := &UploadFile{file,header, GetNewConfer(r,file,sid),nil}
	if err != nil {
		return nil,err
	}

	uploader.file = file
	return uploader,nil
}

func (uploadFile *UploadFile) Error() string {
	return uploadFile.error.Error()
}

func (uploadFile *UploadFile) Judge(config UploadConfig) (bool,error) {
	fmt.Println(uploadFile.header) //判断文件头 是否允许
	size := uploadFile.file.(Size).Size() //判断大小是否超过
	fmt.Println(size)
	return false,uploadFile
}

func (uploadFile *UploadFile) GetMimeType() string {
	return uploadFile.header.Header.Get("Content-Type")
}

func (uploadFile *UploadFile) GetSize() int64  {
	return uploadFile.file.(Size).Size()
}

func (uploadFile *UploadFile) Seek(offset int64, whence int) (int64, error)  {
	return uploadFile.file.Seek(offset,whence)
}

func (uploadFile *UploadFile) Read(p []byte) (n int, err error)  {
	return uploadFile.file.Read(p)
}

func (uploadFile *UploadFile) ToS3() bool {
	return false
}

func (uploadFile *UploadFile) GetFrom() map[string]string {

}