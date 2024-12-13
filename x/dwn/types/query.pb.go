// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dwn/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a31dfed63924cb3, []int{0}
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

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a31dfed63924cb3, []int{1}
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

func (m *QueryParamsResponse) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

// QueryAllocateRequest is the request type for the Allocate RPC method.
type QueryAllocateRequest struct {
}

func (m *QueryAllocateRequest) Reset()         { *m = QueryAllocateRequest{} }
func (m *QueryAllocateRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllocateRequest) ProtoMessage()    {}
func (*QueryAllocateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a31dfed63924cb3, []int{2}
}
func (m *QueryAllocateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllocateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllocateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllocateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllocateRequest.Merge(m, src)
}
func (m *QueryAllocateRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllocateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllocateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllocateRequest proto.InternalMessageInfo

// AllocateResponse is the response type for the Allocate RPC method.
type QueryAllocateResponse struct {
	Success     bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Cid         string `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
	Macaroon    string `protobuf:"bytes,3,opt,name=macaroon,proto3" json:"macaroon,omitempty"`
	PublicUri   string `protobuf:"bytes,4,opt,name=public_uri,json=publicUri,proto3" json:"public_uri,omitempty"`
	ExpiryBlock int64  `protobuf:"varint,5,opt,name=expiry_block,json=expiryBlock,proto3" json:"expiry_block,omitempty"`
}

func (m *QueryAllocateResponse) Reset()         { *m = QueryAllocateResponse{} }
func (m *QueryAllocateResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllocateResponse) ProtoMessage()    {}
func (*QueryAllocateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a31dfed63924cb3, []int{3}
}
func (m *QueryAllocateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllocateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllocateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllocateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllocateResponse.Merge(m, src)
}
func (m *QueryAllocateResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllocateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllocateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllocateResponse proto.InternalMessageInfo

func (m *QueryAllocateResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *QueryAllocateResponse) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *QueryAllocateResponse) GetMacaroon() string {
	if m != nil {
		return m.Macaroon
	}
	return ""
}

func (m *QueryAllocateResponse) GetPublicUri() string {
	if m != nil {
		return m.PublicUri
	}
	return ""
}

func (m *QueryAllocateResponse) GetExpiryBlock() int64 {
	if m != nil {
		return m.ExpiryBlock
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "dwn.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "dwn.v1.QueryParamsResponse")
	proto.RegisterType((*QueryAllocateRequest)(nil), "dwn.v1.QueryAllocateRequest")
	proto.RegisterType((*QueryAllocateResponse)(nil), "dwn.v1.QueryAllocateResponse")
}

func init() { proto.RegisterFile("dwn/v1/query.proto", fileDescriptor_9a31dfed63924cb3) }

var fileDescriptor_9a31dfed63924cb3 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xae, 0xb7, 0x6c, 0xe8, 0x7a, 0x11, 0x5a, 0x0d, 0x01, 0x45, 0x61, 0x37, 0x2a, 0x39, 0xa0,
	0x9e, 0x62, 0xed, 0x72, 0x85, 0x03, 0xfb, 0x04, 0x10, 0x89, 0x0b, 0x1c, 0x56, 0x8e, 0x6b, 0x05,
	0x8b, 0xd4, 0x4e, 0x6d, 0xa7, 0x3f, 0x57, 0x9e, 0x00, 0x89, 0x27, 0xe0, 0x6d, 0xe0, 0x56, 0x89,
	0x0b, 0x47, 0xd4, 0xf2, 0x20, 0x28, 0x76, 0x02, 0x6a, 0x61, 0x2f, 0x51, 0xe6, 0xfb, 0x66, 0xbe,
	0xcf, 0xf3, 0x69, 0x30, 0x4c, 0x97, 0x92, 0x2c, 0x2e, 0xc9, 0xbc, 0xe1, 0x7a, 0x9d, 0xd5, 0x5a,
	0x59, 0x05, 0xc1, 0x74, 0x29, 0xb3, 0xc5, 0x65, 0x7c, 0x5e, 0x2a, 0x55, 0x56, 0x9c, 0xd0, 0x5a,
	0x10, 0x2a, 0xa5, 0xb2, 0xd4, 0x0a, 0x25, 0x8d, 0xef, 0x8a, 0xc3, 0x6e, 0xb2, 0xe4, 0x92, 0x1b,
	0xd1, 0xa1, 0x69, 0x88, 0xe1, 0x75, 0x2b, 0xf5, 0x8a, 0x6a, 0x3a, 0x33, 0x39, 0x9f, 0x37, 0xdc,
	0xd8, 0xf4, 0x05, 0x7e, 0xb0, 0x87, 0x9a, 0x5a, 0x49, 0xc3, 0xe1, 0x29, 0x0e, 0x6a, 0x87, 0x44,
	0x68, 0x8c, 0x26, 0xa7, 0x57, 0xf7, 0x33, 0xef, 0x9c, 0x75, 0x7d, 0x1d, 0x9b, 0x3e, 0xc2, 0xa1,
	0x1b, 0x7f, 0x59, 0x55, 0x8a, 0x51, 0xcb, 0x7b, 0xd9, 0x2f, 0x08, 0x3f, 0x3c, 0x20, 0x3a, 0xe5,
	0x08, 0xdf, 0x35, 0x0d, 0x63, 0xdc, 0x78, 0xe9, 0x51, 0xde, 0x97, 0x70, 0x86, 0x87, 0x4c, 0x4c,
	0xa3, 0xa3, 0x31, 0x9a, 0x9c, 0xe4, 0xed, 0x2f, 0xc4, 0x78, 0x34, 0xa3, 0x8c, 0x6a, 0xa5, 0x64,
	0x34, 0x74, 0xf0, 0x9f, 0x1a, 0x2e, 0x30, 0xae, 0x9b, 0xa2, 0x12, 0xec, 0xa6, 0xd1, 0x22, 0xba,
	0xe3, 0xd8, 0x13, 0x8f, 0xbc, 0xd1, 0x02, 0x9e, 0xe0, 0x7b, 0x7c, 0x55, 0x0b, 0xbd, 0xbe, 0x29,
	0x2a, 0xc5, 0x3e, 0x44, 0xc7, 0x63, 0x34, 0x19, 0xe6, 0xa7, 0x1e, 0xbb, 0x6e, 0xa1, 0xab, 0x6f,
	0x08, 0x1f, 0xbb, 0x37, 0xc2, 0x3b, 0x1c, 0xf8, 0xbd, 0x20, 0xee, 0xf7, 0xfc, 0x37, 0xaa, 0xf8,
	0xf1, 0x7f, 0x39, 0xbf, 0x56, 0x1a, 0x7d, 0xfc, 0xfe, 0xeb, 0xf3, 0x11, 0xc0, 0x19, 0x59, 0xd0,
	0xa6, 0xb2, 0x6d, 0xfc, 0x3e, 0x22, 0x60, 0x78, 0xd4, 0x87, 0x00, 0xe7, 0x7b, 0x12, 0x07, 0xa1,
	0xc5, 0x17, 0xb7, 0xb0, 0x9d, 0x45, 0xec, 0x2c, 0x42, 0x80, 0xbf, 0x16, 0xb4, 0xeb, 0xb9, 0x7e,
	0xfe, 0x75, 0x9b, 0xa0, 0xcd, 0x36, 0x41, 0x3f, 0xb7, 0x09, 0xfa, 0xb4, 0x4b, 0x06, 0x9b, 0x5d,
	0x32, 0xf8, 0xb1, 0x4b, 0x06, 0x6f, 0xd3, 0x52, 0xd8, 0xf7, 0x4d, 0x91, 0x31, 0x35, 0x23, 0x4a,
	0x1a, 0x25, 0x35, 0x71, 0x9f, 0x15, 0x69, 0xaf, 0xc4, 0xae, 0x6b, 0x6e, 0x8a, 0xc0, 0x5d, 0xc8,
	0xb3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x14, 0x99, 0xd4, 0xa2, 0x73, 0x02, 0x00, 0x00,
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
	// Params queries all parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Allocate initializes a Target Vault available for claims with a compatible
	// Authentication mechanism. The default authentication mechanism is WebAuthn.
	Allocate(ctx context.Context, in *QueryAllocateRequest, opts ...grpc.CallOption) (*QueryAllocateResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/dwn.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Allocate(ctx context.Context, in *QueryAllocateRequest, opts ...grpc.CallOption) (*QueryAllocateResponse, error) {
	out := new(QueryAllocateResponse)
	err := c.cc.Invoke(ctx, "/dwn.v1.Query/Allocate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params queries all parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Allocate initializes a Target Vault available for claims with a compatible
	// Authentication mechanism. The default authentication mechanism is WebAuthn.
	Allocate(context.Context, *QueryAllocateRequest) (*QueryAllocateResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Allocate(ctx context.Context, req *QueryAllocateRequest) (*QueryAllocateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Allocate not implemented")
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
		FullMethod: "/dwn.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
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
		FullMethod: "/dwn.v1.Query/Allocate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Allocate(ctx, req.(*QueryAllocateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var Query_serviceDesc = _Query_serviceDesc
var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dwn.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Allocate",
			Handler:    _Query_Allocate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dwn/v1/query.proto",
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
	if m.Params != nil {
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
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllocateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllocateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllocateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAllocateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllocateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllocateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpiryBlock != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ExpiryBlock))
		i--
		dAtA[i] = 0x28
	}
	if len(m.PublicUri) > 0 {
		i -= len(m.PublicUri)
		copy(dAtA[i:], m.PublicUri)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PublicUri)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Macaroon) > 0 {
		i -= len(m.Macaroon)
		copy(dAtA[i:], m.Macaroon)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Macaroon)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0x12
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
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
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllocateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAllocateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Macaroon)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.PublicUri)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.ExpiryBlock != 0 {
		n += 1 + sovQuery(uint64(m.ExpiryBlock))
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
			if m.Params == nil {
				m.Params = &Params{}
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
func (m *QueryAllocateRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllocateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllocateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryAllocateResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllocateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllocateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
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
			m.Cid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Macaroon", wireType)
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
			m.Macaroon = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicUri", wireType)
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
			m.PublicUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiryBlock", wireType)
			}
			m.ExpiryBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiryBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
