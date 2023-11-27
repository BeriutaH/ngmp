package router

import (
	"github.com/gin-gonic/gin"
	"ngmp/router/system"
)

// Routers 路由
func Routers() *gin.Engine {
	var Router = gin.Default()

	// 公共路由
	PublicGroup := Router.Group("/")
	{
		// 公共路由
		system.CommonRouter(PublicGroup)
		// 后台
		system.AdminRouter(PublicGroup)
		// 功能操作
		system.ItemRouter(PublicGroup)

	}
	return Router
}
