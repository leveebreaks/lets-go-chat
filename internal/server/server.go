package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/leveebreaks/lets-go-chat/config"
	"github.com/leveebreaks/lets-go-chat/internal/handlers"
	"net/http"
)

func Start(s config.Settings) {
	addr := fmt.Sprintf("%s:%s", s.ApiHost, s.ApiPort)
	fmt.Printf("Start listening on %s", addr)
	r := getRouter()
	http.ListenAndServe(addr, r)
}

func getRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use()
	r.HandleFunc("/v1/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/v1/user/login", handlers.LoginUser).Methods("POST")

	return r
}
