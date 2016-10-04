from gevent import monkey
monkey.patch_all()

from flask import Flask
from flask_socketio import SocketIO

app = Flask(__name__, static_folder='.')
socketio = SocketIO(app)

ping_count = 0

@app.route('/')
def main():
    print("routing /")
    return app.send_static_file('index.html')


@app.route('/ping')
def ping():
    global ping_count
    ping_count += 1
    socketio.emit('msg', {'count': ping_count}, namespace='/ws')
    return("Pong.")


@socketio.on('connect', namespace='/ws')
def ws_conn():
    global ping_count
    print("ws connect.")
    socketio.emit('msg', {'count': ping_count}, namespace='/ws')


@socketio.on('disconect', namespace='/ws')
def ws_disconn():
    print("ws disconnect.")


if __name__ == '__main__':
    socketio.run(app, "0.0.0.0", port=5000, debug=True)


## Local Variables:
## mode: python
## End:
