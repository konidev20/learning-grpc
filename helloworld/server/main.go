package main

import (
	"context"
	"fmt"
	"helloworld/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedPingerServer
}

func (s *server) Ping(ctx context.Context, in *proto.PingRequest) (*proto.PingReply, error) {
	log.Printf("received; %s", in.Name)
	return &proto.PingReply{Message: fmt.Sprintf("pong: %s", in.Name)}, nil
}

func main() {
	port := 50052
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to start grpc server: %s", err.Error())
	}

	// create new grpc server
	s := grpc.NewServer()

	// register the ping server
	proto.RegisterPingerServer(s, &server{})

	log.Printf("ping server listening at port: %d", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
