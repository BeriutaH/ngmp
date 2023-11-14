package main

import (
	"fmt"
	"ngmp/model"
	"ngmp/router"
)

func main() {

	//加载多个app的路由配置
	router.Include(userdata.Routers, itemcode.Routers)

	// 创建数据库
	model.Init()
	// 初始化路由
	r := router.Init()

	if err := r.Run(":8090"); err != nil {
		fmt.Printf("error message:%v\n", err)
	}
}
