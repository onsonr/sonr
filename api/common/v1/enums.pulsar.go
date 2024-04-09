// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package commonv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: common/v1/enums.proto

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
	return file_common_v1_enums_proto_enumTypes[0].Descriptor()
}

func (CoinType) Type() protoreflect.EnumType {
	return &file_common_v1_enums_proto_enumTypes[0]
}

func (x CoinType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoinType.Descriptor instead.
func (CoinType) EnumDescriptor() ([]byte, []int) {
	return file_common_v1_enums_proto_rawDescGZIP(), []int{0}
}

// IdentifierKind is the type of identifier used to identify a user.
type IdentifierKind int32

const (
	// Unspecified identifier kind.
	IdentifierKind_IDENTIFIER_KIND_UNSPECIFIED IdentifierKind = 0
	// Service handle identifier kind.
	IdentifierKind_IDENTIFIER_KIND_SERVICE_HANDLE IdentifierKind = 1
	// Phone number identifier kind.
	IdentifierKind_IDENTIFIER_KIND_PHONE_NUMBER IdentifierKind = 2
	// Email address identifier kind.
	IdentifierKind_IDENTIFIER_KIND_EMAIL_ADDRESS IdentifierKind = 3
	// OAuth provider identifier kind.
	IdentifierKind_IDENTIFIER_KIND_OAUTH_PROVIDER IdentifierKind = 4
	// External wallet identifier kind.
	IdentifierKind_IDENTIFIER_KIND_EXTERNAL_WALLET IdentifierKind = 5
)

// Enum value maps for IdentifierKind.
var (
	IdentifierKind_name = map[int32]string{
		0: "IDENTIFIER_KIND_UNSPECIFIED",
		1: "IDENTIFIER_KIND_SERVICE_HANDLE",
		2: "IDENTIFIER_KIND_PHONE_NUMBER",
		3: "IDENTIFIER_KIND_EMAIL_ADDRESS",
		4: "IDENTIFIER_KIND_OAUTH_PROVIDER",
		5: "IDENTIFIER_KIND_EXTERNAL_WALLET",
	}
	IdentifierKind_value = map[string]int32{
		"IDENTIFIER_KIND_UNSPECIFIED":     0,
		"IDENTIFIER_KIND_SERVICE_HANDLE":  1,
		"IDENTIFIER_KIND_PHONE_NUMBER":    2,
		"IDENTIFIER_KIND_EMAIL_ADDRESS":   3,
		"IDENTIFIER_KIND_OAUTH_PROVIDER":  4,
		"IDENTIFIER_KIND_EXTERNAL_WALLET": 5,
	}
)

func (x IdentifierKind) Enum() *IdentifierKind {
	p := new(IdentifierKind)
	*p = x
	return p
}

func (x IdentifierKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IdentifierKind) Descriptor() protoreflect.EnumDescriptor {
	return file_common_v1_enums_proto_enumTypes[1].Descriptor()
}

func (IdentifierKind) Type() protoreflect.EnumType {
	return &file_common_v1_enums_proto_enumTypes[1]
}

func (x IdentifierKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IdentifierKind.Descriptor instead.
func (IdentifierKind) EnumDescriptor() ([]byte, []int) {
	return file_common_v1_enums_proto_rawDescGZIP(), []int{1}
}

// PermissionsType is the type of permissions granted by a user.
type PermissionsType int32

const (
	// Unspecified permissions type.
	PermissionsType_PERMISSIONS_TYPE_UNSPECIFIED PermissionsType = 0
	// Anonymous permissions type.
	PermissionsType_PERMISSIONS_TYPE_ANON PermissionsType = 1
	// Read permissions type.
	PermissionsType_PERMISSIONS_TYPE_READ PermissionsType = 2
	// Write permissions type.
	PermissionsType_PERMISSIONS_TYPE_WRITE PermissionsType = 3
	// Ownership permissions type.
	PermissionsType_PERMISSIONS_TYPE_OWN PermissionsType = 4
	// Admin permissions type.
	PermissionsType_PERMISSIONS_TYPE_ADMIN PermissionsType = 5
)

// Enum value maps for PermissionsType.
var (
	PermissionsType_name = map[int32]string{
		0: "PERMISSIONS_TYPE_UNSPECIFIED",
		1: "PERMISSIONS_TYPE_ANON",
		2: "PERMISSIONS_TYPE_READ",
		3: "PERMISSIONS_TYPE_WRITE",
		4: "PERMISSIONS_TYPE_OWN",
		5: "PERMISSIONS_TYPE_ADMIN",
	}
	PermissionsType_value = map[string]int32{
		"PERMISSIONS_TYPE_UNSPECIFIED": 0,
		"PERMISSIONS_TYPE_ANON":        1,
		"PERMISSIONS_TYPE_READ":        2,
		"PERMISSIONS_TYPE_WRITE":       3,
		"PERMISSIONS_TYPE_OWN":         4,
		"PERMISSIONS_TYPE_ADMIN":       5,
	}
)

func (x PermissionsType) Enum() *PermissionsType {
	p := new(PermissionsType)
	*p = x
	return p
}

func (x PermissionsType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PermissionsType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_v1_enums_proto_enumTypes[2].Descriptor()
}

func (PermissionsType) Type() protoreflect.EnumType {
	return &file_common_v1_enums_proto_enumTypes[2]
}

func (x PermissionsType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PermissionsType.Descriptor instead.
func (PermissionsType) EnumDescriptor() ([]byte, []int) {
	return file_common_v1_enums_proto_rawDescGZIP(), []int{2}
}

// TokenKind is the type of token used to represent a currency.
type TokenKind int32

const (
	// Unspecified token kind.
	TokenKind_TOKEN_KIND_UNSPECIFIED TokenKind = 0
	// SONR native token kind.
	TokenKind_TOKEN_KIND_SONR_NATIVE TokenKind = 1
	// SONR staking token kind.
	TokenKind_TOKEN_KIND_SONR_STAKING TokenKind = 2
	// Ethereum native token kind.
	TokenKind_TOKEN_KIND_ETH_NATIVE TokenKind = 3
	// Ethereum staking token kind.
	TokenKind_TOKEN_KIND_ETH_STAKING TokenKind = 4
	// IBC native token kind.
	TokenKind_TOKEN_KIND_IBC_NATIVE TokenKind = 5
	// IBC staking token kind.
	TokenKind_TOKEN_KIND_IBC_STAKING TokenKind = 6
	// Bitcoin native token kind.
	TokenKind_TOKEN_KIND_BTC_NATIVE TokenKind = 7
	// USDC native token kind.
	TokenKind_TOKEN_KIND_USDC_NATIVE TokenKind = 8
)

// Enum value maps for TokenKind.
var (
	TokenKind_name = map[int32]string{
		0: "TOKEN_KIND_UNSPECIFIED",
		1: "TOKEN_KIND_SONR_NATIVE",
		2: "TOKEN_KIND_SONR_STAKING",
		3: "TOKEN_KIND_ETH_NATIVE",
		4: "TOKEN_KIND_ETH_STAKING",
		5: "TOKEN_KIND_IBC_NATIVE",
		6: "TOKEN_KIND_IBC_STAKING",
		7: "TOKEN_KIND_BTC_NATIVE",
		8: "TOKEN_KIND_USDC_NATIVE",
	}
	TokenKind_value = map[string]int32{
		"TOKEN_KIND_UNSPECIFIED":  0,
		"TOKEN_KIND_SONR_NATIVE":  1,
		"TOKEN_KIND_SONR_STAKING": 2,
		"TOKEN_KIND_ETH_NATIVE":   3,
		"TOKEN_KIND_ETH_STAKING":  4,
		"TOKEN_KIND_IBC_NATIVE":   5,
		"TOKEN_KIND_IBC_STAKING":  6,
		"TOKEN_KIND_BTC_NATIVE":   7,
		"TOKEN_KIND_USDC_NATIVE":  8,
	}
)

func (x TokenKind) Enum() *TokenKind {
	p := new(TokenKind)
	*p = x
	return p
}

func (x TokenKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenKind) Descriptor() protoreflect.EnumDescriptor {
	return file_common_v1_enums_proto_enumTypes[3].Descriptor()
}

func (TokenKind) Type() protoreflect.EnumType {
	return &file_common_v1_enums_proto_enumTypes[3]
}

func (x TokenKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenKind.Descriptor instead.
func (TokenKind) EnumDescriptor() ([]byte, []int) {
	return file_common_v1_enums_proto_rawDescGZIP(), []int{3}
}

var File_common_v1_enums_proto protoreflect.FileDescriptor

var file_common_v1_enums_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2a, 0x95, 0x02, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x19, 0x0a, 0x15, 0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f,
	0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x54, 0x4f, 0x4d, 0x10, 0x01, 0x12, 0x14,
	0x0a, 0x10, 0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x58, 0x45, 0x4c,
	0x41, 0x52, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x42, 0x49, 0x54, 0x43, 0x4f, 0x49, 0x4e, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x43,
	0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x54, 0x48, 0x45, 0x52, 0x45, 0x55,
	0x4d, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x45, 0x56, 0x4d, 0x4f, 0x53, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x49, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x43, 0x4f, 0x49, 0x4e, 0x10, 0x06,
	0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x55,
	0x4e, 0x4f, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4f, 0x53, 0x4d, 0x4f, 0x10, 0x08, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x49, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4f, 0x4c, 0x41, 0x4e, 0x41, 0x10, 0x09, 0x12, 0x12,
	0x0a, 0x0e, 0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4f, 0x4e, 0x52,
	0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x52, 0x47, 0x41, 0x5a, 0x45, 0x10, 0x0b, 0x2a, 0xe3, 0x01, 0x0a, 0x0e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1f, 0x0a,
	0x1b, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x4b, 0x49, 0x4e, 0x44,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x22,
	0x0a, 0x1e, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x4b, 0x49, 0x4e,
	0x44, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x45,
	0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52,
	0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x42,
	0x45, 0x52, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49,
	0x45, 0x52, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x41, 0x44,
	0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x49, 0x44, 0x45, 0x4e, 0x54,
	0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x4f, 0x41, 0x55, 0x54, 0x48,
	0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x10, 0x04, 0x12, 0x23, 0x0a, 0x1f, 0x49,
	0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x45,
	0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x57, 0x41, 0x4c, 0x4c, 0x45, 0x54, 0x10, 0x05,
	0x2a, 0xbb, 0x01, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x4f, 0x4e, 0x10,
	0x01, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16,
	0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x45, 0x52, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x57, 0x4e,
	0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e,
	0x53, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x05, 0x2a, 0x85,
	0x02, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x16,
	0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x4f, 0x4b, 0x45,
	0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x53, 0x4f, 0x4e, 0x52, 0x5f, 0x4e, 0x41, 0x54, 0x49,
	0x56, 0x45, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4b, 0x49,
	0x4e, 0x44, 0x5f, 0x53, 0x4f, 0x4e, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x4b, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f,
	0x45, 0x54, 0x48, 0x5f, 0x4e, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16,
	0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x45, 0x54, 0x48, 0x5f, 0x53,
	0x54, 0x41, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x4f, 0x4b, 0x45,
	0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x49, 0x42, 0x43, 0x5f, 0x4e, 0x41, 0x54, 0x49, 0x56,
	0x45, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4b, 0x49, 0x4e,
	0x44, 0x5f, 0x49, 0x42, 0x43, 0x5f, 0x53, 0x54, 0x41, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12,
	0x19, 0x0a, 0x15, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x42, 0x54,
	0x43, 0x5f, 0x4e, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x07, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x4f,
	0x4b, 0x45, 0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x53, 0x44, 0x43, 0x5f, 0x4e, 0x41,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x08, 0x42, 0x92, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x64, 0x61, 0x6f, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x73, 0x6f, 0x6e,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02,
	0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_common_v1_enums_proto_rawDescOnce sync.Once
	file_common_v1_enums_proto_rawDescData = file_common_v1_enums_proto_rawDesc
)

func file_common_v1_enums_proto_rawDescGZIP() []byte {
	file_common_v1_enums_proto_rawDescOnce.Do(func() {
		file_common_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_v1_enums_proto_rawDescData)
	})
	return file_common_v1_enums_proto_rawDescData
}

var file_common_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_common_v1_enums_proto_goTypes = []interface{}{
	(CoinType)(0),        // 0: common.v1.CoinType
	(IdentifierKind)(0),  // 1: common.v1.IdentifierKind
	(PermissionsType)(0), // 2: common.v1.PermissionsType
	(TokenKind)(0),       // 3: common.v1.TokenKind
}
var file_common_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_v1_enums_proto_init() }
func file_common_v1_enums_proto_init() {
	if File_common_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_v1_enums_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_v1_enums_proto_goTypes,
		DependencyIndexes: file_common_v1_enums_proto_depIdxs,
		EnumInfos:         file_common_v1_enums_proto_enumTypes,
	}.Build()
	File_common_v1_enums_proto = out.File
	file_common_v1_enums_proto_rawDesc = nil
	file_common_v1_enums_proto_goTypes = nil
	file_common_v1_enums_proto_depIdxs = nil
}
