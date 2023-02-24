package message_service

import (
	"fmt"
	"qingxunyin/bytedance-tiktok/cache"
	"qingxunyin/bytedance-tiktok/models"
	"strconv"
	"time"
)

type MessageResponse struct {
	MessageList *[]models.Message `json:"message_list"`
}

type MessageService struct {
	Uid      int64
	ToUserId int64
	PreTime  int64
}

func NewMessageService(uid int64, toUserId int64, preTime int64) *MessageService {

	return &MessageService{
		Uid:      uid,
		ToUserId: toUserId,
		PreTime:  preTime,
	}
}

func (m *MessageService) Do() (*MessageResponse, error) {

	//查询用户聊天录
	response := new(MessageResponse)
	var messageList *[]models.Message
	var err error
	var userTime int64
	var toUserTime int64

	if m.PreTime == 0 {
		//初次访问，查找全部聊天记录
		messageList, err = models.GetMessageDao().QueryUserMessage(m.Uid, m.ToUserId)
		if err != nil {
			return nil, err
		}
	} else {
		//查询redis中用户最新时间
		key := fmt.Sprintf("%v:%v", m.Uid, m.ToUserId)
		key2 := fmt.Sprintf("%v:%v", m.ToUserId, m.Uid)
		tomsgTime, _ := cache.GetUserPreMsgTime(key2)
		msgTime, _ := cache.GetUserPreMsgTime(key)
		if msgTime == "" {
			//第一次访问
			err = cache.SetUserPreMsgTime(key, fmt.Sprintf("%v", m.PreTime))
			if err == nil {
				return nil, err
			}
			userTime = m.PreTime
		} else {
			toUserTime, err = strconv.ParseInt(tomsgTime, 10, 64)
			userTime, err = strconv.ParseInt(msgTime, 10, 64)
			if err != nil {
				return nil, err
			}
		}
		if userTime < toUserTime {
			//更新最新时间
			err := cache.SetUserPreMsgTime(key, fmt.Sprintf("%v", toUserTime))
			if err != nil {
				return nil, err
			}
			//err = cache.SetUserPreMsgTime(key2, fmt.Sprintf("%v", m.PreTime))
			//if err != nil {
			//	return nil, err
			//}
			//查询聊天记录
			times := time.Unix(userTime/1000, (userTime%1000)*(1000*1000))
			messageList, err = models.GetMessageDao().QueryUserLastMessage(m.Uid, m.ToUserId, times)
		} else {
			if messageList == nil {
				var list []models.Message
				messageList = &list
			} else {
				response.MessageList = messageList
			}
			return response, nil
		}
	}

	fmt.Sprint(userTime)
	//对create_time做处理
	for index, _ := range *messageList {
		times := (*messageList)[index].CreateTime
		(*messageList)[index].PushTime = (times).UnixNano() / 1e6
		//(*messageList)[index].PushTime = time.Format(time2.Kitchen)
	}
	response.MessageList = messageList
	return response, nil
}
