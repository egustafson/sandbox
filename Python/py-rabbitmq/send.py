#!/usr/bin/env python

import pika

connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
#connection = pika.BlockingConnection(pika.ConnectionParameters('15.185.224.216'))
channel = connection.channel()

channel.queue_declare(queue='hello')

ii = 0

while True:
    ii = ii + 1
    channel.basic_publish(exchange='', routing_key='hello', body="Hello-Message[%d]" % ii)
    print " [x] Sent 'Hello-Message'"

connection.close()
