package app

import (
	"context"
	"log"

	pb "github.com/hollyhox-21/discord_project/user_profile_service/pkg/user_profile"
)

func (i *Implementation) CreateProfile(ctx context.Context, req *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {

	log.Println("CreateProfile")

	return &pb.CreateProfileResponse{}, nil
}
