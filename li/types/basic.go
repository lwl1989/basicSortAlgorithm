package main

import (
	"fmt"
	"math"
)


//type alisa
type (
	byte uint8
	rune int32
	文本 string
)


var strGlobal1 string = "哈哈哈" //全局变量必须显示赋值
var strGlobal2 string  //此时有默认值空字符串

func main()  {


	b := true   // 长度1 取值范围  true false

	i:=1
	var ui uint = 1   // 和运行平台位数相关  32 or 64
	it := int(ui)
	if i == it {
		b = false
	}
	fmt.Println(b,i,ui)

	// int8 1 byte  -128~127  uint8  0~255   uint8 == byte 别名
	// int16 2 byte -32768~32767 uint16 0~65535
	// int32 4 byte  -2^32/2 ~ 2^32/2-1  uint32 0~2^32-1     rune == int32
	// int64 8 byte -2^64/2 ~ 2^64/2-1 uint64 0~2^64-1
	// float32 4 byte
	// float64 8 byte

	// complex64/complex128    8~16 byte  复数类型
	// uintptr  保存指针类型

	// array struct string
	// slice map chan  引用类型（指针）


	// interface  接口类型
	// func  函数类型


	//类型的零值(默认值)
	// 数字  0   bool false  string 空字符串
	var int1 int32
	fmt.Println(int1)
	var bool1 bool
	fmt.Println(bool1)
	var str string
	fmt.Println(str)
	var a []int    //没有大小 切片
	fmt.Println(a)
	var c [1]int   //有大小 才是数组
	fmt.Println(c)


	fmt.Println(math.MaxUint8)

	var strC 文本
	strC = "我的天"
	fmt.Println(strC)


	fmt.Println(strGlobal2,strGlobal1)

	var d,e,f,g int = 1,2,3,4
	// var d,e,f,g = 1,2,3,4  d,e,f,g := 1,2,3,4
	fmt.Println(d,e,f,g)
}