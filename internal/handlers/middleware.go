package handlers

import (
	"avito_test/internal/auth"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type contextKey string

const userTypeKey contextKey = "userType"

func (h *UserHandler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("session_id")
		if err != nil || cookie == nil {
			http.Error(w, "Токен не предоставлен", http.StatusUnauthorized)
			return
		}

		tokenString := cookie.Value

		claims := &auth.TokenClaims{}
		jwt_token, err := jwt.ParseWithClaims(tokenString, claims, func(jwt_token *jwt.Token) (interface{}, error) {
			if _, ok := jwt_token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("неверный метод подписи")
			}
			return auth.SecretKey, nil
		})

		if err != nil || !jwt_token.Valid {
			http.Error(w, "Неверный токен", http.StatusUnauthorized)
			return
		}

		if claims.EXP < time.Now().Unix() {
			http.Error(w, "Токен истек", http.StatusUnauthorized)
			return
		}

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
