package service

import (
	"Bee-work5/dao"
	"Bee-work5/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

//ParasToken 识别token令牌中的账号初始化用户信息到TheUser中
//返回值 错误和Claims结构体中的账号信息
func ParasToken(tokenStr string) (*model.Claims,error) {
	tokenClaims,err := jwt.ParseWithClaims(tokenStr,&model.Claims{},func(token *jwt.Token)(interface{},error){
		return []byte(model.JwtKey),nil
	})
	if err!=nil{
		return nil, err
	}
	if tokenClaims!=nil{                                                                      //代表收到了相关的数据
		if Claims,ok := tokenClaims.Claims.(*model.Claims); ok && tokenClaims.Valid {
			return Claims, nil
		}
	}
	return nil, errors.New("未知错误")
}

func SetAboutCourseWhenToken() error {
	err := dao.SetCourseChosen()                                                                //初始化选课的内容
	return err
}