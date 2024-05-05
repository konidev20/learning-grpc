package main

import (
	"context"
	"helloworld/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := "localhost:50052"
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to grpc server: %s", err.Error())
	}
	defer conn.Close()

	c := proto.NewPingerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.Ping(ctx, &proto.PingRequest{Name: "Koni"})
	if err != nil {
		log.Fatalf("failed to make ping request: %s", err.Error())
	}

	log.Printf("resp: %s", resp.Message)
}
