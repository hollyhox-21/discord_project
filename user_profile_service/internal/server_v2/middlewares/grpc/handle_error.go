package grpc

import (
	"context"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"github.com/bufbuild/protovalidate-go"
	"github.com/hollyhox-21/discord_project/libraries/logger"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandleError interceptor.
func HandleError(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		logger.Infow(ctx, "what happen", "error", err)
	}

	return resp, gRPCError(err)
}

// RPCError - returns codes.InvalidArgument
func gRPCError(err error) error {
	if err == nil {
		return status.Error(codes.OK, "")
	}

	if _, ok := status.FromError(err); ok {
		return err
	}

	switch valErr := err.(type) {
	case *protovalidate.ValidationError:
		st, stErr := status.New(codes.InvalidArgument, codes.InvalidArgument.String()).
			WithDetails(convertProtovalidateValidationErrorToErrDetailsBadRequest(valErr))
		if stErr == nil {
			return st.Err()
		}
	}

	return status.Error(codes.Internal, err.Error())
}

func convertProtovalidateValidationErrorToErrDetailsBadRequest(valErr *protovalidate.ValidationError) *errdetails.BadRequest {
	return &errdetails.BadRequest{
		FieldViolations: protovalidateVialationsToGoogleViolations(valErr.Violations),
	}
}

func protovalidateVialationsToGoogleViolations(vs []*validate.Violation) []*errdetails.BadRequest_FieldViolation {
	res := make([]*errdetails.BadRequest_FieldViolation, len(vs))
	for i, v := range vs {
		res[i] = &errdetails.BadRequest_FieldViolation{
			Field:       v.FieldPath,
			Description: v.Message,
		}
	}
	return res
}
