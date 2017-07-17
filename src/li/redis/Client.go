package redis

import (
	"net"
	"fmt"
	"strconv"
	"io"
	"time"
)

type Client struct {
	Conn net.Conn
	Options *Options
	Expire time.Time
	IsBusy bool
}

func NewClient(options *Options) (*Client,error)  {
	ip := options.Host + ":" + strconv.Itoa(int(options.Port))
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		return nil, err
	}
	client := new(Client)
	client.Conn = conn
	client.Options = options
	return client, nil
}

func (client *Client) DoCommand(commands ...string){
	length := len(commands)
	sendCommand := "*"+strconv.Itoa(length)+"\r\n"
	commandLen := 0
	for _, command := range commands {
		commandLen = len(command)
		sendCommand +="$"+strconv.Itoa(commandLen)+"\r\n"
		sendCommand +=command+"\r\n"
	}
	client.Conn.Write([]byte(sendCommand))
	client.Receive()
	FreeClient(client)
}

func (client *Client) Receive(){
	bufLen := 1024
	buf := make([]byte, bufLen)
	for {
		n, err := client.Conn.Read(buf)
		if err == io.EOF {
			fmt.Println("redis connection lost")
		}
		if err != nil {
			fmt.Println("error message:")
			fmt.Println(err)
			return
		}

		//fmt.Println("redis:", string(buf[0:n]))
		if n<bufLen {
			fmt.Println(string(buf))
			return
		}
	}
}