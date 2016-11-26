package main;

import (
	"li/sort"
	"fmt"
)

func main() {
	arr := []int{9,3,5,6,7,8}
	sort.Shell(arr, len(arr))

/*	selectArr := sort.Select(arr, len(arr))
	for _,value := range selectArr {
		fmt.Printf("%d\r\n", value)
	}*/


	insertArr := sort.Insert(arr, len(arr))
	for key,value := range insertArr {
		fmt.Printf("%d:%d\r\n", key,value)
	}
}
