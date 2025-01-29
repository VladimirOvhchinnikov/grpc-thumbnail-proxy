// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: transport.proto

package transport

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TransportService_SendData_FullMethodName = "/transport.TransportService/SendData"
)

// TransportServiceClient is the client API for TransportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Определение сервиса
type TransportServiceClient interface {
	// RPC метод для отправки данных на сервер
	SendData(ctx context.Context, in *SendDataRequest, opts ...grpc.CallOption) (*SendDataResponse, error)
}

type transportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransportServiceClient(cc grpc.ClientConnInterface) TransportServiceClient {
	return &transportServiceClient{cc}
}

func (c *transportServiceClient) SendData(ctx context.Context, in *SendDataRequest, opts ...grpc.CallOption) (*SendDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendDataResponse)
	err := c.cc.Invoke(ctx, TransportService_SendData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransportServiceServer is the server API for TransportService service.
// All implementations must embed UnimplementedTransportServiceServer
// for forward compatibility.
//
// Определение сервиса
type TransportServiceServer interface {
	// RPC метод для отправки данных на сервер
	SendData(context.Context, *SendDataRequest) (*SendDataResponse, error)
	mustEmbedUnimplementedTransportServiceServer()
}

// UnimplementedTransportServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransportServiceServer struct{}

func (UnimplementedTransportServiceServer) SendData(context.Context, *SendDataRequest) (*SendDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendData not implemented")
}
func (UnimplementedTransportServiceServer) mustEmbedUnimplementedTransportServiceServer() {}
func (UnimplementedTransportServiceServer) testEmbeddedByValue()                          {}

// UnsafeTransportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransportServiceServer will
// result in compilation errors.
type UnsafeTransportServiceServer interface {
	mustEmbedUnimplementedTransportServiceServer()
}

func RegisterTransportServiceServer(s grpc.ServiceRegistrar, srv TransportServiceServer) {
	// If the following call pancis, it indicates UnimplementedTransportServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TransportService_ServiceDesc, srv)
}

func _TransportService_SendData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportServiceServer).SendData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransportService_SendData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportServiceServer).SendData(ctx, req.(*SendDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransportService_ServiceDesc is the grpc.ServiceDesc for TransportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transport.TransportService",
	HandlerType: (*TransportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendData",
			Handler:    _TransportService_SendData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport.proto",
}
