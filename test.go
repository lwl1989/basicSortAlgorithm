package main

import (
	"net/http"
	"io"
	"fmt"
	"log"
)
func Test(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		w.WriteHeader(500)
		io.WriteString(w,"route err")
		return
	}
	r.ParseMultipartForm(90 << 20)
	fmt.Println(r.MultipartForm)
	w.WriteHeader(200)
	io.WriteString(w,"success")
}

func main()  {
	http.HandleFunc("/omp", Test)
	err := http.ListenAndServe("0.0.0.0:9999", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}