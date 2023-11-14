package router

import "github.com/gin-gonic/gin"

// Option 根据定义include函数用来注册子app中定义的路由,init函数用来进行路由的初始化操作
type Option func(*gin.Engine)

var options []Option

// Include 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化

func Init() *gin.Engine {
	r := gin.New()
	for _, opt := range options {
		opt(r)
	}
	return r
}
