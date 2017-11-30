package sort

//import "fmt"

/**
从无序数组中获得一个数字
	对已有序的数据进行比较
		当这个数字大于或者小于（看排序标准）时
			当前下标为这个数字，其他数据进行偏移
9 3 5 6 7 8
	temp = 3 j 0 i 1  3!>9 j = 1
		9 3

 */
func Insert(arr []int, len int) []int  {
	for i := 1; i < len; i++ {
		temp := arr[i];
		j:=0

		for  ; j < i && temp > arr[j]; j++ {
			arr[j+1] = arr[j]
		}

		arr[j+1] = temp
	}
	return arr
}
