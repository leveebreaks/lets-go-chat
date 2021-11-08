package server

import (
	"fmt"
	"github.com/leveebreaks/lets-go-chat/internal/handlers"
	"net/http"
)

func Start() {
	registerEndpoints()
	fmt.Println("Start listening on localhost:8182")
	http.ListenAndServe(":8182", nil)
}

func registerEndpoints() {
	http.HandleFunc("/v1/user", handlers.CreateUser)
	http.HandleFunc("/v1/user/login", handlers.LoginUser)
}
