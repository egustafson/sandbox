#!/usr/bin/env python
# ########################################
#
# Example:  zmq Request/Response
#
import zmq

context = zmq.Context()

print("Connecting to hello world server...")
socket = context.socket(zmq.REQ)
socket.connect("tcp://localhost:5555")

for request in range(10):
    print("Sending request {} ...".format(request))
    socket.send(b"Hello")

    message = socket.recv()
    print("Received reply {} [ {} ]".format(request, message))
