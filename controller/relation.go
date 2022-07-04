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
	fanid := c.Query("to_user_id")
	actionType := c.Query("action_type")

	UID, _ := helper.GetUserIDByToken(token)
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

			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "关注成功"})

		} else if actionType == "2" {
			err := service.UnFollowUser(int64(FID), int64(UID))
			if err != nil {
				return
			}

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
		followList_, _ := service.GetFollowList(int64(UID))
		followList, _ := GenerateFollow(followList_, uint(UID))

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
		fanList_, _ := service.GetFanList(int64(UID))
		fanList, _ := GenerateFollower(fanList_, uint(UID))

		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Get success"},
			UserList: fanList,
		})
	}
}

func GenerateFollower(followerList []model.User, userID uint) ([]User, error) {
	res := make([]User, len(followerList))

	for i, v := range followerList {
		res[i].Id = int64(v.UserID)
		res[i].Name = v.UserName
		res[i].FollowCount, _ = service.GetFollowCount(userID)
		res[i].FollowerCount, _ = service.GetFanCount(userID)
		res[i].IsFollow, _ = service.IsFollow(userID, v.UserID)
	}
	return res, nil
}

func GenerateFollow(followList []model.User, userID uint) ([]User, error) {
	res := make([]User, len(followList))

	for i, v := range followList {
		res[i].Id = int64(v.UserID)
		res[i].Name = v.UserName
		res[i].FollowCount, _ = service.GetFollowCount(userID)
		res[i].FollowerCount, _ = service.GetFanCount(userID)
		res[i].IsFollow, _ = service.IsFollow(userID, v.UserID)
	}
	return res, nil
}
