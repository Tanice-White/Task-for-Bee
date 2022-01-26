package Agreement

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

//Encode :将数据打包防止TCP黏包(头部占4个字节)
//参数 :被json序列化的数据信息
//返回值 :1.写入完成的信息 2.发生的错误
func Encode(message []byte)([]byte,error){
	length := int32(len(message))                      //获取信息的长度
	var pkg = new(bytes.Buffer)                        //准备需要写入的包
	err := binary.Write(pkg,binary.LittleEndian,length)//以二进制写入文件
	if err!=nil{
		fmt.Println("文件封装错误:",err)
		return nil, err
	}
	err = binary.Write(pkg,binary.LittleEndian,message)//将有效信息写入文件
	if err!=nil{
		fmt.Println("文件封装错误:",err)
		return nil, err
	}
	return pkg.Bytes(), err
}

// Decode :解码消息,将被Encode()函数打包的信息读取出来
//参数 :Reader类型的指针
//返回值 :1.解码后的信息 2.发生的错误
func Decode(reader *bufio.Reader) (string, error) {
	byteLength,_ := reader.Peek(4)                         //提取出前四个字节的信息
	buffLength := bytes.NewBuffer(byteLength)                 //转化为可读的形式
	var length int32                                          //记录信息的长度
	err := binary.Read(buffLength,binary.LittleEndian,&length)//将长度读取出来
	if err!=nil{
		return "", err
	}
	if int32(reader.Buffered())<length+4{                    //判断刻度长度是否正常
		return "", err
	}
	var message = make([]byte,length+4)                       //判断无误则准备空间读取信息(只能一次性全部读取)
	_, err= reader.Read(message)
	if err!=nil{
		fmt.Println("信息读取错误: ",err)
		return "", err
	}
	return string(message[4:]), err
}
