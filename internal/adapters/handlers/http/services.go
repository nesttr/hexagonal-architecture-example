package http

import (
	"hexagonal-architecture-example/internal/core/ports"
	"hexagonal-architecture-example/internal/core/services/user"
)

type Services struct {
	UserService *user.Service
}

func NewServices(userRepository ports.UserRepository) *Services {
	return &Services{
		UserService: user.NewUserService(userRepository),
	}
}
