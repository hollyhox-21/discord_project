package server

import "google.golang.org/grpc"

// Config - server config
type Config struct {
	GRPCGatewayPort string
	GRPCPort        string
	INFOport        string

	ChainUnaryInterceptors []grpc.UnaryServerInterceptor
	UnaryInterceptors      []grpc.UnaryServerInterceptor
}
