package main

import (
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Config struct {
	Endpoint string
	Client   *clientv3.Client
}

func initConfig() *Config {
	config := &Config{
		Endpoint: "127.0.0.1:2379",
	}

	dialTimeout := 2 * time.Second
	var err error
	config.Client, err = clientv3.New(clientv3.Config{
		DialTimeout: dialTimeout,
		Endpoints:   []string{config.Endpoint},
	})
	if err != nil {
		if verboseFlag {
			fmt.Printf("etcd client connect failed: %v", err)
		}
		panic(err)
	}

	return config
}
