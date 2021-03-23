#!/bin/env python

#from __future__ import print_function
import logging
import grpc

import service_pb2
import service_pb2_grpc

def run():
    ch = grpc.insecure_channel('localhost:9000')
    stub = service_pb2_grpc.SvcStub(ch)
    req = service_pb2.SvcRequest(req_text="service-request")
    logging.info("--- DoService() --->")
    resp = stub.DoService(req)
    logging.info("<-- received: {}".format(resp.resp_text))


if __name__ == '__main__':
    logging.basicConfig(format="%(asctime)s %(levelname)s: %(message)s", level=logging.DEBUG)
    logging.info('starting client')
    run()
    logging.info('done.')
