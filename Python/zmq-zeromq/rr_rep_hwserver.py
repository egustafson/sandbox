#!/usr/bin/env python
# ########################################
#
# Exmaple:  zmq Request/Response 
#
import zmq

print("libzmq version: {}".format(zmq.zmq_version()))
print(" pyzmq version: {}".format(zmq.__version__))


context = zmq.Context()

print("Creating a hello world server...")
socket = context.socket(zmq.REP)
socket.bind("tcp://127.0.0.1:5555")

while True:
    message = socket.recv()
    print("Received Hello")
    socket.send(b"World")
