// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: lib/proto/slave.proto

package proto

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

// SlaveGrpcServiceClient is the client API for SlaveGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SlaveGrpcServiceClient interface {
	Get(ctx context.Context, in *SlaveGetRequest, opts ...grpc.CallOption) (*SlaveGetResponse, error)
	Set(ctx context.Context, in *SlaveSetRequest, opts ...grpc.CallOption) (*SlaveSetResponse, error)
	Del(ctx context.Context, in *SlaveDelRequest, opts ...grpc.CallOption) (*SlaveDelResponse, error)
	Keys(ctx context.Context, in *SlaveKeysRequest, opts ...grpc.CallOption) (*SlaveKeysResponse, error)
	Exists(ctx context.Context, in *SlaveExistsRequest, opts ...grpc.CallOption) (*SlaveExistsResponse, error)
	Expire(ctx context.Context, in *SlaveExpireRequest, opts ...grpc.CallOption) (*SlaveExpireResponse, error)
	TTL(ctx context.Context, in *SlaveTTLRequest, opts ...grpc.CallOption) (*SlaveTTLResponse, error)
	Heartbeat(ctx context.Context, in *SlaveHeartbeatRequest, opts ...grpc.CallOption) (*SlaveHeartbeatResponse, error)
}

type slaveGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSlaveGrpcServiceClient(cc grpc.ClientConnInterface) SlaveGrpcServiceClient {
	return &slaveGrpcServiceClient{cc}
}

