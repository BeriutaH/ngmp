package system

import (
	"github.com/gin-gonic/gin"
	"ngmp/controllers/system"
	middleware "ngmp/middlewares"
)

// AdminRouter 初始化管理 配置路由信息
func AdminRouter(e *gin.RouterGroup) {
	System := e.Group("/admin").Use(middleware.TokenAuth())
	// 角色操作
	System.GET("/role/list", system.RoleSelect)
	System.POST("/role/add", system.RoleAdd)
	System.PUT("/role/modify/:roleID", system.UpdateRole)
	System.DELETE("/role/delete/:roleID", system.DelRole)

	// 权限操作
	System.GET("/menu/list", system.MenuSelect)
	System.POST("/menu/add", system.MenuAdd)
	System.PUT("/menu/modify/:perID", system.UpdateMenu)
	System.DELETE("/menu/delete/:perID", system.DelMenu)

	// 用户操作
	System.GET("/user/list", system.UserData)
	System.POST("/user/add", system.UserAdd)
	System.PUT("/user/modify/:userID", system.UpdateUser)
	System.DELETE("/user/delete/:userID", system.DelUser)

}
