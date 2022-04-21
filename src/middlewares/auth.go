package middleware

import (
	"github.com/LucaKuechler/StrengthTracker/utils/errors"
	"github.com/LucaKuechler/StrengthTracker/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := helper.ExtractToken(c)
		if err != nil {
			c.Abort()
			c.JSON(err.Status, err)
			return
		}

		token, aerr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(helper.GetJWTToken()), nil
		})

		if aerr != nil {
			c.Abort()
			err := errors.NewBadRequestError("Token is invalid")
			c.JSON(err.Status, err)
			return
		}

		props, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Abort()
			err := errors.NewBadRequestError("Token can not be claimed")
			c.JSON(err.Status, err)
			return
		}

		c.Set("UserID", props["iss"])
	}
}
