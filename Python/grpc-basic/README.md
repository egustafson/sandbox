Python gRPC Example - Basic gRPC
================================

This is an opinionated example of how to create a simple Python gRPC
client and server.  The two key "opinions" in this example are:

1. Create a library like package that includes the gRPC stubs.
2. Place the `.proto` files in a language agnostic directory that
   could be exported or included to/from another git project -- like
   bindings in another language.

Prerequisites
-------------

Have the Python IDL compiler for gRPC installed (aka 'protoc').  You
do not have to install the `protoc` executable.  simply install the
Python packages in `requirements.txt`.

Usage
-----

### Build
```
> make
mkdir -p pb/demo
(cd pb/demo; ln -s ../demo.proto .)
python -m grpc_tools.protoc -Ipb --python_out=. --grpc_python_out=. pb/demo/demo.proto
>
```

### Running
In two separate windows, _client_ and _server_.

<ins>Server:</ins> (press ctl-c to exit)
```
> python server.py
2021-03-18 16:03:49,334 INFO: starting server
```

<ins>Client:</ins>
```
> python client.py
2021-03-18 16:08:26,013 INFO: starting client
2021-03-18 16:08:26,049 INFO: received: heartbeat-note
...
2021-03-18 16:08:26,063 INFO: received: heartbeat-note
2021-03-18 16:08:26,065 INFO: done.
> 
```

