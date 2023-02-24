package models

import (
	"qingxunyin/bytedance-tiktok/util/dbutil"
)

type UserInfo struct {
	Id              int64  `json:"id" gorm:"primary key"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow" gorm:"-"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

type UserInfoDao struct {
}

var userInfoDao UserInfoDao

// GetUserInfoDao 单例返回Dao对象
func GetUserInfoDao() *UserInfoDao {
	return &userInfoDao
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

func (*UserInfoDao) UpdateUserFavor(vid int64, uid int64) error {
	DB := dbutil.GetDB()
	return DB.Exec("insert into user_favor_videos values (?,?)", uid, vid).Error
}

func (*UserInfoDao) UpdateUserFavorCount(vid int64, uid int64) error {
	DB := dbutil.GetDB()
	return DB.Exec("update user_infos set favorite_count = (select count(*) from user_favor_videos where user_info_id = ?)", uid).Error
}

func (*UserInfoDao) UpdateUserTotalCount(vid int64, count int64) error {
	DB := dbutil.GetDB()
	return DB.Exec("update user_infos set total_favorited = ? where id = (select user_info_id from videos where id = ?)", count, vid).Error
}

func (*UserInfoDao) DeleteUserFavor(vid int64, uid int64) error {
	DB := dbutil.GetDB()
	return DB.Exec("delete from  user_favor_videos where user_info_id = ? and video_id = ?", uid, vid).Error
}
