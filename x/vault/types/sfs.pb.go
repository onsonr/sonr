// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/vault/sfs.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// WalletMail is a message that can be sent to a WalletMailbox
type WalletMail struct {
	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	From      string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To        string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Subject   string `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Body      string `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Signature []byte `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	Timestamp int64  `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *WalletMail) Reset()         { *m = WalletMail{} }
func (m *WalletMail) String() string { return proto.CompactTextString(m) }
func (*WalletMail) ProtoMessage()    {}
func (*WalletMail) Descriptor() ([]byte, []int) {
	return fileDescriptor_25fd1d18d000e9f4, []int{0}
}
func (m *WalletMail) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WalletMail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WalletMail.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WalletMail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletMail.Merge(m, src)
}
func (m *WalletMail) XXX_Size() int {
	return m.Size()
}
func (m *WalletMail) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletMail.DiscardUnknown(m)
}

var xxx_messageInfo_WalletMail proto.InternalMessageInfo

func (m *WalletMail) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *WalletMail) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *WalletMail) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *WalletMail) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *WalletMail) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *WalletMail) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *WalletMail) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

// WalletMailbox is a mailbox for an account address that is controlled by a
// Sonr Identity
type WalletMailbox struct {
	Owner   string        `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Address string        `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Inbox   []*WalletMail `protobuf:"bytes,3,rep,name=inbox,proto3" json:"inbox,omitempty"`
}

func (m *WalletMailbox) Reset()         { *m = WalletMailbox{} }
func (m *WalletMailbox) String() string { return proto.CompactTextString(m) }
func (*WalletMailbox) ProtoMessage()    {}
func (*WalletMailbox) Descriptor() ([]byte, []int) {
	return fileDescriptor_25fd1d18d000e9f4, []int{1}
}
func (m *WalletMailbox) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WalletMailbox) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WalletMailbox.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WalletMailbox) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletMailbox.Merge(m, src)
}
func (m *WalletMailbox) XXX_Size() int {
	return m.Size()
}
func (m *WalletMailbox) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletMailbox.DiscardUnknown(m)
}

var xxx_messageInfo_WalletMailbox proto.InternalMessageInfo

func (m *WalletMailbox) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *WalletMailbox) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *WalletMailbox) GetInbox() []*WalletMail {
	if m != nil {
		return m.Inbox
	}
	return nil
}

// VaultAccount represents the IPFS Vault account data the key for this data type would be the associated DID within the registry
type VaultAccount struct {
	N          int64  `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
	P          string `protobuf:"bytes,2,opt,name=p,proto3" json:"p,omitempty"`
	CoinType   uint64 `protobuf:"varint,3,opt,name=coin_type,json=coinType,proto3" json:"coin_type,omitempty"`
	Controller string `protobuf:"bytes,4,opt,name=controller,proto3" json:"controller,omitempty"`
}

func (m *VaultAccount) Reset()         { *m = VaultAccount{} }
func (m *VaultAccount) String() string { return proto.CompactTextString(m) }
func (*VaultAccount) ProtoMessage()    {}
func (*VaultAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_25fd1d18d000e9f4, []int{2}
}
func (m *VaultAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VaultAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VaultAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultAccount.Merge(m, src)
}
func (m *VaultAccount) XXX_Size() int {
	return m.Size()
}
func (m *VaultAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultAccount.DiscardUnknown(m)
}

var xxx_messageInfo_VaultAccount proto.InternalMessageInfo

func (m *VaultAccount) GetN() int64 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *VaultAccount) GetP() string {
	if m != nil {
		return m.P
	}
	return ""
}

func (m *VaultAccount) GetCoinType() uint64 {
	if m != nil {
		return m.CoinType
	}
	return 0
}

func (m *VaultAccount) GetController() string {
	if m != nil {
		return m.Controller
	}
	return ""
}

// VaultKeyshare represents the underlying MPC shard that is used to sequence an
// account. The Key for this structure is based off the Authentication Fragment
// used to claim the account.
type VaultKeyshare struct {
	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Config   []byte `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	CoinType uint32 `protobuf:"varint,3,opt,name=coin_type,json=coinType,proto3" json:"coin_type,omitempty"`
}

