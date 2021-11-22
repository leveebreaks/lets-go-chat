package server

import (
	"fmt"
	"github.com/leveebreaks/lets-go-chat/config"
	"github.com/leveebreaks/lets-go-chat/internal/handlers"
	"net/http"
)

func Start(s config.Settings) {
	registerEndpoints()

	addr := fmt.Sprintf("%s:%s", s.ApiHost, s.ApiPort)
	fmt.Printf("Start listening on %s", addr)
	http.ListenAndServe(addr, nil)
}

func registerEndpoints() {
	http.HandleFunc("/v1/user", handlers.CreateUser)
	http.HandleFunc("/v1/user/login", handlers.LoginUser)
}
