package server

import "google.golang.org/grpc"

// Config - server config
type Config struct {
	GRPCServer        GRPCServer
	GRPCGatewayServer GRPCGatewayServer
	INFOServer        INFOServer
}

type GRPCServer struct {
	Port string

	ChainUnaryInterceptors []grpc.UnaryServerInterceptor
	UnaryInterceptors      []grpc.UnaryServerInterceptor
}

type GRPCGatewayServer struct {
	Port string
}

type INFOServer struct {
	Port string
}
