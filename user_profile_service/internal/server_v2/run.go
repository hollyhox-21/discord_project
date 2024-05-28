package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/hollyhox-21/discord_project/libraries/closer"
	"github.com/hollyhox-21/discord_project/libraries/logger"
)

const (
	GracefulDelay   = 3 * time.Second
	GracefulTimeout = 10 * time.Second
)

func (s *Server) Run() {
	s.runGrpcGateway()
	s.runGRPC()
	s.runInfo()

	// Wait signal and close all server resources
	s.resourceCloser.Wait()
	// Close all other resources from globalCloser
	closer.CloseAll()

	_ = logger.Logger().Sync()
}

func (s *Server) runGRPC() {
	if s.grpc.server != nil {
		go func() {
			logger.Infof(context.Background(), "start serve gRPC to %s", s.grpc.lis.Addr())
			if err := s.grpc.server.Serve(s.grpc.lis); err != nil {
				logger.Infof(context.Background(), fmt.Errorf("grpc: %w", err).Error())
				s.resourceCloser.CloseAll()
			}
		}()
		s.resourceCloser.Add(func() error {
			ctx, cancel := context.WithTimeout(context.Background(), GracefulTimeout)
			defer cancel()

			logger.Infof(ctx, "grpc: waiting stop of traffic")
			time.Sleep(GracefulDelay)
			logger.Infof(ctx, "grpc: shutting down")

			done := make(chan struct{})
			go func() {
				s.grpc.server.GracefulStop()

				close(done)
			}()
			select {
			case <-done:
				logger.Infof(ctx, "grpc: gracefully stopped")
			case <-ctx.Done():
				err := fmt.Errorf("grpc: error during shutdown server: %w", ctx.Err())
				s.grpc.server.Stop()
				return fmt.Errorf("grpc: force stopped: %w", err)
			}
			return nil
		})
	}
}

func (s *Server) runGrpcGateway() {
	if s.grpcGateway.server != nil {
		go func() {
			logger.Infof(context.Background(), "start serve gRPC Gateway to %s", s.grpcGateway.lis.Addr())
			if err := s.grpcGateway.server.Serve(s.grpcGateway.lis); err != nil && !errors.Is(err, http.ErrServerClosed) {
				logger.Infof(context.Background(), fmt.Errorf("grpc gateway: %w", err).Error())
				s.resourceCloser.CloseAll()
			}
		}()
		s.resourceCloser.Add(func() error {
			ctx, cancel := context.WithTimeout(context.Background(), GracefulTimeout)
			defer cancel()

			logger.Infof(ctx, "grpc gateway: waiting stop of traffic")
			time.Sleep(GracefulDelay)
			logger.Infof(ctx, "grpc gateway: shutting down")

			s.grpcGateway.server.SetKeepAlivesEnabled(false)

			if err := s.grpcGateway.server.Shutdown(ctx); err != nil {
				return fmt.Errorf("grpc gateway: error during shutdown: %w", err)
			}
			logger.Infof(ctx, "grpc gateway: gracefully stopped")
			return nil
		})
	}
}

func (s *Server) runInfo() {
	if s.info.server != nil {
		go func() {
			logger.Infof(context.Background(), "start serve INFO to %s", s.info.lis.Addr())
			if err := s.info.server.Serve(s.info.lis); err != nil && !errors.Is(err, http.ErrServerClosed) {
				logger.Infof(context.Background(), fmt.Errorf("info: %w", err).Error())
				s.resourceCloser.CloseAll()
			}
		}()
		s.resourceCloser.Add(func() error {
			ctx, cancel := context.WithTimeout(context.Background(), GracefulTimeout)
			defer cancel()

			logger.Infof(ctx, "info: waiting stop of traffic")
			time.Sleep(GracefulDelay)
			logger.Infof(ctx, "info: shutting down")

			s.info.server.SetKeepAlivesEnabled(false)

			if err := s.info.server.Shutdown(ctx); err != nil {
				return fmt.Errorf("info: error during shutdown: %w", err)
			}
			logger.Infof(ctx, "info: gracefully stopped")
			return nil
		})
	}
}
