package sort


/**
	思想：
		从第一个数字开始进行排序，每次将新进入的数字和之前的数据进行排序 【因为之前的数据一直是有序的，所以可以进行折半查找】
	7,4,8,2,3,5,1
	i = 2  temp = 4  low = high = mid = 0  7>4    high = 0 low = 1                                 不赋值  j=0 high = 0   7 4 8 2 3 5 1
	i = 2 temp = 8 low =0 high =1  mid = 0  7>8 high = 0 low =0  8> 4 high = -1 low = 0            赋值  j = 1 high = -1   8 7 4 2 3 5 1
        i = 3 temp = 2 low = 0 high = 2 mid = 1  7>   low = mid+1 = 2   mid = 2 low =2 high = 2        不赋值  j = 2 high = 2  ......
 */
func Half(arr []int, len int) []int {
	var i,j,temp,low,high,mid int
	for i=1;i<len;i++ {
		temp = arr[i]  /* 保存但前元素 */
		low = 0
		high = i-1
		for ; low<=high; {        /* 在a[low...high]中折半查找有序插入的位置 */

			mid = (low+high)/2   /* 找到中间元素 */

			if(arr[mid]<temp) {  /* 如果中间元素比但前元素大，当前元素要插入到中间元素的左侧 */
				high = mid - 1
			}else {              /* 如果中间元素比当前元素小，但前元素要插入到中间元素的右侧 */
				low = mid + 1
			}
			println(i,arr[mid], temp, low, high,mid )
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
