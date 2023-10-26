package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/egustafson/sandbox/grpc/grpc-metadata/client-go/pb"
)

const (
	listen_addr      = ":9000"
	timeout_duration = time.Second * 10 // 10 seconds
)

func main() {
	fmt.Println("start: gRPC Demo Server")

	var (
		conn *grpc.ClientConn
		err  error
	)

	log.Printf("connecting to localhost:9000\n")
	conn, err = grpc.Dial("localhost:9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	defer conn.Close()

	svc := pb.NewSvcClient(conn)

	doRequest(svc, "client-request")
}

func doRequest(svc pb.SvcClient, msg string) {
	deadline := time.Now().Add(timeout_duration) // <-- timeout
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	// send hostname and pid as grpc metadata
	hostname, _ := os.Hostname()
	pid := fmt.Sprintf("%d", os.Getpid())
	md := metadata.Pairs(
		"app-hostname", hostname,
		"app-pid", pid,
		"app-language", "go",
		"app-language-version", runtime.Version(),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	var header, trailer metadata.MD

	req := &pb.SvcRequest{ReqText: msg}
	resp, err := svc.DoService(ctx, req, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Printf("unexpected ERROR: %v", err)
		return
	}
	log.Printf("response:  %s", resp.RespText)

	log.Print("  Header:")
	for k, vSlice := range header {
		for _, v := range vSlice {
			log.Printf("  - %s: %s", k, v)
		}
	}
	log.Print("  Trailer:")
	for k, vSlice := range trailer {
		for _, v := range vSlice {
			log.Printf("  - %s: %s", k, v)
		}
	}
	// TODO:  print the header and trailer that we pulled out of the gRPC call
}
