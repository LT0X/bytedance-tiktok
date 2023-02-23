package follows_service

import (
	"qingxunyin/bytedance-tiktok/models"
	"qingxunyin/bytedance-tiktok/util/dbutil"
)

type FollowersList struct {
	UserId int64
	List   *[]models.UserInfo
}

func NewFollowersListService(userId int64) *FollowersList {
	return &FollowersList{UserId: userId}
}

func (followersList *FollowersList) Do() (*[]models.UserInfo, error) {
	followsOnce := models.NewFollowsOnce()
	fans, err := followsOnce.GetFans(followersList.UserId)
	var num = 0
	for i := 0; i < len(fans); i++ {
		num++
	}
	db := dbutil.GetDB()
	db.Model(&models.UserInfo{}).Where("user_id=?", followersList.UserId).Update("follower_count", num)
	if err != nil {
		return nil, err
	}
	return &fans, nil
}
