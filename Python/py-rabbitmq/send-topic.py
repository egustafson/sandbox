#!/usr/bin/env python

import pika

connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
#connection = pika.BlockingConnection(pika.ConnectionParameters('10.0.2.15'))
channel = connection.channel()

channel.exchange_declare(exchange='topic_logs', type='topic')

routing_key = 'anonymous.info'

message = 'Message'

channel.basic_publish(exchange='topic_logs', routing_key=routing_key, body=message)

print "Sent."
connection.close()
print "done."
