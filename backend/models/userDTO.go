package models

import (
	"github.com/LucaKuechler/StrengthTracker/utils/errors"
	"strings"
)

type User struct {
    ID uint64 `json:"id"`
    Password string `json:"password"`
    Email string `json:"email"`
}

func (user *User) Validate() *errors.ApiErr {
	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}
