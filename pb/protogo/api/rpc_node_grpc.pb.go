// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.6.1
// source: api/rpc_node.proto

package api

import (
	common "chainmaker.org/chainmaker-sdk-go/pb/protogo/common"
	config "chainmaker.org/chainmaker-sdk-go/pb/protogo/config"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RpcNodeClient is the client API for RpcNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcNodeClient interface {
	// processing transaction message requests
	SendRequest(ctx context.Context, in *common.TxRequest, opts ...grpc.CallOption) (*common.TxResponse, error)
	// processing requests for message subscription
	Subscribe(ctx context.Context, in *common.TxRequest, opts ...grpc.CallOption) (RpcNode_SubscribeClient, error)
	// update debug status (development debugging)
	UpdateDebugConfig(ctx context.Context, in *config.DebugConfigRequest, opts ...grpc.CallOption) (*config.DebugConfigResponse, error)
	// refreshLogLevelsConfig
	RefreshLogLevelsConfig(ctx context.Context, in *config.LogLevelsRequest, opts ...grpc.CallOption) (*config.LogLevelsResponse, error)
	// get chainmaker version
	GetChainMakerVersion(ctx context.Context, in *config.ChainMakerVersionRequest, opts ...grpc.CallOption) (*config.ChainMakerVersionResponse, error)
	// check chain configuration and load new chain dynamically
	CheckNewBlockChainConfig(ctx context.Context, in *config.CheckNewBlockChainConfigRequest, opts ...grpc.CallOption) (*config.CheckNewBlockChainConfigResponse, error)
}

type rpcNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcNodeClient(cc grpc.ClientConnInterface) RpcNodeClient {
	return &rpcNodeClient{cc}
}

func (c *rpcNodeClient) SendRequest(ctx context.Context, in *common.TxRequest, opts ...grpc.CallOption) (*common.TxResponse, error) {
	out := new(common.TxResponse)
	err := c.cc.Invoke(ctx, "/api.RpcNode/SendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcNodeClient) Subscribe(ctx context.Context, in *common.TxRequest, opts ...grpc.CallOption) (RpcNode_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &RpcNode_ServiceDesc.Streams[0], "/api.RpcNode/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcNodeSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RpcNode_SubscribeClient interface {
	Recv() (*common.SubscribeResult, error)
	grpc.ClientStream
}

type rpcNodeSubscribeClient struct {
	grpc.ClientStream
}

func (x *rpcNodeSubscribeClient) Recv() (*common.SubscribeResult, error) {
	m := new(common.SubscribeResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rpcNodeClient) UpdateDebugConfig(ctx context.Context, in *config.DebugConfigRequest, opts ...grpc.CallOption) (*config.DebugConfigResponse, error) {
	out := new(config.DebugConfigResponse)
	err := c.cc.Invoke(ctx, "/api.RpcNode/UpdateDebugConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcNodeClient) RefreshLogLevelsConfig(ctx context.Context, in *config.LogLevelsRequest, opts ...grpc.CallOption) (*config.LogLevelsResponse, error) {
	out := new(config.LogLevelsResponse)
	err := c.cc.Invoke(ctx, "/api.RpcNode/RefreshLogLevelsConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcNodeClient) GetChainMakerVersion(ctx context.Context, in *config.ChainMakerVersionRequest, opts ...grpc.CallOption) (*config.ChainMakerVersionResponse, error) {
	out := new(config.ChainMakerVersionResponse)
	err := c.cc.Invoke(ctx, "/api.RpcNode/GetChainMakerVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcNodeClient) CheckNewBlockChainConfig(ctx context.Context, in *config.CheckNewBlockChainConfigRequest, opts ...grpc.CallOption) (*config.CheckNewBlockChainConfigResponse, error) {
	out := new(config.CheckNewBlockChainConfigResponse)
	err := c.cc.Invoke(ctx, "/api.RpcNode/CheckNewBlockChainConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcNodeServer is the server API for RpcNode service.
// All implementations must embed UnimplementedRpcNodeServer
// for forward compatibility
type RpcNodeServer interface {
	// processing transaction message requests
	SendRequest(context.Context, *common.TxRequest) (*common.TxResponse, error)
	// processing requests for message subscription
	Subscribe(*common.TxRequest, RpcNode_SubscribeServer) error
	// update debug status (development debugging)
	UpdateDebugConfig(context.Context, *config.DebugConfigRequest) (*config.DebugConfigResponse, error)
	// refreshLogLevelsConfig
	RefreshLogLevelsConfig(context.Context, *config.LogLevelsRequest) (*config.LogLevelsResponse, error)
	// get chainmaker version
	GetChainMakerVersion(context.Context, *config.ChainMakerVersionRequest) (*config.ChainMakerVersionResponse, error)
	// check chain configuration and load new chain dynamically
	CheckNewBlockChainConfig(context.Context, *config.CheckNewBlockChainConfigRequest) (*config.CheckNewBlockChainConfigResponse, error)
	mustEmbedUnimplementedRpcNodeServer()
}

// UnimplementedRpcNodeServer must be embedded to have forward compatible implementations.
type UnimplementedRpcNodeServer struct {
}

func (UnimplementedRpcNodeServer) SendRequest(context.Context, *common.TxRequest) (*common.TxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRequest not implemented")
}
func (UnimplementedRpcNodeServer) Subscribe(*common.TxRequest, RpcNode_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedRpcNodeServer) UpdateDebugConfig(context.Context, *config.DebugConfigRequest) (*config.DebugConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDebugConfig not implemented")
}
func (UnimplementedRpcNodeServer) RefreshLogLevelsConfig(context.Context, *config.LogLevelsRequest) (*config.LogLevelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshLogLevelsConfig not implemented")
}
func (UnimplementedRpcNodeServer) GetChainMakerVersion(context.Context, *config.ChainMakerVersionRequest) (*config.ChainMakerVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChainMakerVersion not implemented")
}
func (UnimplementedRpcNodeServer) CheckNewBlockChainConfig(context.Context, *config.CheckNewBlockChainConfigRequest) (*config.CheckNewBlockChainConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckNewBlockChainConfig not implemented")
}
func (UnimplementedRpcNodeServer) mustEmbedUnimplementedRpcNodeServer() {}

// UnsafeRpcNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcNodeServer will
// result in compilation errors.
type UnsafeRpcNodeServer interface {
	mustEmbedUnimplementedRpcNodeServer()
}

func RegisterRpcNodeServer(s grpc.ServiceRegistrar, srv RpcNodeServer) {
	s.RegisterService(&RpcNode_ServiceDesc, srv)
}

func _RpcNode_SendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.TxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcNodeServer).SendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RpcNode/SendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcNodeServer).SendRequest(ctx, req.(*common.TxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcNode_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(common.TxRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RpcNodeServer).Subscribe(m, &rpcNodeSubscribeServer{stream})
}

type RpcNode_SubscribeServer interface {
	Send(*common.SubscribeResult) error
	grpc.ServerStream
}

type rpcNodeSubscribeServer struct {
	grpc.ServerStream
}

func (x *rpcNodeSubscribeServer) Send(m *common.SubscribeResult) error {
	return x.ServerStream.SendMsg(m)
}

func _RpcNode_UpdateDebugConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(config.DebugConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcNodeServer).UpdateDebugConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RpcNode/UpdateDebugConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcNodeServer).UpdateDebugConfig(ctx, req.(*config.DebugConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcNode_RefreshLogLevelsConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(config.LogLevelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcNodeServer).RefreshLogLevelsConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RpcNode/RefreshLogLevelsConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcNodeServer).RefreshLogLevelsConfig(ctx, req.(*config.LogLevelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcNode_GetChainMakerVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(config.ChainMakerVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcNodeServer).GetChainMakerVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RpcNode/GetChainMakerVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcNodeServer).GetChainMakerVersion(ctx, req.(*config.ChainMakerVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcNode_CheckNewBlockChainConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(config.CheckNewBlockChainConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcNodeServer).CheckNewBlockChainConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RpcNode/CheckNewBlockChainConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcNodeServer).CheckNewBlockChainConfig(ctx, req.(*config.CheckNewBlockChainConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RpcNode_ServiceDesc is the grpc.ServiceDesc for RpcNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RpcNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.RpcNode",
	HandlerType: (*RpcNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRequest",
			Handler:    _RpcNode_SendRequest_Handler,
		},
		{
			MethodName: "UpdateDebugConfig",
			Handler:    _RpcNode_UpdateDebugConfig_Handler,
		},
		{
			MethodName: "RefreshLogLevelsConfig",
			Handler:    _RpcNode_RefreshLogLevelsConfig_Handler,
		},
		{
			MethodName: "GetChainMakerVersion",
			Handler:    _RpcNode_GetChainMakerVersion_Handler,
		},
		{
			MethodName: "CheckNewBlockChainConfig",
			Handler:    _RpcNode_CheckNewBlockChainConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _RpcNode_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/rpc_node.proto",
}
