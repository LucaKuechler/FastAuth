package services

import (
	"github.com/LucaKuechler/StrengthTracker/models/DTO"
	"github.com/LucaKuechler/StrengthTracker/utils/errors"
	"github.com/LucaKuechler/StrengthTracker/utils/helper"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

func CreateAuthToken(authorizedUser *dto.AuthorizedUser) (*string, *errors.ApiErr) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(authorizedUser.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
	})

	secret := helper.GetJWTToken()

	token, err := claims.SignedString([]byte(secret))
	if err != nil {
		return nil, errors.NewInternalServerError("Login failed")
	}

	return &token, nil
}
