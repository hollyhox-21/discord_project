package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hollyhox-21/discord_project/libraries/closer"
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

	publicCloser *closer.Closer
}

func NewServer(ctx context.Context, cfg Config, impl *app.Implementation) (*Server, error) {
	srv := &Server{
		impl:         impl,
		publicCloser: closer.New(syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL),
	}

	// Create gRPC server
	srv.grpc.server = initGRPCServer(cfg.GRPCServer, impl)

	// Create info server
	srv.info.server = initInfoServer(cfg.INFOServer)

	// Create gRPC Gateway server
	server, err := initGRPCGatewayServer(ctx, cfg.GRPCGatewayServer, impl)
	if err != nil {
		return nil, fmt.Errorf("server: %v", err)
	}
	srv.grpcGateway.server = server

	// Create a listener on TCP ports
	srv.grpc.lis, srv.grpcGateway.lis, srv.info.lis, err = createListeners(
		cfg.GRPCServer.Port,
		cfg.GRPCGatewayServer.Port,
		cfg.INFOServer.Port,
	)
	if err != nil {
		return nil, fmt.Errorf("server: %v", err)
	}

	return srv, nil
}

func initGRPCServer(cfg GRPCServer, impl *app.Implementation) *grpc.Server {
	grpcServerOptions := unaryInterceptorsToGrpcServerOptions(cfg.UnaryInterceptors...)
	grpcServerOptions = append(grpcServerOptions,
		grpc.ChainUnaryInterceptor(cfg.ChainUnaryInterceptors...),
	)

	grpcSrv := grpc.NewServer(grpcServerOptions...)
	pb.RegisterUserProfileServiceServer(grpcSrv, impl)

	reflection.Register(grpcSrv)

	return grpcSrv
}

func initGRPCGatewayServer(ctx context.Context, cfg GRPCGatewayServer, impl *app.Implementation) (*http.Server, error) {
	mux := runtime.NewServeMux()
	if err := pb.RegisterUserProfileServiceHandlerServer(ctx, mux, impl); err != nil {
		return nil, fmt.Errorf("init grpc gateway server: failed to register handler: %v", err)
	}

	httpRouter := registerRouter(mux)

	httpSrv := &http.Server{
		Handler: httpRouter,
	}

	return httpSrv, nil
}

func initInfoServer(cfg INFOServer) *http.Server {
	infoRouter := chi.NewRouter()

	swaggerRouter := initSwagger()
	infoRouter.Mount("/docs", swaggerRouter)

	infoSrv := &http.Server{
		Handler: infoRouter,
	}

	return infoSrv
}
