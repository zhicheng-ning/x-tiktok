package main

import (
	"github.com/gin-gonic/gin"
	"x-tiktok/controller"
	"x-tiktok/middleware/jwt"
)

func initRouter(r *gin.Engine) {
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// 基础接口
	apiRouter.GET("/feed/", jwt.AuthWithoutLogin(), controller.Feed)
	apiRouter.GET("/user/", jwt.Auth(), controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", jwt.AuthBody(), controller.Publish)
	apiRouter.GET("/publish/list/", jwt.Auth(), controller.PublishList)

	// 互动接口
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// 社交接口
	apiRouter.POST("/relation/action/", jwt.Auth(), controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", jwt.Auth(), controller.FollowList)
	apiRouter.GET("/relation/follower/list/", jwt.Auth(), controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", jwt.Auth(), controller.FriendList)
	apiRouter.GET("/message/chat/", jwt.Auth(), controller.MessageChat)
	apiRouter.POST("/message/action/", jwt.Auth(), controller.MessageAction)

	// 测试接口
	apiRouter.POST("/test/", jwt.Auth(), controller.Test)
}
