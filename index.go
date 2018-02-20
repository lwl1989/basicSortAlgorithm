package main

import (
	"net/http"
	"log"
	"io"
	"basicSortAlgorithm/files"
	"fmt"
)

func Upload(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		w.WriteHeader(500)
		io.WriteString(w,"route err")
		return
	}

	_,err :=files.Init(r)
	if nil != err {
		w.WriteHeader(500)
		fmt.Println(err)
		io.WriteString(w,err.Error())
		return
	}

	w.WriteHeader(200)
	io.WriteString(w,"success")
}
func main()  {
	http.HandleFunc("/f", Upload)
	err := http.ListenAndServe("0.0.0.0:12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}