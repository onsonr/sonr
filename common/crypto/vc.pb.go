// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: sonrhq/common/crypto/vc.proto

package crypto

import (
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

// PubKey represents a public key in bytes format.
type PubKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	KeyType string `protobuf:"bytes,2,opt,name=key_type,json=keyType,proto3" json:"key_type,omitempty"`
}

func (x *PubKey) Reset() {
	*x = PubKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonrhq_common_crypto_vc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubKey) ProtoMessage() {}

func (x *PubKey) ProtoReflect() protoreflect.Message {
	mi := &file_sonrhq_common_crypto_vc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubKey.ProtoReflect.Descriptor instead.
func (*PubKey) Descriptor() ([]byte, []int) {
	return file_sonrhq_common_crypto_vc_proto_rawDescGZIP(), []int{0}
}

func (x *PubKey) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *PubKey) GetKeyType() string {
	if x != nil {
		return x.KeyType
	}
	return ""
}

var File_sonrhq_common_crypto_vc_proto protoreflect.FileDescriptor

var file_sonrhq_common_crypto_vc_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x73, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x42, 0x26, 0x5a, 0x24,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x68,
	0x71, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sonrhq_common_crypto_vc_proto_rawDescOnce sync.Once
	file_sonrhq_common_crypto_vc_proto_rawDescData = file_sonrhq_common_crypto_vc_proto_rawDesc
)

func file_sonrhq_common_crypto_vc_proto_rawDescGZIP() []byte {
	file_sonrhq_common_crypto_vc_proto_rawDescOnce.Do(func() {
		file_sonrhq_common_crypto_vc_proto_rawDescData = protoimpl.X.CompressGZIP(file_sonrhq_common_crypto_vc_proto_rawDescData)
	})
	return file_sonrhq_common_crypto_vc_proto_rawDescData
}

var file_sonrhq_common_crypto_vc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sonrhq_common_crypto_vc_proto_goTypes = []interface{}{
	(*PubKey)(nil), // 0: sonrhq.common.crypto.PubKey
}
var file_sonrhq_common_crypto_vc_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sonrhq_common_crypto_vc_proto_init() }
func file_sonrhq_common_crypto_vc_proto_init() {
	if File_sonrhq_common_crypto_vc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sonrhq_common_crypto_vc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubKey); i {
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
			RawDescriptor: file_sonrhq_common_crypto_vc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sonrhq_common_crypto_vc_proto_goTypes,
		DependencyIndexes: file_sonrhq_common_crypto_vc_proto_depIdxs,
		MessageInfos:      file_sonrhq_common_crypto_vc_proto_msgTypes,
	}.Build()
	File_sonrhq_common_crypto_vc_proto = out.File
	file_sonrhq_common_crypto_vc_proto_rawDesc = nil
	file_sonrhq_common_crypto_vc_proto_goTypes = nil
	file_sonrhq_common_crypto_vc_proto_depIdxs = nil
}
