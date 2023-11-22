package system

import (
	"github.com/gin-gonic/gin"
	"ngmp/controllers/system"
)

// SystemRouter 初始化管理 配置路由信息
func SystemRouter(e *gin.RouterGroup) {

	SystemRouter := e.Group("/system")
	{
		// 角色操作
		SystemRouter.POST("/role/list", system.RoleData)
		SystemRouter.POST("/role/add", system.RoleAdd)
		SystemRouter.POST("/role/modify", system.RoleAdd)
		SystemRouter.GET("/role/delete", system.RoleData)

		// 权限操作
		SystemRouter.GET("/menu/list", system.MenuSelect)
		SystemRouter.POST("/menu/add", system.MenuAdd)
		SystemRouter.POST("/menu/modify", system.RoleAdd)
		SystemRouter.GET("/menu/delete", system.RoleData)
	}

}
