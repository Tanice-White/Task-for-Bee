package api

import (
	"Bee-work5/model"
	"Bee-work5/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToRegister(c *gin.Context){
	var tempU model.TempUser
	err := c.ShouldBind(&tempU)
	if err!=nil {                      //获取输入失败的错误处理
		c.JSON(200,gin.H{
			"err":err.Error(),
		})
		return
	}
	flag,err := service.CanRegister(tempU)   //判断输入是否合法,账号是否重合,并且写入数据库
	if err!=nil{
		c.JSON(200, gin.H{
			"code":500,
			"err": err.Error()})
		return
	}else{
		if flag {                           //flag==true代表可以注册
			c.JSON(200,gin.H{
				"msg":"注册成功",
			})
			return
		}
		c.JSON(200,gin.H{
			"err":"账号重复",
		})
		return
	}
}

func ToLogIn(c *gin.Context) {
	if model.Sum==0{
		c.JSON(http.StatusOK,gin.H{
			"msg":"还未有注册用户",
		})
		c.Abort()
		return
	}
	account := c.DefaultQuery("account","")
	password := c.DefaultQuery("password","")
	if account=="" || password=="" {                      //获取输入失败的错误处理
		c.JSON(http.StatusBadRequest,gin.H{
			"err":"获取失败",
		})
		c.Abort()                       //不允许颁发token令牌
		return
	}
	err := service.HasLogIn(account,password)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"err":err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"msg":"登录成功",
	})
}

func ToChangePassword(c *gin.Context){
	newPassword := c.Query("password")
	if newPassword==model.TheUser.Password{
		c.JSON(200,gin.H{"msg":"无效更改"})
		return
	}
	if newPassword!=""{
		err := service.CanChangePassword(newPassword)
		if err!=nil{
			c.JSON(200,gin.H{"err":err.Error()})
			return
		}
		c.JSON(200,gin.H{"msg":"更改成功,请重新登录"})
		return
	}
	c.JSON(200,gin.H{"msg":"密码不能为空"})
	return
}
