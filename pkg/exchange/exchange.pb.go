// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/protocols/exchange.proto

package exchange

import (
	common "github.com/sonr-io/core/internal/common"
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

// CreateSNameRequest is Message for Signing Request (Hmac Sha256)
type CreateSNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SName     string   `protobuf:"bytes,1,opt,name=sName,proto3" json:"sName,omitempty"`         // SName combined with Device ID and Hashed
	Mnemonic  string   `protobuf:"bytes,2,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`   // Mnemonic Hashed with private key for fingerprint
	DeviceIds []string `protobuf:"bytes,3,rep,name=deviceIds,proto3" json:"deviceIds,omitempty"` // Device IDs for SName
}

func (x *CreateSNameRequest) Reset() {
	*x = CreateSNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_exchange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSNameRequest) ProtoMessage() {}

func (x *CreateSNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_exchange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSNameRequest.ProtoReflect.Descriptor instead.
func (*CreateSNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSNameRequest) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

func (x *CreateSNameRequest) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

func (x *CreateSNameRequest) GetDeviceIds() []string {
	if x != nil {
		return x.DeviceIds
	}
	return nil
}

// CreateSNameResponse is Message for Signing Response (Hmac Sha256)
type CreateSNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool                          `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // If Values were Signed
	Domains []*CreateSNameResponse_Domain `protobuf:"bytes,2,rep,name=domains,proto3" json:"domains,omitempty"`  // Signed Domain TXT Records
	// Resulting Signed Values
	PublicKey     string `protobuf:"bytes,4,opt,name=publicKey,proto3" json:"publicKey,omitempty"`         // Base64 Encoded Public Key
	GivenSName    string `protobuf:"bytes,5,opt,name=givenSName,proto3" json:"givenSName,omitempty"`       // Provided SName
	GivenMnemonic string `protobuf:"bytes,6,opt,name=givenMnemonic,proto3" json:"givenMnemonic,omitempty"` // Provided Mnemonic
}

func (x *CreateSNameResponse) Reset() {
	*x = CreateSNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_exchange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSNameResponse) ProtoMessage() {}

func (x *CreateSNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_exchange_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSNameResponse.ProtoReflect.Descriptor instead.
func (*CreateSNameResponse) Descriptor() ([]byte, []int) {
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSNameResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateSNameResponse) GetDomains() []*CreateSNameResponse_Domain {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *CreateSNameResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *CreateSNameResponse) GetGivenSName() string {
	if x != nil {
		return x.GivenSName
	}
	return ""
}

func (x *CreateSNameResponse) GetGivenMnemonic() string {
	if x != nil {
		return x.GivenMnemonic
	}
	return ""
}

// LookupSNameRequest is Message for Verifying Request (Hmac Sha256)
type LookupSNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SName string `protobuf:"bytes,1,opt,name=sName,proto3" json:"sName,omitempty"` // SName combined with Device ID and Hashed
}

func (x *LookupSNameRequest) Reset() {
	*x = LookupSNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_exchange_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupSNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupSNameRequest) ProtoMessage() {}

func (x *LookupSNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_exchange_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupSNameRequest.ProtoReflect.Descriptor instead.
func (*LookupSNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{2}
}

func (x *LookupSNameRequest) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

// LookupSNameResponse is Message for Verifying Response (Hmac Sha256)
type LookupSNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success   bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`    // If Values were Verified
	PublicKey string `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"` // Base64 Encoded Public Key
}

func (x *LookupSNameResponse) Reset() {
	*x = LookupSNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_exchange_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupSNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupSNameResponse) ProtoMessage() {}

