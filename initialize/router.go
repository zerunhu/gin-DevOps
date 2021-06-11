package initialize

import (
	"gin-DevOps/middleware"
	"gin-DevOps/routers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	Router := gin.Default()
	PublicGroup := Router.Group("api")
	{
		routers.LoginRouters(PublicGroup)
		routers.SwagRouters(PublicGroup)
	}
	PrivateGroup := Router.Group("api")
	PrivateGroup.Use(middleware.JWTAuthMiddleware())
	{
		routers.UserRouters(PrivateGroup)
	}
	return Router
}
