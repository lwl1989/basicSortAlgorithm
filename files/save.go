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
	_id        bson.ObjectId
	module     string
	file_path  string
	cover_path string
	sha1_id    string

	size     string
	playTime int32
	rotation int32

	uploader_seller string
	upload_user     string
	ext_name        string
	file_size       string

	upTime bson.MongoTimestamp
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
		Database:   "small_app_log",
		Collection: "omp_file_meta",
		//	Docs:[]map[string]interface{}
	}

	fmt.Println(insert)
}
