package user

import (
	"context"
	"hexagonal-architecture-example/internal/core/domains/user"
	"hexagonal-architecture-example/internal/core/domains/user/objects"
	"hexagonal-architecture-example/internal/core/ports"
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
type GetListResponse struct {
	Data []user.List
}

type Service struct {
	userRepository ports.UserRepository
}

func (s *Service) GetUserList(ctx context.Context) (GetListResponse, error) {
	list, err := s.userRepository.GetList(ctx)
	return GetListResponse{Data: list}, err
}

func (s *Service) CreateAccount(ctx context.Context, req StoreRequest) (StoreResponse, error) {

	email := objects.Email(req.Email)

	generateNewUser := user.New(email, req.Password)

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
