package main

import (
	"fmt"
	"ngmp/provider"
	"ngmp/router"
)

func main() {

	//加载多个app的路由配置
	//router.Include(system.SystemRouter, system.UserRouter, system.ItemRouter)

	// 创建数据库
	provider.Init()
	// 初始化路由
	//r := router.Init()
	server := router.Routers()

	if err := server.Run(":8090"); err != nil {
		fmt.Printf("error message:%v\n", err)
	}
	//// 服务配置
	//server := &http.Server{
	//	Addr:           ":" + viper.GetString("Server.Port"),
	//	Handler:        router,
	//	ReadTimeout:    30 * time.Second,
	//	WriteTimeout:   30 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
}
