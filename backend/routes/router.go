package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/LucaKuechler/StrengthTracker/controllers"
	"github.com/LucaKuechler/StrengthTracker/middlewares"
)


func Setup(router *gin.Engine) {
	router.POST("/api/v1/register", controller.Register())
	router.POST("/api/v1/login", controller.Login())

	authorized := router.Group("/")
	authorized.Use(middleware.CheckAuth())
	authorized.GET("/api/v1/users", controller.GetUsers())
}
