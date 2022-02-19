package main

import (
	"Bee-work4/Tool"
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

func Consumer(conn net.Conn) {
	defer conn.Close()
	for {
		//开始处理文件内容
		_,_err := os.Stat("qm1.qm")
		if _err!=nil{                                      //如果文件不存在则间隔一秒再直接进行下一次扫描
			time.Sleep(time.Second)
			continue
		}
		file,_err := os.OpenFile("qm1.qm",os.O_RDWR,0666)
		if _err!=nil{
			fmt.Println("文件打开失败: ",_err)
			time.Sleep(time.Second)                        //如果文件打开失败则间隔一秒再直接进行下一次扫描
			continue
		}
		reader := bufio.NewReader(file)
		for {                                              //读取文件所有内容
			strLine,err1 :=reader.ReadString('\n')
			if err1==io.EOF{
				break
			}
			if err1!=nil{
				fmt.Println("文件读取错误: ",err1)
				break
			}
			tempStr,err2 := Tool.Decode(strings.TrimSpace(strLine))//把换行符号去掉后传入解码
			if err2!=nil{
				fmt.Println("解码失败: ",err2)
				break
			}
			byteMsg,flag := Tool.CheckForm(tempStr)
			if !flag{                                        //发现文件传输错误则直接停止读取
				fmt.Println("文件传输错误")
				break
			}
			fmt.Println(byteMsg)
		}
		_ = file.Close()                                     //关闭文件夹
		err := os.Remove("qm1.qm")                     //输出结果后将文件移除
		if err != nil {
			fmt.Println("文件移除失败: ", err)
			return
		}
		time.Sleep(time.Second)
	}
}
