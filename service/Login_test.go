package service

import (
	"fmt"
	"simple-demo/model"
	"testing"
)

func TestLogin(t *testing.T) {
	model.Init()
	username := "zhaorui"
	paseeword := "734056edd0b94138498f5e153c21d6d7"
	userInfo := model.User{UserName: "zhaorui", Password: "734056edd0b94138498f5e153c21d6d7"}
	find := model.DB.Where("username = ? AND password = ? ", username, paseeword).Find(&userInfo)
	if find.RowsAffected == 0 || paseeword != userInfo.Password {
		//if  {
		fmt.Println("用户名或密码错误")
		fmt.Println(find.Error)
		fmt.Println(find.RowsAffected)
		//fmt.Println(int64(userInfo.UserID))
		return
		//}
	} else {
		fmt.Println("用户名和密码正确")
		//fmt.Println(err)
		fmt.Println(int64(userInfo.UserID))
		fmt.Println(find.Error)
		fmt.Println(find.RowsAffected)
	}
}
