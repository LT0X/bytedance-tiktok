package models

import (
	"fmt"
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

func NewOnce() *UserFollowsDao {
	userfollowsOnce.Do(func() {
		userfollowsDao = &UserFollowsDao{} //生成单例对象
	})
	return userfollowsDao
}

// InsertFollows 关注操作
func (*UserFollowsDao) InsertFollows(userId int64, toUserId int64) (int8, error) {
	follows := UserFollows{}
	//先查询表里是否有关注的记录
	db := dbutil.GetDB()
	number := db.Where("user_Id=?", userId).Where("to_user_id=?", toUserId).Take(&follows).RowsAffected
	if number == 0 {
		//查询到0条结果，则新增一条关注数据
		newFollows := UserFollows{UserId: userId, ToUserId: toUserId, ActionType: 1}
		db.Create(newFollows)
		return 0, nil
	}
	//查询到有记录数据，则将数据的ActionType改为1，即关注
	db.Model(&follows).Where("user_Id=?", userId).Where("to_user_id=?", toUserId).Update("action_type", 1)
	return 0, nil
}

// CancelFollows 取消关注
func (*UserFollowsDao) CancelFollows(userId int64, toUserId int64) (int8, error) {
	follows := UserFollows{}
	//先查询表里是否有关注的记录
	db := dbutil.GetDB()
	err := db.Where("user_Id=?", userId).Where("to_user_id=?", toUserId).Take(&follows).Error
	if err != nil && strings.EqualFold(err.Error(), "record not found") {
		//没有查询到时打印record not found
		fmt.Println("error:", err)
		return 0, err
	}
	//查询到有记录数据，则将数据的ActionType改为2，即取消关注
	db.Model(&follows).Where("user_Id=?", userId).Where("to_user_id=?", toUserId).Update("action_type", 2)
	return 0, nil
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
	for id := range idList {
		db.Model(&UserInfo{}).Where("user_id=?", id).Create(&userList)
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
	for id := range idList {
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
