package test

import (
	"fmt"
	"qingxunyin/bytedance-tiktok/models"
	"testing"
)

func TestAdd(t *testing.T) {

	xx := models.UserLogin{
		Id:         12,
		UserInfoId: 1,
		Username:   "xx",
		Password:   "sdsd",
	}

	xxx := models.GetUserLoginDao()
	fmt.Print(xxx, xx)
}

func Pxx() {

}
