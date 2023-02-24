package message_service

import (
	"fmt"
	"qingxunyin/bytedance-tiktok/cache"
	"qingxunyin/bytedance-tiktok/models"
	"time"
)

type ActionMessageService struct {
	FromUserId int64
	ToUserId   int64
	Content    string
}

func NewActionMessageService(tUid int64, fUid int64, content string) *ActionMessageService {
	return &ActionMessageService{
		FromUserId: fUid,
		ToUserId:   tUid,
		Content:    content,
	}
}

func (a *ActionMessageService) Do() error {

	//开始插入聊天数据库,
	msg := models.Message{
		FromUserId: a.FromUserId,
		ToUserId:   a.ToUserId,
		Content:    a.Content,
		CreateTime: time.Now(),
	}
	err := models.GetMessageDao().AddUserMessage(&msg)
	key := fmt.Sprintf("%v:%v", a.FromUserId, a.ToUserId)
	unix := msg.CreateTime.UnixNano() / 1e6
	cache.SetUserPreMsgTime(key, fmt.Sprintf("%v", unix))
	if err != nil {
		return err
	}
	return nil

}
