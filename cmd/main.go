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

	gorillaHandlers "github.com/gorilla/handlers"
)

func main() {
	cfg := storage.PostgresConnConfig{
		DBHost:   "localhost",
		DBPort:   5432,
		DBName:   "avito",
		Username: "avito_admin",
		Password: "qwerty",
		Options:  nil,
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

	corsHandler := gorillaHandlers.CORS(
		gorillaHandlers.AllowedOrigins([]string{"*"}), // Разрешить все домены, но лучше указать конкретные
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		gorillaHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	fmt.Println("Запуск сервера на порту 8080 http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(mux)))
}
