package interfaces

import "fmt"

type Animal interface {
	Grow()
	Move(string) string
}

type Cat struct {
	Name string
	Age uint32
	Address string
}

func (cat *Cat)  Move(new string) string{
	old:=cat.Address
	cat.Address = new
	return old
}

func (cat *Cat) Grow()  {
	fmt.Println("grow")
}

func modify(cat *Cat)  {
	cat.Name = "fff";
}

func Test() {
	return
	// struct Name  声明出来的是一个Cat类型
	cat := Cat{"zhangsan", 2, "555555"}
	animal,ok := interface{}(&cat).(Animal)
	cat.Move("44444")
	fmt.Println(cat)
	modify(&cat)
	fmt.Println(cat)
	fmt.Println(animal, ok)

	// new(struct) 声明出来是一个指针类型
	cat1 := new(Cat)
	cat1.Name="3"
	cat1.Address="3"
	cat1.Age=3
	fmt.Println(cat1)
	modify(cat1)
	fmt.Println(cat1)
}
