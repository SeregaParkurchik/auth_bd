package core

import (
	"avito_test/internal/auth"
	"avito_test/internal/models"
	"avito_test/internal/storage"
	"context"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type Interface interface {
	Register(ctx context.Context, user *models.User) (string, error)
	Login(ctx context.Context, user *models.User) (string, error)
}

type service struct {
	storage storage.Interface
}

func New(storage storage.Interface) Interface {
	return &service{
		storage: storage,
	}
}

func (s *service) Register(ctx context.Context, user *models.User) (string, error) {
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return "", fmt.Errorf("не удалось хэшировать пароль: %w", err)
	}
	user.Password = hashedPassword

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, auth.GenerateTokenClaims(user))
	tokenString, err := jwtToken.SignedString(auth.SecretKey)
	if err != nil {
		return "", fmt.Errorf("не удалось создать токен")
	}
	user.Token = tokenString

	err = s.storage.Register(user)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *service) Login(ctx context.Context, user *models.User) (string, error) {
	foundUser, err := s.storage.Login(user)
	if err != nil {
		return "", err
	}

	if !auth.CheckPasswordHash(user.Password, foundUser.Password) {
		return "", fmt.Errorf("неверный пароль")
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, auth.GenerateTokenClaims(&foundUser))
	tokenString, err := jwtToken.SignedString(auth.SecretKey)
	if err != nil {
		return "", fmt.Errorf("не удалось создать токен")
	}

	err = s.storage.UpdateToken(foundUser.ID, tokenString)
	if err != nil {
		return "", fmt.Errorf("не удалось обновить токен")
	}
	return tokenString, nil
}
