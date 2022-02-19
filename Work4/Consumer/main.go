package main

import (
	"fmt"
	"net"
)

func main() {
	listen,err := net.Listen("tcp","127.0.0.1:8080")
	if err!=nil{
		fmt.Println("监听失败!",err)
	}
	for{
		conn,err0 := listen.Accept()          //建立连接
		if err0!=nil{
			fmt.Println("链接失败!",err0)
			return                          //建立链接失败则退出
		}
		//验证信息是否已经写完
		var n = make([]byte,2)
		_,err0 = conn.Read(n)
		if err0!=nil || string(n)!="OK"{
			continue
		}
		go Consumer(conn)                  //由于此协程不会停止，所以再次发送时会出现文件移除失败的提示以退出之前的协程，无影响
	}
}

