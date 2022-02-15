package message

import (
	"fmt"
	"regexp"
	"strconv"
)

//MemMessage 获取内存使用量
//参数 读取的string语句
//返回值 内存使用量 错误信息
func MemMessage(fileStr string)(float64,error){
	Reg,err := regexp.Compile(`KiB.Mem.*`)            //获取内存信息行
	if err!=nil{
		return -1, err
	}
	StrReg,err := regexp.Compile(`\d{0,6}.used\b`)    //获取已使用内存量字段
	if err!=nil{
		return -1, err
	}
	MsgReg,err := regexp.Compile(`\d{0,6}`)           //获取信息
	if err!=nil{
		return -1, err
	}
	MemLin := Reg.FindString(fileStr)                       //从string信息中找到包含内存信息的句子
	if MemLin==""{
		return -1,nil
	}
	MemStr := StrReg.FindString(MemLin)                     //从包含信息的句子中找到含其他字段的信息
	if MemLin==""{
		return -1,nil
	}
	MemMsg := MsgReg.FindString(MemStr)                     //从字段中提取数字信息
	intMem,err := strconv.ParseFloat(MemMsg,10)
	if err!=nil{
		return -1, err
	}
	tempMem := intMem/1024                                  //从kb转为Mb
	Mem,err := strconv.ParseFloat(fmt.Sprintf("%.1f",tempMem),64)
	if err!=nil{
		return -1, err
	}
	return Mem,nil
}
