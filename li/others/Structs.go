package main

import "fmt"

type Person struct {
	Name    string
	Gender  string
	Age     uint8
	Address string
}

func (per *Person) Move(address string) string {
	old := per.Address
	per.Address = address
	//此時per為指針類型[地址]  地址和外部定義的地方是一致的
	fmt.Printf("%p",per)
	fmt.Println()
	//這個位置是存放per指針的位置
	fmt.Printf("%p",&per)
	fmt.Println()
	return old
}
func (per Person) Test() {
	//這樣傳遞  是進行值傳遞
	fmt.Printf("%p",per)
	fmt.Println()
	//相當於在這個地址新建了一個per對象，然後在函數結束時，釋放掉了
	fmt.Printf("%p",&per)
	fmt.Println()
	fmt.Println(per)
}

type Animal interface {
	Grow()
	Move(string) string
}
type Cat struct{
	Name string
	Age int
	Address string
}
func (cat *Cat) Grow() {
	fmt.Println("cat grow")
}
func (cat *Cat) Move(address string) string {
	cat.Address = address
	return ""
}

func main() {

	myCat := Cat{"Little C", 2, "In the house"}
	pCat := &myCat  //取地址操作  此時 pCat =>  *Cat 類型
	fmt.Println(pCat)
	animal, ok := interface{}(&myCat).(Animal)  //判斷一個對象是否是某一個接口
	fmt.Printf("%v, %v\n", ok, animal)

	p := Person{"Robert", "Male", 33, "Beijing"}
	p1 := Person{Name:"xxx",Gender:"Male",Age:32,Address:"lalal"}
	//p2 := Person{"Robert", "Male", 33}  //不能部分賦值
	p3 := Person{Name:"xxx",Gender:"Male",Age:32}  //帶key則可以部分初始化
	fmt.Printf("%p",&p)
	fmt.Println()
	oldAddress := p.Move("San Francisco")
	p.Test()
	fmt.Printf("%s moved from %s to %s.\n", p.Name, oldAddress, p.Address)

	fmt.Println(p1)
	fmt.Println(p3)
}
