package main

import (
	"Ankipage/db"
	"Ankipage/routes"
)

func main() {
	// 初始化数据库
	db.InitDB()
	// 设置路由
	r := routes.SetupRouter()

	// 启动服务
	r.Run(":8080")
}
