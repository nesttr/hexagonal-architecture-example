package http

import (
	"encoding/json"
	"net/http"
	userDomain "odev-1/internal/core/domains/user"
	userService "odev-1/internal/core/services/user"
)

type UserHandler struct {
	service *userService.Service
}

func NewUserHandler(service *userService.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) StoreUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user userDomain.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	storeRequest := userService.StoreRequest{
		Email:    string(user.Email),
		Password: string(user.Password),
	}
	storeResponse, err := h.service.CreateAccount(r.Context(), storeRequest)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(storeResponse)
	if err != nil {
		return
	}
}
