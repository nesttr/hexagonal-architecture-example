package user

import (
	"context"
	"database/sql"
	"odev-1/internal/core/domains/user"
	"odev-1/internal/core/ports"
	"os"
)

type Repository struct {
	db *sql.DB
}

func (userRepository Repository) Store(ctx context.Context, user user.User) (int64, error) {
	query, err := os.ReadFile("queries/user/create.sql")
	if err != nil {
		return 0, err
	}
	result, err := userRepository.db.ExecContext(ctx, string(query), user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func NewUserRepository(db *sql.DB) ports.UserRepository {
	return &Repository{db: db}
}
