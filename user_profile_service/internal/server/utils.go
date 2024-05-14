package server

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func createListeners(GRPCPort string, GRPCGatewayPort string, INFOport string) (net.Listener, net.Listener, net.Listener, error) {
	listenerHttp, err := net.Listen("tcp", ":"+GRPCGatewayPort)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to listen: %v", err)
	}
	listenerGrpc, err := net.Listen("tcp", ":"+GRPCPort)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to listen: %v", err)
	}
	listenerInfo, err := net.Listen("tcp", ":"+INFOport)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to listen: %v", err)
	}
	return listenerHttp, listenerGrpc, listenerInfo, nil
}

func unaryInterceptorsToGrpcServerOptions(interceptors ...grpc.UnaryServerInterceptor) []grpc.ServerOption {
	opts := make([]grpc.ServerOption, 0, len(interceptors))
	for _, interceptor := range interceptors {
		opts = append(opts, grpc.UnaryInterceptor(interceptor))
	}
	return opts
}
