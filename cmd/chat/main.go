package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/leveebreaks/lets-go-chat/config"
	"github.com/leveebreaks/lets-go-chat/internal/handlers"
	"github.com/leveebreaks/lets-go-chat/internal/repository"
	"github.com/leveebreaks/lets-go-chat/internal/service"
	"github.com/leveebreaks/lets-go-chat/pkg/middlewares"
	"github.com/urfave/negroni"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func main() {
	settings := config.GetSettings()

	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(settings.MongoDbUrl))
	defer client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	db := client.Database("auth")

	authService := service.NewAuth(repository.NewMongoDbAuthRepo(db))
	hUser := handlers.NewUser(authService)

	n := negroni.New()
	n.Use(negroni.NewRecovery())

	r := mux.NewRouter()
	r.Use(middlewares.LogEndPointCalls)
	r.HandleFunc("/v1/user", hUser.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/v1/user/login", hUser.LoginUser).Methods(http.MethodPost)
	n.UseHandler(r)

	addr := fmt.Sprintf("%s:%s", settings.ApiHost, settings.ApiPort)
	fmt.Printf("Start listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, n))
}
