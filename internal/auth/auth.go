package auth

import (
	"avito_test/internal/models"
	"errors"
	"time"
)

type TokenClaims struct {
	User struct {
		Email    string `json:"email"`
		UserType string `json:"user_type"`
	} `json:"user"`
	IAT int64 `json:"iat"`
	EXP int64 `json:"exp"`
}

type RegisterResponse struct {
	AccessToken string `json:"token"`
}

func GenerateTokenClaims(user *models.User) *TokenClaims {
	email := user.Email
	userType := user.UserType

	newTokenClaims := &TokenClaims{
		User: struct {
			Email    string `json:"email"`
			UserType string `json:"user_type"`
		}{
			Email:    email,
			UserType: userType,
		},
		IAT: time.Now().Unix(),
		EXP: time.Now().Add(time.Hour * 12).Unix(),
	}

	return newTokenClaims
}

func (c *TokenClaims) Valid() error {
	currentTime := time.Now().Unix()

	if c.EXP < currentTime {
		return errors.New("токен истек")
	}

	return nil
}

var SecretKey = []byte("mykey")
