package main

import (
	"context"
	_ "errors"
	"fmt"
	"log"
	_ "net"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/egustafson/sandbox/grpc/grpc-error-handling/client-go/pb"
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
		_, ok := status.FromError(err) // <-- demonstrate the way to detect a conversion error (should NEVER happen).
		if !ok {
			log.Fatalf("** FATAL:  gRPC returned error that is not a status.Status ?!?!?")
		}
		logStatus(err)
		return
	}
	log.Printf("response:  %s", resp.RespText)
}

func logStatus(err error) {
	log.Printf("gRPC returned error: %v", err)
	st := status.Convert(err)
	for _, detail := range st.Details() {
		switch t := detail.(type) {
		case *errdetails.ErrorInfo:
			fmt.Println("-- ErrorInfo response:")
			fmt.Printf("     Reason: %s\n", t.Reason)
			fmt.Printf("     Domain: %s\n", t.Domain)
			fmt.Println("     Metadata:")
			for k, v := range t.Metadata {
				fmt.Printf("       %s: %s\n", k, v)
			}
		default:
			fmt.Println("-- Unknown error details.")
		}
	}
}
