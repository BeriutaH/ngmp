package system

import (
	"github.com/gin-gonic/gin"
	"ngmp/controllers"
)

// UserRouters Routers user 配置路由信息
func UserRouters(e *gin.Engine) {
	e.POST("/login", controllers.LoginFunc)
	e.POST("/userdata", controllers.UserData)
	e.POST("/useradd", controllers.UserAdd)
	e.POST("/roleadd", controllers.RoleAdd)
	e.GET("/roledata", controllers.RoleData)
	e.GET("/blog", controllers.BlogComment)

}
