#!/usr/bin/env python

import pika

#credentials = pika.PlainCredentials('collectd', 'Ky0toL1sp')
credentials = pika.PlainCredentials('collectd', 'deadmau5')
parameters  = pika.ConnectionParameters(host='15.185.170.215',
                                        port=5672,
                                        credentials=credentials)

connection = pika.BlockingConnection(parameters)
channel = connection.channel()

channel.exchange_declare(exchange='metrics', type='topic')

result = channel.queue_declare(queue='gustafer', exclusive=True)
queue_name = result.method.queue
 
channel.queue_bind(exchange='metrics', queue=queue_name, routing_key="#")


print ' [*] Waiting for messages.  To exit press ctrl-c'

def callback(ch, method, properties, body):
    print "%r:%r" % (method.routing_key, body,)

channel.basic_consume(callback, queue=queue_name, no_ack=True)

channel.start_consuming()

print "done."
