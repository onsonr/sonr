// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: sonr/service/v1/types.proto

package servicev1

import (
	_ "cosmossdk.io/api/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Params defines the parameters of the module.
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_service_v1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_service_v1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_sonr_service_v1_types_proto_rawDescGZIP(), []int{0}
}

// GenesisState is the state that must be provided at genesis.
type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// params defines all the parameters of the module.
	Params *Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sonr_service_v1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	mi := &file_sonr_service_v1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_sonr_service_v1_types_proto_rawDescGZIP(), []int{1}
}

func (x *GenesisState) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_sonr_service_v1_types_proto protoreflect.FileDescriptor

var file_sonr_service_v1_types_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73,
	0x6f, 0x6e, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x11,
	0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x3a, 0x1a, 0x8a, 0xe7, 0xb0, 0x2a, 0x15, 0x73, 0x6f, 0x6e, 0x72, 0x68, 0x71, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x4a, 0x0a,
	0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x09, 0xc8, 0xde, 0x1f, 0x00, 0xa8, 0xe7, 0xb0, 0x2a,
	0x01, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0xc0, 0x01, 0x0a, 0x13, 0x63, 0x6f,
	0x6d, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72,
	0x68, 0x71, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x78, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x53, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x53, 0x6f, 0x6e, 0x72, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x53, 0x6f, 0x6e, 0x72, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x53, 0x6f, 0x6e,
	0x72, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x53, 0x6f, 0x6e, 0x72, 0x3a,
	0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sonr_service_v1_types_proto_rawDescOnce sync.Once
	file_sonr_service_v1_types_proto_rawDescData = file_sonr_service_v1_types_proto_rawDesc
)

func file_sonr_service_v1_types_proto_rawDescGZIP() []byte {
	file_sonr_service_v1_types_proto_rawDescOnce.Do(func() {
		file_sonr_service_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_sonr_service_v1_types_proto_rawDescData)
	})
	return file_sonr_service_v1_types_proto_rawDescData
}

var file_sonr_service_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sonr_service_v1_types_proto_goTypes = []interface{}{
	(*Params)(nil),       // 0: sonr.service.v1.Params
	(*GenesisState)(nil), // 1: sonr.service.v1.GenesisState
}
var file_sonr_service_v1_types_proto_depIdxs = []int32{
	0, // 0: sonr.service.v1.GenesisState.params:type_name -> sonr.service.v1.Params
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sonr_service_v1_types_proto_init() }
func file_sonr_service_v1_types_proto_init() {
	if File_sonr_service_v1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sonr_service_v1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
		file_sonr_service_v1_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
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
			RawDescriptor: file_sonr_service_v1_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sonr_service_v1_types_proto_goTypes,
		DependencyIndexes: file_sonr_service_v1_types_proto_depIdxs,
		MessageInfos:      file_sonr_service_v1_types_proto_msgTypes,
	}.Build()
	File_sonr_service_v1_types_proto = out.File
	file_sonr_service_v1_types_proto_rawDesc = nil
	file_sonr_service_v1_types_proto_goTypes = nil
	file_sonr_service_v1_types_proto_depIdxs = nil
}
