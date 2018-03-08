package main

import (
	"net/http"
	"log"
	"io"
	"basicSortAlgorithm/files"
	"fmt"
	"basicSortAlgorithm/awsS3"
	"net/url"
	"basicSortAlgorithm/image"
)



func Upload(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		w.WriteHeader(500)
		io.WriteString(w,"route err")
		return
	}

	uploader,err :=files.Init(r)
	fmt.Println(uploader.GetMimeType())  //image/png
	fmt.Println(uploader)				 //%!v(PANIC=runtime error: invalid memory address or nil pointer dereference)
	fmt.Println(&uploader)				 //0xc42000e120
	if nil != err {
		w.WriteHeader(500)
		fmt.Println(err)
		io.WriteString(w,err.Error())
		return
	}

	w.WriteHeader(200)
	io.WriteString(w,"success")

	return
	result,err := awsS3.DoUpload(uploader)
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
	code := 404
	content := ""
	if _, ok := query["params"]; ok {
		fileName := query["params"][0]
		//獲取圖片要生成的尺寸
		//判斷文件存在與否
		//存在生成縮略圖
		code,content = image.DoZoom(fileName)
	}
	w.WriteHeader(code)
	io.WriteString(w,content)
}

func main()  {
	awsS3.Con = &files.Constant{Path:"/www/smart/app/stand/global/config/AwsS3Config.php"}
	awsS3.Con.ReadConstant()
	http.HandleFunc("/omp", Upload)
	http.HandleFunc("/zoom", Zoom)
	err := http.ListenAndServe("0.0.0.0:9998", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}