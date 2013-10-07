#!/usr/bin/env python

import SimpleHTTPServer
import SocketServer

PORT = 8000

Handler = SimpleHTTPServer.SimpleHTTPRequestHandler

class MyHttpHandler(SimpleHTTPServer.SimpleHTTPRequestHandler):

    def do_GET(self):
        print "Request from:"
        print self.client_address
        print self.path
        print self.headers
        self.wfile.write("Response Text.\n")

httpd = SocketServer.TCPServer(("127.0.0.1", PORT), MyHttpHandler)

print "serving at port", PORT
httpd.serve_forever()

print "done."
