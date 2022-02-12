package dao

import (
	"Bee-work5/model"
	"database/sql"
)

//WriteTempUser 将注册的账户写入user库和user_course库中(此时不包含任何课程)
//参数 model.User的结构体
//返回值 错误本身
func WriteTempUser(tempU model.TempUser)error{
	tx,err := DB.Begin()
	if err!=nil{
		return err
	}
	sqlStr := "insert into user (name,account,password) values (?,?,?)"
	_,err = tx.Exec(sqlStr,tempU.Name,tempU.Account,tempU.Password)
	if err!=nil{
		_ = tx.Rollback()
		return err
	}
	sqlStr = "insert into user_course (name,account) values (?,?)"
	_,err = tx.Exec(sqlStr,tempU.Name,tempU.Account)
	if err!=nil{
		_ = tx.Rollback()
		return err
	}
	_ = tx.Commit()                                                 //提交事务
	return nil
}

//FindByAccount 通过账号查询用户--->用于判断账户是否可以被注册()
//顺便初始化TheUser的所有信息包含姓名 账号 密码 学分 选课数量
//参数 用户登录的账号
//返回值 错误本身以及是否查询到重复账号的bool值
func FindByAccount(account string)(bool,error){
	sqlStr := "select name,password,credit,sum_chosen from user where account = ?"
	r:= DB.QueryRow(sqlStr,account)
	err := r.Scan(&model.TheUser.Name,&model.TheUser.Password,&model.TheUser.Credit,&model.TheUser.SumChosen)
	if err==sql.ErrNoRows{
		return false,nil                        //此时视为未发生错误，没有查询到重复账号允许注册
	}else if err!=nil{
		return false, err
	}
	model.TheUser.Account = account
	return true,nil                             //找到了相关数据，初始化完成
}

//ChangePassword 根据账户信息从数据库更改密码
//参数 用户的账号和新的密码
//返回值 错误信息
func ChangePassword(account,newPassword string) error {
	sqlStr := "update user set password=? where account=?"
	_,err := DB.Exec(sqlStr,newPassword,account)
	if err!=nil{
		return err
	}
	return nil
}

//SetCourseChosen 初始化 用户的选课总数 CourseChosen的内容
//返回值 错误信息
func SetCourseChosen() error {
	sqlStr := "select class_1,class_2,class_3,class_4,class_5 from user_course where account = ?"
	row := DB.QueryRow(sqlStr,model.TheUser.Account)
	err := row.Scan(&model.CourseChosen[0],&model.CourseChosen[1],&model.CourseChosen[2],&model.CourseChosen[3],&model.CourseChosen[4])
	if err!=nil{
		return err
	}
	return nil
}
