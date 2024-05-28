package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// ContextErrorUnaryServerInterceptor protects gRPC from returning nil-error in cases when context is alreadt canceled.
func ContextErrorUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if ctx.Err() != nil {
			err = status.FromContextError(ctx.Err()).Err()

			return nil, err
		}

		return resp, err
	}
}

// ContextErrorStreamServerInterceptor protects gRPC from returning nil-error in cases when context is alreadt canceled.
func ContextErrorStreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := ss.Context()

		err := handler(srv, ss)
		if ctx.Err() != nil {
			return status.FromContextError(ctx.Err()).Err()
		}

		return err
	}
}
