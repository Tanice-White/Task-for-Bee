package main

import (
	"Bee-work5/api"
	"Bee-work5/dao"
	"fmt"
)

func main(){
	err := dao.LinkMysql()                           //链接数据库
	if err!=nil{
		fmt.Println("Link Mysql error occurred: ",err)
		return
	}else{
		//添加了默认的六门课程
		//_ = dao.AddCourse(model.C1)
		//_ = dao.AddCourse(model.C2)
		//_ = dao.AddCourse(model.C3)
		//_ = dao.AddCourse(model.C4)
		//_ = dao.AddCourse(model.C5)
		//_ = dao.AddCourse(model.C6)

		api.StartEngine()
	}
}

/*
CREATE TABLE `user` (
`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
`name` VARCHAR(60) NOT NULL,
`account`VARCHAR(60)NOT NULL,
`password` VARCHAR(60) NOT NULL,
`credit` BIGINT(20) DEFAULT 0,
`sum_chosen` BIGINT(20) DEFAULT 0,
PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `user_course` (
`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
`name` VARCHAR(60) NOT NULL,
`account`VARCHAR(60)NOT NULL,
`class_1` BIGINT(20) DEFAULT -1,
`class_2` BIGINT(20) DEFAULT -1,
`class_3` BIGINT(20) DEFAULT -1,
`class_4` BIGINT(20) DEFAULT -1,
`class_5` BIGINT(20) DEFAULT -1,
PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `course` (
`id` BIGINT(20) NOT NULL AUTO_INCREMENT,
`name` VARCHAR(60) NOT NULL,
`credit` BIGINT(20) NOT NULL,
`max_number` BIGINT(20) NOT NULL,
`number` BIGINT(20) DEFAULT 0,
PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
 */