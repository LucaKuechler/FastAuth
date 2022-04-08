package controller

import (
	"github.com/LucaKuechler/StrengthTracker/models"
	"github.com/LucaKuechler/StrengthTracker/services"
	"github.com/LucaKuechler/StrengthTracker/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Status() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Up, running!")
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			err := errors.NewBadRequestError("invalid user json")
			c.JSON(err.Status, err)
			return
		}

		result, err := services.CreateUser(user)

		if err != nil {
			c.JSON(err.Status, err)
			return
		}

		c.JSON(http.StatusOK, result)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
