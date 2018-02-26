package main

import (
	"net/http"
	"log"
	"io"
	"basicSortAlgorithm/files"
	"fmt"
	"basicSortAlgorithm/awsS3"
	"net/url"
)

var con = &files.Constant{Path:"/www/smart/app/stand/global/config/AwsS3Config.php"}

func Upload(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		w.WriteHeader(500)
		io.WriteString(w,"route err")
		return
	}

	uploader,err :=files.Init(r)
	if nil != err {
		w.WriteHeader(500)
		fmt.Println(err)
		io.WriteString(w,err.Error())
		return
	}

	result,err := awsS3.DoUpload(uploader,con)
	if nil != err {
		w.WriteHeader(500)
		fmt.Println(err)
		io.WriteString(w,err.Error())
		return
	}
	fmt.Println(result)
	w.WriteHeader(200)
	io.WriteString(w,"success")
}

func Zoom(w http.ResponseWriter, r *http.Request)  {
	query,_ := url.ParseQuery(r.URL.RawQuery)

	if _, ok := query["param"]; ok {
		fmt.Println(ok)
		fileName := query["param"][0]

		//獲取圖片要生成的尺寸
		//判斷文件存在與否
		//存在生成縮略圖
		fmt.Println(fileName)
	}
}

func main()  {
	con.ReadConstant()
	http.HandleFunc("/omp", Upload)
	http.HandleFunc("/zoom", Zoom)
	err := http.ListenAndServe("0.0.0.0:9999", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}