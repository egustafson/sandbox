#!/usr/bin/env python

import pika

connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
#connection = pika.BlockingConnection(pika.ConnectionParameters('15.185.224.216'))
channel = connection.channel()

channel.queue_declare(queue='hello')

print ' [*] Waiting for messages.  To exit press ctrl-c'

def callback(ch, method, properties, body):
    print " [x] Received %r" % (body,)

channel.basic_consume(callback, queue='hello', no_ack=True)

channel.start_consuming()
