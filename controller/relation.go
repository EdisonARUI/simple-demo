package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simple-demo/helper"
	"simple-demo/model"
	"simple-demo/service"
	"strconv"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction 登录用户对其他用户进行关注或取消关注
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	//userid := c.Query("user_id")
	fanid := c.Query("to_user_id")
	actionType := c.Query("action_type")

	UID, _ := helper.GetUserIDByToken(token)
	//UID, _ := strconv.ParseUint(userid, 10, 32)
	FID, _ := strconv.ParseUint(fanid, 10, 32)

	if token == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "You haven't logged in yet",
		})
	} else {
		if actionType == "1" {
			err := service.FollowUser(int64(FID), int64(UID))
			if err != nil {
				return
			}
			log.Println(UID)
			log.Println(FID)
			log.Println(err)

			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "关注成功"})

		} else if actionType == "2" {
			err := service.UnFollowUser(int64(FID), int64(UID))
			if err != nil {
				return
			}
			log.Println(UID)
			log.Println(FID)
			log.Println(err)

			c.JSON(http.StatusOK, Response{
				StatusCode: 0,
				StatusMsg:  "取关成功",
			})
		}
	}
}

// FollowList 登录用户关注的所有用户列表
func FollowList(c *gin.Context) {
	token := c.Query("token")
	userid := c.Query("user_id")
	UID, _ := strconv.ParseUint(userid, 10, 32)

	if token == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "You haven't logged in yet",
		})
	} else {
		followList_, err := service.GetFollowList(int64(UID))
		followList, err := GenerateFollow(followList_)
		log.Println(followList)
		log.Println(err)
		log.Println(UID)
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Get success"},
			UserList: followList,
		})
	}
}

// FollowerList 登录用户的粉丝列表
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	userid := c.Query("user_id")
	UID, _ := strconv.ParseUint(userid, 10, 32)

	if token == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "You haven't logged in yet",
		})
	} else {
		fanList_, err := service.GetFanList(int64(UID))
		fanList, err := GenerateFollower(fanList_)
		log.Println(fanList)
		log.Println(err)
		log.Println(UID)
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Get success"},
			UserList: fanList,
		})
	}
}

func GenerateFollower(followerList []model.User) ([]User, error) {
	userID := 4
	fancount, _ := service.GetFanCount(uint(userID))
	followcount, _ := service.GetFollowCount(uint(userID))

	res := make([]User, len(followerList))

	for i, v := range followerList {
		res[i].Id = int64(v.UserID)
		res[i].Name = v.UserName
		res[i].FollowCount = followcount
		res[i].FollowerCount = fancount
		res[i].IsFollow, _ = service.IsFollow(uint(userID), v.UserID)
	}
	return res, nil
}

func GenerateFollow(followList []model.User) ([]User, error) {
	userID := 1
	fancount, _ := service.GetFanCount(uint(userID))
	followcount, _ := service.GetFollowCount(uint(userID))

	res := make([]User, len(followList))

	for i, v := range followList {
		res[i].Id = int64(v.UserID)
		res[i].Name = v.UserName
		res[i].FollowCount = followcount
		res[i].FollowerCount = fancount
		res[i].IsFollow, _ = service.IsFollow(uint(userID), v.UserID)
	}
	return res, nil
}
