package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

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
	doConsume(client)
	//
	// Now, sleep for a while in an attempt to timeout the server's socket
	//
	time.Sleep(120 * time.Second)
	log.Print("done.")
}

func doConsume(client pb.DemoServiceClient) {

	hreq := &pb.HeartbeatRequest{
		RequestId: fmt.Sprintf("%d", os.Getpid()),
	}

	stream, err := client.ListenHeartbeat(context.Background(), hreq)
	if err != nil {
		log.Fatalf("client call failed: %s", err)
	}
	ii := 0
	for {
		hb, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Println(hb.Note)
		ii++
		if ii > 10 {
			log.Print("reached limit: returning from doConsume()")
			return
		}
	}
}
