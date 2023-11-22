package system

import (
	"github.com/gin-gonic/gin"
	"ngmp/controllers"
	"ngmp/controllers/system"
)

// SystemRouter 初始化管理 配置路由信息
func SystemRouter(e *gin.RouterGroup) {

	SystemRouter := e.Group("/system")
	{
		// 角色操作
		SystemRouter.POST("/role/list", controllers.RoleData)
		SystemRouter.POST("/role/add", controllers.RoleAdd)
		SystemRouter.POST("/role/modify", controllers.RoleAdd)
		SystemRouter.GET("/role/delete", controllers.RoleData)

		// 权限操作
		SystemRouter.POST("/menu/list", controllers.UserData)
		SystemRouter.POST("/menu/add", system.MenuAdd)
		SystemRouter.POST("/menu/modify", controllers.RoleAdd)
		SystemRouter.GET("/menu/delete", controllers.RoleData)
	}

}
