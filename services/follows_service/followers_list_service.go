package follows_service

import "qingxunyin/bytedance-tiktok/models"

type FollowersList struct {
	UserId int64
	List   *[]models.UserInfo `json:"user_list"`
}

func NewFollowersListService(userId int64) *FollowersList {
	return &FollowersList{UserId: userId}
}

func (followersList *FollowersList) Do() (*[]models.UserInfo, error) {
	followsOnce := models.NewFollowsOnce()
	fans, err := followsOnce.GetFans(followersList.UserId)
	if err != nil {
		return nil, err
	}
	return &fans, nil
}
