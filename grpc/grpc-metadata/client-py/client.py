#!/bin/env python

# from __future__ import print_function
import logging
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
    doRequest(stub, "random-req-msg-python")
    doRequest(stub, "ok-req-msg")
    doRequest(stub, "err-req-msg")
    doRequest(stub, "err-internal-req-msg")
    doRequest(stub, "err-abort-req-msg")
    doRequest(stub, "err-timeout")  # server will delay response for 1 hr


def doRequest(stub: service_pb2_grpc.SvcStub, msg: str) -> None:
    req = service_pb2.SvcRequest(req_text=msg)
    try:
        resp = stub.DoService(req, timeout=TIMEOUT_DURATION)  # <-- timeout
        logging.info("response: {}".format(resp.resp_text))
    except grpc.RpcError as e:
        logStatus(e)
    except BaseException as e:
        logging.warn("unexpected exception:  {}".format(e))
        raise e


def logStatus(e: grpc.RpcError) -> None:
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
