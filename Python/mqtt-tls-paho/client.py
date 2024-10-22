#!/usr/bin/env python

import random
import sys
import time

from paho.mqtt import client as mqtt_client

broker = 'dev0.elfwerks'
port = 8883
topic = "sandbox/python/mqtt"
client_id = f'python-mqtt-{random.randint(0,1000)}'
# username = 'sandbox'
# password = 'sandbox'

def on_connect(client, userdata, flags, rc, properties):
    if rc == 0:
        print("Connected to MQTT broker")
    else:
        print(f"Failed to connect, return code: {rc}")

def on_disconnect(client, userdata, flags, rc, properties):
    print(f"disconnecting, return code: {rc}")

def on_message(client, userdata, msg):
    print(f"recv '{msg.payload.decode()}' from '{msg.topic}' topic")

def connect_mqtt():
    client = mqtt_client.Client(client_id=client_id, callback_api_version=mqtt_client.CallbackAPIVersion.VERSION2)

    # client.username_pw_set(username, password)
    client.on_connect = on_connect
    client.on_disconnect = on_disconnect
    client.tls_set(ca_certs='./ca.pem')
    client.connect(broker, port)
    return client

def publish(client):
    for msg_count in range(5):
        msg = f"message({msg_count})"
        result = client.publish(topic, msg)
        status = result[0]
        if status == 0:
            print(f"sent '{msg}' to topic '{topic}'")
        else:
            print(f"failed to send message to topic '{topic}'")

def subscribe(client):
    client.on_message=on_message
    client.subscribe(topic)


if __name__ == "__main__":

    if len(sys.argv) < 2:
        print("Usage: client [send|recv]")
        sys.exit(1)

    c = connect_mqtt()
    c.loop_start()
    time.sleep(0.1)

    if sys.argv[1] == 'send':
        publish(c)
        time.sleep(1)
    elif sys.argv[1] == 'recv':
        subscribe(c)
        time.sleep(60)
    else:
        print("Usage: client [send|recv]")

    c.loop_stop()
    c.disconnect()
    time.sleep(0.2)
    print("done.")
