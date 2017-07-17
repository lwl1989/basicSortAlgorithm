package redis

import (
	"fmt"
	"sync"
)
var oncePool sync.Once
var pool *Pool
type PoolInterface interface{
	GetClient() (Client,err error)
	lock() bool
	unLock() bool
}

type Pool struct {
	MaxConnection uint8  //最大连接数
	InitConnection uint8 //默认连接数
	Number uint8		 //现在连接数
	FreeNum uint8        //可用连接数
	FreeClient chan *Client
}

func GetPool() *Pool {
	oncePool.Do(func() {
		pool = &Pool{8, 2,0,0,make(chan *Client,8)}
		var i uint8

		for  i=0; i<pool.InitConnection; i++ {
			client,_ := NewClient(GetOption())
			pool.FreeClient <- client
			pool.Number += 1
			pool.FreeNum += 1
		}
	})
	return pool
}

func GetClient(pool *Pool) *Client {
	//fmt.Println(pool)
	if(pool.FreeNum < 1 && pool.Number < pool.MaxConnection) {
		client,err := NewClient(GetOption())
		if(err!=nil) {
			fmt.Println(err)
			return nil
		}
		pool.Number += 1
		return client
	}
	client := <- pool.FreeClient
	pool.FreeNum -= 1
	return client
}

func FreeClient(client *Client)  {

	pool := GetPool()
	pool.FreeClient <- client
	pool.FreeNum += 1
}
