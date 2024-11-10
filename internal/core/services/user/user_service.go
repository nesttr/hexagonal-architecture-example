package user

import (
	"context"
	"odev-1/internal/core/domains/user"
	"odev-1/internal/core/domains/user/objects"
	"odev-1/internal/core/ports"
)

type StoreRequest struct {
	Email    string
	Password string
}

type StoreResponse struct {
	UserID   int64
	Email    string
	Password string
}

type Service struct {
	userRepository ports.UserRepository
}

func (s *Service) CreateAccount(ctx context.Context, req StoreRequest) (StoreResponse, error) {

	email := objects.Email(req.Email)
	password := objects.Password(req.Password)

	generateNewUser := user.New(email, password)

	lastInsertId, err := s.userRepository.Store(ctx, generateNewUser)
	if err != nil {
		return StoreResponse{}, err
	}
	return StoreResponse{
		UserID: lastInsertId,
		Email:  string(email),
	}, nil
}

func NewUserService(repository ports.UserRepository) *Service {
	return &Service{
		userRepository: repository,
	}
}
