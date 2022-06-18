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
	var res []*model.Comment
	res, err := service.GetComment(1)
	// err := service.DeleteComment(2)

	if err != nil {
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Println(*&v)
	}

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
