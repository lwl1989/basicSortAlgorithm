package main

import "fmt"

// & 取地址  * 取值(在類型之前)
type Pet interface {
	Name() string
	Value(name string)
}
type Dog struct {
	name string
	//Name string  如果同樣爲大寫Name 在實現接口的時候會報錯
}
func (dog Dog) Name() string {
	return dog.name
}
func (dog *Dog) Value(name string) {
	dog.name = name
	fmt.Println("in func:")
	fmt.Println(dog)
}

//鏈式寫法
type Link interface {
	NameLink() Dog
	ValueLink(name string) Dog
}

func (dog Dog) NameLink() Dog {
	return dog
}

func (dog Dog) ValueLink(name string) Dog {
	dog.name = name
	fmt.Println("in func:")
	fmt.Println(dog)
	return dog
}

func main(){
	dog := Dog{"TTT"}
	_, interfa := interface{}(&dog).(Pet)  //true 指針類型爲實現
	_, interfa1 := interface{}(dog).(Pet)  //基礎類型不認爲是實現

	_, interfa2 := interface{}(&dog).(Link)  //基礎類型也可以認爲是實現
	_, interfa3 := interface{}(dog).(Link)  //基礎類型也可以認爲是實現

	// =====>>> 一个指针类型拥有它以及她的基底类型为接受类型的所有方法 比如 2、 3
	// =====>>> 而它的基底类型则只拥有它本身为接受类型的方法 比如 1是  2不是 _, interfa1 := interface{}(dog).(Pet)  false
	fmt.Println(interfa, interfa1)
	fmt.Println(interfa2, interfa3)


	dog.Value("FFFF")

	fmt.Println("in main:")
	fmt.Println(dog)
}

