package router

import (
	"github.com/gin-gonic/gin"
	"qingxunyin/bytedance-tiktok/config"
	"qingxunyin/bytedance-tiktok/controllers/user_login"
	"qingxunyin/bytedance-tiktok/middlewares"
)

func InitRouter(r *gin.Engine) {

	//静态资源路径
	r.Static("/static", config.GetConf().Path.StaticPath)

	apiRouter := r.Group("/douyin")

	// 基础接口
	apiRouter.GET("/feed/")
	apiRouter.GET("/user/")
	apiRouter.POST("/user/register/", middlewares.EncryptMiddleWare(), user_login.UserRegisterHandler)
	apiRouter.POST("/user/login/", middlewares.EncryptMiddleWare(), user_login.UserLoginHandler)
	apiRouter.POST("/publish/action/")
	apiRouter.GET("/publish/list/")

	// 互动接口
	apiRouter.POST("/favorite/action/")
	apiRouter.GET("/favorite/list/")
	apiRouter.POST("/comment/action/")
	apiRouter.GET("/comment/list/")

	//社交接口
	apiRouter.POST("/relation/action/")
	apiRouter.GET("/relation/follow/list/")
	apiRouter.GET("/relation/follower/list/")
	apiRouter.GET("/relation/friend/list/")
	apiRouter.GET("/message/chat/")
	apiRouter.POST("/message/action/")

}
