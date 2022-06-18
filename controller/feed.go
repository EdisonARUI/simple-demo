package controller

import (
	"net/http"
	"simple-demo/model"
	"simple-demo/service"
	"simple-demo/define"
	"time"
	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		videoList_, _ := service.GetVideoByNoLoginToken()
		videoList, err := GenerateVideo(videoList_)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1, StatusMsg: "error occur in feeding",
			})
		} else {
			c.JSON(http.StatusOK, FeedResponse{
				Response:       Response{StatusCode: 0},
				VideoList: 		videoList,
				NextTime:       time.Now().Unix(),
			})
		}
	} else {
		videoList_, _ := service.GetVideoByLoginToken(token)
		videoList, err := GenerateVideo(videoList_)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1, StatusMsg: "error occur in feeding",
			})
		} else {
			c.JSON(http.StatusOK, FeedResponse{
				Response:       Response{StatusCode: 0},
				VideoList: 		videoList,
				NextTime:       time.Now().Unix(),
			})
		}
	}
}

func GenerateVideo(videoList []model.Video) ([]Video, error){
	res := make([]Video, len(videoList))
	
	for i, v := range videoList {
		author, _ := GetAuthor(v.UserID)
		res[i].Id = int64(v.VideoID)
		res[i].Author = author
		res[i].PlayUrl = define.VideoRoot + v.PlayUrl
		res[i].CoverUrl = define.ImgRoot + v.CoverUrl
		res[i].FavoriteCount = 0
		res[i].CommentCount = 0
		res[i].IsFavorite = false
		res[i].Title = v.Title
	}
	return res, nil
}

func GetAuthor(userID uint) (user User, err error) {
	if user_, err := service.GetUserByID(userID); err != nil {
		return User{}, err
	} else {
		user.Id = int64(user_.UserID)
		user.Name = user_.UserName
		user.FollowCount = 0
		user.FollowerCount = 0
		user.IsFollow = false
	}
	return user, nil
}