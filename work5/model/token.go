package model

import "github.com/dgrijalva/jwt-go"

//Claims 用于生成token的接口,只包含账户的账号
type Claims struct {
	ID int64
	Account string
	jwt.StandardClaims
}
var JwtKey = "www.unknown.com"//生成token的key值