// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bucket/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgDefineBucket struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Label   string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (m *MsgDefineBucket) Reset()         { *m = MsgDefineBucket{} }
func (m *MsgDefineBucket) String() string { return proto.CompactTextString(m) }
func (*MsgDefineBucket) ProtoMessage()    {}
func (*MsgDefineBucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_3479ee73a3c611d5, []int{0}
}
func (m *MsgDefineBucket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDefineBucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDefineBucket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDefineBucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDefineBucket.Merge(m, src)
}
func (m *MsgDefineBucket) XXX_Size() int {
	return m.Size()
}
func (m *MsgDefineBucket) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDefineBucket.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDefineBucket proto.InternalMessageInfo

func (m *MsgDefineBucket) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDefineBucket) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type MsgDefineBucketResponse struct {
	Status  int32   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	WhereIs *Bucket `protobuf:"bytes,2,opt,name=where_is,json=whereIs,proto3" json:"where_is,omitempty"`
}

func (m *MsgDefineBucketResponse) Reset()         { *m = MsgDefineBucketResponse{} }
func (m *MsgDefineBucketResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDefineBucketResponse) ProtoMessage()    {}
func (*MsgDefineBucketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3479ee73a3c611d5, []int{1}
}
func (m *MsgDefineBucketResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDefineBucketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDefineBucketResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDefineBucketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDefineBucketResponse.Merge(m, src)
}
func (m *MsgDefineBucketResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDefineBucketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDefineBucketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDefineBucketResponse proto.InternalMessageInfo

func (m *MsgDefineBucketResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *MsgDefineBucketResponse) GetWhereIs() *Bucket {
	if m != nil {
		return m.WhereIs
	}
	return nil
}

type MsgGenerateBucket struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgGenerateBucket) Reset()         { *m = MsgGenerateBucket{} }
func (m *MsgGenerateBucket) String() string { return proto.CompactTextString(m) }
func (*MsgGenerateBucket) ProtoMessage()    {}
func (*MsgGenerateBucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_3479ee73a3c611d5, []int{2}
}
func (m *MsgGenerateBucket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGenerateBucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGenerateBucket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGenerateBucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGenerateBucket.Merge(m, src)
}
func (m *MsgGenerateBucket) XXX_Size() int {
	return m.Size()
}
func (m *MsgGenerateBucket) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGenerateBucket.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGenerateBucket proto.InternalMessageInfo

func (m *MsgGenerateBucket) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgGenerateBucketResponse struct {
}

func (m *MsgGenerateBucketResponse) Reset()         { *m = MsgGenerateBucketResponse{} }
func (m *MsgGenerateBucketResponse) String() string { return proto.CompactTextString(m) }
func (*MsgGenerateBucketResponse) ProtoMessage()    {}
func (*MsgGenerateBucketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3479ee73a3c611d5, []int{3}
}
func (m *MsgGenerateBucketResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGenerateBucketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGenerateBucketResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGenerateBucketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGenerateBucketResponse.Merge(m, src)
}
func (m *MsgGenerateBucketResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgGenerateBucketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGenerateBucketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGenerateBucketResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgDefineBucket)(nil), "sonrio.sonr.bucket.MsgDefineBucket")
	proto.RegisterType((*MsgDefineBucketResponse)(nil), "sonrio.sonr.bucket.MsgDefineBucketResponse")
	proto.RegisterType((*MsgGenerateBucket)(nil), "sonrio.sonr.bucket.MsgGenerateBucket")
	proto.RegisterType((*MsgGenerateBucketResponse)(nil), "sonrio.sonr.bucket.MsgGenerateBucketResponse")
}

func init() { proto.RegisterFile("bucket/tx.proto", fileDescriptor_3479ee73a3c611d5) }

