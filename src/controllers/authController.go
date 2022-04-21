package controller

import (
	"github.com/LucaKuechler/StrengthTracker/models/DTO"
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
		var user dto.User

		if err := c.ShouldBindJSON(&user); err != nil {
			err := errors.NewBadRequestError("invalid user json")
			c.JSON(err.Status, err)
			return
		}

		authenticatedUser, err := services.CreateUser(user)

		if err != nil {
			c.JSON(err.Status, err)
			return
		}

		c.JSON(http.StatusOK, authenticatedUser)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user dto.User
		if err := c.ShouldBindJSON(&user); err != nil {
			err := errors.NewBadRequestError("invalid user json")
			c.JSON(err.Status, err)
			return
		}

		dbUser, err := services.GetUserByEmail(user)
		if err != nil {
			c.JSON(err.Status, err)
			return
		}

		authorizedUser, err := services.VailidatePassword(user, *dbUser)
		if err != nil {
			c.JSON(err.Status, err)
			return
		}
		
		token, err := services.CreateAuthToken(authorizedUser)
		if err != nil {
			c.JSON(err.Status, err)
			return
		}

		c.IndentedJSON(http.StatusOK, token)
	}
}
