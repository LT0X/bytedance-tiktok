package test

import (
	"github.com/go-playground/assert/v2"
	"qingxunyin/bytedance-tiktok/models"
	"testing"
)

func TestUserLoginAdd(t *testing.T) {

	xx := models.UserLogin{

		UserInfoId: 1,
		Username:   "test1",
		Password:   "sdsd",
	}
	models.GetUserLoginDao().AddUserLogin(&xx)

}
func TestIsUsernameExist(t *testing.T) {

	bool1 := models.GetUserLoginDao().IsUsernameExist("test1")
	assert.Equal(t, bool1, true)
	bool2 := models.GetUserLoginDao().IsUsernameExist("-1s")
	assert.Equal(t, bool2, false)
}
