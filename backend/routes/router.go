package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/LucaKuechler/StrengthTracker/controller"    
)


func Setup(router *gin.Engine) {
	router.GET("/status", controller.Status())
	router.POST("/register", controller.Register())	
	router.POST("/login", controller.Login())	
}
