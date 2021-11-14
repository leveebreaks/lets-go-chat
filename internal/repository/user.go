package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/leveebreaks/hasher"
)

type user struct {
	userName     string
	passwordHash string
	id           string
}

var users = make(map[string]user)

func CreateUser(userName string, password string) (string, error) {
	_, ok := users[userName]
	if ok {
		return "", errors.New("user with such name already exists")
	}

	passwordHash, err := hasher.HashPassword(password)
	if err != nil {
		return "", err
	}

	id := uuid.NewString()

	users[userName] = user{userName, passwordHash, id}

	return id, nil
}

func CheckUser(userName string, password string) bool {
	u, ok := users[userName]
	if !ok {
		return false
	}

	return hasher.CheckPasswordHash(password, u.passwordHash)
}
