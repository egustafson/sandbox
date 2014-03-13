#!/usr/bin/env python
# ########################################
#
# Exmaple:  zmq Pub/Sub
#
import zmq

print("libzmq version: {}".format(zmq.zmq_version()))
print(" pyzmq version: {}".format(zmq.__version__))


context = zmq.Context()

print("Creating publisher...")
socket = context.socket(zmq.PUB)
socket.connect("tcp://localhost:5557")
#
# Multiple points to publish to by adding connections
#
#socket.connect("tcp://localhost:5558")

seqnum = 0
while True:
    seqnum += 1
    msg = "seq: {}".format(seqnum)
    socket.send_string(msg)
    print(msg)

