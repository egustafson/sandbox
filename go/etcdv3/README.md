Demo etcd v3 Client Lib
=======================

Prerequisite(s)
---------------

An etcd service running (on loopback) is required.  The provided
`start_etcd.sh` script will start an ephemeral etcd server in docker,
assuming docker is installed on the localhost.

Use
---

```
> go run demo.go
```


Citation(s)
-----------

* https://www.compose.com/articles/utilizing-etcd3-with-go/
* https://pkg.go.dev/github.com/LK4D4/etcd/clientv3
* https://github.com/etcd-io/etcd/blob/master/api/etcdserverpb/rpc.proto
* https://github.com/etcd-io/etcd/blob/master/api/etcdserverpb/rpc.pb.go
