package follows_service

import "qingxunyin/bytedance-tiktok/models"

type FollowsList struct {
	UserId int64
	List   *[]models.UserInfo `json:"user_list"`
}

func NewFollowsListService(userId int64) *FollowsList {
	return &FollowsList{UserId: userId}
}

func (followsList *FollowsList) Do() (*[]models.UserInfo, error) {
	followsOnce := models.NewFollowsOnce()
	following, err := followsOnce.GetFollowing(followsList.UserId)
	if err != nil {
		return nil, err
	}
	return &following, nil
}
