package main

import (
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/egustafson/sandbox/go/grpc-client-cancel-recv-stream/pb"
)

type DemoServer struct {
	pb.UnimplementedDemoServiceServer // required by newer versions of protoc
}

func main() {

	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	s := DemoServer{}
	grpcServer := grpc.NewServer()
	pb.RegisterDemoServiceServer(grpcServer, &s)

	log.Print("Server running...")
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (s *DemoServer) ListenHeartbeat(req *pb.HeartbeatRequest, stream pb.DemoService_ListenHeartbeatServer) error {
	log.Printf("Heartbeat started for: %s\n", req.RequestId)
	var ii uint64 = 0
	for {
		h := &pb.Heartbeat{
			Index: ii,
		}
		if err := stream.Send(h); err != nil {
			//
			// Stream was canceled by the client (orderly teardown)
			//
			if status.Code(err) == codes.Canceled {
				log.Printf("stream(%s) Canceled, stream terminating", req.RequestId)
				return nil
			}
			//
			// Connection to client abrubtly closed (unorderly teardown)
			//
			if status.Code(err) == codes.Unavailable {
				log.Printf("stream(%s) Unavailable (transport closing), stream terminating", req.RequestId)
				return nil
			}
			//
			// Unexpected error
			//
			log.Printf("stream(%s) error: %v", req.RequestId, err)
			return err
		}
		ii = ii + 1
		log.Println("tick.")
		time.Sleep(2 * time.Second)
	}
}
