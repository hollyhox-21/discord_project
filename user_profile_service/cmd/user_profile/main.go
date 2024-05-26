package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"github.com/hollyhox-21/discord_project/user_profile_service/internal/app"
	"github.com/hollyhox-21/discord_project/user_profile_service/internal/server"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// init repo
	// init kafka

	// Create Implementation
	impl, err := app.NewImplementation()
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	// delivery
	config := server.Config{
		GRPCGatewayPort:        "8080",
		GRPCPort:               "8081",
		INFOport:               "8082",
		ChainUnaryInterceptors: []grpc.UnaryServerInterceptor{},
	}

	srv, err := server.NewServer(ctx, config, impl)
	if err != nil {
		log.Fatalln(err)
	}

	srv.Run()
}
