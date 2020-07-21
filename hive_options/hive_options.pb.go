// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: hive_options/hive_options.proto

package hive_options

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type HiveMessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A schema is generated for top-level messages in each file,
	// as long as table_name is not blank.
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
}

func (x *HiveMessageOptions) Reset() {
	*x = HiveMessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hive_options_hive_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HiveMessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HiveMessageOptions) ProtoMessage() {}

func (x *HiveMessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_hive_options_hive_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HiveMessageOptions.ProtoReflect.Descriptor instead.
func (*HiveMessageOptions) Descriptor() ([]byte, []int) {
	return file_hive_options_hive_options_proto_rawDescGZIP(), []int{0}
}

func (x *HiveMessageOptions) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

type HiveFieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeOverride string `protobuf:"bytes,1,opt,name=type_override,json=typeOverride,proto3" json:"type_override,omitempty"`
}

func (x *HiveFieldOptions) Reset() {
	*x = HiveFieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hive_options_hive_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HiveFieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HiveFieldOptions) ProtoMessage() {}

func (x *HiveFieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_hive_options_hive_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HiveFieldOptions.ProtoReflect.Descriptor instead.
func (*HiveFieldOptions) Descriptor() ([]byte, []int) {
	return file_hive_options_hive_options_proto_rawDescGZIP(), []int{1}
}

func (x *HiveFieldOptions) GetTypeOverride() string {
	if x != nil {
		return x.TypeOverride
	}
	return ""
}

var file_hive_options_hive_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*HiveMessageOptions)(nil),
		Field:         1021,
		Name:          "hive_options.hive_message_opts",
		Tag:           "bytes,1021,opt,name=hive_message_opts",
		Filename:      "hive_options/hive_options.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*HiveFieldOptions)(nil),
		Field:         1021,
		Name:          "hive_options.hive_field_opts",
		Tag:           "bytes,1021,opt,name=hive_field_opts",
		Filename:      "hive_options/hive_options.proto",
	},
}

// Extension fields to descriptor.MessageOptions.
var (
	// optional hive_options.HiveMessageOptions hive_message_opts = 1021;
	E_HiveMessageOpts = &file_hive_options_hive_options_proto_extTypes[0]
)

// Extension fields to descriptor.FieldOptions.
var (
	// optional hive_options.HiveFieldOptions hive_field_opts = 1021;
	E_HiveFieldOpts = &file_hive_options_hive_options_proto_extTypes[1]
)

var File_hive_options_hive_options_proto protoreflect.FileDescriptor

var file_hive_options_hive_options_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68,
	0x69, 0x76, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x33, 0x0a, 0x12, 0x48, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x10, 0x48, 0x69, 0x76, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x3a,
	0x6e, 0x0a, 0x11, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x6f, 0x70, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfd, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x68,
	0x69, 0x76, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x48, 0x69, 0x76, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0f,
	0x68, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x73, 0x3a,
	0x66, 0x0a, 0x0f, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70,
	0x74, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xfd, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x68, 0x69, 0x76, 0x65, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x48, 0x69, 0x76, 0x65, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x68, 0x69, 0x76, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x73, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x68, 0x69, 0x76,
	0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_hive_options_hive_options_proto_rawDescOnce sync.Once
	file_hive_options_hive_options_proto_rawDescData = file_hive_options_hive_options_proto_rawDesc
)

func file_hive_options_hive_options_proto_rawDescGZIP() []byte {
	file_hive_options_hive_options_proto_rawDescOnce.Do(func() {
		file_hive_options_hive_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_hive_options_hive_options_proto_rawDescData)
	})
	return file_hive_options_hive_options_proto_rawDescData
}

var file_hive_options_hive_options_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_hive_options_hive_options_proto_goTypes = []interface{}{
	(*HiveMessageOptions)(nil),        // 0: hive_options.HiveMessageOptions
	(*HiveFieldOptions)(nil),          // 1: hive_options.HiveFieldOptions
	(*descriptor.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
	(*descriptor.FieldOptions)(nil),   // 3: google.protobuf.FieldOptions
}
var file_hive_options_hive_options_proto_depIdxs = []int32{
	2, // 0: hive_options.hive_message_opts:extendee -> google.protobuf.MessageOptions
	3, // 1: hive_options.hive_field_opts:extendee -> google.protobuf.FieldOptions
	0, // 2: hive_options.hive_message_opts:type_name -> hive_options.HiveMessageOptions
	1, // 3: hive_options.hive_field_opts:type_name -> hive_options.HiveFieldOptions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hive_options_hive_options_proto_init() }
func file_hive_options_hive_options_proto_init() {
	if File_hive_options_hive_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hive_options_hive_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HiveMessageOptions); i {
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
		file_hive_options_hive_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HiveFieldOptions); i {
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
			RawDescriptor: file_hive_options_hive_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_hive_options_hive_options_proto_goTypes,
		DependencyIndexes: file_hive_options_hive_options_proto_depIdxs,
		MessageInfos:      file_hive_options_hive_options_proto_msgTypes,
		ExtensionInfos:    file_hive_options_hive_options_proto_extTypes,
	}.Build()
	File_hive_options_hive_options_proto = out.File
	file_hive_options_hive_options_proto_rawDesc = nil
	file_hive_options_hive_options_proto_goTypes = nil
	file_hive_options_hive_options_proto_depIdxs = nil
}