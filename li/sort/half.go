package sort


/**
	思想：
		折半查找  假设有序（当数组只有一个元素的话，他肯定是有序 且 low = mid = high）
		1.将要排序的第一个元素保存到 比较元素
		2.获取已有序数组的长度
		3.已有序数组的中间位置
		4.递归查找  比较元素和 已有序数组中间元素的大小
		5.如果大则放左边（右边） 小则放右边（左边）
 */
func Half(arr []int, len int) []int {
	var i,j,temp int
	for i=1;i<len;i++ {
		temp = arr[i]  /* 保存但前元素 */
		low := 0
		high := i-1
		for ; low<=high; {        /* 在a[low...high]中折半查找有序插入的位置 */

			mid := (low+high)/2   /* 找到中间元素 */

			if arr[mid]<temp {  /* 如果中间元素比但前元素大，当前元素要插入到中间元素的左侧 */
				high = mid - 1
			}else {              /* 如果中间元素比当前元素小，但前元素要插入到中间元素的右侧 */
				low = mid + 1
			}
		}   /* 找到当前元素的位置，在low和high之间 */
		for j=i-1; j > high ; j-- {
			arr[j+1] = arr[j]
		}

		arr[high+1] = temp
		/*for _,num := range arr {
			fmt.Printf("%d",  num)
		}
		fmt.Println()*/
	}
	return arr
}
