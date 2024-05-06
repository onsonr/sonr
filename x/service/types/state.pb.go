// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service/v1/state.proto

package types

import (
	_ "cosmossdk.io/orm"
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	v1 "github.com/didao-org/sonr/api/common/v1"
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

// ProfileEntity is a profile for a given Service provider.
type ProfileEntity struct {
	Id          string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ZkWitness   string             `protobuf:"bytes,2,opt,name=zk_witness,json=zkWitness,proto3" json:"zk_witness,omitempty"`
	Permissions v1.PermissionsType `protobuf:"varint,3,opt,name=permissions,proto3,enum=common.v1.PermissionsType" json:"permissions,omitempty"`
	Credentials []uint64           `protobuf:"varint,4,rep,packed,name=credentials,proto3" json:"credentials,omitempty"`
	Identifiers []uint64           `protobuf:"varint,5,rep,packed,name=identifiers,proto3" json:"identifiers,omitempty"`
	Wallets     []uint64           `protobuf:"varint,6,rep,packed,name=wallets,proto3" json:"wallets,omitempty"`
}

func (m *ProfileEntity) Reset()         { *m = ProfileEntity{} }
func (m *ProfileEntity) String() string { return proto.CompactTextString(m) }
func (*ProfileEntity) ProtoMessage()    {}
func (*ProfileEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6e3654a2974847, []int{0}
}
func (m *ProfileEntity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProfileEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProfileEntity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProfileEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileEntity.Merge(m, src)
}
func (m *ProfileEntity) XXX_Size() int {
	return m.Size()
}
func (m *ProfileEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileEntity.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileEntity proto.InternalMessageInfo

func (m *ProfileEntity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ProfileEntity) GetZkWitness() string {
	if m != nil {
		return m.ZkWitness
	}
	return ""
}

func (m *ProfileEntity) GetPermissions() v1.PermissionsType {
	if m != nil {
		return m.Permissions
	}
	return v1.PermissionsType_PERMISSIONS_TYPE_UNSPECIFIED
}

func (m *ProfileEntity) GetCredentials() []uint64 {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (m *ProfileEntity) GetIdentifiers() []uint64 {
	if m != nil {
		return m.Identifiers
	}
	return nil
}

func (m *ProfileEntity) GetWallets() []uint64 {
	if m != nil {
		return m.Wallets
	}
	return nil
}

// ServiceRecord is the configuration for a given service provider.
type ServiceRecord struct {
	Sequence     uint64             `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Origin       string             `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	TeamAddress  string             `protobuf:"bytes,3,opt,name=team_address,json=teamAddress,proto3" json:"team_address,omitempty"`
	Metadata     string             `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Permissions  v1.PermissionsType `protobuf:"varint,5,opt,name=permissions,proto3,enum=common.v1.PermissionsType" json:"permissions,omitempty"`
	TldExtension string             `protobuf:"bytes,6,opt,name=tld_extension,json=tldExtension,proto3" json:"tld_extension,omitempty"`
}

func (m *ServiceRecord) Reset()         { *m = ServiceRecord{} }
func (m *ServiceRecord) String() string { return proto.CompactTextString(m) }
func (*ServiceRecord) ProtoMessage()    {}
func (*ServiceRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6e3654a2974847, []int{1}
}
func (m *ServiceRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServiceRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ServiceRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ServiceRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceRecord.Merge(m, src)
}
func (m *ServiceRecord) XXX_Size() int {
	return m.Size()
}
func (m *ServiceRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceRecord proto.InternalMessageInfo

func (m *ServiceRecord) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *ServiceRecord) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *ServiceRecord) GetTeamAddress() string {
	if m != nil {
		return m.TeamAddress
	}
	return ""
}

func (m *ServiceRecord) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *ServiceRecord) GetPermissions() v1.PermissionsType {
	if m != nil {
		return m.Permissions
	}
	return v1.PermissionsType_PERMISSIONS_TYPE_UNSPECIFIED
}

func (m *ServiceRecord) GetTldExtension() string {
	if m != nil {
		return m.TldExtension
	}
	return ""
}

// UserIdentifier is the root sonr user identifier table which contains all sub-identities.
type WebCredential struct {
	// Sequence is the unique identifier for the credential
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// Id is the id of the credential
	Id []byte `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// handle is the handle of the credential
	Handle string `protobuf:"bytes,3,opt,name=handle,proto3" json:"handle,omitempty"`
	// transports is the list of transports supported by the credential
	Transports []string `protobuf:"bytes,4,rep,name=transports,proto3" json:"transports,omitempty"`
	// origin is the origin of the credential
	Origin string `protobuf:"bytes,5,opt,name=origin,proto3" json:"origin,omitempty"`
	// Controller is the address of the owner of the credential
	Controller string `protobuf:"bytes,6,opt,name=controller,proto3" json:"controller,omitempty"`
	// Assertion Type is the type of the credential
	AssertionType string `protobuf:"bytes,7,opt,name=assertion_type,json=assertionType,proto3" json:"assertion_type,omitempty"`
}

func (m *WebCredential) Reset()         { *m = WebCredential{} }
func (m *WebCredential) String() string { return proto.CompactTextString(m) }
func (*WebCredential) ProtoMessage()    {}
func (*WebCredential) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab6e3654a2974847, []int{2}
}
func (m *WebCredential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WebCredential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WebCredential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WebCredential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebCredential.Merge(m, src)
}
func (m *WebCredential) XXX_Size() int {
	return m.Size()
}
func (m *WebCredential) XXX_DiscardUnknown() {
	xxx_messageInfo_WebCredential.DiscardUnknown(m)
}

var xxx_messageInfo_WebCredential proto.InternalMessageInfo

func (m *WebCredential) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *WebCredential) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *WebCredential) GetHandle() string {
	if m != nil {
		return m.Handle
	}
	return ""
}

