package models

import "qingxunyin/bytedance-tiktok/util/dbutil"

type UserInfo struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow" gorm:"-"`
	Avatar        string `json:"avatar"`
}

type UserInfoDao struct {
}

var userInfoDao *UserInfoDao

// GetUserInfoDao 单例返回Dao对象
func GetUserInfoDao() *UserInfoDao {
	return userInfoDao
}

func (*UserInfoDao) AddUserInfoDao(userInfo *UserInfo) error {
	DB := dbutil.GetDB()
	if userInfo == nil {
		return ErrNullPointer
	}
	return DB.Table("user_infos").Create(&userInfo).Error
}

func (*UserInfoDao) QueryUserInfoById(uid int64) (*UserInfo, error) {
	DB := dbutil.GetDB()
	userInfo := new(UserInfo)
	return userInfo, DB.Table("user_infos").
		Where("id =  ?", uid).First(userInfo).Error
}
