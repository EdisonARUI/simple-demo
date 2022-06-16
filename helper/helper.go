package helper

import (
	"crypto/md5"
	"fmt"
	"strconv"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaims struct {
	UserName     string `json:"userName"`
	UserID 		 string `json:"userID"`
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
func GenerateToken(username string, userID int) (string, error) {
	expireTime := time.Now().Add(1 * time.Hour) //过期时间
	nowTime := time.Now()                       //当前时间
	UserClaim := &UserClaims{
		UserName:     username,
		UserID: 	  strconv.Itoa(userID),
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

// AnalyseToken
// 解析 token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwtgo.ParseWithClaims(tokenString, userClaim, func(token *jwtgo.Token) (interface{}, error) {
		return jwtkey, nil
	})
	if err != nil {
		return userClaim, err
	}
	if !claims.Valid {
		return userClaim, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

func GetUserIDByToken(tokenString string) (int, error) {
	claims, err := AnalyseToken(tokenString)
	if err != nil{
		return 0, err
	} else {
		userID, _ := strconv.Atoi(claims.UserID)
		return userID, nil
	}
}