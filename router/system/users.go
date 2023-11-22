package system

import (
	"github.com/gin-gonic/gin"
	"ngmp/controllers"
)

// UserRouter Routers user 配置路由信息
func UserRouter(e *gin.RouterGroup) {

	// 登录登出
	e.POST("/login", controllers.LoginFunc)
	//Router.POST("/login/login", admin.Login)
	//Router.Use(middleware.TokenAuth()).POST("/login/logout", admin.Logout)
	// 用户操作
	UserRouter := e.Group("/user")
	{
		UserRouter.POST("/list", controllers.UserData)
		UserRouter.POST("/add", controllers.UserAdd)
		UserRouter.GET("/blog", controllers.BlogComment)
	}

}
