package app

import (
	"context"
	"log"

	pb "github.com/hollyhox-21/discord_project/user_profile_service/pkg/user_profile"
)

func (i *Implementation) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {

	log.Println("GetProfile")

	return &pb.GetProfileResponse{}, nil
}
