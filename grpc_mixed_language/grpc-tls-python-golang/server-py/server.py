#!/usr/bin/env python

from concurrent import futures
import grpc
import logging
import time

import service_pb2
import service_pb2_grpc

USE_TLS = True

CA_FILE = "../test-ca.pem"
CERT_FILE = "../test-cert.pem"
KEY_FILE = "../test-key.pem"

def read_pem(path):
    with open(path, 'rb') as f:
        data = f.read()
    return data

class SvcService(service_pb2_grpc.SvcServicer):

    def init(self):
        pass

    def DoService(self, request, context):
        logging.info("DoService - received request: {}".format(request.req_text))
        return service_pb2.SvcResponse(resp_text="response-from-python-server")

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=3))
    service_pb2_grpc.add_SvcServicer_to_server(SvcService(), server)

    if USE_TLS:
        creds = grpc.ssl_server_credentials(
            ((read_pem(KEY_FILE), read_pem(CERT_FILE)),),
            root_certificates=read_pem(CA_FILE),
            require_client_auth=False
        )
        server.add_secure_port('localhost:9000', creds)
        logging.info("TLS listener created.")
    else:
        logging.info("insecure listener created.")
        server.add_insecure_port('localhost:9000')


    logging.info('service starting...')
    server.start()
    #server.wait_for_termination()
    while True:
        time.sleep(1)

if __name__ == '__main__':
    logging.basicConfig(format="%(asctime)s %(levelname)s: %(message)s", level=logging.DEBUG)
    serve() ## forever
    logging.info("server exited.")
