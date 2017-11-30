package sort


/**
所谓的“选择”就是在待排序列里，
找出一个最大(小)的元素，
然后将它放在序列某个位置，
这就完成了一次选择过程。
如果将这样的选择循环继续下去，
就是我们所说的选择排序。
这也是选择排序的精髓。
平均效率 lgn
 */
func Select(arr []int,len int) []int{
	var i,max int;
	index := len-1
	for i = 0; i < len-1; i ++ {
		max = i
		for j := i; j < len; j ++ {
			if arr[max]<arr[j] {
				index = j
			}
		}
		if(max != i) {
			arr[index] = arr[i];
			arr[i] = max;
		}
	}
	return arr
}
