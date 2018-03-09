package main

import "fmt"

const (
	PI = 3.14
	C1 = 3213
	C2 = "dasd"
)

var (
	name = "张三"
	age uint = 22
)

type (
	People struct {
		name string
		age uint
	}

	newString string
)


func main() {
	fmt.Println(&People{name:name,age:age})
	fmt.Println(PI,C1,C2)

	var ff newString
	ff = "dasdsa"

	fmt.Println(ff)
}