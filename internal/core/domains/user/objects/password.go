package objects

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

type Password string

var SaltReadError = errors.New("salt read error")

func NewPassword(password string) (Password, error) {
	salt := make([]byte, 16)
	hash := sha256.New()

	_, err := rand.Read(salt)
	if err != nil {
		return "", SaltReadError
	}

	saltString := base64.StdEncoding.EncodeToString(salt)
	hash.Write([]byte(password + saltString))
	return Password(hex.EncodeToString(hash.Sum(nil))), nil
}
