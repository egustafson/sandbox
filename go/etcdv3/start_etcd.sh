#!/bin/sh

REGISTRY=gcr.io/etcd-development/etcd
NODE_NAME=node1
NODEIP=127.0.0.1

docker run \
       -d --rm \
       -p 2379:2379 \
       -p 2380:2380 \
       --name etcd \
       ${REGISTRY} \
       /usr/local/bin/etcd \
       --name ${NODE_NAME} \
       --initial-advertise-peer-urls http://${NODEIP}:2380  --listen-peer-urls http://0.0.0.0:2380 \
       --advertise-client-urls http://${NODEIP}:2379       --listen-client-urls http://0.0.0.0:2379 \
       --initial-cluster ${NODE_NAME}=http://${NODEIP}:2380
