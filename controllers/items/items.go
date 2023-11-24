package items

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShopHello(c *gin.Context) {
	// reData := api.ReturnMsgFunc(200, "success", "None")
	// fmt.Println("数据显示: %d\n", reData)
	//c.JSON(http.StatusOK, utils.ReturnMsgFunc(200, "success", 0))
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcom!!! 项目主体代码逻辑!!!",
	})
}

func ShopComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcom!!! 项目主体代码逻辑!!!",
	})
}
