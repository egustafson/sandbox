#!/bin/env python

#from __future__ import print_function
import logging
import grpc

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

def run():

    creds = grpc.ssl_channel_credentials(
        root_certificates=read_pem(CA_FILE),
        private_key=read_pem(KEY_FILE),
        certificate_chain=read_pem(CERT_FILE)
    )
#     creds = grpc.ssl_channel_credentials(root_certificates=read_pem(CA_FILE))

    if USE_TLS:
        logging.info("using TLS")
        ch = grpc.secure_channel('127.0.0.1:9000', creds)
    else:
        logging.info("TLS Disabled")
        ch = grpc.insecure_channel('localhost:9000')

    stub = service_pb2_grpc.SvcStub(ch)
    req = service_pb2.SvcRequest(req_text="service-request-python")
    logging.info("--- DoService() --->")
    resp = stub.DoService(req, timeout=600)
    logging.info("<-- received: {}".format(resp.resp_text))


if __name__ == '__main__':
    logging.basicConfig(format="%(asctime)s %(levelname)s: %(message)s", level=logging.DEBUG)
    logging.info('starting client')
    run()
    logging.info('done.')
