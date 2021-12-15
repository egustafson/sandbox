package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/egustafson/sandbox/go/grpc-streaming/pb"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial failed to connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewDemoServiceClient(conn)
	doConsume(client)
	time.Sleep(1 * time.Second) // give the gRPC library a moment to clean up.
	log.Print("done.")
}

func doConsume(client pb.DemoServiceClient) {

	reqId := os.Getpid()
	hreq := &pb.HeartbeatRequest{
		RequestId: fmt.Sprintf("%d", reqId),
	}
	log.Printf("requestId: %d", reqId)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := client.ListenHeartbeat(ctx, hreq)
	if err != nil {
		log.Fatalf("client call failed: %s", err)
	}
	ii := 0
	for {
		hb, err := stream.Recv()
		if err != nil {
			//
			// remote (Server) closed the stream (of its own accord).
			//
			if err == io.EOF {
				log.Print("received io.EOF, stream ended (properly)")
				return
			}
			//
			// (this) client canceled the stream
			//   all (local) messages have been consumed
			//
			if status.Code(err) == codes.Canceled {
				log.Print("received grpc Canceled")
				return
			}
			//
			// Unexpected error
			//
			log.Fatalf("stream returned error: %v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println(hb.Note)
		ii++
		if ii > 10 {
			log.Print("reached limit: canceling context")
			//
			// Cancel channel context (orderly shutdown)
			//  keep consuming, there may be messages in-flight.
			//
			cancel()
		}
	}
}
