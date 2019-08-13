#!/usr/bin/env python

from pykafka import KafkaClient

client = KafkaClient(hosts="127.0.0.1:9092")

# implicitly creates the named topic
#
topic = client.topics["eg.test"]

consumer = topic.get_simple_consumer()
for message in consumer:
    if message is not None:
        print message.offset, message.value

print("done.")
