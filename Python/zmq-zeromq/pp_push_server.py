#!/usr/bin/env python
# ########################################
#
# Example:  zmq Push/Pull (Pipeline)
#
import zmq

print("libzmq version: {}".format(zmq.zmq_version()))
print(" pyzmq version: {}".format(zmq.__version__))


context = zmq.Context()

print("Creating a (push) producer...")
socket = context.socket(zmq.PUSH)
socket.bind("tcp://*:5559")

seqnum = 0
while True:
    seqnum += 1
    msg = "seq: {}".format(seqnum)
    socket.send_string(msg)
    print(msg)

