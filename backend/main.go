package main

import (
	"github.com/LucaKuechler/StrengthTracker/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
    router.Use(gin.Logger(), gin.Recovery())

	routes.Setup(router)

	router.Run(":8080")
}
