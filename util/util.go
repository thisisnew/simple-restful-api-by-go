package util

import (
	"log"
	"net/http"
)

func HttpHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.RemoteAddr, " ", r.Proto, " ", r.Method, " ", r.URL)
		handler.ServeHTTP(w, r)
	})
}
