package main

import (
	"Beens-work3/Agreement"
	"Beens-work3/global"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func Server(conn net.Conn)  {
	defer conn.Close()                             //关闭链接
	for{
		reader := bufio.NewReader(conn)            //准备接收
		strMessage, err := Agreement.Decode(reader)//将文件解码
		if err == io.EOF {
			return
		}
		var message global.GetUserReq              //接收到的信息
		_ = json.Unmarshal([]byte(strMessage), &message)
		fmt.Println("收到请求 :",message)
		//从数据库中寻找答案(简化)
		var resp = global.GetUserResp{
			UserID: message.UserID,
			UserName: "Unknown~~",
		}
		byteResp,_ :=json.Marshal(resp)
		byteMessage, err := Agreement.Encode(byteResp)
		_,_ = conn.Write(byteMessage)
	}
}


func main() {
	listen,err := net.Listen("tcp","127.0.0.1:8080")
	if err!=nil{
		fmt.Println("监听失败!",err)
	}
	for{
		conn,err := listen.Accept()       //建立连接
		if err!=nil{
			fmt.Println("链接失败!",err)
			continue                      //一次建立失败则进行下一次,直到建立成功
		}
		go Server(conn)                   //链接成功后则启用协程办理事务
	}
}
