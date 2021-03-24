package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	_ "net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/egustafson/sandbox/_hybrid/grpc-tls-python-golang/client-go/pb"
)

const (
	use_tls     = true
	listen_addr = ":9000"

	cafile   = "../test-ca.pem"
	certfile = "../test-cert.pem"
	keyfile  = "../test-key.pem"
)

func main() {
	fmt.Println("start: gRPC Demo Server")

	var (
		conn *grpc.ClientConn
		err  error
	)
	if use_tls {
		log.Print("using TLS")
		tlsConfig, err := NewTlsConfig()
		if err != nil {
			log.Fatalf("failed to import PEM files for TLS")
		}
		creds := credentials.NewTLS(tlsConfig)
		conn, err = grpc.Dial("localhost:9000", grpc.WithTransportCredentials(creds))
	} else {
		log.Print("TLS disabled")

		conn, err = grpc.Dial("localhost:9000", grpc.WithInsecure())
	}
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	defer conn.Close()

	svc := pb.NewSvcClient(conn)
	req := &pb.SvcRequest{ReqText: "request-text-golang"}
	resp, err := svc.DoService(context.Background(), req)
	if err != nil {
		log.Fatalf("gRPC request failed: %v", err)
	}
	log.Printf("response:  %s", resp.RespText)
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
		Certificates:       []tls.Certificate{cert},
		ClientCAs:          caCertPool,
		RootCAs:            caCertPool,
		ClientAuth:         tls.NoClientCert, // relax TLS checks (no client auth)
		InsecureSkipVerify: true,             // relax TLS checks (no verify cert signed by CA)
	}
	return cfg, nil
}
