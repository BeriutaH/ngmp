package system

import (
	"github.com/gin-gonic/gin"
	"ngmp/controllers/items"
)

// 关于主体项目相关的路由制定路由器

func ItemRouter(e *gin.Engine) {
	OperateRouter := e.Group("/operate")
	{
		OperateRouter.GET("/shop", items.ShopHello)
		OperateRouter.GET("/comment", items.ShopComment)
	}

}
