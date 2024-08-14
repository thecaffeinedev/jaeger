// Copyright (c) 2019 The Jaeger Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.14.0
// source: storage_test.proto

package storageprototest

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

type GetTraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId []byte `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
}

func (x *GetTraceRequest) Reset() {
	*x = GetTraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTraceRequest) ProtoMessage() {}

func (x *GetTraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_storage_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTraceRequest.ProtoReflect.Descriptor instead.
func (*GetTraceRequest) Descriptor() ([]byte, []int) {
	return file_storage_test_proto_rawDescGZIP(), []int{0}
}

func (x *GetTraceRequest) GetTraceId() []byte {
	if x != nil {
		return x.TraceId
	}
	return nil
}

var File_storage_test_proto protoreflect.FileDescriptor

var file_storage_test_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x74, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x49, 0x64, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_storage_test_proto_rawDescOnce sync.Once
	file_storage_test_proto_rawDescData = file_storage_test_proto_rawDesc
)

func file_storage_test_proto_rawDescGZIP() []byte {
	file_storage_test_proto_rawDescOnce.Do(func() {
		file_storage_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_test_proto_rawDescData)
	})
	return file_storage_test_proto_rawDescData
}

var file_storage_test_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storage_test_proto_goTypes = []interface{}{
	(*GetTraceRequest)(nil), // 0: storageprototest.GetTraceRequest
}
var file_storage_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_storage_test_proto_init() }
func file_storage_test_proto_init() {
	if File_storage_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTraceRequest); i {
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
			RawDescriptor: file_storage_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_test_proto_goTypes,
		DependencyIndexes: file_storage_test_proto_depIdxs,
		MessageInfos:      file_storage_test_proto_msgTypes,
	}.Build()
	File_storage_test_proto = out.File
	file_storage_test_proto_rawDesc = nil
	file_storage_test_proto_goTypes = nil
	file_storage_test_proto_depIdxs = nil
}
