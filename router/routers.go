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
		system.CommonRouter(PublicGroup)
	}
	// 后台用户权限路由
	AdminGroup := Router.Group("/admin")
	{
		system.UserRouter(AdminGroup)
		system.AdminRouter(AdminGroup)
	}
	// 功能操作路由
	OperateGroup := Router.Group("/operate")
	{
		system.ItemRouter(OperateGroup)
	}

	return Router
}
