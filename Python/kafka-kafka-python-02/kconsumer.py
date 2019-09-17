#!/usr/bin/env python3

from kafka import KafkaConsumer

consumer = KafkaConsumer('my-topic',
                         group_id='my-group',
                         bootstrap_servers=['localhost:9092'])

for message in consumer:
    print("%s:%d:%d: key=%s value=%s" % (message.topic, message.partition,
                                         message.offset, message.key,
                                         message.value))

KafkaConsumer(auto_offset_reset='earliest', enable_auto_commit=False)

#KafkaConsumer(value_deserializer=lambda m: json.loads(m.decode('ascii')))

#KafkaConsumer(value_deserializer=msgpack.unpackb)

KafkaConsumer(consumer_timeout_ms=1000)

consumer = KafkaConsumer()
consumer.subscribe(pattern='^awesome.*')

# consumer1 = KafkaConsumer('my-topic',
#                           group_id='my-group',
#                           bootstrap_servers='localhost:9092')

print('done.')
