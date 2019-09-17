from flask import Flask
from flask_restplus import Resource, Api

import os

app = Flask(__name__)
api = Api(app)

@api.route('/env')
class Environment(Resource):
    def get(self):
        e = {}
        for (k,v) in os.environ.items():
            e[k] = v
        return e

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000, debug=True)
