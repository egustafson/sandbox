# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import service_pb2 as service__pb2


class SvcStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.DoService = channel.unary_unary(
                '/pb.Svc/DoService',
                request_serializer=service__pb2.SvcRequest.SerializeToString,
                response_deserializer=service__pb2.SvcResponse.FromString,
                )


class SvcServicer(object):
    """Missing associated documentation comment in .proto file."""

    def DoService(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SvcServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'DoService': grpc.unary_unary_rpc_method_handler(
                    servicer.DoService,
                    request_deserializer=service__pb2.SvcRequest.FromString,
                    response_serializer=service__pb2.SvcResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'pb.Svc', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Svc(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def DoService(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/pb.Svc/DoService',
            service__pb2.SvcRequest.SerializeToString,
            service__pb2.SvcResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)