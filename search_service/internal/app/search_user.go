package app

import (
	"context"
	"log"

	pb "github.com/hollyhox-21/discord_project/search_service/pkg/search"
)

func (i *Implementation) SearchUser(ctx context.Context, req *pb.SearchUserRequest) (*pb.SearchUserResponse, error) {

	log.Println("SearchUser")

	return &pb.SearchUserResponse{}, nil
}