func (x *LookupSNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_exchange_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupSNameResponse.ProtoReflect.Descriptor instead.
func (*LookupSNameResponse) Descriptor() ([]byte, []int) {
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{3}
}

func (x *LookupSNameResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LookupSNameResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

// QueryExchangeRequest is Message for searching for Peer
type QueryExchangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SName  string `protobuf:"bytes,1,opt,name=sName,proto3" json:"sName,omitempty"`   // SName combined with Device ID and Hashed
	PeerId string `protobuf:"bytes,2,opt,name=peerId,proto3" json:"peerId,omitempty"` // Peer ID
}

func (x *QueryExchangeRequest) Reset() {
	*x = QueryExchangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_exchange_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryExchangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryExchangeRequest) ProtoMessage() {}

func (x *QueryExchangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_exchange_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryExchangeRequest.ProtoReflect.Descriptor instead.
func (*QueryExchangeRequest) Descriptor() ([]byte, []int) {
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{4}
}

func (x *QueryExchangeRequest) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

func (x *QueryExchangeRequest) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

// UpdateExchangeRequest is Message for updating Peer Data in Exchange
type UpdateExchangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SName  string       `protobuf:"bytes,1,opt,name=sName,proto3" json:"sName,omitempty"`   // SName combined with Device ID and Hashed
	PeerId string       `protobuf:"bytes,2,opt,name=peerId,proto3" json:"peerId,omitempty"` // Peer ID
	Peer   *common.Peer `protobuf:"bytes,3,opt,name=peer,proto3" json:"peer,omitempty"`     // Peer Data
}

func (x *UpdateExchangeRequest) Reset() {
	*x = UpdateExchangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_exchange_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateExchangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExchangeRequest) ProtoMessage() {}

func (x *UpdateExchangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_exchange_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExchangeRequest.ProtoReflect.Descriptor instead.
func (*UpdateExchangeRequest) Descriptor() ([]byte, []int) {
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateExchangeRequest) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

func (x *UpdateExchangeRequest) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *UpdateExchangeRequest) GetPeer() *common.Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

// UpdateExchangeResponse is response for UpdateExchangeRequest
type UpdateExchangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // If Request was Successful
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`      // Error Message if Request was not successful
}

func (x *UpdateExchangeResponse) Reset() {
	*x = UpdateExchangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_exchange_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateExchangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExchangeResponse) ProtoMessage() {}

func (x *UpdateExchangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_exchange_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExchangeResponse.ProtoReflect.Descriptor instead.
func (*UpdateExchangeResponse) Descriptor() ([]byte, []int) {
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateExchangeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdateExchangeResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// VerifySNameRequest is Message for Verifying Request (Hmac Sha256)
type VerifySNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SName    string `protobuf:"bytes,1,opt,name=sName,proto3" json:"sName,omitempty"`       // SName combined with Device ID and Hashed
	Mnemonic string `protobuf:"bytes,2,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"` // Mnemonic Hashed with public key for fingerprint
	DeviceId string `protobuf:"bytes,3,opt,name=deviceId,proto3" json:"deviceId,omitempty"` // Device ID
}

func (x *VerifySNameRequest) Reset() {
	*x = VerifySNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_exchange_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifySNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifySNameRequest) ProtoMessage() {}

func (x *VerifySNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_exchange_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifySNameRequest.ProtoReflect.Descriptor instead.
func (*VerifySNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{7}
}

func (x *VerifySNameRequest) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

func (x *VerifySNameRequest) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

func (x *VerifySNameRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

// VerifySNameResponse is Message for Verifying Response (Hmac Sha256)
type VerifySNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // If Values were Verified
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`      // Error Message
}

func (x *VerifySNameResponse) Reset() {
	*x = VerifySNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_exchange_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifySNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifySNameResponse) ProtoMessage() {}

func (x *VerifySNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_exchange_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifySNameResponse.ProtoReflect.Descriptor instead.
func (*VerifySNameResponse) Descriptor() ([]byte, []int) {
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{8}
}

func (x *VerifySNameResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *VerifySNameResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// Domain TXT Record for single SName Entry
type CreateSNameResponse_Domain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain            string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`                       // Domain Name
	Value             string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`                         // Domain Value
	SignedPrefix      string `protobuf:"bytes,3,opt,name=signedPrefix,proto3" json:"signedPrefix,omitempty"`           // Message for List of Bytes
	SignedFingerprint string `protobuf:"bytes,4,opt,name=signedFingerprint,proto3" json:"signedFingerprint,omitempty"` // Fingerprint Value
}

func (x *CreateSNameResponse_Domain) Reset() {
	*x = CreateSNameResponse_Domain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocols_exchange_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSNameResponse_Domain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSNameResponse_Domain) ProtoMessage() {}

func (x *CreateSNameResponse_Domain) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocols_exchange_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSNameResponse_Domain.ProtoReflect.Descriptor instead.
func (*CreateSNameResponse_Domain) Descriptor() ([]byte, []int) {
	return file_proto_protocols_exchange_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CreateSNameResponse_Domain) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreateSNameResponse_Domain) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *CreateSNameResponse_Domain) GetSignedPrefix() string {
	if x != nil {
		return x.SignedPrefix
	}
	return ""
}

func (x *CreateSNameResponse_Domain) GetSignedFingerprint() string {
	if x != nil {
		return x.SignedFingerprint
	}
	return ""
}

