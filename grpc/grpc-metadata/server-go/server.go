package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/egustafson/sandbox/grpc/grpc-metadata/server-go/pb"
)

const (
	listen_addr   = ":9000"
	timeout_sleep = time.Hour
)

type svc struct {
	pb.UnimplementedSvcServer // required by newer versions of protoc
}

func (s *svc) DoService(ctx context.Context, req *pb.SvcRequest) (*pb.SvcResponse, error) {
	log.Printf("Received message from client: %s", req.ReqText)
	return &pb.SvcResponse{RespText: "response-from-server"}, nil
}

// --  Main  -------------------------------------

func main() {
	fmt.Println("start: gRPC Demo Server")

	var (
		l   net.Listener
		err error
	)

	log.Printf("listening on localhost%s\n", listen_addr)
	l, err = net.Listen("tcp", listen_addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	service := svc{}
	grpcServer := grpc.NewServer()
	pb.RegisterSvcServer(grpcServer, &service)

	log.Print("service configured and running.")
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// never return
}