func (m *VaultKeyshare) Reset()         { *m = VaultKeyshare{} }
func (m *VaultKeyshare) String() string { return proto.CompactTextString(m) }
func (*VaultKeyshare) ProtoMessage()    {}
func (*VaultKeyshare) Descriptor() ([]byte, []int) {
	return fileDescriptor_25fd1d18d000e9f4, []int{3}
}
func (m *VaultKeyshare) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultKeyshare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VaultKeyshare.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VaultKeyshare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultKeyshare.Merge(m, src)
}
func (m *VaultKeyshare) XXX_Size() int {
	return m.Size()
}
func (m *VaultKeyshare) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultKeyshare.DiscardUnknown(m)
}

var xxx_messageInfo_VaultKeyshare proto.InternalMessageInfo

func (m *VaultKeyshare) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VaultKeyshare) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *VaultKeyshare) GetCoinType() uint32 {
	if m != nil {
		return m.CoinType
	}
	return 0
}

func init() {
	proto.RegisterType((*WalletMail)(nil), "core.vault.WalletMail")
	proto.RegisterType((*WalletMailbox)(nil), "core.vault.WalletMailbox")
	proto.RegisterType((*VaultAccount)(nil), "core.vault.VaultAccount")
	proto.RegisterType((*VaultKeyshare)(nil), "core.vault.VaultKeyshare")
}

func init() { proto.RegisterFile("core/vault/sfs.proto", fileDescriptor_25fd1d18d000e9f4) }

