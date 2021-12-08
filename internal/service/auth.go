package service

import (
	"github.com/google/uuid"
	"github.com/leveebreaks/lets-go-chat/internal/repository"
)

type Auth struct {
	repository.Auth
}

func NewAuth(repo repository.Auth) Auth {
	return Auth{repo}
}

// CreateUser ...
func (a *Auth) CreateUser(userName string, password string) (string, error) {
	id, err := a.CreateUser(userName, password)
	return id, err
}

// LoginUser error equal
func (a *Auth) LoginUser(userName string, password string) (string, bool) {
	ok := a.CheckUser(userName, password)
	var token string
	if ok {
		token = uuid.NewString()
	}
	return token, ok
}
