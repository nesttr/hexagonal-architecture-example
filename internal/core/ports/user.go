package ports

import (
	"context"
	"hexagonal-architecture-example/internal/core/domains/user"
)

type UserRepository interface {
	Store(ctx context.Context, user user.User) (int64, error)
	CreateTable(ctx context.Context)
	GetList(ctx context.Context) ([]user.List, error)
}
