package models

import (
	"qingxunyin/bytedance-tiktok/util/dbutil"
	"time"
)

type Comment struct {
	Id         int64     `json:"id"`
	UserInfoId int64     `json:"-" `
	VideoId    int64     `json:"-"`
	UserInfo   UserInfo  `json:"user" `
	Content    string    `json:"content"`
	LikeCount  int       `json:"like_count"`
	TeaseCount int       `json:"tease_count"`
	CreateAt   time.Time `json:"-" `
	PushDate   string    `json:"CreatAt" gorm:"-"`
}

type CommentDao struct {
}

var commentDao CommentDao

func GetCommentDao() *CommentDao {
	return &commentDao
}

func (CommentDao) QueryVideoComment(videoId int64) (*[]Comment, error) {
	DB := dbutil.GetDB()
	commentList := make([]Comment, 30)
	var err error
	err = DB.Table("comments").Where("video_id = ?", videoId).Find(&commentList).Error
	if err != nil {
		return nil, err
	}
	for index, _ := range commentList {
		//查找评论作者信息
		err = DB.Table("user_infos").Where("id = ?", commentList[index].UserInfoId).First(&commentList[index].UserInfo).Error
		if err != nil {
			return nil, err
		}
		//评论时间

		commentList[index].PushDate = commentList[index].CreateAt.Format("01-02")
	}

	return &commentList, nil
}

func (CommentDao) AddVideoComment(comment *Comment, vid int64) error {
	tx := dbutil.GetDB().Begin()
	err := tx.Table("comments").Create(comment).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Exec("update videos set comment_count = comment_count +1 where id =?", vid).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil

}

func (CommentDao) DeleteVideoCommentById(commentId int64, videoId int64) error {

	//开启事务进行多表操作
	tx := dbutil.GetDB()
	tx.Begin()
	err := tx.Exec("delete from comments where id = ?", commentId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Exec("update videos set comment_count = comment_count -1 where id = ?", videoId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil

}
