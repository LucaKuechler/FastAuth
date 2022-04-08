package services

import (
	"github.com/LucaKuechler/StrengthTracker/models"
	"github.com/LucaKuechler/StrengthTracker/utils/errors"
	"github.com/LucaKuechler/StrengthTracker/entity"
	"github.com/LucaKuechler/StrengthTracker/datasource"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) (*models.User, *errors.ApiErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.NewBadRequestError("Failed to encrypt password")
	}

	user.Password = string(hashedPassword[:])

	
	dbUser := entity.User{
		Password: user.Password,
		Email:    user.Email,
	}

	response := db.Client.Create(&dbUser); if response.Error != nil {
		return nil, errors.NewBadRequestError("Failed to create user")
	}

	user.ID = dbUser.ID
	return &user, nil
}
