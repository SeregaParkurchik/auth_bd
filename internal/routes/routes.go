package routes

import (
	"avito_test/internal/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes(userHandler *handlers.UserHandler) *mux.Router {
	api := mux.NewRouter()

	// Открытые маршруты
	api.HandleFunc("/register", userHandler.Register).Methods("POST")
	api.HandleFunc("/login", userHandler.Login).Methods("POST")

	// Защищенные маршруты для авторизованных пользователей
	authHandler := api.PathPrefix("/").Subrouter()
	authHandler.Use(userHandler.AuthMiddleware)
	authHandler.HandleFunc("/flat/create", userHandler.FlatCreate).Methods("POST")

	// Защищенные маршруты только для модераторов
	moderatorHandler := api.PathPrefix("/").Subrouter()
	moderatorHandler.Use(userHandler.AuthMiddleware)
	moderatorHandler.Use(userHandler.ModeratorOnly)
	moderatorHandler.HandleFunc("/house/create", userHandler.HouseCreate).Methods("POST")

	return api
}
