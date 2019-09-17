#!/bin/sh

docker run -it \
  --net=kafka\
  --name=zookeeper \
  -e ZOOKEEPER_CLIENT_PORT=2181 \
  confluentinc/cp-zookeeper:5.3.0

