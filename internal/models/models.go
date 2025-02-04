package models

type User struct {
	ID       int    `json:"user_id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	UserType string `json:"user_type" db:"user_type"` // "client" или "moderator"
	Token    string `json:"token" db:"token"`         // Токен для аутентификации
}
