#!/usr/bin/env python

import json
import pika
import sys
import time


thresh = 10000
counter = 0

vector = []

def callback(ch, method, properties, body):
    global counter
    global vector
    try: 
        counter
    except UnboundLocalError:
        count = 0
        vector = []
    counter += 1
    #print "%r: %r" % (method.routing_key, body,)
    msg = json.loads( body )
    send_time = msg["putval"]["time"] 
    now = time.time()
    delta_time = now - send_time
    #print "delta: %f (%f - %f)" % (delta_time, now, send_time)
    #print json.dumps(msg, indent=2)
    vector.append(delta_time)
    #if delta_time > 60:
    #    print "delta: %f (%f - %f)" % (delta_time, now, send_time)
    #if delta_time < 58:
    #    print "delta: %f (%f - %f)" % (delta_time, now, send_time)

    if counter > thresh: 
        sys.exit(0)







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


channel.basic_consume(callback, queue=queue_name, no_ack=True)
channel.start_consuming()

print "done."
