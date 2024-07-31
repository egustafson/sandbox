package main

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const pageSize = 5

func doPaginate(cmd *cobra.Command, args []string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	etcdCli := config.Client

	keyCount := 0
	totalBytes := 0

	rangeEnd := prefixEnd(matchPrefix)

	opts := []clientv3.OpOption{
		clientv3.WithRange(rangeEnd),
		clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend), // sorted so pagination works
		clientv3.WithLimit(pageSize),
	}
	nextKey := matchPrefix // next key is beginning of prefix range (the prefix)
	firstPass := true
	page, err := etcdCli.KV.Get(ctx, nextKey, opts...)
	for { // bottom tested loop
		if err != nil {
			if verboseFlag {
				fmt.Printf("etcd get error: %v\n", err)
			}
			return err
		}
		nextIndex := 1
		if firstPass {
			firstPass = false
			nextIndex = 0
		}
		for _, kv := range page.Kvs[nextIndex:] {
			if verboseFlag {
				fmt.Printf("counting key: %s\n", kv.Key)
			}
			keyCount += 1
			totalBytes += len(kv.Key) + len(kv.Value)
		}
		if !page.More {
			break
		}
		nextKey = string(page.Kvs[len(page.Kvs)-1].Key)
		if verboseFlag {
			fmt.Printf("next search key: %s\n", nextKey)
		}
		page, err = etcdCli.Get(ctx, nextKey, opts...)
	}

	fmt.Printf("finished, found %d keys of total size: %d bytes \n", keyCount, totalBytes)

	return nil
}

func prefixEnd(prefix string) string {
	end := make([]byte, len(prefix))
	copy(end, []byte(prefix))
	for ii := len(end) - 1; ii >= 0; ii-- {
		if end[ii] < 0xff {
			end[ii] = end[ii] + 1
			return string(end[:ii+1])
		}
	}

	return string([]byte{0})
}
