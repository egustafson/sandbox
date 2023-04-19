package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/egustafson/sandbox/grpc/grpc-tls-python-golang/server-go/pb"
)

const (
	use_tls     = true
	listen_addr = ":9000"

	cafile   = "../test-ca.pem"
	certfile = "../test-cert.pem"
	keyfile  = "../test-key.pem"
)

type svc struct {
	pb.UnimplementedSvcServer // required by newer versions of protoc
}

func (s *svc) DoService(ctx context.Context, req *pb.SvcRequest) (*pb.SvcResponse, error) {
	log.Printf("Received message from client: %s", req.ReqText)
	return &pb.SvcResponse{RespText: "response-text-golang"}, nil
}

func main() {
	fmt.Println("start: gRPC Demo Server")

	var (
		l   net.Listener
		err error
	)
	if use_tls {
		log.Print("using TLS")
		tlsConfig, err := NewTlsConfig()
		tlsConfig.NextProtos = []string{"h2"} // <-- ** Negotiate TLS ALPN (RFC 7540)
		if err != nil {
			log.Fatalf("failed to import PEM files for TLS")
		}
		l, err = tls.Listen("tcp", listen_addr, tlsConfig)
		if err != nil {
			log.Fatalf("failed to form a TLS listener: %v", err)
		}
	} else {
		log.Print("TLS disabled")
		l, err = net.Listen("tcp", listen_addr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
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

// --  loadTlsConfig  ------------------------------------------------

func NewTlsConfig() (cfg *tls.Config, err error) {

	// -- load the certificate + key from PEM
	//
	cert, err := tls.LoadX509KeyPair(certfile, keyfile)
	if err != nil {
		log.Fatalf("problem reading cert/key PEM: %v", err)
		return nil, err
	}

	// -- load the CA from PEM
	caPem, err := ioutil.ReadFile(cafile)
	if err != nil {
		log.Fatalf("problem reading CA PEM file: %v", err)
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caPem) {
		log.Fatal("problem appending CA")
		return nil, errors.New("error appending CA to CertPool")
	}

	cfg = &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    caCertPool,
		RootCAs:      caCertPool,
		//ClientAuth:         tls.NoClientCert, // relax TLS checks (no client auth)
		ClientAuth:         tls.RequireAnyClientCert, // relax TLS checks (no client auth)
		InsecureSkipVerify: true,                     // relax TLS checks (no verify cert signed by CA)
	}
	return cfg, nil
}
