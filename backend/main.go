package main

import (
    "github.com/gin-gonic/gin"
    "github.com/LucaKuechler/StrengthTracker/routes"    
)


func main() {
	router := gin.Default()

    routes.Setup(router)

	router.Run(":8080")
}
