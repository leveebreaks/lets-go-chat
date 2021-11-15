package service

import (
	"github.com/google/uuid"
	"github.com/leveebreaks/lets-go-chat/internal/repository"
)

type AuthService struct {
	repository.AuthRepository
}

// CreateUser ...
func (service *AuthService) CreateUser(userName string, password string) (string, error) {
	id, err := service.AuthRepository.CreateUser(userName, password)
	return id, err
}

// LoginUser error equal
func (service *AuthService) LoginUser(userName string, password string) (string, bool) {
	ok := service.CheckUser(userName, password)
	var token string
	if ok {
		token = uuid.NewString()
	}
	return token, ok
}
