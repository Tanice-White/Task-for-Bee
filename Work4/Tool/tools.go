package Tool

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"strconv"
)

//GiveForm 进行奇校验的封装(一次处理8bit)
//参数 需要校验的信息
//返回值 校验完成的string类型的二进制信息
func GiveForm(message []byte)string{
	var allTemp string                                     //储存一整行的二进制信息
	for i := 0; i < len(message); i++ {
		temp:= strconv.FormatInt(int64(message[i]),2) //将信息转为二进制(占8位)
		allTemp += temp
	}
	count := 0                                              //count用于计算1的个数
	finalMsg := []byte(allTemp)                             //转为byte数组
	for _,v := range finalMsg {                             //对1的个数进行总统计
		if v=='1' {
			count++
		}
	}
	json.Marshal(message)
	if count%2==0{
		finalMsg = append(finalMsg,'1')                //如果有偶数个1则在最后加一个1
	}else{
		finalMsg = append(finalMsg,'0')                //如果有奇数个1则在最后加一个0
	}
	return string(finalMsg)
}

//CheckForm 进行奇校验
//参数 读取出的所有string类型的数据
//返回值 1.真实的信息 2.是否符合规则的bool值
func CheckForm(message []byte)([]byte,bool){
	count := 0                                                    //count用于记录1的个数
	for _,v := range message {
		if v== 49 {                                               //记录1的个数
			count++
		}
	}
	if count%2==0 {
		return nil,false                                          //是偶数代表传输错误
	}
	finalMessage := message[:len(message)-1]                      //顺便把最后一位去掉
	return finalMessage,true
}

//Encode :包的封装,将所发信息的字节大小写入文件头
//参数 :所发信息(string类型)
//返回值 :1.更改后的信息 2.错误类型数据
func Encode(message []byte)([]byte,error){
	length := int64(len(message))                                //读取信息的大小(8个字节),使用binary包时不要使用int
	var pkg = new(bytes.Buffer)
	err := binary.Write(pkg, binary.LittleEndian, length)        //以二进制写文件,此时写入消息头(包的大小)
	if err != nil {
		return nil, err
	}
	err = binary.Write(pkg, binary.LittleEndian, message)//增加需要传输的信息(二进制的形式)
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode :解码消息,将被Encode()函数打包的信息读取出来
//参数 :string类型的数据
//返回值 :1.解码后的信息 2.发生的错误
func Decode(line string) ([]byte, error) {
	temp := []byte(line)
	lengthByte := temp[:8]                                    // 读取前8个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)                 //将其转化为可读的形式
	var length int64
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)//将前8位以编码形式读取得到文件大小
	if err != nil {
			return nil, err
		}
	if int64(len(line)) < length+8 {
			return nil, errors.New("文件出错了")
		}
	return temp[8:], nil                                      //8个字节后的数据才是需要的信息
}