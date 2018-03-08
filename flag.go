package main

import (
	"basicSortAlgorithm/li/interfaces"
	"basicSortAlgorithm/li/redis"
	"fmt"
)


type fff struct {
	Name string
}

func Init() *fff {
	return &fff{"的撒多"}
}
func main() {
	fff1 := Init()
	fmt.Println(fff1.Name)
	fmt.Println(fff1)
	fmt.Println(&fff1)
	return





	option := redis.GetOption()
	option.Port=12345
	option.Host="54.222.155.203"
	pool := redis.GetPool()
	go redis.GetClient(pool).DoCommand("select","4")  //+OK
	go redis.GetClient(pool).DoCommand("get","badge:12")  //66
	go redis.GetClient(pool).DoCommand("get","badge:103") //9
	go redis.GetClient(pool).DoCommand("get","badge:107") //34
	go redis.GetClient(pool).DoCommand("get","badge:108") //24
	go redis.GetClient(pool).DoCommand("get","badge:113") //1
	go redis.GetClient(pool).DoCommand("get","badge:114") //0
	go redis.GetClient(pool).DoCommand("get","badge:115") //0
	go redis.GetClient(pool).DoCommand("get","badge:117") //0
	go redis.GetClient(pool).DoCommand("get","badge:12")  //66
	go redis.GetClient(pool).DoCommand("get","badge:12")  //66
	go redis.GetClient(pool).DoCommand("get","badge:12")  //66

	//fmt.Println(pool.Number)


	interfaces.Test()
	//time.Sleep(3*time.Second)
}
