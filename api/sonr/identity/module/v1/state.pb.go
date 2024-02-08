// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: sonr/identity/module/v1/state.proto

package modulev1

import (
	_ "cosmossdk.io/api/cosmos/orm/v1"
	_ "cosmossdk.io/api/cosmos/orm/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CoinType is the BIP-0044 coin type for each supported coin.
type CoinType int32

const (
	// Bitcoins coin type is 0
	CoinType_COIN_TYPE_UNSPECIFIED CoinType = 0
	// Testnet coin type is 1
	CoinType_COIN_TYPE_ATOM CoinType = 1
	// Litecoin coin type is 2
	CoinType_COIN_TYPE_AXELAR CoinType = 2
	// Dogecoin coin type is 3
	CoinType_COIN_TYPE_BITCOIN CoinType = 3
	// Ethereum coin type is 60
	CoinType_COIN_TYPE_ETHEREUM CoinType = 4
	// Sonr coin type is 703
	CoinType_COIN_TYPE_EVMOS CoinType = 5
	// Cosmos coin type is 118
	CoinType_COIN_TYPE_FILECOIN CoinType = 6
	// Filecoin coin type is 461
	CoinType_COIN_TYPE_JUNO CoinType = 7
	// Handshake coin type is 5353
	CoinType_COIN_TYPE_OSMO CoinType = 8
	// Solana coin type is 501
	CoinType_COIN_TYPE_SOLANA CoinType = 9
	// Ripple coin type is 144
	CoinType_COIN_TYPE_SONR CoinType = 10
	// Stargaze coin type is 1001
	CoinType_COIN_TYPE_STARGAZE CoinType = 11
)

// Enum value maps for CoinType.
var (
	CoinType_name = map[int32]string{
		0:  "COIN_TYPE_UNSPECIFIED",
		1:  "COIN_TYPE_ATOM",
		2:  "COIN_TYPE_AXELAR",
		3:  "COIN_TYPE_BITCOIN",
		4:  "COIN_TYPE_ETHEREUM",
		5:  "COIN_TYPE_EVMOS",
		6:  "COIN_TYPE_FILECOIN",
		7:  "COIN_TYPE_JUNO",
		8:  "COIN_TYPE_OSMO",
		9:  "COIN_TYPE_SOLANA",
		10: "COIN_TYPE_SONR",
		11: "COIN_TYPE_STARGAZE",
	}
	CoinType_value = map[string]int32{
		"COIN_TYPE_UNSPECIFIED": 0,
		"COIN_TYPE_ATOM":        1,
		"COIN_TYPE_AXELAR":      2,
		"COIN_TYPE_BITCOIN":     3,
		"COIN_TYPE_ETHEREUM":    4,
		"COIN_TYPE_EVMOS":       5,
		"COIN_TYPE_FILECOIN":    6,
		"COIN_TYPE_JUNO":        7,
		"COIN_TYPE_OSMO":        8,
		"COIN_TYPE_SOLANA":      9,
		"COIN_TYPE_SONR":        10,
		"COIN_TYPE_STARGAZE":    11,
	}
)

func (x CoinType) Enum() *CoinType {
	p := new(CoinType)
	*p = x
	return p
}

func (x CoinType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CoinType) Descriptor() protoreflect.EnumDescriptor {
	return file_sonr_identity_module_v1_state_proto_enumTypes[0].Descriptor()
}

func (CoinType) Type() protoreflect.EnumType {
	return &file_sonr_identity_module_v1_state_proto_enumTypes[0]
}

func (x CoinType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoinType.Descriptor instead.
func (CoinType) EnumDescriptor() ([]byte, []int) {
	return file_sonr_identity_module_v1_state_proto_rawDescGZIP(), []int{0}
}

// Module is the app config object of the module.
// Learn more: https://docs.cosmos.network/main/building-modules/depinject
type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_identity_module_v1_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_identity_module_v1_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_sonr_identity_module_v1_state_proto_rawDescGZIP(), []int{0}
}

// Account is the root sonr account table which contains all sub-identities.
type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sequence   uint64   `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Controller string   `protobuf:"bytes,2,opt,name=controller,proto3" json:"controller,omitempty"`
	CoinType   CoinType `protobuf:"varint,3,opt,name=coin_type,json=coinType,proto3,enum=sonr.identity.module.v1.CoinType" json:"coin_type,omitempty"`
	PublicKey  []byte   `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Network    string   `protobuf:"bytes,5,opt,name=network,proto3" json:"network,omitempty"`
	Address    string   `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	ChainId    string   `protobuf:"bytes,7,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_identity_module_v1_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_identity_module_v1_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_sonr_identity_module_v1_state_proto_rawDescGZIP(), []int{1}
}

func (x *Account) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *Account) GetController() string {
	if x != nil {
		return x.Controller
	}
	return ""
}

func (x *Account) GetCoinType() CoinType {
	if x != nil {
		return x.CoinType
	}
	return CoinType_COIN_TYPE_UNSPECIFIED
}

