package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/leveebreaks/lets-go-chat/internal/repository"
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

var authService = &service.AuthService{AuthRepository: repository.NewMongoDbAuthRepo("mongodb://localhost:27017")}

// CreateUser /user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req createUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// validate name and pass first
	// ...

	id, err := authService.CreateUser(req.UserName, req.Password)

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
func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req loginUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	token, ok := authService.LoginUser(req.UserName, req.Password)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Invalid username/password", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := loginUserResponse{Url: "url/chat?token=" + token}
	json.NewEncoder(w).Encode(resp)
}
