package main;

import (
//	"li/sort"
	"fmt"
	"math"
)

func main() {
//	arr := []int{9,3,5,6,7,8}
//	sort.Shell(arr, len(arr))

/*	selectArr := sort.Select(arr, len(arr))
	for _,value := range selectArr {
		fmt.Printf("%d\r\n", value)
	}*/


//		insertArr := sort.Insert(arr, len(arr))
// 	for key,value := range insertArr {
//		fmt.Printf("%d:%d\r\n", key,value)
//	}

	y:=1.5   //内存地址1  8个字节 64位
	y++      //内存地址1  
	z := math.Ceil(y)   //申请内存地址2  8个字节64位 复制Y的值放入  申请内存3  8个字节64位 将结果放入 z
	//golang的数组是按传值进行传递的  代价很大  减少使用  尽量使用切片
	//通常会改成传递指针的方式进行 好处: 1使参数传递成本降至最低  2变量的生命周期独立于作用域(理解？？？？)
	fmt.Println(z)
/*	a := 37
	pi := &a
	ppi := &pi
	fmt.Println(a, *pi, **ppi)
	**ppi++;
	fmt.Println(a, *pi, **ppi)
	*/

	i :=5
	j :=9
	product :=0
	//引用传递
	swap(&i,&j,&product)
	fmt.Println(i,j,product)
	i,j,product = swap2(i,j) //推荐第二种写法
	fmt.Println(i,j,product)
}

func swap(x,y,product *int) {
	if *x > *y {
		*x,*y = *y,*x
	}
	*product = *x**y
}

func swap2(x,y int)(int,int,int){
	if x>y {
		x,y = y,x
	}
	return x,y,x*y
}