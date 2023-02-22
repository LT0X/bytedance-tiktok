package follows_service

import (
	"errors"
	"qingxunyin/bytedance-tiktok/cache"
	"qingxunyin/bytedance-tiktok/models"
)

type FollowsService struct {
	UserId     int64
	ToUserId   int64
	ActionType int64
}

func NewFollowsService(userId int64, toUserId int64, actionType int64) *FollowsService {
	return &FollowsService{UserId: userId, ToUserId: toUserId, ActionType: actionType}
}

func (followsService *FollowsService) Do() error {
	//判断自己关注自己的情况
	if followsService.UserId == followsService.ToUserId {
		return errors.New("不能进行此操作")
	}
	//查询用户之间是否有关注记录
	followsOnce := models.NewFollowsOnce()
	relations, _ := followsOnce.FindRelations(followsService.UserId, followsService.ToUserId)
	//当查询到有记录，且要进行关注时
	if relations == 1 && followsService.ActionType == 1 {
		followsOnce.UpdateFollows(followsService.UserId, followsService.ToUserId)
		//更新redis
		cache.NewFollowsRedis().UpdateFollowsAction(followsService.UserId, followsService.ToUserId, 1)
	}
	//当查询到有记录，且要进行取关时
	if relations == 1 && followsService.ActionType == 2 {
		followsOnce.CancelFollows(followsService.UserId, followsService.ToUserId)
		//更新redis
		cache.NewFollowsRedis().UpdateFollowsAction(followsService.UserId, followsService.ToUserId, 2)
	}
	//当没有记录时，插入一条新数据
	if relations == 0 {
		followsOnce.InsertNewFollows(followsService.UserId, followsService.ToUserId)
	}

	return nil
}
