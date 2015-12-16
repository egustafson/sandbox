#!/usr/bin/env python

from pykafka import KafkaClient

client = KafkaClient(hosts="127.0.0.1:9092")

# implicitly creates the topic
#
topic = client.topics["eg.test"]

# Set linger time to 100ms (default is 5s).  Default is to batch
# messages, and lingering 5s causes confusing delay in exit.
#
with topic.get_producer(linger_ms=100) as producer:
    for count in range(100000):
        msg = "m-{}".format(count)
        #print(msg)
        producer.produce(msg)


print("done.")
