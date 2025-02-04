package storage

import (
	"avito_test/internal/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Interface interface {
	Register(user *models.User) error
	Login(user *models.User) (models.User, error)
	UpdateToken(id int, token string) error
}

type PostgresConnConfig struct {
	DBHost   string
	DBPort   uint
	DBName   string
	Username string
	Password string
	Options  map[string]string
}

type AvitoDB struct {
	conn *pgx.Conn
}

func NewAvitoDB(conn *pgx.Conn) *AvitoDB {
	return &AvitoDB{conn: conn}
}

func New(ctx context.Context, cfg PostgresConnConfig) (*pgx.Conn, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.Username, cfg.Password, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	var options string
	if len(cfg.Options) > 0 {
		for key, value := range cfg.Options {
			options += fmt.Sprintf("%s=%s&", key, value)
		}

		options = options[:len(options)-1]
		connStr += options
	}

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to postgres: %w", err)
	}

	return conn, nil
}

func (s *AvitoDB) Register(user *models.User) error {
	// Проверка существования пользователя
	var exists bool
	err := s.conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", user.Email).Scan(&exists)
	if err != nil {
		return fmt.Errorf("ошибка при проверке существования пользователя: %w", err)
	}

	// Если пользователь существует, возвращаем ошибку
	if exists {
		return fmt.Errorf("пользователь с именем %s уже существует", user.Email)
	}

	// Используем уже существующее соединение для вставки нового пользователя
	err = s.conn.QueryRow(context.Background(), "INSERT INTO users (email, password,user_type,token) VALUES ($1, $2, $3, $4) RETURNING id", user.Email, user.Password, user.UserType, user.Token).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("ошибка при вставке результата в таблицу: %w", err)
	}

	return nil
}

func (s *AvitoDB) Login(user *models.User) (models.User, error) {
	var foundUser models.User

	err := s.conn.QueryRow(context.Background(), "SELECT id, email, password, user_type FROM users WHERE email = $1", user.Email).Scan(&foundUser.ID, &foundUser.Email, &foundUser.Password, &foundUser.UserType)
	if err != nil {
		return foundUser, fmt.Errorf("пользователь не найден: %w", err)
	}

	return foundUser, nil
}

func (s *AvitoDB) UpdateToken(id int, token string) error {
	_, err := s.conn.Exec(context.Background(), "UPDATE users SET token = $1 WHERE id = $2", token, id)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении токена: %w", err)
	}
	return nil
}
