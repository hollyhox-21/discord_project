package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/hollyhox-21/discord_project/friendship_service/docs"
)

// @title			FriendshipService
// @version		0.0.1
// @description	Сервис управления взаимосвязями между пользователями
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
