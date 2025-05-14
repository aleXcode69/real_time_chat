package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	UnimplementedCentrifugoProxyServer
}

func (s *server) Subscribe(ctx context.Context, request *SubscribeRequest) (*SubscribeResponse, error) {
	userID := request.User;
	channel := request.Channel;
	if !canUserSubscribe(userID, channel) {
		return &SubscribeResponse{
			Error: &Error{
				Code:    403,
				Message: "Permission denied, you cannot subscribe to this channel",
			},
		}, nil
	}
	return &SubscribeResponse{}, nil
}

func (s *server) Publish(ctx context.Context, request *PublishRequest) (*PublishResponse, error) {
	userID := request.User;
	channel := request.Channel;
	if !canUserPublish(userID, channel) {
		return &PublishResponse{
			Error: &Error{
				Code:    403,
				Message: "Permission denied, you cannot publish into this channel",
			},
		}, nil
	}
	return &PublishResponse{}, nil
}

func main() {
	listener, err := net.Listen("tcp4", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
		return
	}

	grpcServer := grpc.NewServer()
	RegisterCentrifugoProxyServer(grpcServer, &server{})

	fmt.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		fmt.Println("Failed to serve:", err)
	}
}
