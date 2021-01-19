// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: account_service.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type Status int32

const (
	Status_STATUS_NORMAL Status = 0
	Status_STATUS_BANNED Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_NORMAL",
		1: "STATUS_BANNED",
	}
	Status_value = map[string]int32{
		"STATUS_NORMAL": 0,
		"STATUS_BANNED": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_account_service_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_account_service_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_account_service_proto_rawDescGZIP(), []int{0}
}

type Role int32

const (
	Role_ROLE_NORMAL Role = 0
	Role_ROLE_ADMIN  Role = 1
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "ROLE_NORMAL",
		1: "ROLE_ADMIN",
	}
	Role_value = map[string]int32{
		"ROLE_NORMAL": 0,
		"ROLE_ADMIN":  1,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_account_service_proto_enumTypes[1].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_account_service_proto_enumTypes[1]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_account_service_proto_rawDescGZIP(), []int{1}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password     string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Avatar       string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Status       Status `protobuf:"varint,5,opt,name=status,proto3,enum=pb.Status" json:"status,omitempty"`
	Score        int64  `protobuf:"varint,6,opt,name=score,proto3" json:"score,omitempty"`
	BannedReason string `protobuf:"bytes,7,opt,name=banned_reason,json=bannedReason,proto3" json:"banned_reason,omitempty"`
	CreateTime   int64  `protobuf:"varint,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	LastUpdate   int64  `protobuf:"varint,9,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
	Role         Role   `protobuf:"varint,10,opt,name=role,proto3,enum=pb.Role" json:"role,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_account_service_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_NORMAL
}

func (x *User) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *User) GetBannedReason() string {
	if x != nil {
		return x.BannedReason
	}
	return ""
}

func (x *User) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *User) GetLastUpdate() int64 {
	if x != nil {
		return x.LastUpdate
	}
	return 0
}

func (x *User) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_NORMAL
}

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_account_service_proto_rawDescGZIP(), []int{1}
}

func (x *UserRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Ip       string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_account_service_proto_rawDescGZIP(), []int{2}
}

func (x *SignUpRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignUpRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type SignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SignUpResponse) Reset() {
	*x = SignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResponse) ProtoMessage() {}

func (x *SignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResponse.ProtoReflect.Descriptor instead.
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return file_account_service_proto_rawDescGZIP(), []int{3}
}

func (x *SignUpResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_account_service_proto_rawDescGZIP(), []int{4}
}

func (x *SignInRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SignInResponse) Reset() {
	*x = SignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResponse) ProtoMessage() {}

func (x *SignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResponse.ProtoReflect.Descriptor instead.
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return file_account_service_proto_rawDescGZIP(), []int{5}
}

func (x *SignInResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

//通过ids获取user,id为post_id,comment_id,reply_id
type UsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	//0为post 1为comment 2为reply
	IdType int64 `protobuf:"varint,2,opt,name=idType,proto3" json:"idType,omitempty"`
}

func (x *UsersRequest) Reset() {
	*x = UsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersRequest) ProtoMessage() {}

func (x *UsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersRequest.ProtoReflect.Descriptor instead.
func (*UsersRequest) Descriptor() ([]byte, []int) {
	return file_account_service_proto_rawDescGZIP(), []int{6}
}

func (x *UsersRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *UsersRequest) GetIdType() int64 {
	if x != nil {
		return x.IdType
	}
	return 0
}

type UsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UsersResponse) Reset() {
	*x = UsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersResponse) ProtoMessage() {}

func (x *UsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersResponse.ProtoReflect.Descriptor instead.
func (*UsersResponse) Descriptor() ([]byte, []int) {
	return file_account_service_proto_rawDescGZIP(), []int{7}
}

func (x *UsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

//通过user_ids获取users
type UsersByUserIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *UsersByUserIdsRequest) Reset() {
	*x = UsersByUserIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersByUserIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersByUserIdsRequest) ProtoMessage() {}

func (x *UsersByUserIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersByUserIdsRequest.ProtoReflect.Descriptor instead.
func (*UsersByUserIdsRequest) Descriptor() ([]byte, []int) {
	return file_account_service_proto_rawDescGZIP(), []int{8}
}

func (x *UsersByUserIdsRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type UsersByUserIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UsersByUserIdsResponse) Reset() {
	*x = UsersByUserIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersByUserIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersByUserIdsResponse) ProtoMessage() {}

func (x *UsersByUserIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersByUserIdsResponse.ProtoReflect.Descriptor instead.
func (*UsersByUserIdsResponse) Descriptor() ([]byte, []int) {
	return file_account_service_proto_rawDescGZIP(), []int{9}
}

func (x *UsersByUserIdsResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_account_service_proto protoreflect.FileDescriptor

var file_account_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x9d, 0x02, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x22, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x64,
	0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62,
	0x61, 0x6e, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4f, 0x0a, 0x0d, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x20, 0x0a, 0x0e, 0x53,
	0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a,
	0x0d, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x20,
	0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x3c, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x42, 0x02, 0x10,
	0x01, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2f,
	0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22,
	0x2d, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x03, 0x42, 0x02, 0x10, 0x01, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x38,
	0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2a, 0x2e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x52,
	0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x42, 0x41, 0x4e, 0x4e, 0x45, 0x44, 0x10, 0x01, 0x2a, 0x27, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10,
	0x01, 0x32, 0xb8, 0x02, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x31,
	0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x20, 0x0a, 0x08, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x08, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x11, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_service_proto_rawDescOnce sync.Once
	file_account_service_proto_rawDescData = file_account_service_proto_rawDesc
)

func file_account_service_proto_rawDescGZIP() []byte {
	file_account_service_proto_rawDescOnce.Do(func() {
		file_account_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_service_proto_rawDescData)
	})
	return file_account_service_proto_rawDescData
}

var file_account_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_account_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_account_service_proto_goTypes = []interface{}{
	(Status)(0),                    // 0: pb.Status
	(Role)(0),                      // 1: pb.Role
	(*User)(nil),                   // 2: pb.User
	(*UserRequest)(nil),            // 3: pb.UserRequest
	(*SignUpRequest)(nil),          // 4: pb.SignUpRequest
	(*SignUpResponse)(nil),         // 5: pb.SignUpResponse
	(*SignInRequest)(nil),          // 6: pb.SignInRequest
	(*SignInResponse)(nil),         // 7: pb.SignInResponse
	(*UsersRequest)(nil),           // 8: pb.UsersRequest
	(*UsersResponse)(nil),          // 9: pb.UsersResponse
	(*UsersByUserIdsRequest)(nil),  // 10: pb.UsersByUserIdsRequest
	(*UsersByUserIdsResponse)(nil), // 11: pb.UsersByUserIdsResponse
}
var file_account_service_proto_depIdxs = []int32{
	0,  // 0: pb.User.status:type_name -> pb.Status
	1,  // 1: pb.User.role:type_name -> pb.Role
	2,  // 2: pb.UsersResponse.users:type_name -> pb.User
	2,  // 3: pb.UsersByUserIdsResponse.users:type_name -> pb.User
	2,  // 4: pb.Account.GetUser:input_type -> pb.User
	4,  // 5: pb.Account.SignUp:input_type -> pb.SignUpRequest
	2,  // 6: pb.Account.EditUser:input_type -> pb.User
	6,  // 7: pb.Account.SignIn:input_type -> pb.SignInRequest
	8,  // 8: pb.Account.GetUsersByIds:input_type -> pb.UsersRequest
	10, // 9: pb.Account.GetUsersByUserIds:input_type -> pb.UsersByUserIdsRequest
	2,  // 10: pb.Account.GetUser:output_type -> pb.User
	5,  // 11: pb.Account.SignUp:output_type -> pb.SignUpResponse
	2,  // 12: pb.Account.EditUser:output_type -> pb.User
	7,  // 13: pb.Account.SignIn:output_type -> pb.SignInResponse
	9,  // 14: pb.Account.GetUsersByIds:output_type -> pb.UsersResponse
	11, // 15: pb.Account.GetUsersByUserIds:output_type -> pb.UsersByUserIdsResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_account_service_proto_init() }
func file_account_service_proto_init() {
	if File_account_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_account_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_account_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpRequest); i {
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
		file_account_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResponse); i {
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
		file_account_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInRequest); i {
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
		file_account_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResponse); i {
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
		file_account_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersRequest); i {
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
		file_account_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersResponse); i {
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
		file_account_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersByUserIdsRequest); i {
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
		file_account_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersByUserIdsResponse); i {
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
			RawDescriptor: file_account_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_service_proto_goTypes,
		DependencyIndexes: file_account_service_proto_depIdxs,
		EnumInfos:         file_account_service_proto_enumTypes,
		MessageInfos:      file_account_service_proto_msgTypes,
	}.Build()
	File_account_service_proto = out.File
	file_account_service_proto_rawDesc = nil
	file_account_service_proto_goTypes = nil
	file_account_service_proto_depIdxs = nil
}
