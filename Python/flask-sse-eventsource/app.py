from gevent import monkey; monkey.patch_all()

import gevent
from gevent.wsgi import WSGIServer

from flask import Flask, Response

import datetime
import time

class ServerSentEvent(object):

    def __init__(self, data):
        self.data = data
        self.event = None
        self.id = None
        self.desc_map = {
            self.data : "data",
            self.event : "event",
            self.id : "id"
        }

    def encode(self):
        if not self.data:
            return ""
        lines = ["%s: %s" % (v, k)
                 for k, v in self.desc_map.iteritems() if k]
        return "%s\n\n" % "\n".join(lines)

app = Flask(__name__, static_folder=".")


@app.route("/")
def index():
    print("routing /")
    return app.send_static_file('index.html')


@app.route("/subscribe")
def subscribe():
    def gen():
        try:
            while True:
                now = datetime.datetime.now()
                msg = now.isoformat()
                ev = ServerSentEvent(msg)
                yield ev.encode()
                time.sleep(1)
        except GeneratorExit:
            ## exit point for "streaming" SSI route.
            print("subscription caught GeneratorExit.")

    return Response(gen(), mimetype="text/event-stream")


if __name__ == "__main__":
    app.debug = True
    server = WSGIServer(("",5000), app)
    server.serve_forever()
