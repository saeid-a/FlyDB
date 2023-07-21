// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package raftPB

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

// RaftServiceClient is the client API for RaftService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaftServiceClient interface {
	RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error)
	AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error)
	AppendEntriesPipeline(ctx context.Context, opts ...grpc.CallOption) (RaftService_AppendEntriesPipelineClient, error)
	TimeoutNow(ctx context.Context, in *TimeoutNowRequest, opts ...grpc.CallOption) (*TimeoutNowResponse, error)
	InstallSnapshot(ctx context.Context, in *InstallSnapshotRequest, opts ...grpc.CallOption) (*InstallSnapshotResponse, error)
}

type raftServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftServiceClient(cc grpc.ClientConnInterface) RaftServiceClient {
	return &raftServiceClient{cc}
}

func (c *raftServiceClient) RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error) {
	out := new(RequestVoteResponse)
	err := c.cc.Invoke(ctx, "/raft_pb.RaftService/RequestVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftServiceClient) AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error) {
	out := new(AppendEntriesResponse)
	err := c.cc.Invoke(ctx, "/raft_pb.RaftService/AppendEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftServiceClient) AppendEntriesPipeline(ctx context.Context, opts ...grpc.CallOption) (RaftService_AppendEntriesPipelineClient, error) {
	stream, err := c.cc.NewStream(ctx, &RaftService_ServiceDesc.Streams[0], "/raft_pb.RaftService/AppendEntriesPipeline", opts...)
	if err != nil {
		return nil, err
	}
	x := &raftServiceAppendEntriesPipelineClient{stream}
	return x, nil
}

type RaftService_AppendEntriesPipelineClient interface {
	Send(*AppendEntriesRequest) error
	Recv() (*AppendEntriesResponse, error)
	grpc.ClientStream
}

type raftServiceAppendEntriesPipelineClient struct {
	grpc.ClientStream
}

func (x *raftServiceAppendEntriesPipelineClient) Send(m *AppendEntriesRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *raftServiceAppendEntriesPipelineClient) Recv() (*AppendEntriesResponse, error) {
	m := new(AppendEntriesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *raftServiceClient) TimeoutNow(ctx context.Context, in *TimeoutNowRequest, opts ...grpc.CallOption) (*TimeoutNowResponse, error) {
	out := new(TimeoutNowResponse)
	err := c.cc.Invoke(ctx, "/raft_pb.RaftService/TimeoutNow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftServiceClient) InstallSnapshot(ctx context.Context, in *InstallSnapshotRequest, opts ...grpc.CallOption) (*InstallSnapshotResponse, error) {
	out := new(InstallSnapshotResponse)
	err := c.cc.Invoke(ctx, "/raft_pb.RaftService/InstallSnapshot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaftServiceServer is the server API for RaftService service.
// All implementations must embed UnimplementedRaftServiceServer
// for forward compatibility
type RaftServiceServer interface {
	RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error)
	AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error)
	AppendEntriesPipeline(RaftService_AppendEntriesPipelineServer) error
	TimeoutNow(context.Context, *TimeoutNowRequest) (*TimeoutNowResponse, error)
	InstallSnapshot(context.Context, *InstallSnapshotRequest) (*InstallSnapshotResponse, error)
	mustEmbedUnimplementedRaftServiceServer()
}

// UnimplementedRaftServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRaftServiceServer struct {
}

func (UnimplementedRaftServiceServer) RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestVote not implemented")
}
func (UnimplementedRaftServiceServer) AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendEntries not implemented")
}
func (UnimplementedRaftServiceServer) AppendEntriesPipeline(RaftService_AppendEntriesPipelineServer) error {
	return status.Errorf(codes.Unimplemented, "method AppendEntriesPipeline not implemented")
}
func (UnimplementedRaftServiceServer) TimeoutNow(context.Context, *TimeoutNowRequest) (*TimeoutNowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TimeoutNow not implemented")
}
func (UnimplementedRaftServiceServer) InstallSnapshot(context.Context, *InstallSnapshotRequest) (*InstallSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallSnapshot not implemented")
}
func (UnimplementedRaftServiceServer) mustEmbedUnimplementedRaftServiceServer() {}

// UnsafeRaftServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaftServiceServer will
// result in compilation errors.
type UnsafeRaftServiceServer interface {
	mustEmbedUnimplementedRaftServiceServer()
}

func RegisterRaftServiceServer(s grpc.ServiceRegistrar, srv RaftServiceServer) {
	s.RegisterService(&RaftService_ServiceDesc, srv)
}

func _RaftService_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServiceServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft_pb.RaftService/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServiceServer).RequestVote(ctx, req.(*RequestVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftService_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServiceServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft_pb.RaftService/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServiceServer).AppendEntries(ctx, req.(*AppendEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftService_AppendEntriesPipeline_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RaftServiceServer).AppendEntriesPipeline(&raftServiceAppendEntriesPipelineServer{stream})
}

type RaftService_AppendEntriesPipelineServer interface {
	Send(*AppendEntriesResponse) error
	Recv() (*AppendEntriesRequest, error)
	grpc.ServerStream
}

type raftServiceAppendEntriesPipelineServer struct {
	grpc.ServerStream
}

func (x *raftServiceAppendEntriesPipelineServer) Send(m *AppendEntriesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *raftServiceAppendEntriesPipelineServer) Recv() (*AppendEntriesRequest, error) {
	m := new(AppendEntriesRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RaftService_TimeoutNow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeoutNowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServiceServer).TimeoutNow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft_pb.RaftService/TimeoutNow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServiceServer).TimeoutNow(ctx, req.(*TimeoutNowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftService_InstallSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstallSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServiceServer).InstallSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft_pb.RaftService/InstallSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServiceServer).InstallSnapshot(ctx, req.(*InstallSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RaftService_ServiceDesc is the grpc.ServiceDesc for RaftService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RaftService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "raft_pb.RaftService",
	HandlerType: (*RaftServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestVote",
			Handler:    _RaftService_RequestVote_Handler,
		},
		{
			MethodName: "AppendEntries",
			Handler:    _RaftService_AppendEntries_Handler,
		},
		{
			MethodName: "TimeoutNow",
			Handler:    _RaftService_TimeoutNow_Handler,
		},
		{
			MethodName: "InstallSnapshot",
			Handler:    _RaftService_InstallSnapshot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AppendEntriesPipeline",
			Handler:       _RaftService_AppendEntriesPipeline_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "raft.proto",
}