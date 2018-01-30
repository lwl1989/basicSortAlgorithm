package main

import "fmt"

func main()  {

	var numbers [10]int
	numbers[0] = 5
	numbers[3] = 7
	numbers[6] = 3

	fmt.Println(numbers)

	//var numbers1 []int
	//直接赋值会报index out
	//numbers1[0] = 5

	//初始化切片的几种方式 1.存在已有数组  将数组赋值进入，并且可以选择切片空间  起始下标:终止下标:分配的内存（cap） len= 终止下标- 起始下标   //cap 等于 cap值-起始  起始 : len值 ： cap值
 	numbers1 := numbers[3:5:10]
	fmt.Println("初始化切片的几种方式 1.存在已有数组  将数组赋值进入，并且可以选择切片空间  起始下标:终止下标:分配的内存（cap） len= 终止下标- 起始下标")
	fmt.Print("数据长度为：")
	fmt.Println(len(numbers1))
	fmt.Print("占用长度为：")
	fmt.Println(cap(numbers1))
	numbers1[0] = 5

	fmt.Println()
	numbers2 := make([]int,3,10)
	//长度为make的size
	fmt.Println("初始化切片的几种方式 2.make(类型，长度，分配内存)")
	fmt.Print("数据长度为：")
	fmt.Println(len(numbers2))
	fmt.Print("占用长度为：")
	fmt.Println(cap(numbers2))
	numbers2[0] = 5
	fmt.Println("超出len(slice)的会报错:panic: runtime error: index out of range,比如 numbers2[6] = 3")

	fmt.Println()
	fmt.Println("若想动态扩容，请使用append")
	numbers2 = append(numbers2,1,3,4,5)
	fmt.Println(numbers2)
	fmt.Println("还可以合并，类似PHP的array_merge,但是要注意下标的变动")
	numbers2 = append(numbers2, numbers1...)
	fmt.Println("当长度超出原来内存的时候，会自动分配双倍之前的内存（cap）")
	numbers2 = append(numbers2, 6,7,8,9)
	fmt.Print("此时数据长度为：")
	fmt.Println(len(numbers2))
	fmt.Print("占用长度为：")
	fmt.Println(cap(numbers2))
	fmt.Println()

	fmt.Println(numbers1)
	fmt.Println(numbers2)

}
