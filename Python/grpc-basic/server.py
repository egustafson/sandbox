#!/bin/env python

from concurrent import futures
import logging
import grpc

from demo import demo_pb2
from demo import demo_pb2_grpc

class DemoService(demo_pb2_grpc.DemoServiceServicer):

    def init(self):
        pass

    def ListenHeartbeat(self, request, context):
        logging.info("ListenHeartbeat - received request: {}".format(request.request_id))
        return demo_pb2.Heartbeat(note="heartbeat-note")

    def StreamHeartbeat(self, request, context):
        logging.info("StreamHeartbeat - received request: {}".format(request.request_id))
        for ii in range(3):
            resp = demo_pb2.Heartbeat(note="stream-response-{}/{}".format(ii, request.request_id))
            yield resp


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=3))
    demo_pb2_grpc.add_DemoServiceServicer_to_server(DemoService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig(format="%(asctime)s %(levelname)s: %(message)s",level=logging.DEBUG)
    logging.info('starting server')
    serve() ## forever
    logging.info("done.")
