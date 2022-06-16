package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simple-demo/helper"
	"simple-demo/service"
	"strconv"
)

// FavoriteAction 点赞操作
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoid := c.Query("video_id")
	actionType := c.Query("action_type")
	VID, _ := strconv.ParseInt(videoid, 10, 32)
	UID, _ := helper.GetUserIDByToken(token)

	if token == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "You haven't logged in yet",
		})
	} else {
		if actionType == "1" {
			err := service.LikeVideo(int64(UID), VID)
			if err != nil {
				return
			}
			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "点赞成功"})

		} else if actionType == "2" {
			err := service.UnLikeVideo(int64(UID), VID)
			if err != nil {
				return
			}

			c.JSON(http.StatusOK, Response{
				StatusCode: 0,
				StatusMsg:  "取消成功",
			})
		}
	}

}

func FavoriteList(c *gin.Context) {
	userid := c.Query("user_id")
	UID, _ := strconv.ParseUint(userid, 10, 32)
	token := c.Query("token")

	if token == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "You haven't logged in yet",
		})
	} else {
		videoList_, err := service.GetLikeVideo(int64(UID))
		videoList, err := GenerateVideo(videoList_)
		log.Println(videoList)
		log.Println(err)
		log.Println(UID)
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0, StatusMsg: "Get success"},
			VideoList: videoList,
		})
	}

}
