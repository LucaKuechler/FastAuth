package helper

import (
	"github.com/LucaKuechler/StrengthTracker/utils/errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func ExtractToken(c *gin.Context) (string, *errors.ApiErr) {
	bearerToken := c.Request.Header.Get("Authorization")

	if bearerToken == "" {
		err := errors.NewBadRequestError("Authentification token is missing")
		return "", err
		// c.JSON(err.Status, err)
		// return
	}

	strArr := strings.Split(bearerToken, " ")

	if len(strArr) == 2 {
		return strArr[1], nil
	}
	return "", errors.NewBadRequestError("Token is not a bearer token.")
}
