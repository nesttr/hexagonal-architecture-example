package user

import (
	"hexagonal-architecture-example/internal/core/domains/user/objects"
)

type User struct {
	Email    objects.Email
	Password objects.Password
}
type List struct {
	Id    int64
	Email objects.Email
}

func New(
	email objects.Email,
	password string,
) User {
	pass, _ := objects.NewPassword(password)
	return User{
		Email:    email,
		Password: pass,
	}
}
