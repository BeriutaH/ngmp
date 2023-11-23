package system

import (
	"github.com/gin-gonic/gin"
	"ngmp/controllers/system"
)

// UserRouter Routers user 配置路由信息
func UserRouter(e *gin.RouterGroup) {

	//Router.POST("/login/login", admin.Login)
	//Router.Use(middleware.TokenAuth()).POST("/login/logout", admin.Logout)
	// 用户操作
	UserRouter := e.Group("/user")
	{
		UserRouter.GET("/list", system.UserData)
		UserRouter.POST("/add", system.UserAdd)
		UserRouter.GET("/blog", system.BlogComment)
	}

}
