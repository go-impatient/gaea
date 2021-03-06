// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: rpc/helloworld/v1/helloworld.proto

package helloworld_v1

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

type HelloworldEchoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FIXME 请求字段必须写注释
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloworldEchoReq) Reset() {
	*x = HelloworldEchoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_helloworld_v1_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloworldEchoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloworldEchoReq) ProtoMessage() {}

func (x *HelloworldEchoReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_helloworld_v1_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloworldEchoReq.ProtoReflect.Descriptor instead.
func (*HelloworldEchoReq) Descriptor() ([]byte, []int) {
	return file_rpc_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *HelloworldEchoReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type HelloworldEchoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FIXME 响应字段必须写注释
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloworldEchoResp) Reset() {
	*x = HelloworldEchoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_helloworld_v1_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloworldEchoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloworldEchoResp) ProtoMessage() {}

func (x *HelloworldEchoResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_helloworld_v1_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloworldEchoResp.ProtoReflect.Descriptor instead.
func (*HelloworldEchoResp) Descriptor() ([]byte, []int) {
	return file_rpc_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *HelloworldEchoResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_rpc_helloworld_v1_helloworld_proto protoreflect.FileDescriptor

var file_rpc_helloworld_v1_helloworld_proto_rawDesc = []byte{
	0x0a, 0x22, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x76, 0x31, 0x22, 0x25, 0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x26, 0x0a, 0x12, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x32, 0x59, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x12, 0x4b, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x20, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_helloworld_v1_helloworld_proto_rawDescOnce sync.Once
	file_rpc_helloworld_v1_helloworld_proto_rawDescData = file_rpc_helloworld_v1_helloworld_proto_rawDesc
)

func file_rpc_helloworld_v1_helloworld_proto_rawDescGZIP() []byte {
	file_rpc_helloworld_v1_helloworld_proto_rawDescOnce.Do(func() {
		file_rpc_helloworld_v1_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_helloworld_v1_helloworld_proto_rawDescData)
	})
	return file_rpc_helloworld_v1_helloworld_proto_rawDescData
}

var file_rpc_helloworld_v1_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_helloworld_v1_helloworld_proto_goTypes = []interface{}{
	(*HelloworldEchoReq)(nil),  // 0: helloworld.v1.HelloworldEchoReq
	(*HelloworldEchoResp)(nil), // 1: helloworld.v1.HelloworldEchoResp
}
var file_rpc_helloworld_v1_helloworld_proto_depIdxs = []int32{
	0, // 0: helloworld.v1.Helloworld.Echo:input_type -> helloworld.v1.HelloworldEchoReq
	1, // 1: helloworld.v1.Helloworld.Echo:output_type -> helloworld.v1.HelloworldEchoResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_helloworld_v1_helloworld_proto_init() }
func file_rpc_helloworld_v1_helloworld_proto_init() {
	if File_rpc_helloworld_v1_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_helloworld_v1_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloworldEchoReq); i {
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
		file_rpc_helloworld_v1_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloworldEchoResp); i {
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
			RawDescriptor: file_rpc_helloworld_v1_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_helloworld_v1_helloworld_proto_goTypes,
		DependencyIndexes: file_rpc_helloworld_v1_helloworld_proto_depIdxs,
		MessageInfos:      file_rpc_helloworld_v1_helloworld_proto_msgTypes,
	}.Build()
	File_rpc_helloworld_v1_helloworld_proto = out.File
	file_rpc_helloworld_v1_helloworld_proto_rawDesc = nil
	file_rpc_helloworld_v1_helloworld_proto_goTypes = nil
	file_rpc_helloworld_v1_helloworld_proto_depIdxs = nil
}
