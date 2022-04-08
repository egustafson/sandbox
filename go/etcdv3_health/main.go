package main

// Conditions: the example(s) below can be run with a FUNCTIONING etcd
// node on any or all of the listed endpoints.  "Functioning" means
// the node is "in quorem"; this could be a single node configured as
// such, or 2 nodes of a 3 node cluster being up.  etcd will not
// respond if it doesn't have quorem.
//
// Expected behavior:
//
// newExample - regardless if etcd is running or not, the example
// should initialize, NOT panic(), and loop printing HEALTHY or
// UN-HEALTHY based on if etcd is up or not.  etcd can be brought up
// and dropped while the loop is running and the loop should continue,
// printing an appropriate health state.
//
// newExample exemplifies that an etcd client "handle" can be
// constructed and not return an error, even if the etcd cluster is
// not up.
//
// oldExample - prints diagnostic information if etcd is up, or errors
// out if etcd is down.

import (
	"context"
	"fmt"

	//"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	ENDPOINT_1 = "127.0.0.1:2379"
	ENDPOINT_2 = "127.0.0.1:2479"
	ENDPOINT_3 = "127.0.0.1:2579"
)

var (
	dialTimeout = 8 * time.Second
	reqTimeout  = 10 * time.Second
)

func main() {
	newExample()
}

func newExample() {
	endpoints := []string{
		ENDPOINT_1,
		ENDPOINT_2,
		ENDPOINT_3,
	}
	fmt.Printf("attempting to connect to etcd: { %v }\n", endpoints)

	// should never error, even if etcd not present.
	st, err := NewEtcdStore(endpoints)
	if err != nil {
		panic(err)
	}

	for ii := 0; ii < 60; ii++ {
		healthy := st.Healthy()
		if healthy {
			fmt.Println("etcd connection is HEALTHY")
		} else {
			fmt.Println("etcd connection is UN-HEALTHY")
		}
		time.Sleep(time.Second)
	}
}

func origExample() {

	client, err := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{ENDPOINT_1, ENDPOINT_2, ENDPOINT_3},
		//Endpoints:   []string{ENDPOINT_1},
	})
	if err != nil {
		panic(err)
	}
	defer client.Close()
	fmt.Printf("Endpoints after New: %v\n", client.Endpoints())

	// At this point a client object has been constructed, but a
	// connection has may NOT have been established.
	//
	// Reviewing the code, and watching in a debugger: IF a connection
	// can be formed, it will be.  If not, the call will complete with
	// out error.

	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()

	// Dial'ing will attempt to form a connection

	grpcConn, err := client.Dial(ENDPOINT_3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("grpcConn.Target = %v\n", grpcConn.Target())

	// first point a connection to the service is attempted.
	//
	err = client.Sync(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Endpoints after Sync: %v\n", client.Endpoints())

	//
	// List Alarms
	//
	alarmResp, err := client.AlarmList(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("AlarmList has %d items.\n", len(alarmResp.Alarms))

	client.Close()
	time.Sleep(time.Millisecond)
	fmt.Println("done.")
}
