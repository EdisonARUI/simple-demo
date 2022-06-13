package controller

import (
	"fmt"
	"log"
	"simple-demo/helper"
	"simple-demo/model"
	"simple-demo/service"
	"testing"
)

func TestUserInfo(t *testing.T) {

	//token := c.Query("token")
	UserName := model.User{UserName: "zhaorui"}

	userinfo := model.User{UserID: 12, Password: "86985e105f79b95d6bc918fb45ec7727"}

	faninfo := model.User{UserID: 13, Password: "87a38998227cbbc23dcad51cd7f76ab2"}
	token, _ := helper.GenerateToken(UserName.UserName, userinfo.Password)

	UID := userinfo.UserID
	FID := faninfo.UserID

	fanInfo := model.User{UserID: FID}
	userInfo, _ := service.GetUserByID(UID)
	followercount, _ := service.GetFanCount(UID)
	followcount, _ := service.GetFollowCount(UID)
	isfollow, _ := service.IsFollow(fanInfo.UserID, userInfo.UserID)
	log.Println(followercount)
	log.Println(followcount)
	log.Println(isfollow)
	var usersLoginInfo = map[string]User{
		token: {
			Id:            int64(UID),
			Name:          userInfo.UserName,
			FollowCount:   followcount,
			FollowerCount: followercount,
			IsFollow:      isfollow,
		},
	}
	if user, exist := usersLoginInfo[token]; exist {
		fmt.Println(user)
	} else {
		fmt.Println("User doesn't exist")
	}
}