func (c *slaveGrpcServiceClient) Get(ctx context.Context, in *SlaveGetRequest, opts ...grpc.CallOption) (*SlaveGetResponse, error) {
	out := new(SlaveGetResponse)
	err := c.cc.Invoke(ctx, "/slave.SlaveGrpcService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slaveGrpcServiceClient) Set(ctx context.Context, in *SlaveSetRequest, opts ...grpc.CallOption) (*SlaveSetResponse, error) {
	out := new(SlaveSetResponse)
	err := c.cc.Invoke(ctx, "/slave.SlaveGrpcService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slaveGrpcServiceClient) Del(ctx context.Context, in *SlaveDelRequest, opts ...grpc.CallOption) (*SlaveDelResponse, error) {
	out := new(SlaveDelResponse)
	err := c.cc.Invoke(ctx, "/slave.SlaveGrpcService/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slaveGrpcServiceClient) Keys(ctx context.Context, in *SlaveKeysRequest, opts ...grpc.CallOption) (*SlaveKeysResponse, error) {
	out := new(SlaveKeysResponse)
	err := c.cc.Invoke(ctx, "/slave.SlaveGrpcService/Keys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slaveGrpcServiceClient) Exists(ctx context.Context, in *SlaveExistsRequest, opts ...grpc.CallOption) (*SlaveExistsResponse, error) {
	out := new(SlaveExistsResponse)
	err := c.cc.Invoke(ctx, "/slave.SlaveGrpcService/Exists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slaveGrpcServiceClient) Expire(ctx context.Context, in *SlaveExpireRequest, opts ...grpc.CallOption) (*SlaveExpireResponse, error) {
	out := new(SlaveExpireResponse)
	err := c.cc.Invoke(ctx, "/slave.SlaveGrpcService/Expire", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slaveGrpcServiceClient) TTL(ctx context.Context, in *SlaveTTLRequest, opts ...grpc.CallOption) (*SlaveTTLResponse, error) {
	out := new(SlaveTTLResponse)
	err := c.cc.Invoke(ctx, "/slave.SlaveGrpcService/TTL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slaveGrpcServiceClient) Heartbeat(ctx context.Context, in *SlaveHeartbeatRequest, opts ...grpc.CallOption) (*SlaveHeartbeatResponse, error) {
	out := new(SlaveHeartbeatResponse)
	err := c.cc.Invoke(ctx, "/slave.SlaveGrpcService/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SlaveGrpcServiceServer is the server API for SlaveGrpcService service.
// All implementations must embed UnimplementedSlaveGrpcServiceServer
// for forward compatibility
type SlaveGrpcServiceServer interface {
	Get(context.Context, *SlaveGetRequest) (*SlaveGetResponse, error)
	Set(context.Context, *SlaveSetRequest) (*SlaveSetResponse, error)
	Del(context.Context, *SlaveDelRequest) (*SlaveDelResponse, error)
	Keys(context.Context, *SlaveKeysRequest) (*SlaveKeysResponse, error)
	Exists(context.Context, *SlaveExistsRequest) (*SlaveExistsResponse, error)
	Expire(context.Context, *SlaveExpireRequest) (*SlaveExpireResponse, error)
	TTL(context.Context, *SlaveTTLRequest) (*SlaveTTLResponse, error)
	Heartbeat(context.Context, *SlaveHeartbeatRequest) (*SlaveHeartbeatResponse, error)
	mustEmbedUnimplementedSlaveGrpcServiceServer()
}

// UnimplementedSlaveGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSlaveGrpcServiceServer struct {
}

func (UnimplementedSlaveGrpcServiceServer) Get(context.Context, *SlaveGetRequest) (*SlaveGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSlaveGrpcServiceServer) Set(context.Context, *SlaveSetRequest) (*SlaveSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedSlaveGrpcServiceServer) Del(context.Context, *SlaveDelRequest) (*SlaveDelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Del not implemented")
}
func (UnimplementedSlaveGrpcServiceServer) Keys(context.Context, *SlaveKeysRequest) (*SlaveKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Keys not implemented")
}
func (UnimplementedSlaveGrpcServiceServer) Exists(context.Context, *SlaveExistsRequest) (*SlaveExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exists not implemented")
}
func (UnimplementedSlaveGrpcServiceServer) Expire(context.Context, *SlaveExpireRequest) (*SlaveExpireResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Expire not implemented")
}
func (UnimplementedSlaveGrpcServiceServer) TTL(context.Context, *SlaveTTLRequest) (*SlaveTTLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TTL not implemented")
}
func (UnimplementedSlaveGrpcServiceServer) Heartbeat(context.Context, *SlaveHeartbeatRequest) (*SlaveHeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedSlaveGrpcServiceServer) mustEmbedUnimplementedSlaveGrpcServiceServer() {}

// UnsafeSlaveGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SlaveGrpcServiceServer will
// result in compilation errors.
type UnsafeSlaveGrpcServiceServer interface {
	mustEmbedUnimplementedSlaveGrpcServiceServer()
}

func RegisterSlaveGrpcServiceServer(s grpc.ServiceRegistrar, srv SlaveGrpcServiceServer) {
	s.RegisterService(&SlaveGrpcService_ServiceDesc, srv)
}

func _SlaveGrpcService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlaveGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlaveGrpcServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slave.SlaveGrpcService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlaveGrpcServiceServer).Get(ctx, req.(*SlaveGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlaveGrpcService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlaveSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlaveGrpcServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slave.SlaveGrpcService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlaveGrpcServiceServer).Set(ctx, req.(*SlaveSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlaveGrpcService_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlaveDelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlaveGrpcServiceServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slave.SlaveGrpcService/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlaveGrpcServiceServer).Del(ctx, req.(*SlaveDelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlaveGrpcService_Keys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlaveKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlaveGrpcServiceServer).Keys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slave.SlaveGrpcService/Keys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlaveGrpcServiceServer).Keys(ctx, req.(*SlaveKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlaveGrpcService_Exists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlaveExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlaveGrpcServiceServer).Exists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slave.SlaveGrpcService/Exists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlaveGrpcServiceServer).Exists(ctx, req.(*SlaveExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlaveGrpcService_Expire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlaveExpireRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlaveGrpcServiceServer).Expire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slave.SlaveGrpcService/Expire",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlaveGrpcServiceServer).Expire(ctx, req.(*SlaveExpireRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlaveGrpcService_TTL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlaveTTLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlaveGrpcServiceServer).TTL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slave.SlaveGrpcService/TTL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlaveGrpcServiceServer).TTL(ctx, req.(*SlaveTTLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlaveGrpcService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlaveHeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlaveGrpcServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slave.SlaveGrpcService/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlaveGrpcServiceServer).Heartbeat(ctx, req.(*SlaveHeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SlaveGrpcService_ServiceDesc is the grpc.ServiceDesc for SlaveGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SlaveGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "slave.SlaveGrpcService",
	HandlerType: (*SlaveGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SlaveGrpcService_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _SlaveGrpcService_Set_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _SlaveGrpcService_Del_Handler,
		},
		{
			MethodName: "Keys",
			Handler:    _SlaveGrpcService_Keys_Handler,
		},
		{
			MethodName: "Exists",
			Handler:    _SlaveGrpcService_Exists_Handler,
		},
		{
			MethodName: "Expire",
			Handler:    _SlaveGrpcService_Expire_Handler,
		},
		{
			MethodName: "TTL",
			Handler:    _SlaveGrpcService_TTL_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _SlaveGrpcService_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lib/proto/slave.proto",
}
