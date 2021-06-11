package routers

import (
	"gin-DevOps/api"
	"github.com/gin-gonic/gin"
)

//func LoginRouters(e *gin.Engine) {
//	e.POST("/api/login", api.LoginHandler)
//	e.GET("/api/userinfo",api.UserInfoHandler)
//}
func LoginRouters(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("")
	{
		ApiRouter.POST("login", api.Login)
	}
}
func UserRouters(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("")
	{
		ApiRouter.GET("userinfo", api.UserInfo)
		ApiRouter.GET("user", api.ListUser)
		ApiRouter.POST("user", api.CreateUser)
		ApiRouter.DELETE("/user/:id", api.DeleteUser)

		ApiRouter.POST("/group", api.CreateGroup)
		//ApiRouter.POST("/user/:name/*action", api.CreateUser)
	}
}