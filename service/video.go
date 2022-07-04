package service

import (
	"context"
	"math"
	"simple-demo/model"
	"simple-demo/helper"
)

// CreateVideo 创建视频信息
func CreateVideo(ctx context.Context, video *model.Video) error {
	return model.DB.WithContext(ctx).Create(&video).Error
}

// GetVideoListByUserID 根据用户id查视频
func GetVideoListByUserID(userID uint) ([]model.Video, error) {
	var videoList []model.Video
	if err := model.DB.Where("user_id = ?", userID).Find(&videoList).Error; err != nil {
		return nil, err
	}
	return videoList, nil
}

// GetVideoByLoginToken 根据用户token提供视频
func GetVideoByLoginToken(token string) ([]model.Video, error) {
	userID, _ := helper.GetUserIDByToken(token)
	var videoList []model.Video
	if err := model.DB.Raw("SELECT * FROM video WHERE user_id <> ? ORDER BY RAND() LIMIT ? ", userID, 5).Scan(&videoList).Error; err != nil {
		return nil, err
	}
	return videoList, nil
}

// GetVideoByNoLoginToken 给非登录用户提供视频
func GetVideoByNoLoginToken() ([]model.Video, error) {
	var videoList []model.Video
	if err := model.DB.Raw("SELECT * FROM video ORDER BY RAND() LIMIT ? ", 5).Scan(&videoList).Error; err != nil {
		return nil, err
	}
	return videoList, nil
}

// GetLikeCount 返回视频点赞数
func GetLikeCount(videoID int64) (int64, error) {
	video := model.Video{VideoID: uint(videoID)}
	return model.DB.Model(&video).Association("Likes").Count(), nil
}

// GetCommentCount 返回视频评论数
func GetCommentCount(videoID int64) (int64, error) {
	video := model.Video{VideoID: uint(videoID)}
	return model.DB.Model(&video).Association("Comments").Count(), nil
}

// IsFavorite 返回是否点赞
func IsFavorite(videoID int64, userID int64) (bool, error) {
	user := model.User{UserID: uint(userID)}
	return model.DB.Model(&user).Where("`like`.video_id = ?", videoID).Association("Likes").Count() > 0, nil
}

// GetVideoByTime 根据时间戳返回最近count个视频,还需要返回next time
func GetVideoByTime(latestTime int64, count int64) ([]*model.Video, int64, error) {
	var videos []*model.Video
	if err := model.DB.Where("created_at < ?", latestTime).Limit(int(count)).Order("created_at DESC").Find(&videos).Error; err != nil {
		return nil, 0, err
	}
	var nextTime int64 = math.MaxInt32
	if len(videos) != 0 { // 查到了新视频
		nextTime = int64(videos[0].CreatedAt)
	}
	return videos, nextTime, nil
}

// LikeVideo 点赞视频
func LikeVideo(userID int64, videoID int64) error {

	user := model.User{UserID: uint(userID)}
	video := model.Video{VideoID: uint(videoID)}
	err := model.DB.Model(&user).Association("Likes").Append(&video)
	if err != nil {
		return err
	}
	return nil
}

// UnLikeVideo 取消点赞视频
func UnLikeVideo(userID int64, videoID int64) error {
	user := model.User{UserID: uint(userID)}
	video := model.Video{VideoID: uint(videoID)}
	err := model.DB.Model(&user).Association("Likes").Delete(&video)
	if err != nil {
		return err
	}
	return nil
}

func GetLikeVideo(userID int64) ([]model.Video, error) {
	var videoList []model.Video

	model.DB.Table("like").Select("like.video_id,video.user_id,video.title,video.play_url,video.cover_url").
		Where("like.user_id=?", userID).
		Joins("LEFT JOIN video ON like.video_id = video.video_id").
		Find(&videoList)

	return videoList, nil
}

// CreateComment 新增评论,需要dal层返回评论详情
func CreateComment(userID int64, videoID int64, content string) (*model.Comment, error) {
	comment := model.Comment{
		UserID:  uint(userID),
		VideoID: uint(videoID),
		Content: content,
	}
	if err := model.DB.Create(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

// DeleteComment 删除评论
func DeleteComment(commentID int64) error {
	return model.DB.Delete(&model.Comment{}, commentID).Error
}

// GetComment 查询评论,需要dal层返回评论详情,有可能有多条评论
func GetComment(videoID int64) ([]*model.Comment, error) {
	video := model.Video{VideoID: uint(videoID)}
	if err := model.DB.Preload("Comments").Find(&video).Error; err != nil {
		return nil, err
	}
	return video.Comments, nil
}
