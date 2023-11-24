package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginFunc 登录
func LoginFunc(c *gin.Context) {
	// 声明接收的变量
	type Login struct {
		// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
		User     string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	//var loginJson Login
	//if err := c.ShouldBindJSON(&loginJson); err != nil {
	//	c.JSON(http.StatusOK, utils.ReturnMsgFunc(400, "未携带正确信息!", 0))
	//	return
	//}
	// 判断用户名密码是否正确
	//user := model.User{}
	//// 将密码转换成md5
	//pwd := PassMd5(loginJson.Password)
	//db := config.DBDefault
	//db.Find(&user, db.Where("username = ? AND password = ?", loginJson.User, pwd))
	//if pwd == "" {
	//	c.JSON(http.StatusOK, utils.ReturnMsgFunc(400, "用户名或密码错误!", 0))
	//	return
	//} else {
	//	// 校验用户名和密码正确之后,生成token值
	//	token := TokenMd5()
	//	// 把token存到redis里,设置七天时效
	//	//userId := strconv.FormatInt(int64(user.ID), 10)
	//	//if err := model.Redis.Set(userId, token, 7).Err(); err != nil {
	//	//	c.JSON(http.StatusOK, api.ReturnMsgFunc(400, "redis错误!", 0))
	//	//	return
	//	//} else {
	//	//
	//	//}
	//	// 把用户名,用户权限,角色,token一并传入前端
	//	// 模拟权限数据
	//	var per = []map[string]interface{}{
	//		{"id": 1, "name": "create user"},
	//		{"id": 2, "name": "delete user"},
	//		{"id": 3, "name": "edit user"},
	//		{"id": 4, "name": "select user"},
	//		{"id": 5, "name": "create role"},
	//	}
	//	// [{"id": 1, "name": "create user"},{"id": 2, "name": "delete user"},{"id": 3, "name": "edit user"},{"id": 4, "name": "select user"},{"id": 5, "name": "create role"}]
	//	resData := map[string]interface{}{
	//		"username":   user.Username,
	//		"role":       user.RoleId,
	//		"token":      token,
	//		"permission": per,
	//	}
	//	c.JSON(http.StatusOK, utils.ReturnMsgFunc(200, "success", resData))
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功!!!!",
	})

}
