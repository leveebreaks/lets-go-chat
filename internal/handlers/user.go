package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/leveebreaks/lets-go-chat/internal/service"
	"net/http"
)

type createUserRequest struct {
	userName string `json:"userName"`
	password string `json:"password"`
}

type createUserResponse struct {
	id       string `json:"id"`
	userName string `json:"userName"`
}

type loginUserRequest struct {
	userName string `json:"userName"`
	password string `json:"password"`
}

type loginUserResponse struct {
	url string `json:"url"`
}

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

	fmt.Println("Request:", req.userName, req.password)

	// validate name and pass first
	// ...

	id, err := service.CreateUser(req.userName, req.password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&createUserResponse{id, req.userName})
}

// LoginUser /user/login
func LoginUser(w http.ResponseWriter, h *http.Request) {

}
