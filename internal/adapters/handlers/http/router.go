package http

import (
	"net/http"
)

func NewRouter(services *Services) *http.ServeMux {
	mux := http.NewServeMux()

	userHandler := NewUserHandler(services.UserService)

	mux.HandleFunc("/user", userHandler.StoreUser)

	return mux
}
