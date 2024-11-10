package user

import (
	"odev-1/internal/core/domains/user/objects"
)

type User struct {
	Email    objects.Email
	Password objects.Password
}

func New(
	email objects.Email,
	password objects.Password,
) User {
	return User{
		Email:    email,
		Password: password,
	}
}
