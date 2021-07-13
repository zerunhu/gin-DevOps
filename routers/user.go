package routers

import (
	"gin-DevOps/api"
	"github.com/gin-gonic/gin"
)


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
		ApiRouter.GET("users", api.ListUser)
		ApiRouter.POST("users", api.CreateUser)
		ApiRouter.PUT("users/:uid", api.UpdateUser)
		ApiRouter.DELETE("users/:uid", api.DeleteUser)

		ApiRouter.POST("/groups", api.CreateGroup)
		ApiRouter.DELETE("/groups/:gid", api.DeleteGroup)
		ApiRouter.GET("/groups", api.ListGroup)
		ApiRouter.PUT("/groups/:gid", api.UpdateGroup)

		ApiRouter.POST("/groups/:gid/users/:uid", api.CreateGroupUser)
		ApiRouter.DELETE("/groups/:gid/users", api.DeleteGroupUser)
		ApiRouter.GET("/groups/:gid/users", api.ListGroupUser)

		ApiRouter.GET("/permissions", api.ListPermission)
		ApiRouter.GET("/groups/:gid/permissions", api.ListGroupPermission)
		ApiRouter.POST("/groups/:gid/permissions/:pid", api.CreateGroupPermission)
		ApiRouter.DELETE("/groups/:gid/permissions", api.DeleteGroupPermission)
		//ApiRouter.POST("/user/:name/*action", api.CreateUser)
	}
}