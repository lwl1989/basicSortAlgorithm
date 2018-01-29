package sort

import "fmt"
/**
快速排序
外循环（自加1），条件是数组长度-1
	从数组首部取值放入比较单元（临时）
		内循环
		从数组尾部开始循环，条件是已经排序好的外循环次数
		假如当前元素小于 比较单元 swap  当前元素给 比较单元  这样就保证了每次将最小的元素有序放入数组首部
 */
func Quick(arr []int,len int) []int {
	for i:=0;i<len;i++ {
		temp := arr[i]
		for j:=len-1; j>i; j-- {
			if i==2 {
				for offset,num := range arr {
					fmt.Printf("s[%d] == %d  ", offset, num)
				}
				fmt.Println()
				fmt.Println(arr[i]+' '+arr[j])
			}
			if temp>arr[j] {
				arr[i] = arr[j]
				arr[j] = temp
				temp = arr[i]
			}
		}
	}
	return arr
}
