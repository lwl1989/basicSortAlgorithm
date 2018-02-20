package main

import (
	"net/http"
	"log"
	"io"
	"basicSortAlgorithm/files"
	"fmt"
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

	result,err := files.DoUpload(uploader,con)
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

func main()  {
	con.ReadConstant()
	http.HandleFunc("/f", Upload)
	err := http.ListenAndServe("0.0.0.0:9999", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}