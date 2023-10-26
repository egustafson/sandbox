package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

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

	// retrieve and print grpc metadata from the client
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Print("  Received metadata:")
		for k, vSlice := range md {
			for _, v := range vSlice {
				log.Printf("  - %s: %s", k, v)
			}
		}
	} else {
		log.Print("  no metadata in request.")
	}

	// Attach metadata to the response
	hostname, _ := os.Hostname()
	pid := fmt.Sprintf("%d", os.Getpid())
	header := metadata.Pairs(
		"srv-hostname", hostname,
		"srv-pid", pid,
		"srv-language", "go",
		"srv-language-ver", runtime.Version(),
	)
	trailer := metadata.Pairs("app-trailer", "trailer-value")
	grpc.SendHeader(ctx, header)
	grpc.SetTrailer(ctx, trailer)

	return &pb.SvcResponse{RespText: "response-from-server"}, nil
}

// --  Main  -------------------------------------

func main() {
	fmt.Println("start: gRPC Demo Server")

	var (
		l   net.Listener
		err error
	)

	log.Print("\n")
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
