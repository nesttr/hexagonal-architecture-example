package http

import (
	"odev-1/internal/core/ports"
	"odev-1/internal/core/services/user"
)

type Services struct {
	UserService *user.Service
}

func NewServices(userRepository ports.UserRepository) *Services {
	return &Services{
		UserService: user.NewUserService(userRepository),
	}
}
