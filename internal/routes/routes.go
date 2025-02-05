package routes

import (
	"avito_test/internal/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes(userHandler *handlers.UserHandler) *mux.Router {
	api := mux.NewRouter()

	api.HandleFunc("/register", userHandler.Register).Methods("POST")
	api.HandleFunc("/login", userHandler.Login).Methods("POST")

	authHandler := mux.NewRouter()
	authHandler.Use(userHandler.AuthMiddleware)

	authHandler.HandleFunc("/flat/create", userHandler.FlatCreate).Methods("POST")

	moderatorHandler := mux.NewRouter()
	moderatorHandler.Use(userHandler.ModeratorOnly)

	moderatorHandler.HandleFunc("/house/create", userHandler.HouseCreate).Methods("POST")

	return api
}
