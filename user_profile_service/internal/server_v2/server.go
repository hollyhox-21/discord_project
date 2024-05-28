// Package server содержит транспортный уровень запуск и остановка сервера
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
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/hollyhox-21/discord_project/user_profile_service/internal/server_v2/config"
	mwgrpc "github.com/hollyhox-21/discord_project/user_profile_service/internal/server_v2/middlewares/grpc"
)

type ImplementationRegister interface {
	Register() (desc *grpc.ServiceDesc, impl any)
}

type ImplementationRegisterMux interface {
	RegisterMux(ctx context.Context, mux *runtime.ServeMux) error
}

type ImplementationDescription interface {
	ImplementationRegister
	ImplementationRegisterMux
}

type Server struct {
	impls []ImplementationDescription

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

	resourceCloser *closer.Closer
}

func NewServer(ctx context.Context, cfg config.Config, impls ...ImplementationDescription) (*Server, error) {
	srv := &Server{
		impls:          impls,
		resourceCloser: closer.New(syscall.SIGINT, syscall.SIGTERM),
	}

	implGRPC := lo.Map(impls, func(item ImplementationDescription, _ int) ImplementationRegister {
		return ImplementationRegister(item)
	})

	// Create gRPC server
	srv.grpc.server = srv.initGRPCServer(cfg.GRPCServer, implGRPC...)

	// Create info server
	srv.info.server = srv.initInfoServer(cfg.INFOServer)

	implHTTP := lo.Map(impls, func(item ImplementationDescription, _ int) ImplementationRegisterMux {
		return ImplementationRegisterMux(item)
	})
	// Create gRPC Gateway server
	server, err := srv.initGRPCGatewayServer(ctx, cfg.GRPCGatewayServer, implHTTP...)
	if err != nil {
		return nil, fmt.Errorf("server: %v", err)
	}
	srv.grpcGateway.server = server

	// Create a listener on TCP ports
	srv.grpcGateway.lis, srv.grpc.lis, srv.info.lis, err = createListeners(
		cfg.GRPCServer.Addr,
		cfg.GRPCGatewayServer.Addr,
		cfg.INFOServer.Addr,
	)
	if err != nil {
		return nil, fmt.Errorf("server: %v", err)
	}

	return srv, nil
}

func (s *Server) initGRPCServer(cfg config.GRPCServer, impls ...ImplementationRegister) *grpc.Server {
	grpcServerOptions := unaryInterceptorsToGrpcServerOptions([]grpc.UnaryServerInterceptor{}...)
	grpcServerOptions = append(grpcServerOptions,
		grpc.ChainUnaryInterceptor([]grpc.UnaryServerInterceptor{
			mwgrpc.ContextErrorUnaryServerInterceptor(),
			mwgrpc.Recover,
			mwgrpc.HandleError,
		}...),
	)

	grpcSrv := grpc.NewServer(grpcServerOptions...)
	registerService := grpc.ServiceRegistrar(grpcSrv)

	for _, r := range impls {
		registerService.RegisterService(r.Register())
	}

	reflection.Register(grpcSrv)

	return grpcSrv
}

func (s *Server) initGRPCGatewayServer(ctx context.Context, cfg config.GRPCGatewayServer, impl ...ImplementationRegisterMux) (*http.Server, error) {
	mux := runtime.NewServeMux()

	for _, registerMux := range impl {
		if err := registerMux.RegisterMux(ctx, mux); err != nil {
			return nil, fmt.Errorf("init grpc gateway server: failed to register handler: %v", err)
		}
	}

	httpRouter := registerRouter(mux)

	httpSrv := &http.Server{
		Handler: httpRouter,
	}

	return httpSrv, nil
}

func (s *Server) initInfoServer(cfg config.INFOServer) *http.Server {
	infoRouter := chi.NewRouter()

	swaggerRouter := initSwagger()
	infoRouter.Mount("/docs", swaggerRouter)

	infoSrv := &http.Server{
		Handler: infoRouter,
	}

	return infoSrv
}
