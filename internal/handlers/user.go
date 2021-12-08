package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/leveebreaks/lets-go-chat/internal/service"
	"net/http"
)

type createUserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type createUserResponse struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
}

type loginUserRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type loginUserResponse struct {
	Url string `json:"url"`
}

type user struct {
	authService service.Auth
}

func NewUser(as service.Auth) user {
	return user{as}
}

// CreateUser /user
func (u *user) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	id, err := u.authService.CreateUser(req.UserName, req.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createUserResponse{id, req.UserName})
}

// LoginUser /user/login
func (u *user) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req loginUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	token, ok := u.authService.LoginUser(req.UserName, req.Password)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Invalid username/password", http.StatusNotFound)
		return
	}
	service.AddToken(token, req.UserName)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := loginUserResponse{Url: "/ws/echo?token=" + token}
	json.NewEncoder(w).Encode(resp)
}
