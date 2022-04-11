package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("UserID")
		c.JSON(http.StatusOK, user)
	}
}
