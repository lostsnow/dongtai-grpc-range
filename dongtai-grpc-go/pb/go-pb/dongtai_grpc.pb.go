// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package DontaiGRPC

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

// GRPCServiceClient is the client API for GRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GRPCServiceClient interface {
	Send(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type gRPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCServiceClient(cc grpc.ClientConnInterface) GRPCServiceClient {
	return &gRPCServiceClient{cc}
}

func (c *gRPCServiceClient) Send(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/DontaiGRPC.GRPCService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCServiceServer is the server API for GRPCService service.
// All implementations must embed UnimplementedGRPCServiceServer
// for forward compatibility
type GRPCServiceServer interface {
	Send(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedGRPCServiceServer()
}

// UnimplementedGRPCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGRPCServiceServer struct {
}

func (UnimplementedGRPCServiceServer) Send(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedGRPCServiceServer) mustEmbedUnimplementedGRPCServiceServer() {}

// UnsafeGRPCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GRPCServiceServer will
// result in compilation errors.
type UnsafeGRPCServiceServer interface {
	mustEmbedUnimplementedGRPCServiceServer()
}

func RegisterGRPCServiceServer(s grpc.ServiceRegistrar, srv GRPCServiceServer) {
	s.RegisterService(&GRPCService_ServiceDesc, srv)
}

func _GRPCService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DontaiGRPC.GRPCService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCServiceServer).Send(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// GRPCService_ServiceDesc is the grpc.ServiceDesc for GRPCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GRPCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DontaiGRPC.GRPCService",
	HandlerType: (*GRPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _GRPCService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dongtai.proto",
}
