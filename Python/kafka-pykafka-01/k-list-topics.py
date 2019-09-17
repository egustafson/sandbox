#!/usr/bin/env python

from pykafka import KafkaClient

client = KafkaClient(hosts="127.0.0.1:9092")

# Force the client to actually talk to the cluster.
# (it should already, but just in case)
client.update_cluster()

print("Topics:")
for topic in client.topics.iterkeys():
    print("  {}".format(topic))

print("done.")
