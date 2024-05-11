package app

import (
	"context"
	"log"

	pb "github.com/hollyhox-21/discord_project/user_profile_service/pkg/user_profile"
)

func (i *Implementation) DeleteProfile(ctx context.Context, req *pb.DeleteProfileRequest) (*pb.DeleteProfileResponse, error) {

	log.Println("DeleteProfile")

	return &pb.DeleteProfileResponse{}, nil
}
