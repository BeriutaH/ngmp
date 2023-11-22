package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ngmp/utils"
)

func ShopHello(c *gin.Context) {
	// reData := api.ReturnMsgFunc(200, "success", "None")
	// fmt.Println("数据显示: %d\n", reData)
	c.JSON(http.StatusOK, utils.ReturnMsgFunc(200, "success", 0))
}

func ShopComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcom!!! 项目主体代码逻辑!!!",
	})
}
