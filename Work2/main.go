package main

import (
	"Bee-work2/message"
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	fileStr string                                     //储存从文件中读取出的string信息
	CPUf float64                                       //储存单个CPU利用率信息
	Mem1 float64                                       //储存单个的内存使用量信息
	CPU []float64                                      //储存CPU使用率的数组
	Mem []float64                                      //储存内存使用率的数组
)

func main(){
	file,err := os.Open("parameter")
	if err!=nil{
		fmt.Println("an error occurred when opening the file: ",err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for{
		fileStr,err =reader.ReadString('\n')
		if err==io.EOF{                                //读取完成则退出
			break
		}
		if err!=nil{                                   //出现错误
			fmt.Println("an error occurred when reading the file: ",err)
			break
		}
		//获取CPU使用率
		CPUf,err = message.CPUMessage(fileStr)             //获取CPU利用率的信息
		if err!=nil{
			fmt.Println("an error occurred when getting CPU message: ",err)
			return
		}
		if CPUf>=0{
			CPU = append(CPU,CPUf)                     //将信息加入数组
		}
		//获取内存使用量
		Mem1,err = message.MemMessage(fileStr)
		if err!=nil{
			fmt.Println("an error occurred when getting Memory message: ",err)
			return
		}
		if Mem1>=0{
			Mem = append(Mem,Mem1)                     //将信息加入数组
		}
	}
	//方便绘制表格
	fmt.Println("CPU使用率")
	for i := 0; i < len(CPU); i++ {
		fmt.Println(CPU[i],"%")
	}
	fmt.Println("内存占用量")
	for i := 0; i < len(Mem); i++ {
		fmt.Println(Mem[i])
	}

}
