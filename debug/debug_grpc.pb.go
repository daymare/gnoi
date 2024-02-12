// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: debug/debug.proto

package debug

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

// DebugClient is the client API for Debug service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DebugClient interface {
	// Debug will execute the set of commands provided in the request.
	// The command will be executed in the mode provided.
	// All command modes must support an exit code on completion of the
	// command. (e.g. Cli modes must exit after sending a command)
	// Errors:
	//
	//	InvalidArgument: for unspecified mode
	Debug(ctx context.Context, in *DebugRequest, opts ...grpc.CallOption) (Debug_DebugClient, error)
}

type debugClient struct {
	cc grpc.ClientConnInterface
}

func NewDebugClient(cc grpc.ClientConnInterface) DebugClient {
	return &debugClient{cc}
}

func (c *debugClient) Debug(ctx context.Context, in *DebugRequest, opts ...grpc.CallOption) (Debug_DebugClient, error) {
	stream, err := c.cc.NewStream(ctx, &Debug_ServiceDesc.Streams[0], "/gnoi.debug.Debug/Debug", opts...)
	if err != nil {
		return nil, err
	}
	x := &debugDebugClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Debug_DebugClient interface {
	Recv() (*DebugResponse, error)
	grpc.ClientStream
}

type debugDebugClient struct {
	grpc.ClientStream
}

func (x *debugDebugClient) Recv() (*DebugResponse, error) {
	m := new(DebugResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DebugServer is the server API for Debug service.
// All implementations must embed UnimplementedDebugServer
// for forward compatibility
type DebugServer interface {
	// Debug will execute the set of commands provided in the request.
	// The command will be executed in the mode provided.
	// All command modes must support an exit code on completion of the
	// command. (e.g. Cli modes must exit after sending a command)
	// Errors:
	//
	//	InvalidArgument: for unspecified mode
	Debug(*DebugRequest, Debug_DebugServer) error
	mustEmbedUnimplementedDebugServer()
}

// UnimplementedDebugServer must be embedded to have forward compatible implementations.
type UnimplementedDebugServer struct {
}

func (UnimplementedDebugServer) Debug(*DebugRequest, Debug_DebugServer) error {
	return status.Errorf(codes.Unimplemented, "method Debug not implemented")
}
func (UnimplementedDebugServer) mustEmbedUnimplementedDebugServer() {}

// UnsafeDebugServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DebugServer will
// result in compilation errors.
type UnsafeDebugServer interface {
	mustEmbedUnimplementedDebugServer()
}

func RegisterDebugServer(s grpc.ServiceRegistrar, srv DebugServer) {
	s.RegisterService(&Debug_ServiceDesc, srv)
}

func _Debug_Debug_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DebugRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServer).Debug(m, &debugDebugServer{stream})
}

type Debug_DebugServer interface {
	Send(*DebugResponse) error
	grpc.ServerStream
}

type debugDebugServer struct {
	grpc.ServerStream
}

func (x *debugDebugServer) Send(m *DebugResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Debug_ServiceDesc is the grpc.ServiceDesc for Debug service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Debug_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.debug.Debug",
	HandlerType: (*DebugServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Debug",
			Handler:       _Debug_Debug_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "debug/debug.proto",
}
