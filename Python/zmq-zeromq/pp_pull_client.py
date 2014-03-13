#!/usr/bin/env python
# ########################################
#
# Example:  zmq Push/Pull (Pipeline)
#
import zmq

print("libzmq version: {}".format(zmq.zmq_version()))
print(" pyzmq version: {}".format(zmq.__version__))


context = zmq.Context()

print("Creating a (pull) consumer...")
socket = context.socket(zmq.PULL)
socket.connect("tcp://localhost:5559")

while True:
    msg = socket.recv_string()
    print(msg)
