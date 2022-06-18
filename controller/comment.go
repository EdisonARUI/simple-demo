package controller

import (
	"net/http"
	"simple-demo/service"
	"simple-demo/helper"
	"simple-demo/model"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
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
			commentString := c.Query("comment_text")
			cmt, err := service.CreateComment(int64(UID), VID, commentString)
			if err != nil {
				return
			}
			cmtList_ := make([]*model.Comment, 1)
			cmtList_[0] = cmt
			cmtList, _ := GenerateComment(cmtList_)
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{
					StatusCode: 0, 
					StatusMsg: "评论成功",
				},
				Comment: cmtList[0],
			})

		} else if actionType == "2" {
			commentID_ := c.Query("comment_id")
			commentID, _ := strconv.ParseInt(commentID_, 10, 32)
			err := service.DeleteComment(commentID)
			if err != nil {
				return
			}

			c.JSON(http.StatusOK, Response{
				StatusCode: 0,
				StatusMsg:  "删除成功",
			})
		}
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoID := c.Query("video_id")
	token := c.Query("token")
	VID, _ := strconv.ParseInt(videoID, 10, 32)
	videoList_, _ := service.GetComment(VID)
	videoList, _ := GenerateComment(videoList_)

	if token == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "You haven't logged in yet",
		})
	} else {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{StatusCode: 0},
			CommentList: videoList,
		})
	}
}

func GenerateComment(commentList []*model.Comment) ([]Comment, error){
	res := make([]Comment, len(commentList))

	for i, v := range commentList {
		author, _ := GetAuthor((v.UserID))
		createDate, _ := GetCreateDate(int64(v.CreatedAt))
		res[i].Id = int64(v.CommentID)
		res[i].User = author
		res[i].Content = v.Content
		res[i].CreateDate = createDate
	}

	return res, nil
}

func GetCreateDate(unixTime int64) (string, error) {
	timeLayout := "01-02"
	timeString := time.Unix(unixTime, 0).Format(timeLayout)
	return timeString, nil
}