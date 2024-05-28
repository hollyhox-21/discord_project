package server

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-openapi/spec"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hollyhox-21/discord_project/libraries/logger"

	mwhttp "github.com/hollyhox-21/discord_project/user_profile_service/internal/server_v2/middlewares/http"
	"github.com/hollyhox-21/discord_project/user_profile_service/pkg"
	"github.com/hollyhox-21/discord_project/user_profile_service/pkg/dist"
)

func registerRouter(mux *runtime.ServeMux) *chi.Mux {
	r := chi.NewRouter()

	r.Use(mwhttp.Recover)
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

	fs := http.FileServerFS(dist.Dist)
	r.Handle("/*", http.StripPrefix("/docs/", fs))

	r.Get("/router.swagger.json", func(w http.ResponseWriter, r *http.Request) {

		f, err := pkg.SwaggerFile.Open("user_profile/user_profile.swagger.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		data, err := io.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var s spec.Swagger

		if err := json.Unmarshal(data, &s); err != nil {
			logger.Errorf(context.Background(), "failed to unmarshal swaggerDef: %v", err)
			return
		}

		s.Host = "blabla:8081"

		def, err := json.Marshal(s)
		if err != nil {
			logger.Errorf(context.Background(), "failed to marshal swaggerDef: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(def)
	})

	r.Get("/pet.swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, pkg.SwaggerFile, "router/pet.swagger.json")
	})

	return r
}
