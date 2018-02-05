package main

import "fmt"

func main() {
	var number = 5
	if number += 4; 10 > number {

		number += 3
		fmt.Print(number)
	} else if 10 < number {
		number -= 2
		fmt.Print(number)
	}
	fmt.Println()
	fmt.Println(number)   //此时是main函数的作用域
	fmt.Printf("%p\n",&number)
	if number := 3; number < 3 {
		fmt.Println(number)  //此时是if作用域内的number
		fmt.Printf("%p\n",&number)
	}else if number :=4; number!=4 {
		fmt.Println(number)  //此时是if作用域内的number
		fmt.Printf("%p\n",&number)
	}else{
		fmt.Println(number)  //此时是最近if重写的number[配对if else]
		fmt.Printf("%p\n",&number)
	}
	fmt.Println(number)//此时是main函数的作用域
}
