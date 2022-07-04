package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"simple-demo/helper"
	"simple-demo/model"
	"simple-demo/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish 确认token后存储用户上传文件到public目录下
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	userID, _ := helper.GetUserIDByToken(token)
	if userID == 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	user, err := service.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", user.UserID, filename)
	saveFile := filepath.Join("./public/video", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	videoInfo := model.Video{
		UserID:    user.UserID,
		Title:     c.PostForm("title"),
		PlayUrl:   finalName,
		CoverUrl:  "bear.jpg",
		CreatedAt: time.Now().Unix(),
	}

	if err = service.CreateVideo(c, &videoInfo); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList 显示该用户发布过的视频
func PublishList(c *gin.Context) {
	userid := c.Query("user_id")
	UID, _ := strconv.ParseUint(userid, 10, 32)

	if videoList_, err := service.GetVideoListByUserID(uint(UID)); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1, StatusMsg: "error occur in feeding",
		})
	} else {
		videoList, _ := GenerateVideo(videoList_, uint(UID))
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: videoList,
		})
	}
}
