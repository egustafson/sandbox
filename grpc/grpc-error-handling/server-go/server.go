package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/egustafson/sandbox/grpc/grpc-error-handling/server-go/pb"
)

const (
	listen_addr   = ":9000"
	timeout_sleep = time.Hour
)

type svc struct {
	pb.UnimplementedSvcServer // required by newer versions of protoc
}

func (s *svc) DoService(ctx context.Context, req *pb.SvcRequest) (*pb.SvcResponse, error) {
	var (
		respText string
		respErr  error
	)

	log.Printf("Received message from client: %s", req.ReqText)

	switch req.ReqText {
	case "ok-req-msg":
		respText = "ok-resp-msg"
	case "err-req-msg":
		respErr = errors.New("err-resp-msg")
	case "err-internal-req-msg":
		respErr = status.Error(codes.Internal, "err-internal-resp-msg")
	case "err-abort-req-msg":
		respErr = mkErrorInfoErr()
	case "err-timeout":
		time.Sleep(timeout_sleep)
	default:
		respText = "unk-resp-msg"
	}

	if respErr != nil {
		return nil, respErr
	} else {
		return &pb.SvcResponse{RespText: respText}, nil
	}
}

func mkErrorInfoErr() error {
	st := status.New(codes.Aborted, "err-abort-resp-msg")
	errorInfo := &errdetails.ErrorInfo{
		Reason: "reason-string-abort-requested",
		Domain: "kvs.viasat",
		Metadata: map[string]string{
			"key-1": "value-1",
			"key-2": "value-2",
		},
	}
	st, err := st.WithDetails(errorInfo)
	if err != nil {
		panic(fmt.Sprintf("Unexpected error attaching metadata: %v", err))
	}
	return st.Err()
}

// --  Main  -------------------------------------

func main() {
	fmt.Println("start: gRPC Demo Server")

	var (
		l   net.Listener
		err error
	)

	log.Printf("listening on localhost%s\n", listen_addr)
	l, err = net.Listen("tcp", listen_addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	service := svc{}
	grpcServer := grpc.NewServer()
	pb.RegisterSvcServer(grpcServer, &service)

	log.Print("service configured and running.")
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// never return
}
