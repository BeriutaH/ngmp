package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"ngmp/config"
	"ngmp/model"
	"ngmp/provider"
	"ngmp/utils"
	"ngmp/utils/response"
	"strings"
)

// LoginFunc 登录
func LoginFunc(c *gin.Context) {
	/*
		1. 校验参数
		2. 解析密码
		3. 生成token
		4. 存入redis，设置七天有效，键值对为token为键，UserID为值
		5. 把token返回，token仅仅是token，内部没有任何意义
	*/
	// 声明接收的变量
	type Login struct {
		// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	var login Login
	if err := c.ShouldBindJSON(&login); err != nil {
		response.ValidatorFailedJson(err, c)
		return
	}

	// 判断用户名密码是否正确,将密文密码跟key传入解密方法
	userModel := model.NewUser()
	userStruct, err := userModel.FindUserByName(login.Username)
	if err != nil {
		response.InvalidArgumentJSON("查找用户失败: "+err.Error(), c)
		return
	}
	secreteText, err := utils.DecryptByAes(userStruct.Password, userStruct.SecretCode)
	if err != nil {
		response.InvalidArgumentJSON("密码解密失败: "+err.Error(), c)
		return
	}
	if login.Password == secreteText {
		token := utils.TokenMd5()
		//把token存到redis里,设置七天时效
		err = provider.SetRedisKey(config.RedisDefault, token, userStruct.ID, 7)
		if err != nil {
			response.InvalidArgumentJSON("数据库操作失败: "+err.Error(), c)
			return
		}
		userInfo := map[string]interface{}{"identity": "Bearer " + token}
		response.SuccessJSON(userInfo, "", c)
		return
	} else {
		response.InvalidArgumentJSON("用户名或密码错误", c)
		return
	}
}

// LogoutFunc 登出
func LogoutFunc(c *gin.Context) {
	accessToken := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	log.Println(accessToken)
	userId, err := provider.GetRedisKey(config.RedisDefault, accessToken)
	if err != nil {
		response.InvalidArgumentJSON("当前平台无此用户: "+err.Error(), c)
		return
	}
	// 删除当前用户的登录状态
	err = provider.DelData(config.RedisDefault, accessToken)
	if err != nil {
		response.InvalidArgumentJSON("删除用户状态失败: "+err.Error(), c)
		return
	}
	resp := map[string]string{"id": userId}
	response.SuccessJSON(resp, "", c)

}