func (m *WebCredential) GetTransports() []string {
	if m != nil {
		return m.Transports
	}
	return nil
}

func (m *WebCredential) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *WebCredential) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

func (m *WebCredential) GetAssertionType() string {
	if m != nil {
		return m.AssertionType
	}
	return ""
}

func init() {
	proto.RegisterType((*ProfileEntity)(nil), "service.v1.ProfileEntity")
	proto.RegisterType((*ServiceRecord)(nil), "service.v1.ServiceRecord")
	proto.RegisterType((*WebCredential)(nil), "service.v1.WebCredential")
}

func init() { proto.RegisterFile("service/v1/state.proto", fileDescriptor_ab6e3654a2974847) }

var fileDescriptor_ab6e3654a2974847 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6b, 0x14, 0x31,
	0x18, 0xc6, 0x9b, 0xdd, 0xed, 0xb6, 0x7d, 0xbb, 0xb3, 0xd4, 0xa0, 0x35, 0x2c, 0x38, 0xac, 0xad,
	0x42, 0x0b, 0xba, 0x43, 0xf5, 0x56, 0xbc, 0x68, 0xe9, 0xbd, 0x8c, 0x42, 0xc1, 0xcb, 0x92, 0x9d,
	0xbc, 0x6d, 0x43, 0x67, 0x92, 0x35, 0x49, 0xb7, 0x7f, 0xbe, 0x82, 0x20, 0x7e, 0x02, 0x3f, 0x8f,
	0xc7, 0x82, 0x17, 0x8f, 0xd2, 0xde, 0x3c, 0x7a, 0x17, 0x64, 0x92, 0xd9, 0xdd, 0xb1, 0x07, 0xc1,
	0xe3, 0xf3, 0x7b, 0x9f, 0x99, 0xe4, 0x79, 0x92, 0xc0, 0xba, 0x45, 0x33, 0x91, 0x19, 0x26, 0x93,
	0x9d, 0xc4, 0x3a, 0xee, 0x70, 0x30, 0x36, 0xda, 0x69, 0x0a, 0x15, 0x1f, 0x4c, 0x76, 0x7a, 0x0f,
	0x33, 0x6d, 0x0b, 0x6d, 0x13, 0x6d, 0x8a, 0xd2, 0xa6, 0x4d, 0x11, 0x4c, 0xbd, 0xfb, 0x99, 0x2e,
	0x0a, 0xad, 0x4a, 0x38, 0xe2, 0xb6, 0xfa, 0xb4, 0xf7, 0x60, 0x4e, 0x51, 0x9d, 0x15, 0x36, 0xe0,
	0x8d, 0x9f, 0x04, 0xa2, 0x03, 0xa3, 0x8f, 0x64, 0x8e, 0xfb, 0xca, 0x49, 0x77, 0x49, 0xbb, 0xd0,
	0x90, 0x82, 0x91, 0x3e, 0xd9, 0x5a, 0x49, 0x1b, 0x52, 0xd0, 0x47, 0x00, 0x57, 0xa7, 0xc3, 0x73,
	0xe9, 0x14, 0x5a, 0xcb, 0x1a, 0x9e, 0xaf, 0x5c, 0x9d, 0x1e, 0x06, 0x40, 0x5f, 0xc1, 0xea, 0x18,
	0x4d, 0x21, 0xad, 0x95, 0x5a, 0x59, 0xd6, 0xec, 0x93, 0xad, 0xee, 0x8b, 0xde, 0x20, 0xac, 0x36,
	0x98, 0xec, 0x0c, 0x0e, 0xe6, 0xd3, 0x77, 0x97, 0x63, 0x4c, 0xeb, 0x76, 0xda, 0x87, 0xd5, 0xcc,
	0xa0, 0x40, 0xe5, 0x24, 0xcf, 0x2d, 0x6b, 0xf5, 0x9b, 0x5b, 0xad, 0xb4, 0x8e, 0x4a, 0x87, 0xf4,
	0xe2, 0x48, 0xa2, 0xb1, 0x6c, 0x31, 0x38, 0x6a, 0x88, 0x32, 0x58, 0x3a, 0xe7, 0x79, 0x8e, 0xce,
	0xb2, 0xb6, 0x9f, 0x4e, 0xe5, 0x6e, 0xf7, 0xd7, 0x97, 0x6f, 0x9f, 0x9a, 0xcb, 0xd0, 0x0a, 0x91,
	0x36, 0x3e, 0x36, 0x20, 0x7a, 0x1b, 0x1a, 0x4c, 0x31, 0xd3, 0x46, 0xd0, 0x1e, 0x2c, 0x5b, 0xfc,
	0x70, 0x86, 0x2a, 0x43, 0x1f, 0xb9, 0x95, 0xce, 0x34, 0x5d, 0x87, 0xb6, 0x36, 0xf2, 0x58, 0xaa,
	0x2a, 0x74, 0xa5, 0xe8, 0x63, 0xe8, 0x38, 0xe4, 0xc5, 0x90, 0x0b, 0x61, 0xca, 0x4a, 0x9a, 0x7e,
	0xba, 0x5a, 0xb2, 0xd7, 0x01, 0x95, 0xbf, 0x2d, 0xd0, 0x71, 0xc1, 0x1d, 0x67, 0x2d, 0x3f, 0x9e,
	0xe9, 0xbb, 0x85, 0x2d, 0xfe, 0x5f, 0x61, 0x9b, 0x10, 0xb9, 0x5c, 0x0c, 0xf1, 0xc2, 0xa1, 0x2a,
	0x09, 0x6b, 0xfb, 0xdf, 0x77, 0x5c, 0x2e, 0xf6, 0xa7, 0x6c, 0xf7, 0x89, 0xcf, 0x1d, 0x43, 0x67,
	0x9e, 0x6e, 0x8d, 0xd0, 0xce, 0x34, 0xcf, 0x1a, 0x61, 0x84, 0x35, 0x36, 0x7e, 0x13, 0x88, 0x0e,
	0x71, 0xb4, 0x37, 0x2b, 0xfb, 0x9f, 0x6d, 0x84, 0x6b, 0x51, 0x36, 0xd1, 0xf1, 0xd7, 0x62, 0x1d,
	0xda, 0x27, 0x5c, 0x89, 0x1c, 0xab, 0xfc, 0x95, 0xa2, 0x31, 0x80, 0x33, 0x5c, 0xd9, 0xb1, 0x36,
	0x2e, 0x1c, 0xe8, 0x4a, 0x5a, 0x23, 0xb5, 0x56, 0x17, 0xff, 0x6a, 0x35, 0x06, 0xc8, 0xb4, 0x72,
	0x46, 0xe7, 0x39, 0x9a, 0x2a, 0x55, 0x8d, 0xd0, 0xa7, 0xd0, 0xe5, 0xd6, 0xa2, 0x71, 0x52, 0xab,
	0xa1, 0xbb, 0x1c, 0x23, 0x5b, 0xf2, 0x9e, 0x68, 0x46, 0xcb, 0xb2, 0x76, 0xb7, 0x7d, 0xf4, 0xcd,
	0x3b, 0xd1, 0xef, 0x41, 0x14, 0x96, 0x79, 0x16, 0x76, 0xb9, 0x46, 0x58, 0xeb, 0xcd, 0xde, 0xd7,
	0x9b, 0x98, 0x5c, 0xdf, 0xc4, 0xe4, 0xc7, 0x4d, 0x4c, 0x3e, 0xdf, 0xc6, 0x0b, 0xd7, 0xb7, 0xf1,
	0xc2, 0xf7, 0xdb, 0x78, 0xe1, 0xfd, 0xf6, 0xb1, 0x74, 0x27, 0x67, 0xa3, 0xf2, 0x4c, 0x12, 0x21,
	0x05, 0xd7, 0xcf, 0xb5, 0x39, 0x4e, 0xac, 0x56, 0x26, 0xb9, 0x48, 0xa6, 0x4f, 0xb3, 0xdc, 0x83,
	0x1d, 0xb5, 0xfd, 0x33, 0x7a, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x6d, 0x43, 0x65, 0xb2,
	0x03, 0x00, 0x00,
}

