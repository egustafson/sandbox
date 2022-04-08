package main

// The Store (store.go) abstracts a storage endpoint, an etcd cluster.
// The Store does not assume the endpoint is present or healthy,
// mearly that it is configured and may be now, or in the future,
// accessible (i.e. available and healthy).
//
// The Connector (connector.go) abstracts a connection to the store
// and is capable of reporting the health state of the connection, and
// store by proxy.

type Store interface {
	State() *StoreState
	Healthy() bool
	GetConnector() (Connector, error)
	Close()
}

// StoreState ideally implements some more abstract 'State' interface
type StoreState struct {
	Healthy bool
	Err     error
}

type etcdStore struct {
	endpoints []string
	connector *etcdConnector
}

func NewEtcdStore(endpoints []string) (Store, error) {
	var err error
	st := &etcdStore{endpoints: endpoints}
	st.connector, err = newEtcdConnector(endpoints)
	return st, err
}

func (st *etcdStore) State() *StoreState {
	return st.connector.State()
}

func (st *etcdStore) Healthy() bool {
	return st.connector.Healthy()
}

func (st *etcdStore) GetConnector() (Connector, error) {
	return st.connector, nil
}

func (st *etcdStore) Close() {
	// no-op
}
