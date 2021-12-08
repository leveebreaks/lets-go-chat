package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

func LogEndPointCalls(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logText := fmt.Sprintf("[request] RemoteAddr: %s, method: %s, RequestURI: %s", r.RemoteAddr, r.Method, r.RequestURI)
		log.Println(logText)
		nextHandler.ServeHTTP(w, r)
	})
}
