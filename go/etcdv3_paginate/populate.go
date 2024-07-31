package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	clientv3 "go.etcd.io/etcd/client/v3"
	"golang.org/x/exp/rand"
)

func doPopulate(cmd *cobra.Command, args []string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) // <- time for _all_ operations
	defer cancel()
	etcdCli := config.Client

	matchByteSize := 0
	for i := 0; i < matchCount; i++ {
		k := matchPrefix + uuid.New().String()
		v := randValue()
		matchByteSize += len(k) + len(v)
		put(ctx, k, v, etcdCli)
	}

	put(ctx, "0_before", "placeholder-value", etcdCli)
	put(ctx, "z_after", "placeholder-value", etcdCli)

	fmt.Printf("Added %d kv's of total size: %d bytes\n", matchCount, matchByteSize)

	return nil
}

func put(ctx context.Context, key, value string, etcdCli *clientv3.Client) {
	resp, err := etcdCli.Put(ctx, key, value)
	if err != nil {
		if verboseFlag {
			fmt.Printf("put failed: %v\n", err)
			panic(err)
		}
	}
	if verboseFlag {
		fmt.Printf("put response: %v\n", resp)
	}
}

var alphaNums = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

func randValue() string {
	size := 16 + rand.Int31n(128)
	b := make([]byte, size)
	for i := range b {
		b[i] = alphaNums[rand.Intn(len(alphaNums))]
	}
	return string(b)
}