var fileDescriptor_25fd1d18d000e9f4 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x41, 0x8b, 0xd4, 0x30,
	0x14, 0xc7, 0x27, 0xd3, 0x99, 0x59, 0xe7, 0x39, 0xe3, 0x21, 0x2c, 0x4b, 0x40, 0x29, 0xa5, 0x78,
	0xe8, 0x41, 0x5a, 0xd0, 0xbb, 0xa0, 0x57, 0xf1, 0x52, 0x16, 0x05, 0x2f, 0xd2, 0xa6, 0x99, 0x4e,
	0xa4, 0xcd, 0xab, 0x49, 0xaa, 0xd3, 0x6f, 0xe1, 0xe7, 0xf0, 0x93, 0x78, 0xdc, 0xa3, 0x47, 0x99,
	0xf9, 0x22, 0x92, 0xb4, 0x4b, 0x17, 0xf6, 0x96, 0xff, 0x2f, 0x2f, 0xef, 0xff, 0xe7, 0xe5, 0xc1,
	0x35, 0x47, 0x2d, 0xb2, 0x1f, 0x45, 0xdf, 0xd8, 0xcc, 0x1c, 0x4c, 0xda, 0x69, 0xb4, 0x48, 0xc1,
	0xd1, 0xd4, 0xd3, 0xf8, 0x37, 0x01, 0xf8, 0x5c, 0x34, 0x8d, 0xb0, 0x1f, 0x0b, 0xd9, 0xd0, 0x67,
	0xb0, 0x94, 0x15, 0x23, 0x11, 0x49, 0xb6, 0xf9, 0x52, 0x56, 0x94, 0xc2, 0xea, 0xa0, 0xb1, 0x65,
	0x4b, 0x4f, 0xfc, 0xd9, 0xd5, 0x58, 0x64, 0xc1, 0x58, 0x63, 0x91, 0x32, 0xb8, 0x32, 0x7d, 0xf9,
	0x4d, 0x70, 0xcb, 0x56, 0x1e, 0xde, 0x4b, 0xf7, 0xba, 0xc4, 0x6a, 0x60, 0xeb, 0xf1, 0xb5, 0x3b,
	0xd3, 0x17, 0xb0, 0x35, 0xb2, 0x56, 0x85, 0xed, 0xb5, 0x60, 0x9b, 0x88, 0x24, 0xbb, 0x7c, 0x06,
	0xee, 0xd6, 0xca, 0x56, 0x18, 0x5b, 0xb4, 0x1d, 0xbb, 0x8a, 0x48, 0x12, 0xe4, 0x33, 0x88, 0x5b,
	0xd8, 0xcf, 0x59, 0x4b, 0x3c, 0xd1, 0x6b, 0x58, 0xe3, 0x4f, 0x25, 0xf4, 0x94, 0x78, 0x14, 0x2e,
	0x50, 0x51, 0x55, 0x5a, 0x18, 0x33, 0xe5, 0xbe, 0x97, 0xf4, 0x15, 0xac, 0xa5, 0x2a, 0xf1, 0xc4,
	0x82, 0x28, 0x48, 0x9e, 0xbe, 0xbe, 0x49, 0xe7, 0x49, 0xa4, 0x73, 0xe7, 0x7c, 0x2c, 0x8a, 0x6b,
	0xd8, 0x7d, 0x72, 0x57, 0xef, 0x38, 0xc7, 0x5e, 0x59, 0xba, 0x03, 0xa2, 0xbc, 0x53, 0x90, 0x13,
	0xe5, 0x54, 0x37, 0xf5, 0x27, 0x1d, 0x7d, 0x0e, 0x5b, 0x8e, 0x52, 0x7d, 0xb5, 0x43, 0x27, 0xfc,
	0x6c, 0x56, 0xf9, 0x13, 0x07, 0x6e, 0x87, 0x4e, 0xd0, 0x10, 0x80, 0xa3, 0xb2, 0x1a, 0x9b, 0x46,
	0xe8, 0x69, 0x48, 0x0f, 0x48, 0x7c, 0x0b, 0x7b, 0x6f, 0xf4, 0x41, 0x0c, 0xe6, 0x58, 0x68, 0xf1,
	0xe8, 0x1b, 0x6e, 0x60, 0xc3, 0x51, 0x1d, 0x64, 0xed, 0x0d, 0x77, 0xf9, 0xa4, 0x1e, 0xbb, 0xee,
	0x67, 0xd7, 0xf7, 0x6f, 0xff, 0x9c, 0x43, 0x72, 0x77, 0x0e, 0xc9, 0xbf, 0x73, 0x48, 0x7e, 0x5d,
	0xc2, 0xc5, 0xdd, 0x25, 0x5c, 0xfc, 0xbd, 0x84, 0x8b, 0x2f, 0x2f, 0x6b, 0x69, 0x8f, 0x7d, 0x99,
	0x72, 0x6c, 0x33, 0x83, 0x4a, 0x1f, 0xbf, 0x67, 0x7e, 0x51, 0x4e, 0xd3, 0xaa, 0xb8, 0x76, 0xa6,
	0xdc, 0xf8, 0x6d, 0x79, 0xf3, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x56, 0xe8, 0x5c, 0x45, 0x02,
	0x00, 0x00,
}

