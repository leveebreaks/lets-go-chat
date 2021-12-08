package handlers

import (
	"github.com/gorilla/websocket"
	"github.com/leveebreaks/lets-go-chat/internal/service"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func Echo(w http.ResponseWriter, r *http.Request) {
	token, ok := r.URL.Query()["token"]
	if !ok || len(token) == 0 {
		http.Error(w, "no token provided", http.StatusBadRequest)
	}
	if !service.RevokeToken(token[0]) {
		http.Error(w, "token is not valid", http.StatusBadRequest)
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
