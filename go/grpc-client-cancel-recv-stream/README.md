grpc-client-cancel-recv-stream
------------------------------

This example demonstrates a gRPC client blocked on Recv() in a stream
response and a separate thread invoking Cancel() on the stream's
context.

Expected results: The Recv() call unblocks and returns grpc
codes.Canceled.  -- this code does confirm this.
