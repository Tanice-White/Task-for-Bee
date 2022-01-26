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

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080") //和服务器建立连接
	if err != nil {
		fmt.Println("链接错误: ", err)
		return
	}
	defer conn.Close()                                             // 关闭连接
	for i:=0;i<1;i++{                                              //只请求一次
		//获取输入(这里直接简化了)
		var Req = global.GetUserReq{
			UserID: 1,
		}
		request,_:=json.Marshal(Req)
		byteMessage,err := Agreement.Encode(request)
		if err!=nil{
			fmt.Println("包的封装错误: ",err)
			break
		}
		_,err = conn.Write(byteMessage)                           //发送数据
		if err!=nil{
			fmt.Println("发送出错了: ",err)
			break
		}

		//接收信息
		reader := bufio.NewReader(conn)                            //准备包接受数据
		strMessage, err := Agreement.Decode(reader)                //将文件解码
		if err == io.EOF {
			return
		}
		var message global.GetUserResp                             //接收到的信息
		_ = json.Unmarshal([]byte(strMessage), &message)
		fmt.Printf("收到数据: %#v",message)
	}
}
