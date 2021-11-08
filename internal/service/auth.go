package service

import (
	"github.com/google/uuid"
	"github.com/leveebreaks/lets-go-chat/internal/repository"
)

func CreateUser(userName string, password string) (string, error) {
	id, err := repository.CreateUser(userName, password)
	return id, err
}

func LoginUser(userName string, password string) (string, bool) {
	ok := repository.CheckUser(userName, password)
	var token string
	if ok {
		token = uuid.NewString()
	}
	return token, ok
}
