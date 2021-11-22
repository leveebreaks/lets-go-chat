package main

import (
	"fmt"
	"github.com/leveebreaks/lets-go-chat/config"
	"github.com/leveebreaks/lets-go-chat/internal/repository"
	"github.com/leveebreaks/lets-go-chat/internal/server"
)

func main() {
	fmt.Println("here we go")
	settings := config.GetSettings()

	repository.InitAuthRepository(repository.NewMongoDbAuthRepo(settings.MongoDbUrl))

	server.Start(settings)
}
