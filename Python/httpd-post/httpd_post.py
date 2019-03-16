#!/usr/bin/env python3

import os
import shutil
import http.server

class PostHandler(http.server.BaseHTTPRequestHandler):
    def dump_Request(self):
        print("----------------------------------------")
        print("Received a {} command.".format(self.command))
        print(" from: {}:{}".format(str(self.client_address[0]), str(self.client_address[1])))
        print(" path: {}".format(self.path))
        for k in self.headers.keys():
            print("  {}: {}".format(k, self.headers[k]))

    def do_GET(self):
        self.send_response(501)
        return

    def do_PUT(self):
        self.send_response(501)
        return

    def do_POST(self):
        self.dump_Request()
        size = int(self.headers['content-length'])
        print(str(self.rfile.read(size), 'utf-8'))
        self.send_response(200, "OK")
        self.end_headers()
        return ""

if __name__ == '__main__':
    httpdAddress = ('', 8000)
    httpdHandler = PostHandler
    httpd = http.server.HTTPServer(httpdAddress, httpdHandler)
    httpd.serve_forever()
