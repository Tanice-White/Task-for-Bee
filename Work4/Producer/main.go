package main

import (
	"Bee-work4/Tool"
	"bufio"
	"fmt"
	"net"
	"os"
)

var Task string                                 //用于接收命令语句
var allTask []string                            //生成最后的指令块

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("链接错误: ", err)
		return
	}
	defer conn.Close()
	file,err := os.OpenFile("qm1.qm",os.O_CREATE|os.O_APPEND,0666)
	if err!=nil {
		fmt.Println("文件创建错误:",err)
		return
	}
	defer file.Close()
	fmt.Println("请输入待执行的命令,输入q结束")
	Scanner := bufio.NewScanner(os.Stdin)           //确保能读入空格
	for Scanner.Scan() {
		Task = Scanner.Text()
		if Task=="q"{
			break
		}
		allTask = append(allTask,Task)               //生成代码块
	}
	for i := 0; i < len(allTask); i++ {              //对每一句命令操作
		tempStr := Tool.GiveForm([]byte(allTask[i])) //奇校验的封装
		message,_ := Tool.Encode([]byte(tempStr))    //防黏包协议的封装
		_,_ =file.Write(message)
		_,_ = file.WriteString("\n")              //每封装完一句命令换行
	}
	//发送信息验证
	var message = "OK"
	_,err = conn.Write([]byte(message))
	if err!=nil{
		fmt.Println("状态错误")
		return
	}
}