package routes

import (
	"avito_test/internal/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes(userHandler *handlers.UserHandler) *mux.Router {
	api := mux.NewRouter()
	api.Handle("/api/", api)

	api.HandleFunc("/api/register", userHandler.Register).Methods("POST")
	api.HandleFunc("/api/login", userHandler.Login).Methods("POST")
	return api
}
