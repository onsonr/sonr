// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: sonr/service/v1/query.proto

package servicev1

import (
	reflect "reflect"
	sync "sync"

	_ "cosmossdk.io/api/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	v1 "github.com/sonrhq/sonr/api/sonr/service/module/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueryParamsRequest) Reset() {
	*x = QueryParamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_service_v1_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParamsRequest) ProtoMessage() {}

func (x *QueryParamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_service_v1_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryParamsRequest.ProtoReflect.Descriptor instead.
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return file_sonr_service_v1_query_proto_rawDescGZIP(), []int{0}
}

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// params defines the parameters of the module.
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *QueryParamsResponse) Reset() {
	*x = QueryParamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_service_v1_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParamsResponse) ProtoMessage() {}

func (x *QueryParamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_service_v1_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryParamsResponse.ProtoReflect.Descriptor instead.
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return file_sonr_service_v1_query_proto_rawDescGZIP(), []int{1}
}

func (x *QueryParamsResponse) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

// QueryCredentialsRequest is the request type for the Query/Credentials RPC method.
type QueryCredentialsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address is the address of the credentials to query.
	Origin string `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
	// identifier is the owner of the credentials to query.
	Identifier string `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// type is the type of the credentials to query.
	ParamsType ParamsType `protobuf:"varint,3,opt,name=params_type,json=paramsType,proto3,enum=sonr.service.v1.ParamsType" json:"params_type,omitempty"`
}

func (x *QueryCredentialsRequest) Reset() {
	*x = QueryCredentialsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_service_v1_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryCredentialsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryCredentialsRequest) ProtoMessage() {}

func (x *QueryCredentialsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_service_v1_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryCredentialsRequest.ProtoReflect.Descriptor instead.
func (*QueryCredentialsRequest) Descriptor() ([]byte, []int) {
	return file_sonr_service_v1_query_proto_rawDescGZIP(), []int{2}
}

func (x *QueryCredentialsRequest) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *QueryCredentialsRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *QueryCredentialsRequest) GetParamsType() ParamsType {
	if x != nil {
		return x.ParamsType
	}
	return ParamsType_PARAMS_TYPE_UNSPECIFIED
}

// QueryCredentialsResponse is the response type for the Query/Credentials RPC method.
type QueryCredentialsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// attestation_options is the attestation options of the credentials.
	AttestationOptions string `protobuf:"bytes,1,opt,name=attestation_options,json=attestationOptions,proto3" json:"attestation_options,omitempty"`
	// assertion_options is the assertion options of the credentials.
	AssertionOptions string `protobuf:"bytes,2,opt,name=assertion_options,json=assertionOptions,proto3" json:"assertion_options,omitempty"`
	// origin is the service record of the origin of the credentials.
	Origin string `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (x *QueryCredentialsResponse) Reset() {
	*x = QueryCredentialsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_service_v1_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryCredentialsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryCredentialsResponse) ProtoMessage() {}

func (x *QueryCredentialsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_service_v1_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryCredentialsResponse.ProtoReflect.Descriptor instead.
func (*QueryCredentialsResponse) Descriptor() ([]byte, []int) {
	return file_sonr_service_v1_query_proto_rawDescGZIP(), []int{3}
}

func (x *QueryCredentialsResponse) GetAttestationOptions() string {
	if x != nil {
		return x.AttestationOptions
	}
	return ""
}

func (x *QueryCredentialsResponse) GetAssertionOptions() string {
	if x != nil {
		return x.AssertionOptions
	}
	return ""
}

func (x *QueryCredentialsResponse) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

type QueryServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Origin string `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (x *QueryServiceRequest) Reset() {
	*x = QueryServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_service_v1_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryServiceRequest) ProtoMessage() {}