var File_proto_protocols_exchange_proto protoreflect.FileDescriptor

var file_proto_protocols_exchange_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x73, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x64, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x73, 0x22, 0xed, 0x02, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x4d, 0x0a, 0x07, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x73, 0x6f,
	0x6e, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2e, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x52, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x69, 0x76, 0x65, 0x6e,
	0x53, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x69, 0x76,
	0x65, 0x6e, 0x53, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x67, 0x69, 0x76, 0x65, 0x6e,
	0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x67, 0x69, 0x76, 0x65, 0x6e, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x1a, 0x88, 0x01,
	0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x64, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x46, 0x69, 0x6e,
	0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x22, 0x2a, 0x0a, 0x12, 0x4c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x53, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x13, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x53, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x22, 0x44, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x22, 0x6a, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52,
	0x04, 0x70, 0x65, 0x65, 0x72, 0x22, 0x48, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x62, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_protocols_exchange_proto_rawDescOnce sync.Once
	file_proto_protocols_exchange_proto_rawDescData = file_proto_protocols_exchange_proto_rawDesc
)

func file_proto_protocols_exchange_proto_rawDescGZIP() []byte {
	file_proto_protocols_exchange_proto_rawDescOnce.Do(func() {
		file_proto_protocols_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_protocols_exchange_proto_rawDescData)
	})
	return file_proto_protocols_exchange_proto_rawDescData
}

var file_proto_protocols_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_protocols_exchange_proto_goTypes = []interface{}{
	(*CreateSNameRequest)(nil),         // 0: sonr.protocols.exchange.CreateSNameRequest
	(*CreateSNameResponse)(nil),        // 1: sonr.protocols.exchange.CreateSNameResponse
	(*LookupSNameRequest)(nil),         // 2: sonr.protocols.exchange.LookupSNameRequest
	(*LookupSNameResponse)(nil),        // 3: sonr.protocols.exchange.LookupSNameResponse
	(*QueryExchangeRequest)(nil),       // 4: sonr.protocols.exchange.QueryExchangeRequest
	(*UpdateExchangeRequest)(nil),      // 5: sonr.protocols.exchange.UpdateExchangeRequest
	(*UpdateExchangeResponse)(nil),     // 6: sonr.protocols.exchange.UpdateExchangeResponse
	(*VerifySNameRequest)(nil),         // 7: sonr.protocols.exchange.VerifySNameRequest
	(*VerifySNameResponse)(nil),        // 8: sonr.protocols.exchange.VerifySNameResponse
	(*CreateSNameResponse_Domain)(nil), // 9: sonr.protocols.exchange.CreateSNameResponse.Domain
	(*common.Peer)(nil),                // 10: sonr.core.Peer
}
var file_proto_protocols_exchange_proto_depIdxs = []int32{
	9,  // 0: sonr.protocols.exchange.CreateSNameResponse.domains:type_name -> sonr.protocols.exchange.CreateSNameResponse.Domain
	10, // 1: sonr.protocols.exchange.UpdateExchangeRequest.peer:type_name -> sonr.core.Peer
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_protocols_exchange_proto_init() }
func file_proto_protocols_exchange_proto_init() {
	if File_proto_protocols_exchange_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_protocols_exchange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSNameRequest); i {
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
		file_proto_protocols_exchange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSNameResponse); i {
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
		file_proto_protocols_exchange_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupSNameRequest); i {
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
		file_proto_protocols_exchange_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupSNameResponse); i {
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
		file_proto_protocols_exchange_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryExchangeRequest); i {
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
		file_proto_protocols_exchange_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateExchangeRequest); i {
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
		file_proto_protocols_exchange_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateExchangeResponse); i {
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
		file_proto_protocols_exchange_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifySNameRequest); i {
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
		file_proto_protocols_exchange_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifySNameResponse); i {
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
		file_proto_protocols_exchange_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSNameResponse_Domain); i {
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
			RawDescriptor: file_proto_protocols_exchange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_protocols_exchange_proto_goTypes,
		DependencyIndexes: file_proto_protocols_exchange_proto_depIdxs,
		MessageInfos:      file_proto_protocols_exchange_proto_msgTypes,
	}.Build()
	File_proto_protocols_exchange_proto = out.File
	file_proto_protocols_exchange_proto_rawDesc = nil
	file_proto_protocols_exchange_proto_goTypes = nil
	file_proto_protocols_exchange_proto_depIdxs = nil
}
