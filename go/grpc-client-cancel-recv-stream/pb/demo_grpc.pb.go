// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DemoServiceClient is the client API for DemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DemoServiceClient interface {
	ListenHeartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (DemoService_ListenHeartbeatClient, error)
}

type demoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDemoServiceClient(cc grpc.ClientConnInterface) DemoServiceClient {
	return &demoServiceClient{cc}
}

func (c *demoServiceClient) ListenHeartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (DemoService_ListenHeartbeatClient, error) {
	stream, err := c.cc.NewStream(ctx, &DemoService_ServiceDesc.Streams[0], "/pb.DemoService/ListenHeartbeat", opts...)
	if err != nil {
		return nil, err
	}
	x := &demoServiceListenHeartbeatClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DemoService_ListenHeartbeatClient interface {
	Recv() (*Heartbeat, error)
	grpc.ClientStream
}

type demoServiceListenHeartbeatClient struct {
	grpc.ClientStream
}

func (x *demoServiceListenHeartbeatClient) Recv() (*Heartbeat, error) {
	m := new(Heartbeat)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DemoServiceServer is the server API for DemoService service.
// All implementations must embed UnimplementedDemoServiceServer
// for forward compatibility
type DemoServiceServer interface {
	ListenHeartbeat(*HeartbeatRequest, DemoService_ListenHeartbeatServer) error
	mustEmbedUnimplementedDemoServiceServer()
}

// UnimplementedDemoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDemoServiceServer struct {
}

func (UnimplementedDemoServiceServer) ListenHeartbeat(*HeartbeatRequest, DemoService_ListenHeartbeatServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenHeartbeat not implemented")
}
func (UnimplementedDemoServiceServer) mustEmbedUnimplementedDemoServiceServer() {}

// UnsafeDemoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DemoServiceServer will
// result in compilation errors.
type UnsafeDemoServiceServer interface {
	mustEmbedUnimplementedDemoServiceServer()
}

func RegisterDemoServiceServer(s grpc.ServiceRegistrar, srv DemoServiceServer) {
	s.RegisterService(&DemoService_ServiceDesc, srv)
}

func _DemoService_ListenHeartbeat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HeartbeatRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DemoServiceServer).ListenHeartbeat(m, &demoServiceListenHeartbeatServer{stream})
}

type DemoService_ListenHeartbeatServer interface {
	Send(*Heartbeat) error
	grpc.ServerStream
}

type demoServiceListenHeartbeatServer struct {
	grpc.ServerStream
}

func (x *demoServiceListenHeartbeatServer) Send(m *Heartbeat) error {
	return x.ServerStream.SendMsg(m)
}

// DemoService_ServiceDesc is the grpc.ServiceDesc for DemoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DemoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DemoService",
	HandlerType: (*DemoServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListenHeartbeat",
			Handler:       _DemoService_ListenHeartbeat_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "demo.proto",
}