# -*- coding: utf-8 -*-

# This is a simple flask application that will act as the service
# under test.

from flask import Flask, jsonify

app = Flask(__name__)

@app.route("/")
def landing_page():
    return jsonify(status='ok')
