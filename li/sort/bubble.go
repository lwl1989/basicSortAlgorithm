package sort
/**
粗暴排序 冒泡
    逐个比较元素，然后进行交换
 */

func Bubble(arr []int, len int) []int  {
	for i:=0;i<len;i++ {
		for j:=0; j<len-1; j++ {
			if arr[j]>arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}
