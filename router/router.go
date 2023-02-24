package router

import (
	"github.com/gin-gonic/gin"
	"qingxunyin/bytedance-tiktok/config"
	"qingxunyin/bytedance-tiktok/controllers/message"
	"qingxunyin/bytedance-tiktok/controllers/user_follow"
	"qingxunyin/bytedance-tiktok/controllers/user_info"
	"qingxunyin/bytedance-tiktok/controllers/user_login"
	"qingxunyin/bytedance-tiktok/controllers/video"
	"qingxunyin/bytedance-tiktok/middlewares"
)

func InitRouter(r *gin.Engine) {

	//静态资源路径
	r.Static("/static", config.GetConf().Path.StaticPath)

	apiRouter := r.Group("/douyin")

	// 基础接口
	apiRouter.GET("/feed/", video.UserFeedController)
	apiRouter.GET("/user/", middlewares.JWTMiddleWare(), user_info.UserInfoController)
	apiRouter.POST("/user/register/", middlewares.EncryptMiddleWare(), user_login.UserRegisterHandler)
	apiRouter.POST("/user/login/", middlewares.LoginMiddleWare(), user_login.UserLoginHandler)
	apiRouter.POST("/publish/action/", middlewares.JWTMiddleWare(), video.PublishVideoController)
	apiRouter.GET("/publish/list/", middlewares.JWTMiddleWare(), user_info.PublishListController)

	// 互动接口
	apiRouter.POST("/favorite/action/", middlewares.JWTMiddleWare(), video.FavoriteActionController)
	apiRouter.GET("/favorite/list/", middlewares.JWTMiddleWare(), user_info.FavoriteListController)
	apiRouter.POST("/comment/action/", middlewares.JWTMiddleWare(), video.CommentActionController)
	apiRouter.GET("/comment/list/", middlewares.JWTMiddleWare(), video.CommentListController)

	//社交接口
	apiRouter.POST("/relation/action/", middlewares.JWTMiddleWare(), user_follow.FollowsAction)
	apiRouter.GET("/relation/follow/list/", middlewares.JWTMiddleWare(), user_follow.GetFollowsList)
	apiRouter.GET("/relation/follower/list/", middlewares.JWTMiddleWare(), user_follow.GetFollowersList)
	apiRouter.GET("/relation/friend/list/", middlewares.JWTMiddleWare(), user_follow.GetFriendsList)
	apiRouter.GET("/message/chat/", middlewares.JWTMiddleWare(), message.UserMessageController)
	apiRouter.POST("/message/action/", middlewares.JWTMiddleWare(), message.ActionMessageController)

}
