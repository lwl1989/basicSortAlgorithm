package files

import (
	"net/http"
	"basicSortAlgorithm/awsS3"
)

type FileMapping map[string]map[string]string

var OmpFileMapping = FileMapping{
	"msg"	:	{
		"cover"	:	"omp/msg/image/cover",
		"video"	:	"omp/msg/video",
		//...
	},
	//...
}

type UploadConfig interface {
	SaveName() string
	SavePath() string
	Sha1()	string
	SaveBucket() string
	GetKey() string
}


type UploadConfer struct {
	name string
	path string
	sha1 string
	bucket string
	key string
	confType string
}

func GetNewConfer(r *http.Request) *UploadConfer {
	//sha1 := r.Form.Get("sha1")
	return &UploadConfer{
		bucket: awsS3.Con.GetConstant("S3_REGION"),
	}
}

func (uploadConfer *UploadConfer) SaveName() string {

}




