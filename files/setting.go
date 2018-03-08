package files

import (
	"errors"
	"strconv"
)

type FileMapping map[string]map[string]string

func GetOmpFileMapping() *FileMapping {
	return &FileMapping{
		"msg": {
			"cover": "omp/msg/image/cover",
			"video": "omp/msg/video",
			//...
		},
		//...
	}
}

type Setting struct {
	FileMapping *FileMapping
	MaxSize     int64
	AllowMime   []string
	SaveExt     string
	Bucket      string
}

func GetOmpSetting() *Setting {
	return &Setting{
		GetOmpFileMapping(),
		5 << 20,
		[]string{"image/png", "image/jpeg"},
		"png",
		"dev-smart-app",
	}
}


func (setting *Setting) JudgeMime(mime string) error {
	for _,v := range setting.AllowMime {
		if v == mime {
			return nil
		}
	}
	return errors.New("MimeType not allow")
}

func (setting *Setting) JudgeSize(size int64) error {
	if size > setting.MaxSize {
		err := "upload size "+strconv.FormatInt(size,10)+" > MaxSize "+strconv.FormatInt(setting.MaxSize,10)
		return errors.New(err)
	}
	return nil
}
