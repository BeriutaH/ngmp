package system

import (
	"github.com/gin-gonic/gin"
	"ngmp/controllers"
	middleware "ngmp/middlewares"
)

// CommonRouter 公共路由
func CommonRouter(e *gin.RouterGroup) {

	// 登录登出
	e.POST("/login", controllers.LoginFunc)
	e.POST("/logout", controllers.LogoutFunc).Use(middleware.TokenAuth())

}
