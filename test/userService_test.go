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

// func TestGetOne(t *testing.T) {
// 	var res *model.User
// 	res, err := service.GetUserIDByToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFt6ImVkaXNvbiIsInVzZXJQYXNzd29yZCI6IjcxYjM3NzM5ODgzZjJlMzA1OWI1YzQ2NWU1Yjg1ZjJlIiwiZXhwIjoxNjU1MTMyMjE2LCJpYXQiOjE2NTUxMjg2MTYsImlzcyI6IkRvdXlpbiIsInN1YiI6InVzZXJUb2tlbiJ9.hBVXXBhnryDzxCdDhkwfWrJNg6zXd_Ovclg2qOX1W2g")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(*res)

// }

func TestVideoList(t *testing.T) {
	var res []model.Video
	res, err := service.GetVideoByLoginToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6InRlc3QiLCJ1c2VyUGFzc3dvcmQiOiJlMTBhZGMzOTQ5YmE1OWFiYmU1NmUwNTdmMjBmODgzZSIsImV4cCI6MTY1NTA5OTgwOSwiaWF0IjoxNjU1MDk2MjA5LCJpc3MiOiJEb3V5aW4iLCJzdWIiOiJ1c2VyVG9rZW4ifQ.ON8_TppDyzjblmR0Oel7jiqkz4NpNZ830UPcmoaMeTA")
	if err != nil {
		fmt.Println(err)
	}
		for _, v := range res {
		fmt.Println(v)
	}

}