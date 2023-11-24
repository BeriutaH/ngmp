package system

import (
	"github.com/gin-gonic/gin"
	"ngmp/controllers/system"
)

// AdminRouter 初始化管理 配置路由信息
func AdminRouter(e *gin.RouterGroup) {

	// 角色操作
	e.GET("/role/list", system.RoleSelect)
	e.POST("/role/add", system.RoleAdd)
	e.PUT("/role/modify/:roleID", system.UpdateRole)
	e.DELETE("/role/delete/:roleID", system.DelRole)

	// 权限操作
	e.GET("/menu/list", system.MenuSelect)
	e.POST("/menu/add", system.MenuAdd)
	e.POST("/menu/modify", system.RoleAdd)
	e.POST("/menu/delete", system.RoleSelect)

}
