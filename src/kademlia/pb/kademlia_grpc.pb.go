// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: kademlia.proto

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

// FullNodeClient is the client API for FullNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FullNodeClient interface {
	Ping(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Node, error)
	Store(ctx context.Context, opts ...grpc.CallOption) (FullNode_StoreClient, error)
	FindNode(ctx context.Context, in *TargetID, opts ...grpc.CallOption) (*KBucket, error)
	FindValue(ctx context.Context, in *TargetID, opts ...grpc.CallOption) (FullNode_FindValueClient, error)
}

type fullNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewFullNodeClient(cc grpc.ClientConnInterface) FullNodeClient {
	return &fullNodeClient{cc}
}

func (c *fullNodeClient) Ping(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/kademlia.FullNode/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fullNodeClient) Store(ctx context.Context, opts ...grpc.CallOption) (FullNode_StoreClient, error) {
	stream, err := c.cc.NewStream(ctx, &FullNode_ServiceDesc.Streams[0], "/kademlia.FullNode/Store", opts...)
	if err != nil {
		return nil, err
	}
	x := &fullNodeStoreClient{stream}
	return x, nil
}

type FullNode_StoreClient interface {
	Send(*Data) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type fullNodeStoreClient struct {
	grpc.ClientStream
}

func (x *fullNodeStoreClient) Send(m *Data) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fullNodeStoreClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fullNodeClient) FindNode(ctx context.Context, in *TargetID, opts ...grpc.CallOption) (*KBucket, error) {
	out := new(KBucket)
	err := c.cc.Invoke(ctx, "/kademlia.FullNode/FindNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fullNodeClient) FindValue(ctx context.Context, in *TargetID, opts ...grpc.CallOption) (FullNode_FindValueClient, error) {
	stream, err := c.cc.NewStream(ctx, &FullNode_ServiceDesc.Streams[1], "/kademlia.FullNode/FindValue", opts...)
	if err != nil {
		return nil, err
	}
	x := &fullNodeFindValueClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FullNode_FindValueClient interface {
	Recv() (*FindValueResponse, error)
	grpc.ClientStream
}

type fullNodeFindValueClient struct {
	grpc.ClientStream
}

func (x *fullNodeFindValueClient) Recv() (*FindValueResponse, error) {
	m := new(FindValueResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FullNodeServer is the server API for FullNode service.
// All implementations must embed UnimplementedFullNodeServer
// for forward compatibility
type FullNodeServer interface {
	Ping(context.Context, *Node) (*Node, error)
	Store(FullNode_StoreServer) error
	FindNode(context.Context, *TargetID) (*KBucket, error)
	FindValue(*TargetID, FullNode_FindValueServer) error
	mustEmbedUnimplementedFullNodeServer()
}

// UnimplementedFullNodeServer must be embedded to have forward compatible implementations.
type UnimplementedFullNodeServer struct {
}

func (UnimplementedFullNodeServer) Ping(context.Context, *Node) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedFullNodeServer) Store(FullNode_StoreServer) error {
	return status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (UnimplementedFullNodeServer) FindNode(context.Context, *TargetID) (*KBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNode not implemented")
}
func (UnimplementedFullNodeServer) FindValue(*TargetID, FullNode_FindValueServer) error {
	return status.Errorf(codes.Unimplemented, "method FindValue not implemented")
}
func (UnimplementedFullNodeServer) mustEmbedUnimplementedFullNodeServer() {}

// UnsafeFullNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FullNodeServer will
// result in compilation errors.
type UnsafeFullNodeServer interface {
	mustEmbedUnimplementedFullNodeServer()
}

func RegisterFullNodeServer(s grpc.ServiceRegistrar, srv FullNodeServer) {
	s.RegisterService(&FullNode_ServiceDesc, srv)
}

func _FullNode_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FullNodeServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kademlia.FullNode/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FullNodeServer).Ping(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _FullNode_Store_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FullNodeServer).Store(&fullNodeStoreServer{stream})
}

type FullNode_StoreServer interface {
	SendAndClose(*Response) error
	Recv() (*Data, error)
	grpc.ServerStream
}

type fullNodeStoreServer struct {
	grpc.ServerStream
}

func (x *fullNodeStoreServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fullNodeStoreServer) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FullNode_FindNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TargetID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FullNodeServer).FindNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kademlia.FullNode/FindNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FullNodeServer).FindNode(ctx, req.(*TargetID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FullNode_FindValue_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TargetID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FullNodeServer).FindValue(m, &fullNodeFindValueServer{stream})
}

type FullNode_FindValueServer interface {
	Send(*FindValueResponse) error
	grpc.ServerStream
}

type fullNodeFindValueServer struct {
	grpc.ServerStream
}

func (x *fullNodeFindValueServer) Send(m *FindValueResponse) error {
	return x.ServerStream.SendMsg(m)
}

// FullNode_ServiceDesc is the grpc.ServiceDesc for FullNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FullNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kademlia.FullNode",
	HandlerType: (*FullNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _FullNode_Ping_Handler,
		},
		{
			MethodName: "FindNode",
			Handler:    _FullNode_FindNode_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Store",
			Handler:       _FullNode_Store_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindValue",
			Handler:       _FullNode_FindValue_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kademlia.proto",
}
