package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/egustafson/sandbox/go/grpc-basic/api"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial failed to connect: %s", err)
	}
	defer conn.Close()

	c := api.NewChatServiceClient(conn)

	resp, err := c.SayHello(context.Background(), &api.Message{Body: "Hello From Client."})
	if err != nil {
		log.Fatalf("Error grpc'ing SayHello: %s", err)
	}
	log.Printf("response:  %s", resp.Body)
}
