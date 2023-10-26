#!/bin/env python

# from __future__ import print_function
import logging
import os
import socket
import sys

import grpc
from grpc_status import rpc_status

from google.rpc import error_details_pb2

import service_pb2
import service_pb2_grpc


LISTEN_ADDR = 'localhost:9000'
TIMEOUT_DURATION = 10  # 10 seconds


def run():

    logging.info("connecting to {}".format(LISTEN_ADDR))
    ch = grpc.insecure_channel(LISTEN_ADDR)

    stub = service_pb2_grpc.SvcStub(ch)
    doRequest(stub, "client-request-py")
    ch.close()


def doRequest(stub: service_pb2_grpc.SvcStub, msg: str) -> None:
    """Invoke `SvcRequest` with `msg`, adding call headers and logging response headers"""
    hostname = socket.gethostname()
    pid = "{}".format(os.getpid())
    md = (  # call metadata
        ('hostname', hostname),
        ('client-pid', pid),
        ('python-version', sys.version)
    )

    req = service_pb2.SvcRequest(req_text=msg)
    try:
        resp, call = stub.DoService.with_call(req, metadata=md, timeout=TIMEOUT_DURATION)  # <-- timeout

        logging.info("response: {}".format(resp.resp_text))
        logging.info("  Header:")
        for key, value in call.initial_metadata():
            logging.info("  - {}: {}".format(key, value))
        logging.info("  Footer:")
        for key, value in call.trailing_metadata():
            logging.info("  - {}: {}".format(key, value))
    except grpc.RpcError as e:
        logStatus(e)
    except BaseException as e:
        logging.warn("unexpected exception:  {}".format(e))
        raise e


def logStatus(e: grpc.RpcError) -> None:
    """Dissect the response grpc error and log the error's component parts"""
    logging.info("----------------------------------")
    # logging.info("gRPC returned error: {}".format(e))
    # logging.info("  trailing_metadata:  {}".format(e.trailing_metadata()))
    print("----")
    code = e.code()
    print("  code:    {} ({})".format(code.name, code.value[0]))
    details = e.details()
    print("  details: {}".format(details))
    err_string = e.debug_error_string()
    print("  err str: {}".format(err_string))
    status = rpc_status.from_call(e)
    if status is not None:
        for detail in status.details:
            if detail.Is(error_details_pb2.ErrorInfo.DESCRIPTOR):
                errinfo = error_details_pb2.ErrorInfo()
                detail.Unpack(errinfo)
                print("  ErrorInfo:")
                print("    Reason: {}".format(errinfo.reason))
                print("    Domain: {}".format(errinfo.domain))
                print("    Metadata:")
                for k, v in errinfo.metadata.items():
                    print("      {}: {}".format(k, v))
                # logging.info("  ErrorInfo: {}".format(errinfo))
            else:
                print("  -- unknown details type.")


if __name__ == '__main__':
    logging.basicConfig(format="%(asctime)s %(levelname)s: %(message)s",
                        level=logging.DEBUG)
    logging.info('starting client')
    run()
    logging.info('done.')
