package service_go

import (
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	UnimplementedControlServiceServer
}

func (s *server) StreamControl(stream ControlService_StreamControlServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received command: %s with args: %v", in.CommandName, in.Args)
		// Echo back the received message as an example response
		if err := stream.Send(in); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterControlServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
