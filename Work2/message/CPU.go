package message

import (
	"regexp"
	"strconv"
)

//CPUMessage 获取CPU使用率的数据
//参数 读取的string数据
//返回值 读取出的数据 错误信息
func CPUMessage(fileStr string) (float64,error) {
	CPUReg,err := regexp.Compile(`\d\d\.\d.id\b`)   //获取CPU的未利用率包括后方的id字段
	if err!=nil{
		return -1,err
	}
	unCPUReg,err := regexp.Compile(`\d\d\.\d`)      //找到所有符合利用率的字段
	if err!=nil{
		return -1,err
	}
	CPUStr := CPUReg.FindString(fileStr)                  //从文件中的句子找到 CPUReg 相关字段
	unCPUStr := unCPUReg.FindString(CPUStr)               //从 CPUReg相关字段找到为利用率的字段
	if unCPUStr==""{                                      //数据不存在则返回负数
		return -1,nil
	}
	unCPUMsg,err := strconv.ParseFloat(unCPUStr,10)//将信string信息化为float64
	if err!=nil{
		return -1,err
	}
	return (1000-unCPUMsg*10)/10,nil                      //防止精度缺失
}