func (x *Account) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Account) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Account) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Account) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

// Blockchain is the configuration table for connected blockchains
type Blockchain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index     uint64   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	ChainId   string   `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ChainCode uint32   `protobuf:"varint,3,opt,name=chain_code,json=chainCode,proto3" json:"chain_code,omitempty"`
	Name      string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Hrp       string   `protobuf:"bytes,5,opt,name=hrp,proto3" json:"hrp,omitempty"`
	DidMethod string   `protobuf:"bytes,6,opt,name=did_method,json=didMethod,proto3" json:"did_method,omitempty"`
	Denoms    []string `protobuf:"bytes,7,rep,name=denoms,proto3" json:"denoms,omitempty"`
	ChannelId string   `protobuf:"bytes,8,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *Blockchain) Reset() {
	*x = Blockchain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_identity_module_v1_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blockchain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blockchain) ProtoMessage() {}

func (x *Blockchain) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_identity_module_v1_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blockchain.ProtoReflect.Descriptor instead.
func (*Blockchain) Descriptor() ([]byte, []int) {
	return file_sonr_identity_module_v1_state_proto_rawDescGZIP(), []int{2}
}

func (x *Blockchain) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Blockchain) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

func (x *Blockchain) GetChainCode() uint32 {
	if x != nil {
		return x.ChainCode
	}
	return 0
}

func (x *Blockchain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Blockchain) GetHrp() string {
	if x != nil {
		return x.Hrp
	}
	return ""
}

func (x *Blockchain) GetDidMethod() string {
	if x != nil {
		return x.DidMethod
	}
	return ""
}

func (x *Blockchain) GetDenoms() []string {
	if x != nil {
		return x.Denoms
	}
	return nil
}

func (x *Blockchain) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

// Identifier is a psuedo-anonomyous representation of a unique id on the Sonr blockchain. Used as
// authorizer to the underlying wallet interface.
type Accumulator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index      uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Controller string `protobuf:"bytes,2,opt,name=controller,proto3" json:"controller,omitempty"`
	Key        string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value      string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Accumulator) Reset() {
	*x = Accumulator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_identity_module_v1_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Accumulator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Accumulator) ProtoMessage() {}

func (x *Accumulator) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_identity_module_v1_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Accumulator.ProtoReflect.Descriptor instead.
func (*Accumulator) Descriptor() ([]byte, []int) {
	return file_sonr_identity_module_v1_state_proto_rawDescGZIP(), []int{3}
}

func (x *Accumulator) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Accumulator) GetController() string {
	if x != nil {
		return x.Controller
	}
	return ""
}

