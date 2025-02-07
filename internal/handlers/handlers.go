package handlers

import (
	"avito_test/internal/auth"
	"avito_test/internal/core"
	"avito_test/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandler struct {
	service core.Interface
}

func NewUserHandler(service core.Interface) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Невалидные данные", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	tokenString, err := h.service.Register(r.Context(), &newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	response := auth.RegisterResponse{AccessToken: tokenString}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Невалидные данные", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	tokenString, err := h.service.Login(r.Context(), &newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	response := auth.RegisterResponse{AccessToken: tokenString}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) FlatCreate(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) HouseCreate(w http.ResponseWriter, r *http.Request) {
	var newHouse models.House

	// Декодируем JSON из тела запроса
	if err := json.NewDecoder(r.Body).Decode(&newHouse); err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}
	house, _ := h.service.HouseCreate(r.Context(), &newHouse)
	fmt.Println(house)

}
