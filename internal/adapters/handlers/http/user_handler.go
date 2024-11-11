package http

import (
	"encoding/json"
	userDomain "hexagonal-architecture-example/internal/core/domains/user"
	userService "hexagonal-architecture-example/internal/core/services/user"
	"net/http"
)

type UserHandler struct {
	service *userService.Service
}

func NewUserHandler(service *userService.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userList, _ := h.service.GetUserList(r.Context())

	err := json.NewEncoder(w).Encode(userList)
	if err != nil {
		return
	}
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
