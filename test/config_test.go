package test

import (
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

	models.GetUserLoginDao().AddUserLogin(&xx)
}

func Pxx() {

}
