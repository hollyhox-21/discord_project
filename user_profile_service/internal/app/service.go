package app

import (
	"context"
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	pb "github.com/hollyhox-21/discord_project/user_profile_service/pkg/user_profile"
)

type Implementation struct {
	pb.UnimplementedUserProfileServiceServer

	Validator *protovalidate.Validator
}

func NewImplementation() (*Implementation, error) {
	srv := &Implementation{}

	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(
			// Добавляем сюда все запросы наши
			&pb.CreateProfileRequest{},
			&pb.GetProfileRequest{},
			&pb.UpdateProfileRequest{},
			&pb.DeleteProfileRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize validator: %w", err)
	}

	srv.Validator = validator

	return srv, nil
}

func (i *Implementation) Register() (*grpc.ServiceDesc, any) {
	return &pb.UserProfileService_ServiceDesc, i
}

func (i *Implementation) RegisterMux(ctx context.Context, mux *runtime.ServeMux) error {
	return pb.RegisterUserProfileServiceHandlerServer(ctx, mux, i)
}
