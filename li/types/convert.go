package main

import (
	"fmt"
	"strconv"
)

const (
	a = 1
	b
	c
)

const (
	e,f = 1,"2"
	g,h
)

const  (
	ddd = iota
	ccc = 'a'
	bbb
	fff= iota
)

func main()  {
	//ccc=97 bbb=97 ddd=0 fff=3   iota和你在出现的第几个位置有关
	fmt.Println(ccc,bbb,ddd,fff)
	// a=1 b=1 c=1
	fmt.Println(a,b,c)
	// e=1 f="2" g=1 h="2"
	fmt.Println(e,f,g,h)
	a := 3.8
	b := int(a)
	fmt.Println(b)

	//b1 := bool(a)  //can't convert num => bool
	//fmt.Println(b1)

	//边界问题
	var aa int = 129
	bb := strconv.Itoa(aa)
	aa,_ = strconv.Atoi(bb)
	cc := int64(aa)
	dd := int8(aa)  //   127 -128 -127
	ff := 257
	ee := int8(ff) //  255 0 1
	fmt.Println(aa,bb,cc,dd,ee)

	m :=0
	if m!=0 && 10/m > 1 {
		fmt.Println("ok")
	}

}