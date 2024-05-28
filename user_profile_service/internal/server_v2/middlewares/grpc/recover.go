package grpc

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/hollyhox-21/discord_project/libraries/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func recoverFrom(ctx context.Context, p any) error {
	logger.Errorw(
		ctx, fmt.Sprintf("recovered from panic: %v", p),
		"stack_trace", string(debug.Stack()),
		"panic", true,
		"component", "grpc_recover_middleware",
	)

	return status.Errorf(codes.Internal, "recover: unexpected server error")
}

// Recover interceptor.
func Recover(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = recoverFrom(ctx, p)
		}
	}()
	return handler(ctx, req)
}

// RecoverStream interceptor recovers from panic.
func RecoverStream(srv interface{}, stream grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = recoverFrom(stream.Context(), p)
		}
	}()

	return handler(srv, stream)
}
