#!/usr/bin/env python

import threading
import time

import SocketServer

##
##
##

class TcpRequestHandler(SocketServer.BaseRequestHandler):

    def old_handle(self):
        # for TCP self.request is a socket object
        #
        data = self.request.recv(4096)
        response = "bye"
        self.request.sendall(response)
        self.request.close()
        print("exiting handler.")

    def handle(self):
        # for TCP self.request is a socket object
        #
        data = self.request.recv(4096)
        while data:
            self.request.sendall(data)
            data = self.request.recv(4096)

class ThreadedTCPServer(SocketServer.ThreadingMixIn, SocketServer.TCPServer):
    pass

def start_server(sockAddr):
    server = ThreadedTCPServer(sockAddr, TcpRequestHandler)

    server_thread = threading.Thread(target=server.serve_forever)
    server_thread.daemon = True
    server_thread.start()
    return server

##
## Main
##
if __name__ == '__main__':
    sockAddr = ('', 10001)
    server = start_server(sockAddr)
    ii = 0
    while ii < 6:
        time.sleep(10)
        print("tick.")
        ii += 1
    print("shutting down.")
    server.shutdown()
    print("done.")

