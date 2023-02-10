package test

import (
	"fmt"
	"qingxunyin/bytedance-tiktok/models"
	"testing"
	"time"
)

func TestAddMessage(t *testing.T) {

	message := &models.Message{
		ToUserId:   2,
		FromUserId: 1,
		Content:    "okok",
		CreateTime: time.Now(),
	}

	dao := models.GetMessageDao()
	dao1 := models.GetUserLoginDao()
	dao.AddUserMessage(message)
	fmt.Print(dao1)

}