func (x *QueryServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_service_v1_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryServiceRequest.ProtoReflect.Descriptor instead.
func (*QueryServiceRequest) Descriptor() ([]byte, []int) {
	return file_sonr_service_v1_query_proto_rawDescGZIP(), []int{4}
}

func (x *QueryServiceRequest) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

type QueryServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service *v1.ServiceRecord `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *QueryServiceResponse) Reset() {
	*x = QueryServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_service_v1_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryServiceResponse) ProtoMessage() {}

func (x *QueryServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_service_v1_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryServiceResponse.ProtoReflect.Descriptor instead.
func (*QueryServiceResponse) Descriptor() ([]byte, []int) {
	return file_sonr_service_v1_query_proto_rawDescGZIP(), []int{5}
}

func (x *QueryServiceResponse) GetService() *v1.ServiceRecord {
	if x != nil {
		return x.Service
	}
	return nil
}

var File_sonr_service_v1_query_proto protoreflect.FileDescriptor

var file_sonr_service_v1_query_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73,
	0x6f, 0x6e, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x22,
	0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61,
	0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x51, 0x0a, 0x13,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x09, 0xc8, 0xde,
	0x1f, 0x00, 0xa8, 0xe7, 0xb0, 0x2a, 0x01, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22,
	0x8f, 0x01, 0x0a, 0x17, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x90, 0x01, 0x0a, 0x18, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x13, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x2b, 0x0a, 0x11, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x73, 0x73, 0x65,
	0x72, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x22, 0x2d, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x22, 0x57, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73,
	0x6f, 0x6e, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0xa8, 0x03, 0x0a,
	0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x76, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x23, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0xa0,
	0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x28,
	0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x12, 0x34, 0x2f, 0x73, 0x6f,
	0x6e, 0x72, 0x68, 0x71, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x7d, 0x2f, 0x7b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x7d, 0x12, 0x83, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x25, 0x12, 0x23, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x7b,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x7d, 0x42, 0xc0, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x68, 0x71,
	0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x78, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x53, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x53, 0x6f, 0x6e, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x53, 0x6f, 0x6e, 0x72, 0x5c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x53, 0x6f, 0x6e, 0x72, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x53, 0x6f, 0x6e, 0x72, 0x3a, 0x3a, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sonr_service_v1_query_proto_rawDescOnce sync.Once
	file_sonr_service_v1_query_proto_rawDescData = file_sonr_service_v1_query_proto_rawDesc
)

func file_sonr_service_v1_query_proto_rawDescGZIP() []byte {
	file_sonr_service_v1_query_proto_rawDescOnce.Do(func() {
		file_sonr_service_v1_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_sonr_service_v1_query_proto_rawDescData)
	})
	return file_sonr_service_v1_query_proto_rawDescData
}

var file_sonr_service_v1_query_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sonr_service_v1_query_proto_goTypes = []interface{}{
	(*QueryParamsRequest)(nil),       // 0: sonr.service.v1.QueryParamsRequest
	(*QueryParamsResponse)(nil),      // 1: sonr.service.v1.QueryParamsResponse
	(*QueryCredentialsRequest)(nil),  // 2: sonr.service.v1.QueryCredentialsRequest
	(*QueryCredentialsResponse)(nil), // 3: sonr.service.v1.QueryCredentialsResponse
	(*QueryServiceRequest)(nil),      // 4: sonr.service.v1.QueryServiceRequest
	(*QueryServiceResponse)(nil),     // 5: sonr.service.v1.QueryServiceResponse
	(*Params)(nil),                   // 6: sonr.service.v1.Params
	(ParamsType)(0),                  // 7: sonr.service.v1.ParamsType
	(*v1.ServiceRecord)(nil),         // 8: sonr.service.module.v1.ServiceRecord
}
var file_sonr_service_v1_query_proto_depIdxs = []int32{
	6, // 0: sonr.service.v1.QueryParamsResponse.params:type_name -> sonr.service.v1.Params
	7, // 1: sonr.service.v1.QueryCredentialsRequest.params_type:type_name -> sonr.service.v1.ParamsType
	8, // 2: sonr.service.v1.QueryServiceResponse.service:type_name -> sonr.service.module.v1.ServiceRecord
	0, // 3: sonr.service.v1.Query.Params:input_type -> sonr.service.v1.QueryParamsRequest
	2, // 4: sonr.service.v1.Query.Credentials:input_type -> sonr.service.v1.QueryCredentialsRequest
	4, // 5: sonr.service.v1.Query.Service:input_type -> sonr.service.v1.QueryServiceRequest
	1, // 6: sonr.service.v1.Query.Params:output_type -> sonr.service.v1.QueryParamsResponse
	3, // 7: sonr.service.v1.Query.Credentials:output_type -> sonr.service.v1.QueryCredentialsResponse
	5, // 8: sonr.service.v1.Query.Service:output_type -> sonr.service.v1.QueryServiceResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sonr_service_v1_query_proto_init() }
func file_sonr_service_v1_query_proto_init() {
	if File_sonr_service_v1_query_proto != nil {
		return
	}
	file_sonr_service_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sonr_service_v1_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParamsRequest); i {
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
		file_sonr_service_v1_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParamsResponse); i {
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
		file_sonr_service_v1_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryCredentialsRequest); i {
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
		file_sonr_service_v1_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryCredentialsResponse); i {
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
		file_sonr_service_v1_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryServiceRequest); i {
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
		file_sonr_service_v1_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryServiceResponse); i {
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
			RawDescriptor: file_sonr_service_v1_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sonr_service_v1_query_proto_goTypes,
		DependencyIndexes: file_sonr_service_v1_query_proto_depIdxs,
		MessageInfos:      file_sonr_service_v1_query_proto_msgTypes,
	}.Build()
	File_sonr_service_v1_query_proto = out.File
	file_sonr_service_v1_query_proto_rawDesc = nil
	file_sonr_service_v1_query_proto_goTypes = nil
	file_sonr_service_v1_query_proto_depIdxs = nil
}
