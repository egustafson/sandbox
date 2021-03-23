package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/egustafson/sandbox/_hybrid/grpc-tls-python-golang/server-go/pb"
)

type svc struct{}

func (s *svc) DoService(ctx context.Context, req *pb.SvcRequest) (*pb.SvcResponse, error) {
	log.Printf("Received message from client: %s", req.ReqText)
	return &pb.SvcResponse{RespText: "response-text"}, nil
}

func main() {
	fmt.Println("start: gRPC Demo Server")

	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	service := svc{}
	grpcServer := grpc.NewServer()
	pb.RegisterSvcServer(grpcServer, &service)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// never return
}
