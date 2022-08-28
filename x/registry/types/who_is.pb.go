// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: registry/who_is.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// WhoIsType is the type of DIDDocument stored in the registry module
type WhoIsType int32

const (
	// User is the type of the registered name
	WhoIsType_USER WhoIsType = 0
	// Application is the type of the registered name
	WhoIsType_APPLICATION WhoIsType = 1
)

var WhoIsType_name = map[int32]string{
	0: "USER",
	1: "APPLICATION",
}

var WhoIsType_value = map[string]int32{
	"USER":        0,
	"APPLICATION": 1,
}

func (x WhoIsType) String() string {
	return proto.EnumName(WhoIsType_name, int32(x))
}

func (WhoIsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9d9bdfc8d37d9424, []int{0}
}

type WhoIs struct {
	// Alias is the list of registered `alsoKnownAs` identifiers of the User or Application
	Alias []*Alias `protobuf:"bytes,1,rep,name=alias,proto3" json:"alias,omitempty"`
	// Owner is the top level DID of the User or Application derived from the multisignature wallet.
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// DIDDocument is the bytes representation of DIDDocument within the WhoIs. Initially marshalled as JSON.
	DidDocument *DIDDocument `protobuf:"bytes,3,opt,name=did_document,json=didDocument,proto3" json:"did_document,omitempty"`
	// Credentials are the biometric info of the registered name and account encoded with public key
	Controllers []string `protobuf:"bytes,4,rep,name=controllers,proto3" json:"controllers,omitempty"`
	// Type is the kind of the entity. Possible values are: "user", "application"
	Type WhoIsType `protobuf:"varint,5,opt,name=type,proto3,enum=sonrio.sonr.registry.WhoIsType" json:"type,omitempty"`
	// Timestamp is the time of the last update of the DID Document
	Timestamp int64 `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// IsActive is the status of the DID Document
	IsActive bool `protobuf:"varint,7,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// Metadata is a map of key-value pairs that can be used to store additional information about the DID Document
	Metadata map[string]string `protobuf:"bytes,8,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *WhoIs) Reset()         { *m = WhoIs{} }
func (m *WhoIs) String() string { return proto.CompactTextString(m) }
func (*WhoIs) ProtoMessage()    {}
func (*WhoIs) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d9bdfc8d37d9424, []int{0}
}
func (m *WhoIs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WhoIs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WhoIs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WhoIs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhoIs.Merge(m, src)
}
func (m *WhoIs) XXX_Size() int {
	return m.Size()
}
func (m *WhoIs) XXX_DiscardUnknown() {
	xxx_messageInfo_WhoIs.DiscardUnknown(m)
}

var xxx_messageInfo_WhoIs proto.InternalMessageInfo

func (m *WhoIs) GetAlias() []*Alias {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *WhoIs) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *WhoIs) GetDidDocument() *DIDDocument {
	if m != nil {
		return m.DidDocument
	}
	return nil
}

func (m *WhoIs) GetControllers() []string {
	if m != nil {
		return m.Controllers
	}
	return nil
}

func (m *WhoIs) GetType() WhoIsType {
	if m != nil {
		return m.Type
	}
	return WhoIsType_USER
}

func (m *WhoIs) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *WhoIs) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *WhoIs) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Alias is a message detailing a known "alsoKnownAs" identifier of a DIDDocument and contains properties for transfer/exchange
type Alias struct {
	// Name is the string name of an Alias
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// IsForSale is the boolean value indicating if the Alias is for sale
	IsForSale bool `protobuf:"varint,2,opt,name=is_for_sale,json=isForSale,proto3" json:"is_for_sale,omitempty"`
	// Amount is the amount listed for purchasing the Alias from the User/Application
	Amount int32 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *Alias) Reset()         { *m = Alias{} }
func (m *Alias) String() string { return proto.CompactTextString(m) }
func (*Alias) ProtoMessage()    {}
func (*Alias) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d9bdfc8d37d9424, []int{1}
}
func (m *Alias) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Alias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Alias.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Alias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alias.Merge(m, src)
}
func (m *Alias) XXX_Size() int {
	return m.Size()
}
func (m *Alias) XXX_DiscardUnknown() {
	xxx_messageInfo_Alias.DiscardUnknown(m)
}

var xxx_messageInfo_Alias proto.InternalMessageInfo

func (m *Alias) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Alias) GetIsForSale() bool {
	if m != nil {
		return m.IsForSale
	}
	return false
}

func (m *Alias) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterEnum("sonrio.sonr.registry.WhoIsType", WhoIsType_name, WhoIsType_value)
	proto.RegisterType((*WhoIs)(nil), "sonrio.sonr.registry.WhoIs")
	proto.RegisterMapType((map[string]string)(nil), "sonrio.sonr.registry.WhoIs.MetadataEntry")
	proto.RegisterType((*Alias)(nil), "sonrio.sonr.registry.Alias")
}

func init() { proto.RegisterFile("registry/who_is.proto", fileDescriptor_9d9bdfc8d37d9424) }

var fileDescriptor_9d9bdfc8d37d9424 = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0xc7, 0x3b, 0x9b, 0xa6, 0x26, 0x27, 0x7e, 0x94, 0x61, 0x5d, 0xc2, 0xae, 0xc4, 0xb8, 0x17,
	0x12, 0x05, 0x53, 0xec, 0xde, 0x88, 0x5e, 0x55, 0x5b, 0xa1, 0xe0, 0x47, 0x99, 0xae, 0x08, 0x22,
	0x84, 0xd9, 0x66, 0xdc, 0x0e, 0x26, 0x99, 0x32, 0x33, 0xdd, 0x35, 0x6f, 0xe1, 0x83, 0xf8, 0x20,
	0x5e, 0xee, 0xa5, 0x97, 0xd2, 0xbe, 0x88, 0x64, 0x9a, 0x46, 0x85, 0xb2, 0x57, 0x33, 0xff, 0xc9,
	0xef, 0x9c, 0x30, 0xbf, 0x39, 0x70, 0x57, 0xb2, 0x73, 0xae, 0xb4, 0x2c, 0x7b, 0x97, 0x73, 0x91,
	0x70, 0x15, 0x2f, 0xa4, 0xd0, 0x02, 0xef, 0x2b, 0x51, 0x48, 0x2e, 0xe2, 0x6a, 0x89, 0xb7, 0xc8,
	0x21, 0x6e, 0xe0, 0x94, 0xa7, 0x1b, 0xf2, 0xf8, 0x87, 0x05, 0xf6, 0xc7, 0xb9, 0x18, 0x2b, 0xfc,
	0x14, 0x6c, 0x9a, 0x71, 0xaa, 0x7c, 0x14, 0x5a, 0x91, 0xd7, 0x3f, 0x8a, 0x77, 0xf5, 0x88, 0x07,
	0x15, 0x42, 0x36, 0x24, 0xde, 0x07, 0x5b, 0x5c, 0x16, 0x4c, 0xfa, 0x7b, 0x21, 0x8a, 0x5c, 0xb2,
	0x09, 0x78, 0x08, 0x37, 0x53, 0x9e, 0x26, 0xa9, 0x98, 0x2d, 0x73, 0x56, 0x68, 0xdf, 0x0a, 0x51,
	0xe4, 0xf5, 0x1f, 0xec, 0xee, 0x37, 0x1c, 0x0f, 0x87, 0x35, 0x48, 0xbc, 0x94, 0xa7, 0xdb, 0x80,
	0x43, 0xf0, 0x66, 0xa2, 0xd0, 0x52, 0x64, 0x19, 0x93, 0xca, 0x6f, 0x87, 0x56, 0xe4, 0x92, 0x7f,
	0x8f, 0xf0, 0x09, 0xb4, 0x75, 0xb9, 0x60, 0xbe, 0x1d, 0xa2, 0xe8, 0x76, 0xff, 0xfe, 0xee, 0xfe,
	0xe6, 0x6e, 0xa7, 0xe5, 0x82, 0x11, 0x03, 0xe3, 0x7b, 0xe0, 0x6a, 0x9e, 0x33, 0xa5, 0x69, 0xbe,
	0xf0, 0x3b, 0x21, 0x8a, 0x2c, 0xf2, 0xf7, 0x00, 0x1f, 0x81, 0xcb, 0x55, 0x42, 0x67, 0x9a, 0x5f,
	0x30, 0xff, 0x46, 0x88, 0x22, 0x87, 0x38, 0x5c, 0x0d, 0x4c, 0xc6, 0x23, 0x70, 0x72, 0xa6, 0x69,
	0x4a, 0x35, 0xf5, 0x1d, 0xe3, 0xe8, 0xd1, 0x35, 0xff, 0x8c, 0xdf, 0xd6, 0xec, 0xa8, 0xd0, 0xb2,
	0x24, 0x4d, 0xe9, 0xe1, 0x0b, 0xb8, 0xf5, 0xdf, 0x27, 0xdc, 0x05, 0xeb, 0x2b, 0x2b, 0x7d, 0x64,
	0x1c, 0x56, 0xdb, 0xca, 0xeb, 0x05, 0xcd, 0x96, 0x6c, 0xeb, 0xd5, 0x84, 0xe7, 0x7b, 0xcf, 0xd0,
	0xf1, 0x14, 0x6c, 0xf3, 0x02, 0x18, 0x43, 0xbb, 0xa0, 0x39, 0xab, 0xab, 0xcc, 0x1e, 0x07, 0xe0,
	0x71, 0x95, 0x7c, 0x11, 0x32, 0x51, 0x34, 0xdb, 0x14, 0x3b, 0xc4, 0xe5, 0xea, 0xb5, 0x90, 0x53,
	0x9a, 0x31, 0x7c, 0x00, 0x1d, 0x9a, 0x8b, 0x65, 0xfd, 0x24, 0x36, 0xa9, 0xd3, 0xe3, 0x87, 0xe0,
	0x36, 0x9a, 0xb0, 0x03, 0xed, 0x0f, 0xd3, 0x11, 0xe9, 0xb6, 0xf0, 0x1d, 0xf0, 0x06, 0x93, 0xc9,
	0x9b, 0xf1, 0xab, 0xc1, 0xe9, 0xf8, 0xfd, 0xbb, 0x2e, 0x7a, 0xf9, 0xf9, 0xe7, 0x2a, 0x40, 0x57,
	0xab, 0x00, 0xfd, 0x5e, 0x05, 0xe8, 0xfb, 0x3a, 0x68, 0x5d, 0xad, 0x83, 0xd6, 0xaf, 0x75, 0xd0,
	0x82, 0x83, 0xad, 0x83, 0xca, 0xb1, 0x6a, 0x4c, 0x4c, 0xd0, 0xa7, 0xe8, 0x9c, 0xeb, 0xf9, 0xf2,
	0x2c, 0x9e, 0x89, 0xbc, 0x57, 0x11, 0x4f, 0xb8, 0x30, 0x6b, 0xef, 0x5b, 0xaf, 0x99, 0x46, 0x53,
	0x74, 0xd6, 0x31, 0x03, 0x79, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x96, 0x32, 0x07, 0xd3,
	0x02, 0x00, 0x00,
}

func (m *WhoIs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WhoIs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WhoIs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		for k := range m.Metadata {
			v := m.Metadata[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintWhoIs(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintWhoIs(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintWhoIs(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x42
		}
	}
	if m.IsActive {
		i--
		if m.IsActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.Timestamp != 0 {
		i = encodeVarintWhoIs(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x30
	}
	if m.Type != 0 {
		i = encodeVarintWhoIs(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Controllers) > 0 {
		for iNdEx := len(m.Controllers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Controllers[iNdEx])
			copy(dAtA[i:], m.Controllers[iNdEx])
			i = encodeVarintWhoIs(dAtA, i, uint64(len(m.Controllers[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.DidDocument != nil {
		{
			size, err := m.DidDocument.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWhoIs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintWhoIs(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Alias) > 0 {
		for iNdEx := len(m.Alias) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Alias[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintWhoIs(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Alias) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Alias) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Alias) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintWhoIs(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if m.IsForSale {
		i--
		if m.IsForSale {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintWhoIs(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWhoIs(dAtA []byte, offset int, v uint64) int {
	offset -= sovWhoIs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WhoIs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Alias) > 0 {
		for _, e := range m.Alias {
			l = e.Size()
			n += 1 + l + sovWhoIs(uint64(l))
		}
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovWhoIs(uint64(l))
	}
	if m.DidDocument != nil {
		l = m.DidDocument.Size()
		n += 1 + l + sovWhoIs(uint64(l))
	}
	if len(m.Controllers) > 0 {
		for _, s := range m.Controllers {
			l = len(s)
			n += 1 + l + sovWhoIs(uint64(l))
		}
	}
	if m.Type != 0 {
		n += 1 + sovWhoIs(uint64(m.Type))
	}
	if m.Timestamp != 0 {
		n += 1 + sovWhoIs(uint64(m.Timestamp))
	}
	if m.IsActive {
		n += 2
	}
	if len(m.Metadata) > 0 {
		for k, v := range m.Metadata {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovWhoIs(uint64(len(k))) + 1 + len(v) + sovWhoIs(uint64(len(v)))
			n += mapEntrySize + 1 + sovWhoIs(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *Alias) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovWhoIs(uint64(l))
	}
	if m.IsForSale {
		n += 2
	}
	if m.Amount != 0 {
		n += 1 + sovWhoIs(uint64(m.Amount))
	}
	return n
}

func sovWhoIs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWhoIs(x uint64) (n int) {
	return sovWhoIs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WhoIs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhoIs
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
			return fmt.Errorf("proto: WhoIs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WhoIs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alias", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhoIs
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
				return ErrInvalidLengthWhoIs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWhoIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alias = append(m.Alias, &Alias{})
			if err := m.Alias[len(m.Alias)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhoIs
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
				return ErrInvalidLengthWhoIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhoIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidDocument", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhoIs
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
				return ErrInvalidLengthWhoIs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWhoIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DidDocument == nil {
				m.DidDocument = &DIDDocument{}
			}
			if err := m.DidDocument.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controllers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhoIs
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
				return ErrInvalidLengthWhoIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhoIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controllers = append(m.Controllers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhoIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= WhoIsType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhoIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhoIs
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
			m.IsActive = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhoIs
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
				return ErrInvalidLengthWhoIs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWhoIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowWhoIs
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWhoIs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthWhoIs
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthWhoIs
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWhoIs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthWhoIs
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthWhoIs
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipWhoIs(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthWhoIs
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Metadata[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWhoIs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhoIs
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
func (m *Alias) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhoIs
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
			return fmt.Errorf("proto: Alias: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Alias: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhoIs
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
				return ErrInvalidLengthWhoIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhoIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsForSale", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhoIs
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
			m.IsForSale = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhoIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWhoIs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhoIs
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
func skipWhoIs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWhoIs
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
					return 0, ErrIntOverflowWhoIs
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
					return 0, ErrIntOverflowWhoIs
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
				return 0, ErrInvalidLengthWhoIs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWhoIs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWhoIs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWhoIs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWhoIs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWhoIs = fmt.Errorf("proto: unexpected end of group")
)
