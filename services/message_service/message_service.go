package message_service

import (
	"qingxunyin/bytedance-tiktok/models"
	time2 "time"
)

type MessageResponse struct {
	MessageList *[]models.Message `json:"message_list"`
}

type MessageService struct {
	Uid      int64
	ToUserId int64
}

func NewMessageService(uid int64, toUserId int64) *MessageService {

	return &MessageService{
		Uid:      uid,
		ToUserId: toUserId,
	}
}

func (m *MessageService) Do() (*MessageResponse, error) {
	//查询用户聊天录
	response := new(MessageResponse)
	messageList, err := models.GetMessageDao().QueryUserMessage(m.Uid, m.ToUserId)

	//将time.Time 转为int64
	for index, _ := range *messageList {
		time := (*messageList)[index].CreateTime

		//(*messageList)[index].Temp = (time).UnixNano() / 1e6
		(*messageList)[index].PushTime = time.Format(time2.Kitchen)
	}
	if err != nil {
		return nil, err
	}

	response.MessageList = messageList
	return response, nil
}
