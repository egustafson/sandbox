#!/usr/bin/env python

import socket

HOST = '127.0.0.1'
PORT = 10001

if __name__ == '__main__':
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.connect((HOST, PORT))
    s.sendall('hello')
    data = s.recv(4096)
    s.close()
    print( repr(data) )
    print("done.")
