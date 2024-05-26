package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/swaggest/swgui/v5emb"
)

func registerRouter(mux *runtime.ServeMux) *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.AllowAll().Handler)

	initProbes(r)

	r.Mount("/", mux)

	return r
}

func initProbes(r *chi.Mux) {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Get("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
}

func initSwagger() http.Handler {
	r := chi.NewRouter()

	r.Mount("/", v5emb.NewHandler("API Definition",
		"/docs/user_profile.swagger.json",
		"/docs"))

	r.Get("/user_profile.swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pkg/user_profile/user_profile.swagger.json")
	})

	return r
}
