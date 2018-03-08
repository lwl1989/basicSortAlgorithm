package files

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/henrylee2cn/pholcus/common/mgo"
	"fmt"
	"reflect"
)

//_id,module,file_path,cover_path,size,play_time,rotation
//user keyId=? and source=? and project = ? and type = ?
type fileSave struct {
	id       bson.ObjectId `_id`
	module   string
	filePath string	`file_path`
	coverPath	string	`cover_path`
	sha1     string	`sha1_id`

	size	 string
	playTime int32
	rotation int32

	sid      string `uploader_seller`
	uid		 string `upload_user`
	extName  string `ext_name`
	fileSize string `file_size`

	upTime   bson.MongoTimestamp
}

func (fileSave *fileSave) insert() {
	docs := map[string]interface{}{

	}


	value := reflect.TypeOf(fileSave)
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}

	fmt.Println(docs)
	insert := mgo.Insert{
		Database:"small_app_log",
		Collection:"omp_file_meta",
	//	Docs:[]map[string]interface{}
	}

	fmt.Println(insert)
}
