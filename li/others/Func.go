package main

import (
	"sync/atomic"
	"strconv"
	"fmt"
)

//定義函數類型  员工生成器
type EmployeeIdGenerator func(company string, department string, sn uint32) string

var company = "Gophers"
var sn uint32

//函数类型  当参数传入
func generateId(generator EmployeeIdGenerator, department string) (string, bool) {
	//atomic包用于保证原子性
	newSn := atomic.AddUint32(&sn, 1)
	return generator(company, department, newSn),true
}

func generateId1(generator EmployeeIdGenerator, department string) (string, bool) {
	//atomic包用于保证原子性
	newSn := atomic.AddUint32(&sn, 1)
	return generator(company, department, newSn),true
}


func main() {

	//类似PHP匿名回调的实现方式
	var generator EmployeeIdGenerator = func(company string, department string, sn uint32) string {
		return company + department + strconv.FormatUint(uint64(sn), 10)
	}

	fmt.Println(generateId(generator, "SS"))
	fmt.Println(generateId(generator, "SS"))
	fmt.Println(generateId(generator, "SS"))
	fmt.Println(generateId(generator, "SS"))

	//可以根据参数进行不同的匿名实现
	var generator1 EmployeeIdGenerator = func(c string, d string, sn uint32) string {
		return "fff" +strconv.FormatUint(uint64(sn), 10)
	}

	fmt.Println(generateId1(generator1,"dd"))


}










