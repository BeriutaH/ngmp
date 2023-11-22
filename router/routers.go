package router

import (
	"github.com/gin-gonic/gin"
	"ngmp/router/system"
)

// Routers 路由
func Routers() *gin.Engine {
	var Router = gin.Default()

	//// 公共路由
	//PublicGroup := Router.Group("/")
	//{
	//	router.InitTestRouter(PublicGroup)
	//}

	// 后台路由
	AdminGroup := Router.Group("/admin")
	{
		system.UserRouter(AdminGroup)
		system.SystemRouter(AdminGroup)
	}

	return Router
}
