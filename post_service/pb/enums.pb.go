// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: enums.proto

package pb

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

type UserStatusEnum_Status int32

const (
	UserStatusEnum_NORMAL UserStatusEnum_Status = 0
	UserStatusEnum_BANNED UserStatusEnum_Status = 1
)

// Enum value maps for UserStatusEnum_Status.
var (
	UserStatusEnum_Status_name = map[int32]string{
		0: "NORMAL",
		1: "BANNED",
	}
	UserStatusEnum_Status_value = map[string]int32{
		"NORMAL": 0,
		"BANNED": 1,
	}
)

func (x UserStatusEnum_Status) Enum() *UserStatusEnum_Status {
	p := new(UserStatusEnum_Status)
	*p = x
	return p
}

func (x UserStatusEnum_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserStatusEnum_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[0].Descriptor()
}

func (UserStatusEnum_Status) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[0]
}

func (x UserStatusEnum_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserStatusEnum_Status.Descriptor instead.
func (UserStatusEnum_Status) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{0, 0}
}

type UserRoleEnum_Role int32

const (
	UserRoleEnum_NORMAL UserRoleEnum_Role = 0
	UserRoleEnum_ADMIN  UserRoleEnum_Role = 1
)

// Enum value maps for UserRoleEnum_Role.
var (
	UserRoleEnum_Role_name = map[int32]string{
		0: "NORMAL",
		1: "ADMIN",
	}
	UserRoleEnum_Role_value = map[string]int32{
		"NORMAL": 0,
		"ADMIN":  1,
	}
)

func (x UserRoleEnum_Role) Enum() *UserRoleEnum_Role {
	p := new(UserRoleEnum_Role)
	*p = x
	return p
}

func (x UserRoleEnum_Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserRoleEnum_Role) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[1].Descriptor()
}

func (UserRoleEnum_Role) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[1]
}

func (x UserRoleEnum_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserRoleEnum_Role.Descriptor instead.
func (UserRoleEnum_Role) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{1, 0}
}

type PostStatusEnum_Status int32

const (
	PostStatusEnum_NORMAL PostStatusEnum_Status = 0
	PostStatusEnum_DELETE PostStatusEnum_Status = 1
)

// Enum value maps for PostStatusEnum_Status.
var (
	PostStatusEnum_Status_name = map[int32]string{
		0: "NORMAL",
		1: "DELETE",
	}
	PostStatusEnum_Status_value = map[string]int32{
		"NORMAL": 0,
		"DELETE": 1,
	}
)

func (x PostStatusEnum_Status) Enum() *PostStatusEnum_Status {
	p := new(PostStatusEnum_Status)
	*p = x
	return p
}

func (x PostStatusEnum_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PostStatusEnum_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_proto_enumTypes[2].Descriptor()
}

func (PostStatusEnum_Status) Type() protoreflect.EnumType {
	return &file_enums_proto_enumTypes[2]
}

func (x PostStatusEnum_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PostStatusEnum_Status.Descriptor instead.
func (PostStatusEnum_Status) EnumDescriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{2, 0}
}

type UserStatusEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserStatusEnum) Reset() {
	*x = UserStatusEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStatusEnum) ProtoMessage() {}

func (x *UserStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStatusEnum.ProtoReflect.Descriptor instead.
func (*UserStatusEnum) Descriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{0}
}

type UserRoleEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserRoleEnum) Reset() {
	*x = UserRoleEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRoleEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRoleEnum) ProtoMessage() {}

func (x *UserRoleEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRoleEnum.ProtoReflect.Descriptor instead.
func (*UserRoleEnum) Descriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{1}
}

type PostStatusEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostStatusEnum) Reset() {
	*x = PostStatusEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostStatusEnum) ProtoMessage() {}

func (x *PostStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_enums_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostStatusEnum.ProtoReflect.Descriptor instead.
func (*PostStatusEnum) Descriptor() ([]byte, []int) {
	return file_enums_proto_rawDescGZIP(), []int{2}
}

var File_enums_proto protoreflect.FileDescriptor

var file_enums_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x22, 0x32, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x42, 0x41, 0x4e, 0x4e, 0x45, 0x44, 0x10, 0x01, 0x22, 0x2d, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x6f, 0x6c, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x1d, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x01, 0x22, 0x32, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_proto_rawDescOnce sync.Once
	file_enums_proto_rawDescData = file_enums_proto_rawDesc
)

func file_enums_proto_rawDescGZIP() []byte {
	file_enums_proto_rawDescOnce.Do(func() {
		file_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_proto_rawDescData)
	})
	return file_enums_proto_rawDescData
}

var file_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_enums_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_enums_proto_goTypes = []interface{}{
	(UserStatusEnum_Status)(0), // 0: enums.UserStatusEnum.Status
	(UserRoleEnum_Role)(0),     // 1: enums.UserRoleEnum.Role
	(PostStatusEnum_Status)(0), // 2: enums.PostStatusEnum.Status
	(*UserStatusEnum)(nil),     // 3: enums.UserStatusEnum
	(*UserRoleEnum)(nil),       // 4: enums.UserRoleEnum
	(*PostStatusEnum)(nil),     // 5: enums.PostStatusEnum
}
var file_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_proto_init() }
func file_enums_proto_init() {
	if File_enums_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStatusEnum); i {
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
		file_enums_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRoleEnum); i {
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
		file_enums_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostStatusEnum); i {
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
			RawDescriptor: file_enums_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_proto_goTypes,
		DependencyIndexes: file_enums_proto_depIdxs,
		EnumInfos:         file_enums_proto_enumTypes,
		MessageInfos:      file_enums_proto_msgTypes,
	}.Build()
	File_enums_proto = out.File
	file_enums_proto_rawDesc = nil
	file_enums_proto_goTypes = nil
	file_enums_proto_depIdxs = nil
}
