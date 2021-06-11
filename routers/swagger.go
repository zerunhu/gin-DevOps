package routers

import (
	_ "gin-DevOps/docs" // 千万不要忘了导入把你上一步生成的docs

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)
func SwagRouters(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("")
	{
		ApiRouter.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	}
}