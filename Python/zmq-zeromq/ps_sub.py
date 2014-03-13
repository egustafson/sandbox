#!/usr/bin/env python
# ########################################
#
# Exmaple:  zmq Pub/Sub
#
import zmq

print("libzmq version: {}".format(zmq.zmq_version()))
print(" pyzmq version: {}".format(zmq.__version__))


context = zmq.Context()

print("Creating subscriber...")
socket = context.socket(zmq.SUB)
socket.connect("tcp://localhost:5556")

socket.setsockopt_string(zmq.SUBSCRIBE, u"")

while True:
    msg = socket.recv_string()
    print(msg)
