package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/define"
	"simple-demo/helper"
	"simple-demo/model"
	"simple-demo/service"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个
func Feed(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		videoList_, _ := service.GetVideoByNoLoginToken()
		videoList, err := GenerateVideo(videoList_, 0)

		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1, StatusMsg: "error occur in feeding",
			})
		} else {
			c.JSON(http.StatusOK, FeedResponse{
				Response:  Response{StatusCode: 0},
				VideoList: videoList,
				NextTime:  time.Now().Unix(),
			})
		}
	} else {
		UID, _ := helper.GetUserIDByToken(token)
		videoList_, _ := service.GetVideoByLoginToken(token)
		videoList, err := GenerateVideo(videoList_, uint(UID))
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1, StatusMsg: "error occur in feeding",
			})
		} else {
			c.JSON(http.StatusOK, FeedResponse{
				Response:  Response{StatusCode: 0},
				VideoList: videoList,
				NextTime:  time.Now().Unix(),
			})
		}
	}
}

func GenerateVideo(videoList []model.Video, userid uint) ([]Video, error) {
	res := make([]Video, len(videoList))

	for i, v := range videoList {
		author, _ := GetAuthor(v.UserID, userid)
		res[i].Id = int64(v.VideoID)
		res[i].Author = author
		res[i].PlayUrl = define.VideoRoot + v.PlayUrl
		res[i].CoverUrl = define.ImgRoot + v.CoverUrl
		res[i].FavoriteCount, _ = service.GetLikeCount(int64(v.VideoID))
		res[i].CommentCount, _ = service.GetCommentCount(int64(v.VideoID))
		res[i].IsFavorite, _ = service.IsFavorite(int64(v.VideoID), int64(userid))
		res[i].Title = v.Title
	}
	return res, nil
}

func GetAuthor(AuthorID uint, userid uint) (user User, err error) {
	if user_, err := service.GetUserByID(AuthorID); err != nil {
		return User{}, err
	} else {
		user.Id = int64(user_.UserID)
		user.Name = user_.UserName
		user.FollowCount, _ = service.GetFollowCount(AuthorID)
		user.FollowerCount, _ = service.GetFanCount(AuthorID)
		user.IsFollow, _ = service.IsFollow(userid, AuthorID)

	}
	return user, nil
}
