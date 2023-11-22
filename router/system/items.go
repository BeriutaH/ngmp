package system

import (
	"github.com/gin-gonic/gin"
	"ngmp/controllers"
)

// 关于主体项目相关的路由制定路由器

func ItemRouters(e *gin.Engine) {
	e.GET("/shop", controllers.ShopHello)
	e.GET("/comment", controllers.ShopComment)
}
