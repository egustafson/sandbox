package main

import (
	"fmt"
	"log"
	"net"

	"github.com/egustafson/sandbox/go/grpc-basic/api"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go gRPC Tutorial")

	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := api.Server{}
	grpcServer := grpc.NewServer()
	api.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
