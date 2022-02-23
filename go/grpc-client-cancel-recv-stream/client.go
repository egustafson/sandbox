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

	"github.com/egustafson/sandbox/go/grpc-client-cancel-recv-stream/pb"
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
		if ii > 4 {
			log.Print("blocking on stream.Recv()")
		}
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
		log.Printf("beat(%d)", hb.Index)
		ii++
		if ii > 4 {
			log.Print("reached limit: initiating context.Cancel()")
			//
			// Defer canceling the context for 1 second.  Allow time
			// for the Recv() call to block.
			//
			go func() {
				log.Print("sleeping before cancel.")
				time.Sleep(time.Second)
				log.Print("canceling context")
				cancel()
			}()
		}
	}
}