func (x *Accumulator) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Accumulator) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Controller is the root sonr controller table which contains all sub-identities.
type Controller struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sequence       uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	PeerId         string `protobuf:"bytes,2,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	Address        string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	PublicKey      []byte `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Ipns           string `protobuf:"bytes,5,opt,name=ipns,proto3" json:"ipns,omitempty"`
	AccumulatorKey []byte `protobuf:"bytes,6,opt,name=accumulator_key,json=accumulatorKey,proto3" json:"accumulator_key,omitempty"`
	Network        string `protobuf:"bytes,7,opt,name=network,proto3" json:"network,omitempty"`
	ChainId        string `protobuf:"bytes,8,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (x *Controller) Reset() {
	*x = Controller{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_identity_module_v1_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Controller) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Controller) ProtoMessage() {}

func (x *Controller) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_identity_module_v1_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Controller.ProtoReflect.Descriptor instead.
func (*Controller) Descriptor() ([]byte, []int) {
	return file_sonr_identity_module_v1_state_proto_rawDescGZIP(), []int{4}
}

func (x *Controller) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *Controller) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *Controller) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Controller) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Controller) GetIpns() string {
	if x != nil {
		return x.Ipns
	}
	return ""
}

func (x *Controller) GetAccumulatorKey() []byte {
	if x != nil {
		return x.AccumulatorKey
	}
	return nil
}

func (x *Controller) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Controller) GetChainId() string {
	if x != nil {
		return x.ChainId
	}
	return ""
}

var File_sonr_identity_module_v1_state_proto protoreflect.FileDescriptor

var file_sonr_identity_module_v1_state_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x3a, 0x2a, 0x82, 0x9f, 0xd3, 0x8e, 0x03, 0x24, 0x0a, 0x22, 0x08, 0x01, 0x12, 0x1e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac,
	0x02, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73, 0x6f, 0x6e, 0x72,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x63, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x3a, 0x37, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x31, 0x0a, 0x0c, 0x0a, 0x08,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x10, 0x01, 0x18, 0x01, 0x12, 0x10, 0x0a, 0x0a, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x10, 0x02, 0x18, 0x01, 0x18, 0x01, 0x22, 0x89, 0x02,
	0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x68, 0x72, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x68,
	0x72, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x3a, 0x2f, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x29,
	0x0a, 0x09, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x08, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x10, 0x01, 0x18, 0x01, 0x12, 0x0a, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x10, 0x02, 0x18, 0x01, 0x18, 0x03, 0x22, 0x96, 0x01, 0x0a, 0x0b, 0x41, 0x63,
	0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x29, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x23, 0x0a,
	0x09, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x0e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2c, 0x6b, 0x65, 0x79, 0x10, 0x01, 0x18, 0x01,
	0x18, 0x04, 0x22, 0xc0, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x70, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69,
	0x70, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x61, 0x63,
	0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49,
	0x64, 0x3a, 0x52, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x4c, 0x0a, 0x0c, 0x0a, 0x08, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x10, 0x01, 0x18, 0x01, 0x12, 0x10, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x10, 0x02, 0x18, 0x01, 0x12, 0x0d, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x10, 0x03, 0x18, 0x01, 0x12, 0x0a, 0x0a, 0x04, 0x69, 0x70, 0x6e, 0x73, 0x10,
	0x04, 0x18, 0x01, 0x18, 0x02, 0x2a, 0x95, 0x02, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x54, 0x4f, 0x4d, 0x10,
	0x01, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41,
	0x58, 0x45, 0x4c, 0x41, 0x52, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x49, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x49, 0x54, 0x43, 0x4f, 0x49, 0x4e, 0x10, 0x03, 0x12, 0x16,
	0x0a, 0x12, 0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x54, 0x48, 0x45,
	0x52, 0x45, 0x55, 0x4d, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x45, 0x56, 0x4d, 0x4f, 0x53, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x43,
	0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x43, 0x4f, 0x49,
	0x4e, 0x10, 0x06, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4a, 0x55, 0x4e, 0x4f, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x49, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x53, 0x4d, 0x4f, 0x10, 0x08, 0x12, 0x14, 0x0a, 0x10, 0x43,
	0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4f, 0x4c, 0x41, 0x4e, 0x41, 0x10,
	0x09, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53,
	0x4f, 0x4e, 0x52, 0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x47, 0x41, 0x5a, 0x45, 0x10, 0x0b, 0x42, 0xf0, 0x01,
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2f, 0x73,
	0x6f, 0x6e, 0x72, 0x2f, 0x78, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x49, 0x4d, 0xaa, 0x02, 0x17, 0x53, 0x6f, 0x6e, 0x72,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x53, 0x6f, 0x6e, 0x72, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x23,
	0x53, 0x6f, 0x6e, 0x72, 0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5c, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x53, 0x6f, 0x6e, 0x72, 0x3a, 0x3a, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sonr_identity_module_v1_state_proto_rawDescOnce sync.Once
	file_sonr_identity_module_v1_state_proto_rawDescData = file_sonr_identity_module_v1_state_proto_rawDesc
)

func file_sonr_identity_module_v1_state_proto_rawDescGZIP() []byte {
	file_sonr_identity_module_v1_state_proto_rawDescOnce.Do(func() {
		file_sonr_identity_module_v1_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_sonr_identity_module_v1_state_proto_rawDescData)
	})
	return file_sonr_identity_module_v1_state_proto_rawDescData
}

var file_sonr_identity_module_v1_state_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sonr_identity_module_v1_state_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sonr_identity_module_v1_state_proto_goTypes = []interface{}{
	(CoinType)(0),       // 0: sonr.identity.module.v1.CoinType
	(*State)(nil),       // 1: sonr.identity.module.v1.State
	(*Account)(nil),     // 2: sonr.identity.module.v1.Account
	(*Blockchain)(nil),  // 3: sonr.identity.module.v1.Blockchain
	(*Accumulator)(nil), // 4: sonr.identity.module.v1.Accumulator
	(*Controller)(nil),  // 5: sonr.identity.module.v1.Controller
}
var file_sonr_identity_module_v1_state_proto_depIdxs = []int32{
	0, // 0: sonr.identity.module.v1.Account.coin_type:type_name -> sonr.identity.module.v1.CoinType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sonr_identity_module_v1_state_proto_init() }
func file_sonr_identity_module_v1_state_proto_init() {
	if File_sonr_identity_module_v1_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sonr_identity_module_v1_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sonr_identity_module_v1_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sonr_identity_module_v1_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blockchain); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sonr_identity_module_v1_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Accumulator); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sonr_identity_module_v1_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Controller); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sonr_identity_module_v1_state_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sonr_identity_module_v1_state_proto_goTypes,
		DependencyIndexes: file_sonr_identity_module_v1_state_proto_depIdxs,
		EnumInfos:         file_sonr_identity_module_v1_state_proto_enumTypes,
		MessageInfos:      file_sonr_identity_module_v1_state_proto_msgTypes,
	}.Build()
	File_sonr_identity_module_v1_state_proto = out.File
	file_sonr_identity_module_v1_state_proto_rawDesc = nil
	file_sonr_identity_module_v1_state_proto_goTypes = nil
	file_sonr_identity_module_v1_state_proto_depIdxs = nil
}
