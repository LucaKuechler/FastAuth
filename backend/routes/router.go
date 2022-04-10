package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/LucaKuechler/StrengthTracker/controllers"
)


func Setup(router *gin.Engine) {
	router.POST("/api/v1/register", controller.Register())	
	router.POST("/api/v1/login", controller.Login())	
	router.GET("/api/v1/users", controller.GetUsers())	
}
