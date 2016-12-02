package sort


func Heap(arr []int, len int) []int  {
	for o:=len/2; o>=0; o=o-1{
		defer func(arr []int, start int,end int) {
			temp := arr[start]

			for i := 2*start + 1; i<=end; i*=2 {
			if(i<end && arr[i]<arr [i+1]){ //左右孩子的比较
				i=i+1
			}
			if(temp > arr[i]){//左右孩子中获胜者与父亲的比较
				break;
			}
			//将孩子结点上位，则以孩子结点的位置进行下一轮的筛选
			arr[start] = arr[i]
			start = i
		}

		arr[start] = temp
		}(arr,o,len)

	}

	/*for i:=0;i<len;i++ {
		max := i;
		for j:=0;j<len-i-1;j++ {
			if(arr[max] < arr[j]) {
				max = j;
			}
		}
		for j:=len-i-1;j>0;j++ {

		}

	}*/
	return arr
}

func heapCreate(arr []int, start int,end int){
	temp := arr[start]

	for i:= 2*start+1;i<=end;i*=2 {
		if(i<end && arr[i]<arr [i+1]){ //左右孩子的比较
		//	++i;//i为较大的记录的下标
			i=i+1
		}
		if(temp > arr[i]){//左右孩子中获胜者与父亲的比较
			break;
		}
		//将孩子结点上位，则以孩子结点的位置进行下一轮的筛选
		arr[start] = arr[i]
		start = i
	}

	arr[start] = temp
}

