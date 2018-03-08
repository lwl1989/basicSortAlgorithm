package files

import (
	"net/http"
	"mime/multipart"
	"fmt"
)

type UploadFile struct {
	file multipart.File
	header *multipart.FileHeader
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
	defer file.Close()
	if err != nil {
		return nil,err
	}
	//sid := r.Header.Get("x-seller-long")
	var uploader *UploadFile = nil
	uploader = &UploadFile{file,header,nil}

	fmt.Println(uploader.GetMimeType())  //image/png
	fmt.Println(uploader)				 //%!v(PANIC=runtime error: invalid memory address or nil pointer dereference)
	fmt.Println(&uploader)				 //0xc42000e120


	err = uploader.Judge(GetOmpSetting())
	if err != nil {
		return nil,err
	}

	uploader.file = file
	return uploader,nil
}

func (uploadFile *UploadFile) Error() string {
	return uploadFile.error.Error()
}

func (uploadFile *UploadFile) Judge(setting *Setting) error {
	err := setting.JudgeMime(uploadFile.GetMimeType())
	if err != nil {
		return err
	}
	err = setting.JudgeSize(uploadFile.file.(Size).Size())
	if err != nil {
		return err
	}
	return nil
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
