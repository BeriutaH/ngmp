package middleware

import (
	"ngmp/config"
	"ngmp/model"
	"ngmp/provider"
	"ngmp/utils/response"

	"strings"

	"github.com/gin-gonic/gin"
)

// TokenAuth Token授权
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if accessToken == "" {
			response.InvalidArgumentJSON("未携带token，请先登录", c)
			c.Abort()
			return
		}

		// 从redis中Token
		userId, err := provider.GetRedisKey(config.RedisDefault, accessToken)
		if err != nil {
			response.InvalidArgumentJSON("token已过期，请重新登录", c)
			c.Abort()
			return
		}
		// 查找当前用户，并把用户存入
		userStruct, err := model.NewUser().FindUserById(userId)
		if err != nil {
			response.InvalidArgumentJSON("无此用户，请重新登录", c)
			c.Abort()
			return
		}
		// 将ID合并到请求中（注意：请求参数不要使用到该参数）
		c.Set("identity", userStruct)
		c.Next()
		return
	}
}

// GetTokenAuthInfo 获取令牌信息
func GetTokenAuthInfo(c *gin.Context) *model.User {
	user, exist := c.Get("identity")
	if !exist {
		return model.NewUser()
	}
	if u, ok := user.(*model.User); ok {
		return u
	} else {
		return model.NewUser()
	}
}
