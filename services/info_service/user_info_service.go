package info_service

import (
	"qingxunyin/bytedance-tiktok/cache"
	"qingxunyin/bytedance-tiktok/models"
)

type InfoResponse struct {
	UserInfo *models.UserInfo `json:"user"`
}

type UserInfoService struct {
	Uid int64
	*InfoResponse
}

func NewUserInfoService(uid int64) *UserInfoService {
	return &UserInfoService{
		Uid:          uid,
		InfoResponse: &InfoResponse{},
	}
}

func (u *UserInfoService) Do() (*InfoResponse, error) {
	// 查询用户信息
	info, err := models.GetUserInfoDao().QueryUserInfoById(u.Uid)
	if err != nil {
		return nil, err
	}
	u.InfoResponse.UserInfo = info
	//查询是否关注
	cache.GetIsFollow()
	//返回处理信息
	return u.InfoResponse, nil
}
