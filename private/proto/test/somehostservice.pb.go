// This file defines simple interfaces used for testing. These interfaces are
// only intended to be used internally and by SPIRE. See /private/README.md.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: test/somehostservice.proto

package test

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

var File_test_somehostservice_proto protoreflect.FileDescriptor

var file_test_somehostservice_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x6f, 0x6d, 0x65, 0x68, 0x6f, 0x73, 0x74, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x4b, 0x0a, 0x0f, 0x53, 0x6f, 0x6d, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0f, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x11, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2d, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_test_somehostservice_proto_goTypes = []interface{}{
	(*EchoRequest)(nil),  // 0: test.EchoRequest
	(*EchoResponse)(nil), // 1: test.EchoResponse
}
var file_test_somehostservice_proto_depIdxs = []int32{
	0, // 0: test.SomeHostService.HostServiceEcho:input_type -> test.EchoRequest
	1, // 1: test.SomeHostService.HostServiceEcho:output_type -> test.EchoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_somehostservice_proto_init() }
func file_test_somehostservice_proto_init() {
	if File_test_somehostservice_proto != nil {
		return
	}
	file_test_echo_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_somehostservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_somehostservice_proto_goTypes,
		DependencyIndexes: file_test_somehostservice_proto_depIdxs,
	}.Build()
	File_test_somehostservice_proto = out.File
	file_test_somehostservice_proto_rawDesc = nil
	file_test_somehostservice_proto_goTypes = nil
	file_test_somehostservice_proto_depIdxs = nil
}
