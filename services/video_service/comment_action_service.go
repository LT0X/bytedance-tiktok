package video_service

import (
	"errors"
	"qingxunyin/bytedance-tiktok/models"
	"time"
)

type ActionResponse struct {
	Comment *models.Comment `json:"comment"`
}

type CommentActionService struct {
	Uid         int64
	VideoId     int64
	ActionType  string
	CommentTest string
	CommentId   int64
}

func NewCommentActionService(uid int64, videoId int64,
	action string, commentTest string, commentId int64) *CommentActionService {
	return &CommentActionService{
		Uid:         uid,
		VideoId:     videoId,
		ActionType:  action,
		CommentTest: commentTest,
		CommentId:   commentId,
	}
}

func (c *CommentActionService) Do() (*ActionResponse, error) {

	actionResponse := new(ActionResponse)

	if c.ActionType == "1" {
		//发布评论
		comment := &models.Comment{
			UserInfoId: c.Uid,
			VideoId:    c.VideoId,
			Content:    c.CommentTest,
			CreateAt:   time.Now(),
		}
		//开始插入视频评论
		err := models.GetCommentDao().AddVideoComment(comment, c.VideoId)
		if err != nil {
			return nil, err
		}
		user, err := models.GetUserInfoDao().QueryUserInfoById(c.Uid)
		if err != nil {
			return nil, err
		}
		comment.UserInfo = *user
		actionResponse.Comment = comment
		return actionResponse, nil

	} else if c.ActionType == "2" {
		//删除视频评论
		err := models.GetCommentDao().DeleteVideoCommentById(c.CommentId, c.VideoId)
		if err != nil {
			return nil, err
		}
		return actionResponse, nil
	}

	return nil, errors.New("错误操作")
}
