package services

import (
	"github.com/LucaKuechler/StrengthTracker/database"
	"github.com/LucaKuechler/StrengthTracker/models/DAO"
	"github.com/LucaKuechler/StrengthTracker/models/DTO"
	"github.com/LucaKuechler/StrengthTracker/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user dto.User) (*dto.AuthorizedUser, *errors.ApiErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.NewBadRequestError("Failed to encrypt password")
	}

	user.Password = string(hashedPassword[:])

	dbUser := dao.User{
		Password: user.Password,
		Email:    user.Email,
	}

	response := db.Client.Create(&dbUser)
	if response.Error != nil {
		return nil, errors.NewBadRequestError("Failed to create user")
	}

	authorizedUser := dto.AuthorizedUser{
		ID: dbUser.ID,
		Email: user.Email,
		Password: user.Password,
	}

	return &authorizedUser, nil
}

func GetUserByEmail(user dto.User) (*dao.User, *errors.ApiErr) {
	var dbUser dao.User
	db.Client.Where("email = ?", user.Email).First(&dbUser)

	if dbUser.ID == 0 {
		return nil, errors.NewBadRequestError("Failed to get user")
	}
	
	return &dbUser, nil
}

func VailidatePassword(loginUserRequest dto.User, dbUser dao.User) (*dto.AuthorizedUser, *errors.ApiErr) {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(dbUser.Password), []byte(loginUserRequest.Password)); err != nil {
		return nil, errors.NewBadRequestError("Failed to validate password")
	}

	authorizedUser := dto.AuthorizedUser{
		ID: dbUser.ID,
		Email: loginUserRequest.Email,
		Password: loginUserRequest.Password,
	}

	return &authorizedUser, nil
}
