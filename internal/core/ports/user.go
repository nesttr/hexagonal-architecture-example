package ports

import (
	"context"
	"odev-1/internal/core/domains/user"
)

type UserRepository interface {
	Store(ctx context.Context, user user.User) (int64, error)
}
