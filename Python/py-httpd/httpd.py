#!/usr/bin/env python

import os
import shutil
import BaseHTTPServer

class MyHandler(BaseHTTPServer.BaseHTTPRequestHandler):
    def dump_Request(self):
        print "----------------------------------------"
        print "Received a %s command." % self.command
        print " from: %s:%s" % (str(self.client_address[0]), str(self.client_address[1]))
        print " path: %s" % self.path
        for k in self.headers.keys():
            print "  %s: %s" % (k, self.headers[k])
    
    def do_GET(self):
        # This routine has no exception handling so that it is
        # easy to debug.
        #
        self.dump_Request()
        filename = self.path.strip("/")
        print "Opening '%s'." % filename
        f = open(filename, 'rb')
        size = str(os.fstat(f.fileno())[6])
        print "'%s'is %s bytes." % (filename, size)
        self.send_response(200)
        self.send_header("Content-Length", size)
        self.end_headers()
        shutil.copyfileobj(f, self.wfile)
        f.close()
        return

    def do_PUT(self):
        self.dump_Request()
        block_size = 4096
        size = int(self.headers['content-length'])
        filename = self.path.strip("/")
        print "Opening '%s' for %d bytes." % (filename, size)
        bread = 0
        f = open(filename, 'wb')
        while bread < size:
            rsize = min(size-bread, block_size)
            block = self.rfile.read(rsize)
            bread += len(block)
            f.write(block)
        f.close()
        
        self.send_response(200)
        return

    def do_POST(self):
        self.send_error(501)
        return

if __name__ == '__main__':
    httpdAddress = ('', 8000)
    httpdHandler = MyHandler
    httpd = BaseHTTPServer.HTTPServer(httpdAddress, httpdHandler)
    httpd.serve_forever()
