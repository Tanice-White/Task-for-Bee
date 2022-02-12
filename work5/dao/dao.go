package dao

import (
	"Bee-work5/model"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB               			     //启动数据库的指针

//LinkMysql 创建和数据库的链接
//返回值 错误本身
func LinkMysql()(err error){
	dsn:="root:@tcp(127.0.0.1:3306)/text_for_bee"
	DB,err = sql.Open("mysql",dsn)
	if err!=nil {
		return err
	}
	err = DB.Ping()
	if err!=nil {
		return err
	}
	r,_ := DB.Query("select `id` from user")
	if err!=nil{
		return err
	}
	for r.Next() {
		_ = r.Scan(&model.Sum)               //获取总人数，初始化Sum
	}
	_ = r.Close()
	return err
}
