package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/LucaKuechler/StrengthTracker/utils/errors"
	"github.com/LucaKuechler/StrengthTracker/utils/helper"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := helper.ExtractToken(c)
		if err != nil {
			c.JSON(err.Status, err)
			return
		}

	}
}
