package objects

import (
	"errors"
	"net/mail"
)

type Email string

func NewEmail(email string) (Email, error) {
	ErrorEmailInvalid := errors.New("email is invalid")
	_, err := mail.ParseAddress(email)
	if err != nil {
		return "", ErrorEmailInvalid
	}
	return Email(email), nil

}
