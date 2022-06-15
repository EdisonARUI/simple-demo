package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simple-demo/model"
	"testing"
)

// 建表
func TestTableCreate(t *testing.T){
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:3306)/simple_demo?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	err = db.AutoMigrate(&model.User{})
	err = db.AutoMigrate(&model.Video{})
	if err != nil {
		t.Fatal(err)
	}
}