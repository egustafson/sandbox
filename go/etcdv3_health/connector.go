package main

// The Connector (connector.go) abstracts a connection to the Store
// and is capable of reporting the health state of the connection, and
// store by proxy.

import (
	"context"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	etcdTimeout = 2 * time.Second
)

type Connector interface {
	State() *StoreState
	Healthy() bool
	//
	// Other, actual operations on the Connector
	//
}

type etcdConnector struct {
	endpoints []string
	client    *clientv3.Client
}

func newEtcdConnector(endpoints []string) (*etcdConnector, error) {
	var err error
	conn := &etcdConnector{endpoints: endpoints}
	conn.client, err = clientv3.New(clientv3.Config{
		Endpoints: endpoints,
	})
	return conn, err
}

func (c *etcdConnector) State() (st *StoreState) {
	st = new(StoreState)
	ctx, cancel := context.WithTimeout(context.TODO(), etcdTimeout)
	defer cancel()

	st.Err = c.client.Sync(ctx)
	st.Healthy = (st.Err == nil)
	return st
}

func (c *etcdConnector) Healthy() bool {
	st := c.State()
	return st.Healthy
}
