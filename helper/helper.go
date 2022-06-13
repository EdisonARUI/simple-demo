package helper

import (
	"crypto/md5"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaims struct {
	UserName     string `json:"userName"`
	UserPassword string `json:"userPassword"`
	jwtgo.StandardClaims
}

var jwtkey = []byte("JjUhqZteNUhtDQfvXH9uCHhdKDmUDyAm")

// GetMd5
// 生成 md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// GenerateToken
// 生成 token
func GenerateToken(username string, userpassword string) (string, error) {
	expireTime := time.Now().Add(1 * time.Hour) //过期时间
	nowTime := time.Now()                       //当前时间
	UserClaim := &UserClaims{
		UserName:     username,
		UserPassword: userpassword,
		StandardClaims: jwtgo.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间戳
			IssuedAt:  nowTime.Unix(),    //当前时间戳
			Issuer:    "Douyin",          //颁发者签名
			Subject:   "userToken",       //签名主题
		},
	}
	tokenStruct := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, UserClaim)
	tokenString, err := tokenStruct.SignedString(jwtkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
