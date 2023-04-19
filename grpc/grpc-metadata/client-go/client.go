package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

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

	doRequest(svc, "random-req-msg-golang")
	doRequest(svc, "ok-req-msg")
	doRequest(svc, "err-req-msg")
	doRequest(svc, "err-internal-req-msg")
	doRequest(svc, "err-abort-req-msg") // emits an ErrorInfo
	doRequest(svc, "err-timeout")       // cause service to delay 1 hr
}

func doRequest(svc pb.SvcClient, msg string) {
	deadline := time.Now().Add(timeout_duration) // <-- timeout
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	req := &pb.SvcRequest{ReqText: msg}
	resp, err := svc.DoService(ctx, req)
	if err != nil {
		log.Printf("unexpected ERROR: %v", err)
		return
	}
	log.Printf("response:  %s", resp.RespText)
}
