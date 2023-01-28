package router

import (
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {

	//静态资源路径
	r.Static("", "")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/")
	apiRouter.GET("/user/")
	apiRouter.POST("/user/register/")
	apiRouter.POST("/user/login/")
	apiRouter.POST("/publish/action/")
	apiRouter.GET("/publish/list/")

	// extra apis - I
	apiRouter.POST("/favorite/action/")
	apiRouter.GET("/favorite/list/")
	apiRouter.POST("/comment/action/")
	apiRouter.GET("/comment/list/")

	// extra apis - II
	apiRouter.POST("/relation/action/")
	apiRouter.GET("/relation/follow/list/")
	apiRouter.GET("/relation/follower/list/")
	apiRouter.GET("/relation/friend/list/")
	apiRouter.GET("/message/chat/")
	apiRouter.POST("/message/action/")

}
