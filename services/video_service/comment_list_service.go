package video_service

import "qingxunyin/bytedance-tiktok/models"

type CommentResponse struct {
	CommentList *[]models.Comment `json:"comment_list"`
}

type CommentListService struct {
	VideoId int64
}

func NewCommentListService(videoId int64) *CommentListService {
	return &CommentListService{
		VideoId: videoId,
	}
}

func (c CommentListService) Do() (*CommentResponse, error) {

	response := new(CommentResponse)
	//开始查询评论列表
	comment, err := models.GetCommentDao().QueryVideoComment(c.VideoId)
	if err != nil {
		return nil, err
	}
	response.CommentList = comment
	return response, err
}