var fileDescriptor_3479ee73a3c611d5 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2a, 0x4d, 0xce,
	0x4e, 0x2d, 0xd1, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2a, 0xce, 0xcf,
	0x2b, 0xca, 0xcc, 0xd7, 0x03, 0x51, 0x7a, 0x10, 0x49, 0x29, 0x61, 0xa8, 0x22, 0x08, 0x05, 0x51,
	0xa8, 0xe4, 0xc8, 0xc5, 0xef, 0x5b, 0x9c, 0xee, 0x92, 0x9a, 0x96, 0x99, 0x97, 0xea, 0x04, 0x96,
	0x10, 0x92, 0xe0, 0x62, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x92, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x82, 0x71, 0x85, 0x44, 0xb8, 0x58, 0x73, 0x12, 0x93, 0x52, 0x73, 0x24, 0x98, 0xc0,
	0xe2, 0x10, 0x8e, 0x52, 0x06, 0x97, 0x38, 0x9a, 0x11, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5,
	0xa9, 0x42, 0x62, 0x5c, 0x6c, 0xc5, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0x60, 0x93, 0x58, 0x83, 0xa0,
	0x3c, 0x21, 0x53, 0x2e, 0x8e, 0xf2, 0x8c, 0xd4, 0xa2, 0xd4, 0xf8, 0xcc, 0x62, 0xb0, 0x59, 0xdc,
	0x46, 0x52, 0x7a, 0x98, 0x2e, 0xd6, 0x83, 0x9a, 0xc6, 0x0e, 0x56, 0xeb, 0x59, 0xac, 0xa4, 0xcb,
	0x25, 0xe8, 0x5b, 0x9c, 0xee, 0x9e, 0x9a, 0x97, 0x5a, 0x94, 0x58, 0x42, 0xd0, 0xb9, 0x4a, 0xd2,
	0x5c, 0x92, 0x18, 0xca, 0x61, 0x4e, 0x33, 0x3a, 0xcf, 0xc8, 0xc5, 0xec, 0x5b, 0x9c, 0x2e, 0x94,
	0xc0, 0xc5, 0x83, 0xe2, 0x7b, 0x65, 0x6c, 0x0e, 0x41, 0xf3, 0x9f, 0x94, 0x36, 0x11, 0x8a, 0xe0,
	0x81, 0x90, 0xc6, 0xc5, 0x87, 0xe6, 0x64, 0x55, 0x1c, 0xda, 0x51, 0x95, 0x49, 0xe9, 0x12, 0xa5,
	0x0c, 0x66, 0x8f, 0x53, 0xc4, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x70, 0x89,
	0xc0, 0xcc, 0x28, 0xa9, 0x2c, 0x48, 0x2d, 0x86, 0x9a, 0x14, 0xc0, 0x18, 0xa5, 0x96, 0x9e, 0x59,
	0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x92, 0xd7, 0xcd, 0xcc, 0x07, 0xd3, 0xfa,
	0x15, 0xfa, 0xb0, 0x04, 0x05, 0xd2, 0x90, 0xc4, 0x06, 0x4e, 0x2b, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x48, 0xff, 0x19, 0xde, 0x67, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	DefineBucket(ctx context.Context, in *MsgDefineBucket, opts ...grpc.CallOption) (*MsgDefineBucketResponse, error)
	GenerateBucket(ctx context.Context, in *MsgGenerateBucket, opts ...grpc.CallOption) (*MsgGenerateBucketResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) DefineBucket(ctx context.Context, in *MsgDefineBucket, opts ...grpc.CallOption) (*MsgDefineBucketResponse, error) {
	out := new(MsgDefineBucketResponse)
	err := c.cc.Invoke(ctx, "/sonrio.sonr.bucket.Msg/DefineBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GenerateBucket(ctx context.Context, in *MsgGenerateBucket, opts ...grpc.CallOption) (*MsgGenerateBucketResponse, error) {
	out := new(MsgGenerateBucketResponse)
	err := c.cc.Invoke(ctx, "/sonrio.sonr.bucket.Msg/GenerateBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	DefineBucket(context.Context, *MsgDefineBucket) (*MsgDefineBucketResponse, error)
	GenerateBucket(context.Context, *MsgGenerateBucket) (*MsgGenerateBucketResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) DefineBucket(ctx context.Context, req *MsgDefineBucket) (*MsgDefineBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DefineBucket not implemented")
}
func (*UnimplementedMsgServer) GenerateBucket(ctx context.Context, req *MsgGenerateBucket) (*MsgGenerateBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateBucket not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_DefineBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDefineBucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DefineBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrio.sonr.bucket.Msg/DefineBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DefineBucket(ctx, req.(*MsgDefineBucket))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GenerateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGenerateBucket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GenerateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrio.sonr.bucket.Msg/GenerateBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GenerateBucket(ctx, req.(*MsgGenerateBucket))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonrio.sonr.bucket.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DefineBucket",
			Handler:    _Msg_DefineBucket_Handler,
		},
		{
			MethodName: "GenerateBucket",
			Handler:    _Msg_GenerateBucket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bucket/tx.proto",
}

func (m *MsgDefineBucket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDefineBucket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDefineBucket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Label)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDefineBucketResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDefineBucketResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDefineBucketResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WhereIs != nil {
		{
			size, err := m.WhereIs.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Status != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgGenerateBucket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGenerateBucket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGenerateBucket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgGenerateBucketResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGenerateBucketResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGenerateBucketResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgDefineBucket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDefineBucketResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovTx(uint64(m.Status))
	}
	if m.WhereIs != nil {
		l = m.WhereIs.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgGenerateBucket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgGenerateBucketResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgDefineBucket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDefineBucket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDefineBucket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDefineBucketResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDefineBucketResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDefineBucketResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WhereIs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.WhereIs == nil {
				m.WhereIs = &Bucket{}
			}
			if err := m.WhereIs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgGenerateBucket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgGenerateBucket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGenerateBucket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgGenerateBucketResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgGenerateBucketResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGenerateBucketResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
