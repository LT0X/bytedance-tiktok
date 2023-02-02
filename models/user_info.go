package models

import "qingxunyin/bytedance-tiktok/util/dbutil"

type UserInfo struct {
	Id            int64
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool `json:"is_follow" gorm:"-"`
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
	return DB.Table("user_info").Create(&userInfo).Error
}
