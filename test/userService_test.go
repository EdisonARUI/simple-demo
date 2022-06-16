package test

import (
	"fmt"
	"simple-demo/service"
	"simple-demo/model"
	"testing"
)

// func TestInsertRecord(t *testing.T) {
// 	row, err := service.UserInsertRecord("edison", "12345")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(row)
// }

// func TestGetAllRecord(t *testing.T) {
// 	res := make([]*model.User, 0)
// 	res, err := service.UserSelectAll()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for _, v := range res {
// 		fmt.Println(*&v.ID)
// 	}
// }

func TestGetOne(t *testing.T) {
	var res *model.User
	res, err := service.GetUserByID(5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*res)

}

// func TestVideoList(t *testing.T) {
// 	var res []model.Video
// 	res, err := service.GetVideoListByUserID(1)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 		for _, v := range res {
// 		fmt.Println(v)
// 	}

// }

