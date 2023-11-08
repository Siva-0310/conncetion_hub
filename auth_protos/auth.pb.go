// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: auth_protos/auth.proto

package auth_protos

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

type JwtToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *JwtToken) Reset() {
	*x = JwtToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_protos_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JwtToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtToken) ProtoMessage() {}

func (x *JwtToken) ProtoReflect() protoreflect.Message {
	mi := &file_auth_protos_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtToken.ProtoReflect.Descriptor instead.
func (*JwtToken) Descriptor() ([]byte, []int) {
	return file_auth_protos_auth_proto_rawDescGZIP(), []int{0}
}

func (x *JwtToken) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Exists bool  `protobuf:"varint,2,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_protos_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_auth_protos_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_auth_protos_auth_proto_rawDescGZIP(), []int{1}
}

func (x *UserId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserId) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

var File_auth_protos_auth_proto protoreflect.FileDescriptor

var file_auth_protos_auth_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x1c, 0x0a, 0x08, 0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6a, 0x77, 0x74, 0x22, 0x30, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x32, 0x41, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x39, 0x0a,
	0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x1a, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x75, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_protos_auth_proto_rawDescOnce sync.Once
	file_auth_protos_auth_proto_rawDescData = file_auth_protos_auth_proto_rawDesc
)

func file_auth_protos_auth_proto_rawDescGZIP() []byte {
	file_auth_protos_auth_proto_rawDescOnce.Do(func() {
		file_auth_protos_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_protos_auth_proto_rawDescData)
	})
	return file_auth_protos_auth_proto_rawDescData
}

var file_auth_protos_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_protos_auth_proto_goTypes = []interface{}{
	(*JwtToken)(nil), // 0: auth_protos.JwtToken
	(*UserId)(nil),   // 1: auth_protos.UserId
}
var file_auth_protos_auth_proto_depIdxs = []int32{
	0, // 0: auth_protos.Auth.CheckUser:input_type -> auth_protos.JwtToken
	1, // 1: auth_protos.Auth.CheckUser:output_type -> auth_protos.UserId
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_protos_auth_proto_init() }
func file_auth_protos_auth_proto_init() {
	if File_auth_protos_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_protos_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JwtToken); i {
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
		file_auth_protos_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserId); i {
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
			RawDescriptor: file_auth_protos_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_protos_auth_proto_goTypes,
		DependencyIndexes: file_auth_protos_auth_proto_depIdxs,
		MessageInfos:      file_auth_protos_auth_proto_msgTypes,
	}.Build()
	File_auth_protos_auth_proto = out.File
	file_auth_protos_auth_proto_rawDesc = nil
	file_auth_protos_auth_proto_goTypes = nil
	file_auth_protos_auth_proto_depIdxs = nil
}