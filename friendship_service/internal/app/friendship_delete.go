package app

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/hollyhox-21/discord_project/friendship_service/pkg/friendship"
)

func (i *Implementation) FriendshipDelete(ctx context.Context, req *pb.FriendshipDeleteRequest) (*emptypb.Empty, error) {

	log.Println("DeleteFriendship")

	return &emptypb.Empty{}, nil
}
