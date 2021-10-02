/// This file contains service for the Node RPC Server

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/node/highway.proto

package node

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_node_highway_proto protoreflect.FileDescriptor

var file_proto_node_highway_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x68, 0x69, 0x67,
	0x68, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x6f, 0x6e, 0x72,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xdc, 0x01, 0x0a, 0x0e,
	0x48, 0x69, 0x67, 0x68, 0x77, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48,
	0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x6f,
	0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x16, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x1a, 0x2e, 0x73, 0x6f, 0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x6f,
	0x6e, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6e, 0x72, 0x2d, 0x69, 0x6f,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_node_highway_proto_goTypes = []interface{}{
	(*AuthorizeRequest)(nil),  // 0: sonr.node.AuthorizeRequest
	(*LinkRequest)(nil),       // 1: sonr.node.LinkRequest
	(*RegisterRequest)(nil),   // 2: sonr.node.RegisterRequest
	(*AuthorizeResponse)(nil), // 3: sonr.node.AuthorizeResponse
	(*LinkResponse)(nil),      // 4: sonr.node.LinkResponse
	(*RegisterResponse)(nil),  // 5: sonr.node.RegisterResponse
}
var file_proto_node_highway_proto_depIdxs = []int32{
	0, // 0: sonr.node.HighwayService.Authorize:input_type -> sonr.node.AuthorizeRequest
	1, // 1: sonr.node.HighwayService.Link:input_type -> sonr.node.LinkRequest
	2, // 2: sonr.node.HighwayService.Register:input_type -> sonr.node.RegisterRequest
	3, // 3: sonr.node.HighwayService.Authorize:output_type -> sonr.node.AuthorizeResponse
	4, // 4: sonr.node.HighwayService.Link:output_type -> sonr.node.LinkResponse
	5, // 5: sonr.node.HighwayService.Register:output_type -> sonr.node.RegisterResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_node_highway_proto_init() }
func file_proto_node_highway_proto_init() {
	if File_proto_node_highway_proto != nil {
		return
	}
	file_proto_node_api_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_node_highway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_node_highway_proto_goTypes,
		DependencyIndexes: file_proto_node_highway_proto_depIdxs,
	}.Build()
	File_proto_node_highway_proto = out.File
	file_proto_node_highway_proto_rawDesc = nil
	file_proto_node_highway_proto_goTypes = nil
	file_proto_node_highway_proto_depIdxs = nil
}
