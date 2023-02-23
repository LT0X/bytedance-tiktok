package models

import (
	"qingxunyin/bytedance-tiktok/util/dbutil"
	"strings"
	"sync"
)

type UserFollows struct {
	Id         int64
	UserId     int64
	ToUserId   int64
	ActionType int8
}

type UserFollowsDao struct {
}

var (
	userfollowsDao  *UserFollowsDao
	userfollowsOnce sync.Once //用来限定userfollowsDao为单例
)

func NewFollowsOnce() *UserFollowsDao {
	userfollowsOnce.Do(func() {
		userfollowsDao = &UserFollowsDao{} //生成单例对象
	})
	return userfollowsDao
}

// FindRelations 查询用户之间是否有关注记录
func (*UserFollowsDao) FindRelations(userId int64, toUserId int64) (int8, error) {
	follows := UserFollows{}
	//先查询表里是否有关注的记录
	db := dbutil.GetDB()
	err := db.Where("user_id=?", userId).Where("to_user_id=?", toUserId).Take(&follows).Error
	if err != nil && strings.EqualFold(err.Error(), "record not found") {
		return 0, err
	}
	//查询到有记录，返回1,nil
	return 1, nil
}

// InsertNewFollows 当没有关注记录时，可以插入一条新记录
func (*UserFollowsDao) InsertNewFollows(userId int64, toUserId int64) bool {
	db := dbutil.GetDB()
	newFollows := UserFollows{UserId: userId, ToUserId: toUserId, ActionType: 1}
	db.Create(newFollows)
	var num int64
	var num1 int64
	db.Model(&UserInfo{}).Select("follow_count").Where("user_id=?", userId).Find(&num)
	num = num + 1
	db.Model(&UserInfo{}).Select("follower_count").Where("user_id=?", toUserId).Find(&num1)
	num1 = num1 + 1
	db.Model(&UserInfo{}).Where("user_id=?", userId).Update("follow_count", num)
	db.Model(&UserInfo{}).Where("user_id=?", toUserId).Update("follower_count", num1)
	return true
}

// UpdateFollows 更新关注操作
func (*UserFollowsDao) UpdateFollows(userId int64, toUserId int64) bool {
	db := dbutil.GetDB()
	db.Model(&UserFollows{}).Where("user_id=?", userId).Where("to_user_id=?", toUserId).Update("action_type", 1)
	var num int64
	var num1 int64
	db.Model(&UserInfo{}).Select("follow_count").Where("user_id=?", userId).Find(&num)
	num = num + 1
	db.Model(&UserInfo{}).Select("follower_count").Where("user_id=?", toUserId).Find(&num1)
	num1 = num1 + 1
	db.Model(&UserInfo{}).Where("user_id=?", userId).Update("follow_count", num)
	db.Model(&UserInfo{}).Where("user_id=?", toUserId).Update("follower_count", num1)
	return true
}

// CancelFollows 取消关注
func (*UserFollowsDao) CancelFollows(userId int64, toUserId int64) bool {
	db := dbutil.GetDB()
	db.Model(&UserFollows{}).Where("user_id=?", userId).Where("to_user_id=?", toUserId).Update("action_type", 2)
	var num int64
	var num1 int64
	db.Model(&UserInfo{}).Select("follow_count").Where("user_id=?", userId).Find(&num)
	num = num - 1
	db.Model(&UserInfo{}).Select("follower_count").Where("user_id=?", toUserId).Find(&num1)
	num1 = num1 - 1
	db.Model(&UserInfo{}).Where("user_id=?", userId).Update("follow_count", num)
	db.Model(&UserInfo{}).Where("user_id=?", toUserId).Update("follower_count", num1)
	return true
}

// GetFans 获取粉丝列表
func (*UserFollowsDao) GetFans(userId int64) ([]UserInfo, error) {
	var idList []int64
	var userList []UserInfo
	//获取粉丝列表id
	db := dbutil.GetDB()
	err := db.Model(&UserFollows{}).Select("user_id").Where("to_user_id=?", userId).Find(&idList).Error
	//没有粉丝时
	if err != nil && strings.EqualFold(err.Error(), "record not found") {
		return nil, err
	}
	//根据粉丝列表id获取粉丝信息
	for _, id := range idList {
		db.Model(&UserInfo{}).Where("user_id=?", id).Create(&userList)
	}
	for _, user := range userList {
		var isf bool
		db.Model(&UserFollows{}).Select("action_type").Where("user_id=?", userId).Where("to_user_id=?", user.Id).Take(&isf)
		if isf == true {
			user.IsFollow = true
		} else {
			user.IsFollow = false
		}
	}
	return userList, nil
}

// GetFollowing 获取关注列表
func (*UserFollowsDao) GetFollowing(userId int64) ([]UserInfo, error) {
	var idList []int64
	var userList []UserInfo
	//获取关注列表id
	db := dbutil.GetDB()
	err := db.Model(&UserFollows{}).Select("to_user_id").Where("user_id=?", userId).Where("action_type=?", 1).Find(&idList).Error
	//没有关注列表时
	if err != nil && strings.EqualFold(err.Error(), "record not found") {
		return nil, err
	}
	//根据关注列表id获取关注信息
	for _, id := range idList {
		db.Model(&UserInfo{}).Where("user_id=?", id).Create(&userList)
	}
	return userList, nil
}

// FriendsList 获取好友列表
func (*UserFollowsDao) FriendsList(userId int64) ([]UserInfo, error) {
	var idList []int64
	var friendIdList []int64
	var userList []UserInfo
	//获取关注列表id
	db := dbutil.GetDB()
	err := db.Model(&UserFollows{}).Select("to_user_id").Where("user_id=?", userId).Where("action_type=?", 1).Find(&idList).Error
	//没有关注列表时
	if err != nil && strings.EqualFold(err.Error(), "record not found") {
		return nil, err
	}
	//获取互关的id列表
	for id := range idList {
		db.Model(&UserFollows{}).Select("to_user_id").Where("user_id=?", id).Where("action_type=?", 1).Find(&friendIdList)
	}
	//根据互关的id列表获取用户信息
	for usId := range friendIdList {
		db.Model(&UserInfo{}).Where("user_id=?", usId).Create(&userList)
	}
	return userList, nil
}
