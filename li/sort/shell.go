package sort
/****
希尔排序
gap 3
	i = 3
		j = 0
			arr[j] > arr[j+gap]
 */
func Shell(arr []int,len int) []int  {
	var gap,i,j,temp int
	for gap=len/2; gap>0; gap /=2 { /* 设置排序的步长，步长gap每次减半，直到减到*/
		for i=gap; i<len; i++ {   /* 定位到每一个元素 */
			for j = i-gap; (j >= 0) && (arr[j] > arr[j+gap]) ; j -= gap { /* 比较相距gap远的两个元素的大小，根据排序方向决定如何调换 */
				temp = arr[j]
				arr[j] = arr[j+gap]
				arr[j+gap] = temp
			}
		}
	}
	return arr;
}
