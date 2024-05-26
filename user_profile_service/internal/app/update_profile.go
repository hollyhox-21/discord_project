package app

import (
	"context"
	"log"

	pb "github.com/hollyhox-21/discord_project/user_profile_service/pkg/user_profile"
)

func (i *Implementation) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {

	log.Println("UpdateProfile")

	return &pb.UpdateProfileResponse{}, nil
}
