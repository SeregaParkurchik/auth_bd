package models

import "time"

type User struct {
	ID       int    `json:"user_id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	UserType string `json:"user_type" db:"user_type"` // "client" или "moderator"
	Token    string `json:"token" db:"token"`         // Токен для аутентификации
}

type House struct {
	ID        int       `json:"id"`
	Address   string    `json:"address"`
	Year      int       `json:"year"`
	Developer string    `json:"developer,omitempty"` // Указатель, чтобы можно было не указывать застройщика
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
