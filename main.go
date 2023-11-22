package main

import (
	"fmt"
	"ngmp/provider"
	"ngmp/router"
	"ngmp/router/system"
)

func main() {

	//加载多个app的路由配置
	router.Include(system.UserRouters, system.ItemRouters)

	// 创建数据库
	provider.Init()
	// 初始化路由
	r := router.Init()

	if err := r.Run(":8090"); err != nil {
		fmt.Printf("error message:%v\n", err)
	}
}
