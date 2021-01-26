package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/coreos/etcd/clientv3"
	//clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	ETCD_ENDPOINT = "127.0.0.1:2379"
)

var (
	dialTimeout = 2 * time.Second
	reqTimeout  = 10 * time.Second
)

func main() {

	// Create a context to be used within the etcd (kv) client.
	//
	ctx, cancel := context.WithTimeout(context.Background(), reqTimeout)
	defer cancel()

	// Configure and allocate an etcd client
	//
	client, err := clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{ETCD_ENDPOINT},
	})
	checkErr("new etcd clientv3 failed", err)
	defer client.Close()

	// kvc := clientv3.NewKV(client) // returnes a narrow-ed interface of 'client'
	// _ = kvc

	printClusterInfo(ctx, client)
	GetSingleValueDemo(ctx, client)
	GetMultipleValuesWithPaginationDemo(ctx, client)
}

func checkErr(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func GetSingleValueDemo(ctx context.Context, kv clientv3.KV) {

	fmt.Println(" ---  Basic get/put/del Demo  ---")

	// Delete
	//
	dr, err := kv.Delete(ctx, "key", clientv3.WithPrefix()) // clear all keys beginning with 'key'
	checkErr("delete('key*') failed", err)
	fmt.Printf("del('key*') succeeded -> rev: %d\n", dr.Header.Revision)

	// Put
	//
	pr, err := kv.Put(ctx, "key", "444")
	checkErr("put('key', '444') failed", err)

	cluster_id := pr.Header.ClusterId
	member_id := pr.Header.MemberId
	rev := pr.Header.Revision
	raft_term := pr.Header.RaftTerm

	fmt.Printf("put('key', '444') succeeded -> rev: %d\n", rev)
	fmt.Printf("cluster_id: %d,   member_id: %d,   raft_term: %d\n", cluster_id, member_id, raft_term)

	// Get
	//
	gr, err := kv.Get(ctx, "key")
	checkErr("get('key') failed", err)
	fmt.Printf("get('key') -> '%s'  (rev: %d)\n", string(gr.Kvs[0].Value), gr.Header.Revision)
	prev_rev := gr.Header.Revision // used in the final get, to get an older revision

	// Put
	//
	pr, err = kv.Put(ctx, "key", "555")
	checkErr("put('key', '555') failed", err)
	fmt.Printf("put('key', '555') succeeded -> rev: %d\n", pr.Header.Revision)

	// Get
	//
	gr, err = kv.Get(ctx, "key")
	checkErr("get('key') failed", err)
	fmt.Printf("get('key') -> '%s'  (rev: %d)\n", string(gr.Kvs[0].Value), gr.Header.Revision)

	// Get previous revision value
	//
	gr, err = kv.Get(ctx, "key", clientv3.WithRev(prev_rev))
	checkErr("get('key', prev_rev) failed", err)
	fmt.Printf("get('key', rev: %d) -> %s\n", prev_rev, string(gr.Kvs[0].Value))
}

func GetMultipleValuesWithPaginationDemo(ctx context.Context, kv clientv3.KV) {

	fmt.Println(" ---  Pagination Demo  ---")

	// delete all keys
	_, err := kv.Delete(ctx, "key", clientv3.WithPrefix())
	checkErr("del('key*') failed", err)

	// Insert 20 unique keys
	for ii := 0; ii < 20; ii++ {
		k := fmt.Sprintf("key_%02d", ii)
		_, err = kv.Put(ctx, k, strconv.Itoa(ii))
		checkErr("put('key_??') failed", err)
	}

	opts := []clientv3.OpOption{
		clientv3.WithPrefix(), // ==> get a range
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend), // sorted so pagination works
		clientv3.WithLimit(5), // limit "page" to 5 in order to demo paging
	}

	var nextkey = "key"
	var firstPass = true
	gr, err := kv.Get(ctx, nextkey, opts...)
	checkErr("get('key') <page 1> failed", err)
	for {
		var firstIndex = 1
		if firstPass {
			firstIndex = 0
			opts = append(opts, clientv3.WithFromKey()) // apply after the initial Get()
			firstPass = false
		}
		fmt.Println("--- page ---")
		for _, item := range gr.Kvs[firstIndex:] {
			fmt.Println(string(item.Key), string(item.Value))
		}
		if !gr.More { // bottom tested loop
			break
		}
		nextkey = string(gr.Kvs[len(gr.Kvs)-1].Key)
		gr, err = kv.Get(ctx, nextkey, opts...)
		checkErr("get('key...') <next page> failed", err)
	}
	fmt.Println("--- end ---")
}

func printClusterInfo(ctx context.Context, client *clientv3.Client) {

	fmt.Println(" ---  Cluster Info  ---")

	for _, item := range client.Endpoints() {
		fmt.Printf("ep: %s\n", item)
	}
	conn := client.ActiveConnection()
	if conn != nil {
		fmt.Printf("connection state:   %s\n", conn.GetState().String())
		fmt.Printf("connection target:  %s\n", conn.Target()) // marked 'EXPERIMENTAL' in the docs.
	} else {
		fmt.Println("connection state: <no connection>")
	}

	// Print cluster member list
	fmt.Println("Members:")
	ml, err := client.MemberList(ctx)
	checkErr("MemberList() failed", err)
	for _, m := range ml.Members {
		fmt.Printf("  %s(id: %d)\n", m.GetName(), m.GetID())
	}
}
