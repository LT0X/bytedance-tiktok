package follows_service

import (
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/util/dbutil"
)

type FollowsList struct {
	UserId int64
	List   *[]models.UserInfo
}

func NewFollowsListService(userId int64) *FollowsList {
	return &FollowsList{UserId: userId}
}

func (followsList *FollowsList) Do() (*[]models.UserInfo, error) {
	followsOnce := models.NewFollowsOnce()
	following, err := followsOnce.GetFollowing(followsList.UserId)
	var num = 0
	for i := 0; i < len(following); i++ {
		num++
	}
	db := dbutil.GetDB()
	db.Model(&models.UserInfo{}).Where("user_id=?", followsList.UserId).Update("follow_count", num)
	if err != nil {
		return nil, err
	}
	return &following, nil
}
