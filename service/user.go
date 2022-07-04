package service

import (
	"simple-demo/model"
)

// CreateUser 根据所给user信息创建用户,并返回用户ID
func CreateUser(userInfo *model.User) (int64, error) {
	err := model.DB.Create(userInfo).Error
	if err != nil {
		return 0, err
	}
	return int64(userInfo.UserID), nil
}

// UserLogin 根据用户输入信息确认user,并返回用户ID
func UserLogin(userInfo model.User) (int64, int64) {
	find := model.DB.Where("username = ? AND password = ? ", userInfo.UserName, userInfo.Password).Find(&userInfo)
	if find.RowsAffected == 0 {
		return 0, find.RowsAffected
	} else {
		return int64(userInfo.UserID), find.RowsAffected
	}

}

//GetUserByID 需要通过用户ID查询用户信息
func GetUserByID(userID uint) (*model.User, error) {
	var user model.User
	if err := model.DB.Where("userid = ?", userID).First(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

// GetUserByName 需要通过用户名查询用户信息
func GetUserByName(userName string) (*model.User, error) {
	var user model.User
	if err := model.DB.Where("username = ?", userName).First(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

//FollowUser 根据传入的两个用户id  执行用户A关注用户B操作
func FollowUser(fanID, userID int64) error {
	fan := model.User{UserID: uint(fanID)}
	user := model.User{UserID: uint(userID)}
	return model.DB.Model(&user).Association("Fans").Append(&fan)
}

//UnFollowUser 根据传入的两个用户id  执行用户A取消关注用户B操作
func UnFollowUser(fanID, userID int64) error {
	fan := model.User{UserID: uint(fanID)}
	user := model.User{UserID: uint(userID)}
	return model.DB.Model(&user).Association("Fans").Delete(&fan)
}

// GetFanList 根据userID 获取该用户的粉丝列表
func GetFanList(userID int64) ([]model.User, error) {
	var followerList []model.User

	model.DB.Table("follow").Select("follow.fan_id,user.userid,user.username").
		Where("follow.user_id=?", userID).
		Joins("LEFT JOIN user ON follow.fan_id = user.userid").
		Find(&followerList)
	return followerList, nil
}

// GetFollowList 根据userID 获取这个用户所关注的用户列表
func GetFollowList(fanID int64) ([]model.User, error) {
	var followList []model.User

	model.DB.Table("follow").Select("follow.user_id,user.userid,user.username").
		Where("follow.fan_id=?", fanID).
		Joins("LEFT JOIN user ON follow.user_id = user.userid").
		Find(&followList)
	return followList, nil
}

//GetFanCount 传入用户ID 查询粉丝用户数量
func GetFanCount(userID uint) (int64, error) {
	user := model.User{UserID: userID}
	return model.DB.Model(&user).Association("Fans").Count(), nil
}

//GetFollowCount 传入用户ID 查询关注用户ID集合
func GetFollowCount(fanID uint) (int64, error) {
	var count int64
	if err := model.DB.Model(&model.Follow{}).Where("fan_id = ?", fanID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

//IsFollow 根据传入两个用户ID 在follow表中查询用户A是否关注用户B
func IsFollow(fanID, userID uint) (bool, error) {
	user := model.User{UserID: userID}
	return model.DB.Model(&user).Where("fan_id = ?", fanID).Association("Fans").Count() > 0, nil
}
