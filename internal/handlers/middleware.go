package handlers

import (
	"avito_test/internal/auth"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

type contextKey string

const userTypeKey contextKey = "userType"

func (h *UserHandler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Извлечение токена из заголовка Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Токен не предоставлен", http.StatusUnauthorized)
			return
		}

		// Проверка формата токена
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // Если токен не был найден
			http.Error(w, "Неверный формат токена", http.StatusUnauthorized)
			return
		}

		// Далее идет проверка токена...
		claims := &auth.TokenClaims{}
		jwtToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("неверный метод подписи")
			}
			return auth.SecretKey, nil
		})

		if err != nil || !jwtToken.Valid {
			http.Error(w, "Неверный токен", http.StatusUnauthorized)
			return
		}

		// Если токен валиден, добавляем информацию о пользователе в контекст
		ctx := context.WithValue(r.Context(), userTypeKey, claims.User.UserType)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (h *UserHandler) ModeratorOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Извлекаем информацию о пользователе из контекста
		userType := r.Context().Value(userTypeKey)
		if userType == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Приводим значение к строке
		userTypeStr, ok := userType.(string)
		if !ok || userTypeStr != "moderator" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Если пользователь модератор, продолжаем обработку запроса
		next.ServeHTTP(w, r)
	})
}
