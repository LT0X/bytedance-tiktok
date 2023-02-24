package follows_service

import "qingxunyin/bytedance-tiktok/models"

type FriendsList struct {
	UserId int64
	List   *[]models.UserInfo
}

func NewFriendsListService(userId int64) *FriendsList {
	return &FriendsList{UserId: userId}
}

func (friendsList *FriendsList) Do() (*[]models.UserInfo, error) {
	followsOnce := models.NewFollowsOnce()
	list, err := followsOnce.FriendsList(friendsList.UserId)
	if err != nil {
		return nil, err
	}
	return &list, nil
}
