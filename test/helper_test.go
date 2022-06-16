package test

import (
	"fmt"
	"simple-demo/helper"
	"testing"
)

func TestHelper(t *testing.T) {
	res, _ := helper.AnalyseToken("FvU")
	fmt.Println(res.UserID)
	fmt.Println(res.UserName)
}