package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/egustafson/sandbox/go/grpc-streaming/pb"
)

type DemoServer struct{}

func main() {

	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	s := DemoServer{}
	grpcServer := grpc.NewServer()
	pb.RegisterDemoServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (s *DemoServer) ListenHeartbeat(req *pb.HeartbeatRequest, stream pb.DemoService_ListenHeartbeatServer) error {
	log.Printf("Heartbeat started for: %s\n", req.RequestId)
	var ii int = 0
	for {
		h := &pb.Heartbeat{
			Note: fmt.Sprintf("beat: %d", ii),
		}
		if err := stream.Send(h); err != nil {
			return err
		}
		ii = ii + 1
		log.Println("tick.")
		time.Sleep(time.Second)
	}
	return nil
}
