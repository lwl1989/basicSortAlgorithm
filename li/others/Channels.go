package main

import (
	"fmt"
	"time"
)
type Sender chan<- int

type Receiver <-chan int

func main() {
	ch2 := make(chan string, 1)
	// 下面就是传说中的通过启用一个Goroutine来并发的执行代码块的方法。
	// 关键字 go 后跟的就是需要被并发执行的代码块，它由一个匿名函数代表。
	// 对于 go 关键字以及函数编写方法，我们后面再做专门介绍。
	// 在这里，我们只要知道在花括号中的就是将要被并发执行的代码就可以了。
	go func() {
		ch2 <- "已到达！"
	}()


	var value string = "数据"
	values,ok := <-ch2  //返回值,狀態
	value = value+values
	fmt.Println(values,ok)

	fmt.Println("sss")


		//var myChannel = make(chan int, 0)  //size是0  會阻塞住等待輸出
		//var number = 6
		//go func() {
		//	var sender Sender = myChannel
		//	sender <- number
		//	fmt.Println("Sent!")
		//}()
		//go func() {
		//	var receiver Receiver = myChannel
		//	fmt.Println("Received!", <-receiver)
		//}()
		//// 让main函数执行结束的时间延迟1秒，
		//// 以使上面两个代码块有机会被执行。
		//time.Sleep(time.Second)

		var channels = make(chan int,2)
		var num =0;
		for ;num<6;num++ {

			go func() {
				var sender Sender = channels
				sender <- num
				fmt.Println(num)
			}()
		}
	fmt.Println("阻塞了？")
		go func() {

			var receiver Receiver = channels
			dd := <-receiver
			fmt.Println("釋放了一個",dd)
		}()
		time.Sleep(time.Second)
	}
