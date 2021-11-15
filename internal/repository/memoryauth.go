package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/leveebreaks/hasher"
)

type inMemoryAuthRepo struct {
	users map[string]user
}

type user struct {
	userName     string
	passwordHash string
	id           string
}

func NewInMemoryAuthRepo() *inMemoryAuthRepo {
	users := make(map[string]user)
	return &inMemoryAuthRepo{users}
}

func (repo *inMemoryAuthRepo) CreateUser(userName, password string) (string, error) {
	_, ok := repo.users[userName]
	if ok {
		return "", errors.New("user with such name already exists")
	}

	passwordHash, err := hasher.HashPassword(password)
	if err != nil {
		return "", err
	}

	id := uuid.NewString()

	repo.users[userName] = user{userName, passwordHash, id}

	return id, nil
}

func (repo *inMemoryAuthRepo) CheckUser(userName, password string) bool {
	u, ok := repo.users[userName]
	if !ok {
		return false
	}

	return hasher.CheckPasswordHash(password, u.passwordHash)
}
