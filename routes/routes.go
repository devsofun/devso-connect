package routes

import (
	"myapp/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", controllers.Index)
	router.GET("/login", controllers.ShowLoginForm)
	router.POST("/login", controllers.ProcessLogin)
}
