package controller

import (
	"log"
	"net/http"
	"simple-demo/helper"
	"simple-demo/model"
	"simple-demo/service"
	"strconv"
	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserLoginResponse struct {
	Response
	UserId   int64  `json:"user_id,omitempty"`
	Token    string `json:"token"`
	Username string `json:"username"`
}

type UserResponse struct {
	Response
	User User `json:"user"`

}

// type UserRegisterService struct {
// 	ctx context.Context
// }

func Register(c *gin.Context) {
	username := c.Query("username")
	password := helper.GetMd5(c.Query("password"))
 
	userInfo := model.User{UserName: username, Password: password} 
	token, _ := helper.GenerateToken(userInfo.UserName, int(userInfo.UserID))
 
	_, err := service.GetUserByName(userInfo.UserName)
	if err == nil {
	   c.JSON(http.StatusOK, UserLoginResponse{
		  Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
	   })
	} else {
		userInfo := model.User{UserName: username, Password: password, Token: token}
		userid, _ := service.CreateUser(&userInfo)

		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userid,
			Token:    token,
			Username: username,
		})
	}
 }

func Login(c *gin.Context) {
	username := c.Query("username")
	password := helper.GetMd5(c.Query("password"))

	userInfo := model.User{UserName: username, Password: password}

	userid, RowsAffected := service.UserLogin(userInfo)

	token, _ := helper.GenerateToken(userInfo.UserName, int(userid))

	if RowsAffected != 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "User exist"},
			UserId:   userid,
			Token:    token,
			Username: username,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	userid := c.Query("user_id")
	fanid := c.Query("fan_id")
	token := c.Query("token")

	UID, _ := strconv.ParseUint(userid, 10, 32)
	FID, _ := strconv.ParseUint(fanid, 10, 32)

	fanInfo := model.User{UserID: uint(FID)}
	userInfo, _ := service.GetUserByID(uint(UID))
	followercount, _ := service.GetFanCount(uint(UID))
	followcount, _ := service.GetFollowCount(uint(UID))
	isfollow, _ := service.IsFollow(fanInfo.UserID, userInfo.UserID)
	log.Println(UID)
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
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,

		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
