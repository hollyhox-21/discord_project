package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/hollyhox-21/discord_project/user_profile_service/internal/app"
	pb "github.com/hollyhox-21/discord_project/user_profile_service/pkg/user_profile"
)

type Server struct {
	impl *app.Implementation

	grpc struct {
		lis    net.Listener
		server *grpc.Server
	}
	grpcGateway struct {
		lis    net.Listener
		server *http.Server
	}
	info struct {
		lis    net.Listener
		server *http.Server
	}

	publicCloser *Closer
}

func NewServer(ctx context.Context, cfg Config, impl *app.Implementation) (*Server, error) {
	srv := &Server{
		impl:         impl,
		publicCloser: New(syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL),
	}

	// Create gRPC server
	{
		grpcServerOptions := unaryInterceptorsToGrpcServerOptions(cfg.UnaryInterceptors...)
		grpcServerOptions = append(grpcServerOptions,
			grpc.ChainUnaryInterceptor(cfg.ChainUnaryInterceptors...),
		)

		grpcSrv := grpc.NewServer()
		pb.RegisterUserProfileServiceServer(grpcSrv, impl)
		reflection.Register(grpcSrv)

		srv.grpc.server = grpcSrv
	}

	// Create gRPC Gateway server
	{
		mux := runtime.NewServeMux()
		if err := pb.RegisterUserProfileServiceHandlerServer(ctx, mux, impl); err != nil {
			return nil, fmt.Errorf("server: failed to register handler: %v", err)
		}

		httpRouter := registerRouter(mux)

		httpSrv := &http.Server{
			Handler: httpRouter,
		}

		srv.grpcGateway.server = httpSrv
	}

	// Create info server
	{
		infoRouter := chi.NewRouter()

		swaggerRouter := initSwagger()
		infoRouter.Mount("/docs", swaggerRouter)

		infoSrv := &http.Server{
			Handler: infoRouter,
		}

		srv.info.server = infoSrv
	}

	// Create a listener on TCP ports
	listenerGrpcGateway, listenerGrpc, listenerInfo, err := createListeners(cfg.GRPCPort, cfg.GRPCGatewayPort, cfg.INFOport)
	if err != nil {
		return nil, fmt.Errorf("server: %v", err)
	}
	srv.grpc.lis = listenerGrpc
	srv.grpcGateway.lis = listenerGrpcGateway
	srv.info.lis = listenerInfo

	return srv, nil
}