func (m *ProfileEntity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProfileEntity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProfileEntity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Wallets) > 0 {
		dAtA2 := make([]byte, len(m.Wallets)*10)
		var j1 int
		for _, num := range m.Wallets {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintState(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Identifiers) > 0 {
		dAtA4 := make([]byte, len(m.Identifiers)*10)
		var j3 int
		for _, num := range m.Identifiers {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintState(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Credentials) > 0 {
		dAtA6 := make([]byte, len(m.Credentials)*10)
		var j5 int
		for _, num := range m.Credentials {
			for num >= 1<<7 {
				dAtA6[j5] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j5++
			}
			dAtA6[j5] = uint8(num)
			j5++
		}
		i -= j5
		copy(dAtA[i:], dAtA6[:j5])
		i = encodeVarintState(dAtA, i, uint64(j5))
		i--
		dAtA[i] = 0x22
	}
	if m.Permissions != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Permissions))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ZkWitness) > 0 {
		i -= len(m.ZkWitness)
		copy(dAtA[i:], m.ZkWitness)
		i = encodeVarintState(dAtA, i, uint64(len(m.ZkWitness)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintState(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ServiceRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ServiceRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TldExtension) > 0 {
		i -= len(m.TldExtension)
		copy(dAtA[i:], m.TldExtension)
		i = encodeVarintState(dAtA, i, uint64(len(m.TldExtension)))
		i--
		dAtA[i] = 0x32
	}
	if m.Permissions != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Permissions))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintState(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TeamAddress) > 0 {
		i -= len(m.TeamAddress)
		copy(dAtA[i:], m.TeamAddress)
		i = encodeVarintState(dAtA, i, uint64(len(m.TeamAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Origin) > 0 {
		i -= len(m.Origin)
		copy(dAtA[i:], m.Origin)
		i = encodeVarintState(dAtA, i, uint64(len(m.Origin)))
		i--
		dAtA[i] = 0x12
	}
	if m.Sequence != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *WebCredential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WebCredential) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WebCredential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AssertionType) > 0 {
		i -= len(m.AssertionType)
		copy(dAtA[i:], m.AssertionType)
		i = encodeVarintState(dAtA, i, uint64(len(m.AssertionType)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintState(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Origin) > 0 {
		i -= len(m.Origin)
		copy(dAtA[i:], m.Origin)
		i = encodeVarintState(dAtA, i, uint64(len(m.Origin)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Transports) > 0 {
		for iNdEx := len(m.Transports) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Transports[iNdEx])
			copy(dAtA[i:], m.Transports[iNdEx])
			i = encodeVarintState(dAtA, i, uint64(len(m.Transports[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Handle) > 0 {
		i -= len(m.Handle)
		copy(dAtA[i:], m.Handle)
		i = encodeVarintState(dAtA, i, uint64(len(m.Handle)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintState(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if m.Sequence != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProfileEntity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.ZkWitness)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.Permissions != 0 {
		n += 1 + sovState(uint64(m.Permissions))
	}
	if len(m.Credentials) > 0 {
		l = 0
		for _, e := range m.Credentials {
			l += sovState(uint64(e))
		}
		n += 1 + sovState(uint64(l)) + l
	}
	if len(m.Identifiers) > 0 {
		l = 0
		for _, e := range m.Identifiers {
			l += sovState(uint64(e))
		}
		n += 1 + sovState(uint64(l)) + l
	}
	if len(m.Wallets) > 0 {
		l = 0
		for _, e := range m.Wallets {
			l += sovState(uint64(e))
		}
		n += 1 + sovState(uint64(l)) + l
	}
	return n
}

func (m *ServiceRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovState(uint64(m.Sequence))
	}
	l = len(m.Origin)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.TeamAddress)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.Permissions != 0 {
		n += 1 + sovState(uint64(m.Permissions))
	}
	l = len(m.TldExtension)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func (m *WebCredential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovState(uint64(m.Sequence))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Handle)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if len(m.Transports) > 0 {
		for _, s := range m.Transports {
			l = len(s)
			n += 1 + l + sovState(uint64(l))
		}
	}
	l = len(m.Origin)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.AssertionType)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProfileEntity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: ProfileEntity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProfileEntity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZkWitness", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ZkWitness = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			m.Permissions = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Permissions |= v1.PermissionsType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowState
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Credentials = append(m.Credentials, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowState
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthState
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthState
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Credentials) == 0 {
					m.Credentials = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowState
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Credentials = append(m.Credentials, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Credentials", wireType)
			}
		case 5:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowState
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Identifiers = append(m.Identifiers, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowState
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthState
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthState
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Identifiers) == 0 {
					m.Identifiers = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowState
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Identifiers = append(m.Identifiers, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifiers", wireType)
			}
		case 6:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowState
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Wallets = append(m.Wallets, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowState
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthState
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthState
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Wallets) == 0 {
					m.Wallets = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowState
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Wallets = append(m.Wallets, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Wallets", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *ServiceRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: ServiceRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Origin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TeamAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TeamAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			m.Permissions = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Permissions |= v1.PermissionsType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TldExtension", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TldExtension = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *WebCredential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: WebCredential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WebCredential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transports", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transports = append(m.Transports, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Origin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssertionType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssertionType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)