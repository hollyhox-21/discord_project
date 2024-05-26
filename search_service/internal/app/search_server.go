package app

import (
	"context"
	"log"

	pb "github.com/hollyhox-21/discord_project/search_service/pkg/search"
)

func (i *Implementation) SearchServer(ctx context.Context, req *pb.SearchServerRequest) (*pb.SearchServerResponse, error) {

	log.Println("SearchServer")

	return &pb.SearchServerResponse{}, nil
}
