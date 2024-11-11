package user

import (
	"context"
	"database/sql"
	"fmt"
	"hexagonal-architecture-example/internal/core/domains/user"
	"hexagonal-architecture-example/internal/core/ports"
	"os"
)

type Repository struct {
	db *sql.DB
}

func (userRepository Repository) GetList(ctx context.Context) ([]user.List, error) {
	query, err := os.ReadFile("./internal/adapters/repositories/queries/get_list.sql")
	if err != nil {
		return nil, err
	}
	rows, err := userRepository.db.QueryContext(ctx, string(query))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []user.List
	for rows.Next() {
		var u user.List
		err := rows.Scan(&u.Id, &u.Email)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (userRepository Repository) Store(ctx context.Context, user user.User) (int64, error) {
	query, err := os.ReadFile("./internal/adapters/repositories/queries/store.sql")
	if err != nil {
		return 0, err
	}
	var lastInsertId int64
	err = userRepository.db.QueryRowContext(ctx, string(query), user.Email, user.Password).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func (userRepository Repository) CreateTable(ctx context.Context) {
	query, err := os.ReadFile("./internal/adapters/repositories/queries/create_table.sql")
	if err != nil {
		return
	}
	_, err = userRepository.db.ExecContext(ctx, string(query))
	if err != nil {
		return
	}
}

func NewUserRepository(db *sql.DB) ports.UserRepository {
	return &Repository{db: db}
}
