package main

import "fmt"

func main() {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}  //分配的连续内存
	slice5 := numbers4[4:6:8]
	length := 2
	capacity := 4   //cap 等于 cap值-起始  起始 : len值 ： cap值
	fmt.Println(&numbers4[4])   //0xc420094020
	fmt.Println(&slice5[0])   //0xc420094020  5 6  地址一直未改变过

	fmt.Printf("%v, %v\n", length == len(slice5), capacity == cap(slice5))
	slice5 = slice5[:cap(slice5)]  //将len改变成4 cap还是4
	fmt.Println(&slice5[0])  //5 6 7 8  并没有改变数组的指针 0xc420094020
	slice5 = append(slice5, 11, 12, 13)
	length = 5
	fmt.Printf("%v\n", length == len(slice5))



	slice6 := []int{0, 0, 0}
	//slice5 5 6 7 8 11 12 13
	copy(slice5, slice6)
	//slice5 0 0 0 8 11 12 13
	fmt.Println(&numbers4[0])
	fmt.Println(&slice6[0])  //copy过来 产生的是新地址
	fmt.Println(cap(slice6),len(slice6))  //长度等于容量等于初始化时候的长度
	fmt.Println(slice6)      //

	e2 := 0
	e3 := 8
	e4 := 11
	fmt.Printf("%v, %v, %v\n", e2 == slice5[2], e3 == slice5[3], e4 == slice5[4])
}