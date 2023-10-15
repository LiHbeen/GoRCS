// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: api/server/server.proto

package server

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

const (
	Server_CreateServer_FullMethodName = "/api.server.Server/CreateServer"
	Server_UpdateServer_FullMethodName = "/api.server.Server/UpdateServer"
	Server_DeleteServer_FullMethodName = "/api.server.Server/DeleteServer"
	Server_GetServer_FullMethodName    = "/api.server.Server/GetServer"
	Server_ListServer_FullMethodName   = "/api.server.Server/ListServer"
)

// ServerClient is the client API for Server service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerClient interface {
	CreateServer(ctx context.Context, in *CreateServerRequest, opts ...grpc.CallOption) (*CreateServerReply, error)
	UpdateServer(ctx context.Context, in *UpdateServerRequest, opts ...grpc.CallOption) (*UpdateServerReply, error)
	DeleteServer(ctx context.Context, in *DeleteServerRequest, opts ...grpc.CallOption) (*DeleteServerReply, error)
	GetServer(ctx context.Context, in *GetServerRequest, opts ...grpc.CallOption) (*GetServerReply, error)
	ListServer(ctx context.Context, in *ListServerRequest, opts ...grpc.CallOption) (*ListServerReply, error)
}

type serverClient struct {
	cc grpc.ClientConnInterface
}

func NewServerClient(cc grpc.ClientConnInterface) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) CreateServer(ctx context.Context, in *CreateServerRequest, opts ...grpc.CallOption) (*CreateServerReply, error) {
	out := new(CreateServerReply)
	err := c.cc.Invoke(ctx, Server_CreateServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) UpdateServer(ctx context.Context, in *UpdateServerRequest, opts ...grpc.CallOption) (*UpdateServerReply, error) {
	out := new(UpdateServerReply)
	err := c.cc.Invoke(ctx, Server_UpdateServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) DeleteServer(ctx context.Context, in *DeleteServerRequest, opts ...grpc.CallOption) (*DeleteServerReply, error) {
	out := new(DeleteServerReply)
	err := c.cc.Invoke(ctx, Server_DeleteServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) GetServer(ctx context.Context, in *GetServerRequest, opts ...grpc.CallOption) (*GetServerReply, error) {
	out := new(GetServerReply)
	err := c.cc.Invoke(ctx, Server_GetServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverClient) ListServer(ctx context.Context, in *ListServerRequest, opts ...grpc.CallOption) (*ListServerReply, error) {
	out := new(ListServerReply)
	err := c.cc.Invoke(ctx, Server_ListServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerServer is the server API for Server service.
// All implementations must embed UnimplementedServerServer
// for forward compatibility
type ServerServer interface {
	CreateServer(context.Context, *CreateServerRequest) (*CreateServerReply, error)
	UpdateServer(context.Context, *UpdateServerRequest) (*UpdateServerReply, error)
	DeleteServer(context.Context, *DeleteServerRequest) (*DeleteServerReply, error)
	GetServer(context.Context, *GetServerRequest) (*GetServerReply, error)
	ListServer(context.Context, *ListServerRequest) (*ListServerReply, error)
	mustEmbedUnimplementedServerServer()
}

// UnimplementedServerServer must be embedded to have forward compatible implementations.
type UnimplementedServerServer struct {
}

func (UnimplementedServerServer) CreateServer(context.Context, *CreateServerRequest) (*CreateServerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServer not implemented")
}
func (UnimplementedServerServer) UpdateServer(context.Context, *UpdateServerRequest) (*UpdateServerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServer not implemented")
}
func (UnimplementedServerServer) DeleteServer(context.Context, *DeleteServerRequest) (*DeleteServerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServer not implemented")
}
func (UnimplementedServerServer) GetServer(context.Context, *GetServerRequest) (*GetServerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServer not implemented")
}
func (UnimplementedServerServer) ListServer(context.Context, *ListServerRequest) (*ListServerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServer not implemented")
}
func (UnimplementedServerServer) mustEmbedUnimplementedServerServer() {}

// UnsafeServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerServer will
// result in compilation errors.
type UnsafeServerServer interface {
	mustEmbedUnimplementedServerServer()
}

func RegisterServerServer(s grpc.ServiceRegistrar, srv ServerServer) {
	s.RegisterService(&Server_ServiceDesc, srv)
}

func _Server_CreateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).CreateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Server_CreateServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).CreateServer(ctx, req.(*CreateServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_UpdateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).UpdateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Server_UpdateServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).UpdateServer(ctx, req.(*UpdateServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_DeleteServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).DeleteServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Server_DeleteServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).DeleteServer(ctx, req.(*DeleteServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_GetServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).GetServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Server_GetServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).GetServer(ctx, req.(*GetServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Server_ListServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).ListServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Server_ListServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).ListServer(ctx, req.(*ListServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Server_ServiceDesc is the grpc.ServiceDesc for Server service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Server_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.server.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateServer",
			Handler:    _Server_CreateServer_Handler,
		},
		{
			MethodName: "UpdateServer",
			Handler:    _Server_UpdateServer_Handler,
		},
		{
			MethodName: "DeleteServer",
			Handler:    _Server_DeleteServer_Handler,
		},
		{
			MethodName: "GetServer",
			Handler:    _Server_GetServer_Handler,
		},
		{
			MethodName: "ListServer",
			Handler:    _Server_ListServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/server/server.proto",
}
