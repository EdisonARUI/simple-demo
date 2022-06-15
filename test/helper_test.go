package test

import (
	"fmt"
	"simple-demo/helper"
	"testing"
)

func TestHelper(t *testing.T) {
	res, _ := helper.GenerateToken("resr", "dsfsf")
	fmt.Println(res)
}