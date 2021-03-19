#!/bin/env python

from __future__ import print_function
import logging
import grpc

from demo import demo_pb2
from demo import demo_pb2_grpc

def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = demo_pb2_grpc.DemoServiceStub(channel)
        for ii in range(2) :
            ## Req / Resp (single objects)
            req = demo_pb2.HeartbeatRequest(request_id="req-id-{}".format(ii))
            logging.info("--- ListenHeartbeat() --------->")
            resp = stub.ListenHeartbeat(req)
            logging.info("<-- received: {}".format(resp.note))

        ## 1-Req / Stream Resp
        logging.info("--- StreamHeartbeat() --------->")
        req = demo_pb2.HeartbeatRequest(request_id="stream-1")
        for resp in stub.StreamHeartbeat(req):
            logging.info("<-- received: {}".format(resp.note))


if __name__ == '__main__':
    logging.basicConfig(format="%(asctime)s %(levelname)s: %(message)s", level=logging.DEBUG)
    logging.info('starting client')
    run()
    logging.info('done.')
