package service

import (
	"Bee-work5/dao"
	"Bee-work5/model"
	"Bee-work5/tool"
	"errors"
)

//CanRegister
//参数 注册者的账号和密码以tempUser结构体传输
//返回值 错误本身以及是否能够注册的bool值
func CanRegister(tempU model.TempUser)(bool,error) {                                        //要求不含关键字并且账号为10位，密码不小于6位
	if tool.CanUse(tempU.Account)&&tool.CanUse(tempU.Password)&&tool.CanUse(tempU.Name)&&len(tempU.Account)==10&&len(tempU.Password)>=6{
		flag,err := dao.FindByAccount(tempU.Account)                                        //是否能找到重复账户
		if err!=nil{
			return false,err
		}
		if flag {                                                                           //flag==true表示找到了重复账户
			return false,nil
		}
		err = dao.WriteTempUser(tempU)                                                      //将信息写入数据库
		if err!=nil{                                                                        //出错，不允许注册
			return false,err
		}
		model.Sum++                                                                         //同步人数
		return true,nil
	}else{
		err := errors.New("输入不合法")
		return false,err
	}
}
//HasLogIn 判断是否登录成功
//参数 tempUser结构体内的信息
//返回值 错误
func HasLogIn(account,password string) error {
	if tool.CanUse(account)&&tool.CanUse(password) {         //判断输入是否合法
		flag, err := dao.FindByAccount(account)              //找到此用户的数据并且调用
		if err!=nil{
			return err
		}
		if flag && model.TheUser.Password==password {        //flag==true表示找到了账户
			err = dao.SetCourseChosen()                      //初始化选课的内容
			return nil
		}
		return errors.New("密码错误")                     //没有出现错误，没有找到用户
	}else{
		return errors.New("输入不合法")
	}
}

//CanChangePassword 重置密码是否成功
//参数 更改的新密码
//返回值 错误信息
func CanChangePassword(newPassword string) error {
	if model.TheUser.Account==""{
		return errors.New("登录状态异常,请尝试重新登录")
	}
	if len(newPassword)<6{
		return errors.New("密码长度不足")
	}
	err := dao.ChangePassword(model.TheUser.Account,newPassword)
	if err!=nil{
		return err
	}
	model.TheUser.Password = newPassword
	return nil
}