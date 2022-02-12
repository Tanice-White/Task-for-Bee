package api

import "github.com/gin-gonic/gin"

func GetHelp(c *gin.Context){
	c.JSON(200,gin.H{
		"注册账号":"127.0.0.1:8080/register",
		"登录账号":"127.0.0.1:8080/login",
		"更改密码":"127.0.0.1:8080/user/change",
		"选择课程":"127.0.0.1:8080/user/choose",
		"删除课程":"127.0.0.1:8080/user/delete",
	})
}
