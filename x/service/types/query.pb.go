// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/service/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	_ "sonr.io/core/x/identity/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cf4b2348245f9e3, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cf4b2348245f9e3, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryGetServiceRecordRequest is the request type for the Query/ServiceRecord
// RPC method.
type QueryGetServiceRecordRequest struct {
	Origin string `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (m *QueryGetServiceRecordRequest) Reset()         { *m = QueryGetServiceRecordRequest{} }
func (m *QueryGetServiceRecordRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetServiceRecordRequest) ProtoMessage()    {}
func (*QueryGetServiceRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cf4b2348245f9e3, []int{2}
}
func (m *QueryGetServiceRecordRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetServiceRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetServiceRecordRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetServiceRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetServiceRecordRequest.Merge(m, src)
}
func (m *QueryGetServiceRecordRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetServiceRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetServiceRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetServiceRecordRequest proto.InternalMessageInfo

func (m *QueryGetServiceRecordRequest) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

// QueryGetServiceRecordResponse is the response type for the
// Query/ServiceRecord RPC method.
type QueryGetServiceRecordResponse struct {
	ServiceRecord ServiceRecord `protobuf:"bytes,1,opt,name=ServiceRecord,proto3" json:"ServiceRecord"`
}

func (m *QueryGetServiceRecordResponse) Reset()         { *m = QueryGetServiceRecordResponse{} }
func (m *QueryGetServiceRecordResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetServiceRecordResponse) ProtoMessage()    {}
func (*QueryGetServiceRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cf4b2348245f9e3, []int{3}
}
func (m *QueryGetServiceRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetServiceRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetServiceRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetServiceRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetServiceRecordResponse.Merge(m, src)
}
func (m *QueryGetServiceRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetServiceRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetServiceRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetServiceRecordResponse proto.InternalMessageInfo

func (m *QueryGetServiceRecordResponse) GetServiceRecord() ServiceRecord {
	if m != nil {
		return m.ServiceRecord
	}
	return ServiceRecord{}
}

// QueryAllServiceRecordRequest is the request type for the
// Query/ServiceRecordAll RPC method.
type QueryAllServiceRecordRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllServiceRecordRequest) Reset()         { *m = QueryAllServiceRecordRequest{} }
func (m *QueryAllServiceRecordRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllServiceRecordRequest) ProtoMessage()    {}
func (*QueryAllServiceRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cf4b2348245f9e3, []int{4}
}
func (m *QueryAllServiceRecordRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllServiceRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllServiceRecordRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllServiceRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllServiceRecordRequest.Merge(m, src)
}
func (m *QueryAllServiceRecordRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllServiceRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllServiceRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllServiceRecordRequest proto.InternalMessageInfo

func (m *QueryAllServiceRecordRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryAllServiceRecordResponse is the response type for the
// Query/ServiceRecordAll RPC method.
type QueryAllServiceRecordResponse struct {
	ServiceRecord []ServiceRecord     `protobuf:"bytes,1,rep,name=ServiceRecord,proto3" json:"ServiceRecord"`
	Pagination    *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllServiceRecordResponse) Reset()         { *m = QueryAllServiceRecordResponse{} }
func (m *QueryAllServiceRecordResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllServiceRecordResponse) ProtoMessage()    {}
func (*QueryAllServiceRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cf4b2348245f9e3, []int{5}
}
func (m *QueryAllServiceRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllServiceRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllServiceRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllServiceRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllServiceRecordResponse.Merge(m, src)
}
func (m *QueryAllServiceRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllServiceRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllServiceRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllServiceRecordResponse proto.InternalMessageInfo

func (m *QueryAllServiceRecordResponse) GetServiceRecord() []ServiceRecord {
	if m != nil {
		return m.ServiceRecord
	}
	return nil
}

func (m *QueryAllServiceRecordResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "core.service.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "core.service.QueryParamsResponse")
	proto.RegisterType((*QueryGetServiceRecordRequest)(nil), "core.service.QueryGetServiceRecordRequest")
	proto.RegisterType((*QueryGetServiceRecordResponse)(nil), "core.service.QueryGetServiceRecordResponse")
	proto.RegisterType((*QueryAllServiceRecordRequest)(nil), "core.service.QueryAllServiceRecordRequest")
	proto.RegisterType((*QueryAllServiceRecordResponse)(nil), "core.service.QueryAllServiceRecordResponse")
}

func init() { proto.RegisterFile("core/service/query.proto", fileDescriptor_5cf4b2348245f9e3) }

var fileDescriptor_5cf4b2348245f9e3 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0xb3, 0x2d, 0x44, 0xc2, 0x50, 0x81, 0x4c, 0x08, 0x21, 0x84, 0xa5, 0xec, 0x01, 0x50,
	0x91, 0x6c, 0x35, 0x48, 0xbd, 0xb7, 0x07, 0x22, 0x6e, 0x65, 0xb9, 0x71, 0x73, 0x12, 0xb3, 0x58,
	0x6c, 0x77, 0xb6, 0xb6, 0x5b, 0x11, 0x21, 0x2e, 0x15, 0x07, 0x8e, 0x48, 0x3c, 0x09, 0x6f, 0xd1,
	0x63, 0x25, 0x2e, 0x9c, 0x10, 0x4a, 0x78, 0x10, 0xb4, 0xf6, 0x14, 0x62, 0xba, 0x0b, 0xa8, 0xb7,
	0x5d, 0xcf, 0x3f, 0x33, 0xdf, 0x78, 0x7e, 0x93, 0xde, 0x04, 0xb4, 0xe4, 0x46, 0xea, 0x43, 0x35,
	0x91, 0x7c, 0xff, 0x40, 0xea, 0x19, 0x2b, 0x35, 0x58, 0xa0, 0x57, 0xaa, 0x08, 0xc3, 0x48, 0xff,
	0xa6, 0xd3, 0xa9, 0xa9, 0x2c, 0xac, 0xb2, 0x33, 0x3e, 0x55, 0x53, 0x2f, 0xeb, 0xdf, 0x0a, 0x0a,
	0x94, 0x42, 0x8b, 0x3d, 0x53, 0x1b, 0xd2, 0x72, 0x02, 0xfa, 0x34, 0x6b, 0x63, 0x02, 0x66, 0x0f,
	0x0c, 0x1f, 0x0b, 0x83, 0x5d, 0xf9, 0xe1, 0xe6, 0x58, 0x5a, 0xb1, 0xc9, 0x4b, 0x91, 0xa9, 0x42,
	0x58, 0x05, 0x05, 0x6a, 0x3b, 0x19, 0x64, 0xe0, 0x3e, 0x79, 0xf5, 0x85, 0xa7, 0x83, 0x0c, 0x20,
	0xcb, 0x25, 0x17, 0xa5, 0xe2, 0xa2, 0x28, 0xc0, 0xba, 0x14, 0x6c, 0x9d, 0x74, 0x08, 0x7d, 0x56,
	0x55, 0xdd, 0x75, 0x3c, 0xa9, 0xdc, 0x3f, 0x90, 0xc6, 0x26, 0x4f, 0xc9, 0xf5, 0xe0, 0xd4, 0x94,
	0x50, 0x18, 0x49, 0x87, 0xa4, 0xed, 0xb9, 0x7b, 0xd1, 0x7a, 0xf4, 0xf0, 0xf2, 0xb0, 0xc3, 0x96,
	0x47, 0x67, 0x5e, 0xbd, 0x73, 0xe1, 0xf8, 0xdb, 0xdd, 0x56, 0x8a, 0xca, 0x64, 0x8b, 0x0c, 0x5c,
	0xa9, 0x91, 0xb4, 0xcf, 0xbd, 0x2e, 0x75, 0xf3, 0x61, 0x2b, 0xda, 0x25, 0x6d, 0xd0, 0x2a, 0x53,
	0x85, 0xab, 0x79, 0x29, 0xc5, 0xbf, 0xe4, 0x15, 0xb9, 0xd3, 0x90, 0x87, 0x30, 0x23, 0xb2, 0x16,
	0x04, 0x90, 0xe9, 0x76, 0xc8, 0x14, 0x48, 0x10, 0x2d, 0xcc, 0x4b, 0x5e, 0x22, 0xe1, 0x76, 0x9e,
	0xd7, 0x12, 0x3e, 0x21, 0xe4, 0xf7, 0x55, 0x63, 0x97, 0xfb, 0xcc, 0xef, 0x85, 0x55, 0x7b, 0x61,
	0xde, 0x0d, 0xb8, 0x17, 0xb6, 0x2b, 0x32, 0x89, 0xb9, 0xe9, 0x52, 0x66, 0xf2, 0x39, 0xc2, 0x91,
	0xce, 0x36, 0x6a, 0x1e, 0x69, 0xf5, 0x3c, 0x23, 0xd1, 0x51, 0x80, 0xbc, 0xe2, 0x90, 0x1f, 0xfc,
	0x13, 0xd9, 0x53, 0x2c, 0x33, 0x0f, 0xdf, 0xaf, 0x92, 0x8b, 0x8e, 0x99, 0xbe, 0x26, 0x6d, 0xbf,
	0x5f, 0xba, 0x1e, 0xe2, 0x9c, 0xb5, 0x4f, 0xff, 0xde, 0x5f, 0x14, 0xbe, 0x49, 0x32, 0x38, 0xfa,
	0xf2, 0xe3, 0xd3, 0x4a, 0x97, 0x76, 0xb8, 0xf3, 0xbe, 0x37, 0xcb, 0xe9, 0x13, 0xa0, 0x1f, 0xa2,
	0x3f, 0x6e, 0x82, 0x6e, 0xd4, 0x94, 0x6c, 0xb0, 0x54, 0xff, 0xd1, 0x7f, 0x69, 0x11, 0x24, 0x76,
	0x20, 0x3d, 0xda, 0xe5, 0xc1, 0x23, 0x7c, 0xeb, 0x6d, 0xf8, 0x8e, 0x1e, 0x45, 0xe4, 0x5a, 0x90,
	0xb9, 0x9d, 0xe7, 0xb5, 0x34, 0x0d, 0xf6, 0xa9, 0xa5, 0x69, 0x72, 0x40, 0x72, 0xc3, 0xd1, 0x5c,
	0xa5, 0x6b, 0x01, 0xcd, 0xce, 0xd6, 0xf1, 0x3c, 0x8e, 0x4e, 0xe6, 0x71, 0xf4, 0x7d, 0x1e, 0x47,
	0x1f, 0x17, 0x71, 0xeb, 0x64, 0x11, 0xb7, 0xbe, 0x2e, 0xe2, 0xd6, 0x8b, 0x81, 0x81, 0x42, 0x33,
	0x05, 0x5e, 0xff, 0xe6, 0x17, 0xbf, 0x9d, 0x95, 0xd2, 0x8c, 0xdb, 0xee, 0x91, 0x3f, 0xfe, 0x19,
	0x00, 0x00, 0xff, 0xff, 0x4d, 0x4c, 0xe0, 0x3b, 0xbd, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// ServiceRecord queries a list of ServiceRecord items.
	ServiceRecord(ctx context.Context, in *QueryGetServiceRecordRequest, opts ...grpc.CallOption) (*QueryGetServiceRecordResponse, error)
	// ServiceRecordAll queries all ServiceRecord items.
	ServiceRecordAll(ctx context.Context, in *QueryAllServiceRecordRequest, opts ...grpc.CallOption) (*QueryAllServiceRecordResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/core.service.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ServiceRecord(ctx context.Context, in *QueryGetServiceRecordRequest, opts ...grpc.CallOption) (*QueryGetServiceRecordResponse, error) {
	out := new(QueryGetServiceRecordResponse)
	err := c.cc.Invoke(ctx, "/core.service.Query/ServiceRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ServiceRecordAll(ctx context.Context, in *QueryAllServiceRecordRequest, opts ...grpc.CallOption) (*QueryAllServiceRecordResponse, error) {
	out := new(QueryAllServiceRecordResponse)
	err := c.cc.Invoke(ctx, "/core.service.Query/ServiceRecordAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// ServiceRecord queries a list of ServiceRecord items.
	ServiceRecord(context.Context, *QueryGetServiceRecordRequest) (*QueryGetServiceRecordResponse, error)
	// ServiceRecordAll queries all ServiceRecord items.
	ServiceRecordAll(context.Context, *QueryAllServiceRecordRequest) (*QueryAllServiceRecordResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) ServiceRecord(ctx context.Context, req *QueryGetServiceRecordRequest) (*QueryGetServiceRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceRecord not implemented")
}
func (*UnimplementedQueryServer) ServiceRecordAll(ctx context.Context, req *QueryAllServiceRecordRequest) (*QueryAllServiceRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceRecordAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
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
		FullMethod: "/core.service.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ServiceRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetServiceRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ServiceRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.service.Query/ServiceRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ServiceRecord(ctx, req.(*QueryGetServiceRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ServiceRecordAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllServiceRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ServiceRecordAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.service.Query/ServiceRecordAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ServiceRecordAll(ctx, req.(*QueryAllServiceRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "core.service.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ServiceRecord",
			Handler:    _Query_ServiceRecord_Handler,
		},
		{
			MethodName: "ServiceRecordAll",
			Handler:    _Query_ServiceRecordAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/service/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetServiceRecordRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetServiceRecordRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetServiceRecordRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Origin) > 0 {
		i -= len(m.Origin)
		copy(dAtA[i:], m.Origin)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Origin)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetServiceRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetServiceRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetServiceRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ServiceRecord.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllServiceRecordRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllServiceRecordRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllServiceRecordRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllServiceRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllServiceRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllServiceRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ServiceRecord) > 0 {
		for iNdEx := len(m.ServiceRecord) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ServiceRecord[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetServiceRecordRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Origin)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetServiceRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ServiceRecord.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllServiceRecordRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllServiceRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ServiceRecord) > 0 {
		for _, e := range m.ServiceRecord {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetServiceRecordRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetServiceRecordRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetServiceRecordRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Origin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetServiceRecordResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetServiceRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetServiceRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceRecord", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ServiceRecord.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllServiceRecordRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllServiceRecordRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllServiceRecordRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllServiceRecordResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllServiceRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllServiceRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceRecord", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceRecord = append(m.ServiceRecord, ServiceRecord{})
			if err := m.ServiceRecord[len(m.ServiceRecord)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
