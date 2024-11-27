// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: dwn/v1/query.proto

package dwnv1

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
	Query_Params_FullMethodName   = "/dwn.v1.Query/Params"
	Query_Schema_FullMethodName   = "/dwn.v1.Query/Schema"
	Query_Allocate_FullMethodName = "/dwn.v1.Query/Allocate"
	Query_Sync_FullMethodName     = "/dwn.v1.Query/Sync"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Query provides defines the gRPC querier service.
type QueryClient interface {
	// Params queries all parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Schema queries the DID document by its id. And returns the required PKL
	// information
	Schema(ctx context.Context, in *QuerySchemaRequest, opts ...grpc.CallOption) (*QuerySchemaResponse, error)
	// Allocate initializes a Target Vault available for claims with a compatible
	// Authentication mechanism. The default authentication mechanism is WebAuthn.
	Allocate(ctx context.Context, in *QueryAllocateRequest, opts ...grpc.CallOption) (*QueryAllocateResponse, error)
	// Sync queries the DID document by its id. And returns the required PKL
	// information
	Sync(ctx context.Context, in *QuerySyncRequest, opts ...grpc.CallOption) (*QuerySyncResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Schema(ctx context.Context, in *QuerySchemaRequest, opts ...grpc.CallOption) (*QuerySchemaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QuerySchemaResponse)
	err := c.cc.Invoke(ctx, Query_Schema_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Allocate(ctx context.Context, in *QueryAllocateRequest, opts ...grpc.CallOption) (*QueryAllocateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryAllocateResponse)
	err := c.cc.Invoke(ctx, Query_Allocate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Sync(ctx context.Context, in *QuerySyncRequest, opts ...grpc.CallOption) (*QuerySyncResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QuerySyncResponse)
	err := c.cc.Invoke(ctx, Query_Sync_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility.
//
// Query provides defines the gRPC querier service.
type QueryServer interface {
	// Params queries all parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Schema queries the DID document by its id. And returns the required PKL
	// information
	Schema(context.Context, *QuerySchemaRequest) (*QuerySchemaResponse, error)
	// Allocate initializes a Target Vault available for claims with a compatible
	// Authentication mechanism. The default authentication mechanism is WebAuthn.
	Allocate(context.Context, *QueryAllocateRequest) (*QueryAllocateResponse, error)
	// Sync queries the DID document by its id. And returns the required PKL
	// information
	Sync(context.Context, *QuerySyncRequest) (*QuerySyncResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQueryServer struct{}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Schema(context.Context, *QuerySchemaRequest) (*QuerySchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Schema not implemented")
}
func (UnimplementedQueryServer) Allocate(context.Context, *QueryAllocateRequest) (*QueryAllocateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Allocate not implemented")
}
func (UnimplementedQueryServer) Sync(context.Context, *QuerySyncRequest) (*QuerySyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}
func (UnimplementedQueryServer) testEmbeddedByValue()               {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	// If the following call pancis, it indicates UnimplementedQueryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Schema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Schema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Schema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Schema(ctx, req.(*QuerySchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Allocate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllocateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Allocate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Allocate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Allocate(ctx, req.(*QueryAllocateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Sync(ctx, req.(*QuerySyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dwn.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Schema",
			Handler:    _Query_Schema_Handler,
		},
		{
			MethodName: "Allocate",
			Handler:    _Query_Allocate_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _Query_Sync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dwn/v1/query.proto",
}