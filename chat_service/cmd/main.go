package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// @title			ChatService
// @version		0.0.1
// @description	Сервис отправки сообщений
// @BasePath		/api
func main() {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Get("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
