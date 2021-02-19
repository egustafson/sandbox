package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"

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

	hreq := &pb.HeartbeatRequest{
		RequestId: fmt.Sprintf("%d", os.Getpid()),
	}

	stream, err := client.ListenHeartbeat(context.Background(), hreq)
	if err != nil {
		log.Fatalf("client call failed: %s", err)
	}
	for {
		hb, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println(hb.Note)
	}

	log.Print("done.")
}
