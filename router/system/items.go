package system

import (
	"github.com/gin-gonic/gin"
	"ngmp/controllers/items"
)

// 关于主体项目相关的路由制定路由器

func ItemRouter(e *gin.RouterGroup) {
	e.GET("/shop", items.ShopHello)
	e.GET("/comment", items.ShopComment)
}
