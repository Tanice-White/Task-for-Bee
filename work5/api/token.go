package api

import (
	"Bee-work5/dao"
	"Bee-work5/model"
	"Bee-work5/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//CheckAuthorizationBeforeLogIn 在登录前检测是否已有token令牌
func CheckAuthorizationBeforeLogIn(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")   //获取头部Authorization下的内容
	if tokenStr==""{                                //没有token则直接登录
		c.Next()
		return
	}
	_,err := service.ParasToken(tokenStr)          //获取token内的内容
	if err!=nil {                                  //token失效或发生错误则重新登录
		c.Next()
		return
	}
	c.JSON(200,gin.H{"msg":"您已登陆"})        //存在token时,不允许重复登录
	c.Abort()
	return
}

//GenerateToken 在登录后生成token令牌
//返回值 token令牌和错误信息
func GenerateToken(c *gin.Context){
	expireTime := time.Now().Add(3600*time.Second)//有效时间是创建后的1小时
	Claims := &model.Claims{
		ID: 10086,
		Account: c.Query("account"),
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),       //过期时间
			IssuedAt: time.Now().Unix(),        //颁发时间
			Issuer: "127.0.0.1",                //颁发者-->这里写的是本地地址
			Subject: "user token",              //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,Claims)//选择SigningMethodHS256加密方式生成token
	tokenStr,err := token.SignedString([]byte(model.JwtKey))
	if err!=nil{                                             //出现错误则返回错误并且不允许登录
		c.JSON(200,gin.H{
			"err":err,
		})
		c.Abort()
		return
	}
	c.JSON(200,gin.H{
		"token":tokenStr,
	})
}

//CheckAuthorization 获取token并且初始化用户信息
func CheckAuthorization(c *gin.Context){
	tokenStr := c.GetHeader("Authorization")   //获取头部Authorization下的内容
	if tokenStr==""{
		c.JSON(http.StatusUnauthorized,gin.H{
			"msg":"请先登录",
		})
		c.Abort()
		return
	}
	Claims,err := service.ParasToken(tokenStr)    //获取token内的内容
	if err!=nil {
		c.JSON(200,gin.H{
			"err":err.Error(),
		})
		c.Abort()
		return
	}
	flag,err:=dao.FindByAccount(Claims.Account)
	if err!=nil{
		c.JSON(200,gin.H{
			"err":err,
		})
		c.Abort()
		return
	}
	if !flag {
		c.JSON(200,gin.H{
			"err":"解析错误",
		})
		c.Abort()
		return
	}
	err = service.SetAboutCourseWhenToken()
	if err!=nil{
		c.JSON(200,gin.H{"err":err.Error()})
	}
	c.Next()
	return
}