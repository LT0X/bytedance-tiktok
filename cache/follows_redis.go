package cache

import (
	"context"
	"fmt"
)

var ctx = context.Background()

type FollowsRedis struct {
}

var followsRedis FollowsRedis

func GetFollowsRedis() *FollowsRedis {
	return &followsRedis
}

// UpdateFollowsAction 更新关注状态，1关注，2取关
func (f *FollowsRedis) UpdateFollowsAction(userId int64, toUserId int64, actionType int64) {
	sprintf := fmt.Sprintf("%d", userId)
	if actionType == 1 {
		rdb.SAdd(sprintf, toUserId)
		return
	}
	//不是关注的话就移除
	rdb.SRem(sprintf, toUserId)
}

// GetFollowsType 获取关注状态
func (f *FollowsRedis) GetFollowsType(userId int64, toUserId int64) bool {
	sprintf := fmt.Sprintf("%d", userId)
	val := rdb.SIsMember(sprintf, toUserId).Val()
	return val
}
