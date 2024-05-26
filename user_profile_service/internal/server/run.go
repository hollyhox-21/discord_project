package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
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
	s.publicCloser.Wait()
	// Close all other resources from globalCloser
	CloseAll()
}

func (s *Server) runGRPC() {
	if s.grpc.server != nil {
		go func() {
			log.Println("start serve gRPC to", s.grpc.lis.Addr())
			if err := s.grpc.server.Serve(s.grpc.lis); err != nil {
				log.Println(fmt.Errorf("grpc: %w", err).Error())
				s.publicCloser.CloseAll()
			}
		}()
		s.publicCloser.Add(func() error {
			ctx, cancel := context.WithTimeout(context.Background(), GracefulTimeout)
			defer cancel()

			log.Println("grpc: waiting stop of traffic")
			time.Sleep(GracefulDelay)
			log.Println("grpc: shutting down")

			done := make(chan struct{})
			go func() {
				s.grpc.server.GracefulStop()
				close(done)
			}()
			select {
			case <-done:
				log.Println("grpc: gracefully stopped")
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
			log.Println("start serve gRPC Gateway to", s.grpcGateway.lis.Addr())
			if err := s.grpcGateway.server.Serve(s.grpcGateway.lis); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Println(fmt.Errorf("grpc gateway: %w", err).Error())
				s.publicCloser.CloseAll()
			}
		}()
		s.publicCloser.Add(func() error {
			ctx, cancel := context.WithTimeout(context.Background(), GracefulTimeout)
			defer cancel()

			log.Println("grpc gateway: waiting stop of traffic")
			time.Sleep(GracefulDelay)
			log.Println("grpc gateway: shutting down")

			s.grpcGateway.server.SetKeepAlivesEnabled(false)

			if err := s.grpcGateway.server.Shutdown(ctx); err != nil {
				return fmt.Errorf("grpc gateway: error during shutdown: %w", err)
			}
			log.Println("grpc gateway: gracefully stopped")
			return nil
		})
	}
}

func (s *Server) runInfo() {
	if s.info.server != nil {
		go func() {
			log.Println("start serve INFO to", s.info.lis.Addr())
			if err := s.info.server.Serve(s.info.lis); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Println(fmt.Errorf("info: %w", err).Error())
				s.publicCloser.CloseAll()
			}
		}()
		s.publicCloser.Add(func() error {
			ctx, cancel := context.WithTimeout(context.Background(), GracefulTimeout)
			defer cancel()

			log.Println("info: waiting stop of traffic")
			time.Sleep(GracefulDelay)
			log.Println("info: shutting down")

			s.info.server.SetKeepAlivesEnabled(false)

			if err := s.info.server.Shutdown(ctx); err != nil {
				return fmt.Errorf("info: error during shutdown: %w", err)
			}
			log.Println("info: gracefully stopped")
			return nil
		})
	}
}
