package models

import (
	"qingxunyin/bytedance-tiktok/util/dbutil"
	"time"
)

type Message struct {
	Id         int64     `json:"id"`
	ToUserId   int64     `json:"to_user_id"`
	FromUserId int64     `json:"from_user_id"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"-"`
	PushTime   int64     `json:"create_time" gorm:"-"`
}

type Msg struct {
	Id         int64  `json:"id"`
	ToUserId   int64  `json:"to_user_id"`
	FromUserId int64  `json:"from_user_id"`
	Content    string `json:"content"`
}

type MessageDao struct {
}

var messageDao MessageDao

func GetMessageDao() *MessageDao {
	return &messageDao
}

// QueryUserMessage 返回用户和好友的聊天记录
func (MessageDao) QueryUserMessage(uid int64, toUserId int64) (*[]Message, error) {
	DB := dbutil.GetDB()
	messageList := make([]Message, 30)
	return &messageList, DB.Table("messages").Order("create_time asc").
		Where("to_user_id = ? and from_user_id = ? or "+
			"to_user_id = ? and from_user_id = ?", uid, toUserId, toUserId, uid).
		Find(&messageList).Error
}

// QueryUserMsg 返回用户和好友的聊天记录
func (MessageDao) QueryUserMsg(uid int64, toUserId int64) (*[]Msg, error) {
	DB := dbutil.GetDB()
	messageList := make([]Msg, 30)
	return &messageList, DB.Table("messages").Order("create_time asc").
		Where("to_user_id = ? and from_user_id = ? or "+
			"to_user_id = ? and from_user_id = ?", uid, toUserId, toUserId, uid).
		Find(&messageList).Error
}

// AddUserMessage 添加用户聊天记录
func (MessageDao) AddUserMessage(message *Message) error {
	DB := dbutil.GetDB()
	return DB.Table("messages").Create(message).Error
}

func (MessageDao) QueryUserMsgByTime(uid int64, toUserId int64, preTime time.Time) (*[]Message, error) {

	DB := dbutil.GetDB()
	messageList := new([]Message)
	return messageList, DB.Table("messages").Order("create_time asc").
		Where("create_time > ? and to_user_id = ? and from_user_id = ? or "+
			"to_user_id = ? and from_user_id = ? and create_time > ? ", uid, toUserId,
			toUserId, uid, preTime, preTime).Find(messageList).Error
}

func (MessageDao) QueryUserLastMessage(uid int64, toUserId int64, preTime time.Time) (*[]Message, error) {

	DB := dbutil.GetDB()
	messageList := new([]Message)
	return messageList, DB.Table("messages").Where("create_time > ", preTime).Find(messageList).Error
}
