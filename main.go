package main

import (
	"myapp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	// 加载路由
	routes.SetupRoutes(router)

	// 启动服务
	router.Run(":8080")
}