func (m *WalletMail) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WalletMail) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WalletMail) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		i = encodeVarintSfs(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintSfs(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintSfs(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintSfs(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintSfs(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintSfs(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintSfs(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WalletMailbox) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WalletMailbox) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WalletMailbox) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Inbox) > 0 {
		for iNdEx := len(m.Inbox) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Inbox[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSfs(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSfs(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintSfs(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VaultAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Controller) > 0 {
		i -= len(m.Controller)
		copy(dAtA[i:], m.Controller)
		i = encodeVarintSfs(dAtA, i, uint64(len(m.Controller)))
		i--
		dAtA[i] = 0x22
	}
	if m.CoinType != 0 {
		i = encodeVarintSfs(dAtA, i, uint64(m.CoinType))
		i--
		dAtA[i] = 0x18
	}
	if len(m.P) > 0 {
		i -= len(m.P)
		copy(dAtA[i:], m.P)
		i = encodeVarintSfs(dAtA, i, uint64(len(m.P)))
		i--
		dAtA[i] = 0x12
	}
	if m.N != 0 {
		i = encodeVarintSfs(dAtA, i, uint64(m.N))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *VaultKeyshare) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultKeyshare) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultKeyshare) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CoinType != 0 {
		i = encodeVarintSfs(dAtA, i, uint64(m.CoinType))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Config) > 0 {
		i -= len(m.Config)
		copy(dAtA[i:], m.Config)
		i = encodeVarintSfs(dAtA, i, uint64(len(m.Config)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintSfs(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSfs(dAtA []byte, offset int, v uint64) int {
	offset -= sovSfs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WalletMail) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSfs(uint64(l))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovSfs(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovSfs(uint64(l))
	}
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovSfs(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovSfs(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovSfs(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovSfs(uint64(m.Timestamp))
	}
	return n
}

func (m *WalletMailbox) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovSfs(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSfs(uint64(l))
	}
	if len(m.Inbox) > 0 {
		for _, e := range m.Inbox {
			l = e.Size()
			n += 1 + l + sovSfs(uint64(l))
		}
	}
	return n
}

func (m *VaultAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.N != 0 {
		n += 1 + sovSfs(uint64(m.N))
	}
	l = len(m.P)
	if l > 0 {
		n += 1 + l + sovSfs(uint64(l))
	}
	if m.CoinType != 0 {
		n += 1 + sovSfs(uint64(m.CoinType))
	}
	l = len(m.Controller)
	if l > 0 {
		n += 1 + l + sovSfs(uint64(l))
	}
	return n
}

func (m *VaultKeyshare) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSfs(uint64(l))
	}
	l = len(m.Config)
	if l > 0 {
		n += 1 + l + sovSfs(uint64(l))
	}
	if m.CoinType != 0 {
		n += 1 + sovSfs(uint64(m.CoinType))
	}
	return n
}

func sovSfs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSfs(x uint64) (n int) {
	return sovSfs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WalletMail) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSfs
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
			return fmt.Errorf("proto: WalletMail: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WalletMail: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSfs
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
				return ErrInvalidLengthSfs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSfs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSfs
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
				return ErrInvalidLengthSfs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSfs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSfs
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
				return ErrInvalidLengthSfs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSfs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSfs
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
				return ErrInvalidLengthSfs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSfs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSfs
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
				return ErrInvalidLengthSfs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSfs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSfs
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
				return ErrInvalidLengthSfs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSfs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSfs
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
		default:
			iNdEx = preIndex
			skippy, err := skipSfs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSfs
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
func (m *WalletMailbox) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSfs
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
			return fmt.Errorf("proto: WalletMailbox: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WalletMailbox: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSfs
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
				return ErrInvalidLengthSfs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSfs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSfs
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
				return ErrInvalidLengthSfs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSfs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inbox", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSfs
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
				return ErrInvalidLengthSfs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSfs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Inbox = append(m.Inbox, &WalletMail{})
			if err := m.Inbox[len(m.Inbox)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSfs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSfs
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
func (m *VaultAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSfs
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
			return fmt.Errorf("proto: VaultAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field N", wireType)
			}
			m.N = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSfs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.N |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field P", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSfs
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
				return ErrInvalidLengthSfs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSfs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.P = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinType", wireType)
			}
			m.CoinType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSfs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoinType |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Controller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSfs
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
				return ErrInvalidLengthSfs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSfs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Controller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSfs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSfs
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
func (m *VaultKeyshare) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSfs
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
			return fmt.Errorf("proto: VaultKeyshare: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultKeyshare: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSfs
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
				return ErrInvalidLengthSfs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSfs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSfs
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
				return ErrInvalidLengthSfs
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSfs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Config = append(m.Config[:0], dAtA[iNdEx:postIndex]...)
			if m.Config == nil {
				m.Config = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoinType", wireType)
			}
			m.CoinType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSfs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoinType |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSfs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSfs
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
func skipSfs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSfs
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
					return 0, ErrIntOverflowSfs
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
					return 0, ErrIntOverflowSfs
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
				return 0, ErrInvalidLengthSfs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSfs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSfs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSfs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSfs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSfs = fmt.Errorf("proto: unexpected end of group")
)
