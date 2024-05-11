package app

import (
	"context"
	"log"

	pb "github.com/hollyhox-21/discord_project/friendship_service/pkg/friendship"
)

func (i *Implementation) FriendshipList(ctx context.Context, req *pb.FriendshipListRequest) (*pb.FriendshipListResponse, error) {

	log.Println("ListFriendships")

	return &pb.FriendshipListResponse{}, nil
}
