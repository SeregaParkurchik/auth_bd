package main

import (
	"avito_test/internal/core"
	"avito_test/internal/handlers"
	"avito_test/internal/routes"
	"avito_test/internal/storage"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := storage.PostgresConnConfig{
		DBHost:   "localhost", //localhost, host.docker.internal
		DBPort:   5432,
		DBName:   "avito",
		Username: "avito_admin",
		Password: "qwerty",
		Options:  nil, // или добавьте опции, если необходимо
	}

	// Создание соединения с базой данных
	conn, err := storage.New(context.Background(), cfg)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer conn.Close(context.Background())

	avitoDB := storage.NewAvitoDB(conn)
	authService := core.New(avitoDB)
	userHandler := handlers.NewUserHandler(authService)

	mux := routes.InitRoutes(userHandler)
	fmt.Println("Запуск сервера на порту 8080 http://localhost:8080/")
	http.ListenAndServe(":8080", mux)
}
