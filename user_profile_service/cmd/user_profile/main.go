package main

import (
	"context"
	"log"

	"github.com/hollyhox-21/discord_project/user_profile_service/internal/app"
	"github.com/hollyhox-21/discord_project/user_profile_service/internal/config"
	server_v2 "github.com/hollyhox-21/discord_project/user_profile_service/internal/server_v2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.Parse(ctx)

	// init repo
	// init kafka

	// Create Implementation
	impl, err := app.NewImplementation()
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	srv, err := server_v2.NewServer(ctx, cfg.ServerConfig, impl)
	if err != nil {
		log.Fatalln(err)
	}

	srv.Run()
}
