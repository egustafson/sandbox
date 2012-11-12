#!/usr/bin/env python

import pika

connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
#connection = pika.BlockingConnection(pika.ConnectionParameters('10.0.2.15'))
channel = connection.channel()

channel.exchange_declare(exchange='topic_logs', type='topic')

result = channel.queue_declare(exclusive=True)
queue_name = result.method.queue

binding_keys = "#"
channel.queue_bind(exchange='topic_logs', queue=queue_name, routing_key=binding_keys)

print "Waiting ..."

def callback(ch, method, properties, body):
    print "%r: %r" % (method.routing_key, body,)

channel.basic_consume(callback, queue=queue_name, no_ack=True)

channel.start_consuming()
