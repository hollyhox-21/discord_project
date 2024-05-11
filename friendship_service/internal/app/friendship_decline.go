package app

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/hollyhox-21/discord_project/friendship_service/pkg/friendship"
)

func (i *Implementation) FriendshipDecline(ctx context.Context, req *pb.FriendshipDeclineRequest) (*emptypb.Empty, error) {

	log.Println("DeclineFriendship")

	return &emptypb.Empty{}, nil
}
