// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: user_profile/messages.proto

package user_profile

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ProfileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname         string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email            string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	AboutUser        string `protobuf:"bytes,3,opt,name=about_user,json=aboutUser,proto3" json:"about_user,omitempty"`
	AvatarProfileUrl string `protobuf:"bytes,4,opt,name=avatar_profile_url,json=avatarProfileUrl,proto3" json:"avatar_profile_url,omitempty"`
}

func (x *ProfileInfo) Reset() {
	*x = ProfileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileInfo) ProtoMessage() {}

func (x *ProfileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileInfo.ProtoReflect.Descriptor instead.
func (*ProfileInfo) Descriptor() ([]byte, []int) {
	return file_user_profile_messages_proto_rawDescGZIP(), []int{0}
}

func (x *ProfileInfo) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *ProfileInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ProfileInfo) GetAboutUser() string {
	if x != nil {
		return x.AboutUser
	}
	return ""
}

func (x *ProfileInfo) GetAvatarProfileUrl() string {
	if x != nil {
		return x.AvatarProfileUrl
	}
	return ""
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user_id - уникальный идентификатор пользователя
	UserId uint64 `protobuf:"varint,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
	// info - информация о пользователе
	Info *Profile `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_user_profile_messages_proto_rawDescGZIP(), []int{1}
}

func (x *Profile) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Profile) GetInfo() *Profile {
	if x != nil {
		return x.Info
	}
	return nil
}

// GetProfileRequest - запрос информации о пользователе
type GetProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user_id - уникальный идентификатор пользователя
	UserId uint64 `protobuf:"varint,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
}

func (x *GetProfileRequest) Reset() {
	*x = GetProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileRequest) ProtoMessage() {}

func (x *GetProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileRequest.ProtoReflect.Descriptor instead.
func (*GetProfileRequest) Descriptor() ([]byte, []int) {
	return file_user_profile_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GetProfileRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// GetProfileResponse - ответ с информацией о пользователе
type GetProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *ProfileInfo `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *GetProfileResponse) Reset() {
	*x = GetProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileResponse) ProtoMessage() {}

func (x *GetProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileResponse.ProtoReflect.Descriptor instead.
func (*GetProfileResponse) Descriptor() ([]byte, []int) {
	return file_user_profile_messages_proto_rawDescGZIP(), []int{3}
}

func (x *GetProfileResponse) GetProfile() *ProfileInfo {
	if x != nil {
		return x.Profile
	}
	return nil
}

// CreateProfileRequest - запрос на создание профиля
type CreateProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname       string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email          string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password       string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	AboutUser      string `protobuf:"bytes,4,opt,name=about_user,json=about,proto3" json:"about_user,omitempty"`
	AvatarPhotoUrl string `protobuf:"bytes,5,opt,name=avatar_photo_url,proto3" json:"avatar_photo_url,omitempty"`
}

func (x *CreateProfileRequest) Reset() {
	*x = CreateProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfileRequest) ProtoMessage() {}

func (x *CreateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfileRequest.ProtoReflect.Descriptor instead.
func (*CreateProfileRequest) Descriptor() ([]byte, []int) {
	return file_user_profile_messages_proto_rawDescGZIP(), []int{4}
}

func (x *CreateProfileRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *CreateProfileRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateProfileRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateProfileRequest) GetAboutUser() string {
	if x != nil {
		return x.AboutUser
	}
	return ""
}

func (x *CreateProfileRequest) GetAvatarPhotoUrl() string {
	if x != nil {
		return x.AvatarPhotoUrl
	}
	return ""
}

// CreateProfileResponse - ответ на создание профиля
type CreateProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
}

func (x *CreateProfileResponse) Reset() {
	*x = CreateProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfileResponse) ProtoMessage() {}

func (x *CreateProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfileResponse.ProtoReflect.Descriptor instead.
func (*CreateProfileResponse) Descriptor() ([]byte, []int) {
	return file_user_profile_messages_proto_rawDescGZIP(), []int{5}
}

func (x *CreateProfileResponse) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// UpdateProfileRequest - запрос на обновление профиля
type UpdateProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         uint64 `protobuf:"varint,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
	Nickname       string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email          string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	AboutUser      string `protobuf:"bytes,4,opt,name=about_user,proto3" json:"about_user,omitempty"`
	AvatarPhotoUrl string `protobuf:"bytes,5,opt,name=avatar_photo_url,proto3" json:"avatar_photo_url,omitempty"`
}

func (x *UpdateProfileRequest) Reset() {
	*x = UpdateProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileRequest) ProtoMessage() {}

func (x *UpdateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileRequest.ProtoReflect.Descriptor instead.
func (*UpdateProfileRequest) Descriptor() ([]byte, []int) {
	return file_user_profile_messages_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateProfileRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateProfileRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UpdateProfileRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateProfileRequest) GetAboutUser() string {
	if x != nil {
		return x.AboutUser
	}
	return ""
}

func (x *UpdateProfileRequest) GetAvatarPhotoUrl() string {
	if x != nil {
		return x.AvatarPhotoUrl
	}
	return ""
}

// UpdateProfileResponse - ответ на обновление профиля
type UpdateProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *UpdateProfileResponse) Reset() {
	*x = UpdateProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileResponse) ProtoMessage() {}

func (x *UpdateProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileResponse.ProtoReflect.Descriptor instead.
func (*UpdateProfileResponse) Descriptor() ([]byte, []int) {
	return file_user_profile_messages_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateProfileResponse) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

// DeleteProfileRequest - запрос на удаление профиля
type DeleteProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
}

func (x *DeleteProfileRequest) Reset() {
	*x = DeleteProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProfileRequest) ProtoMessage() {}

func (x *DeleteProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProfileRequest.ProtoReflect.Descriptor instead.
func (*DeleteProfileRequest) Descriptor() ([]byte, []int) {
	return file_user_profile_messages_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteProfileRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// DeleteProfileResponse - ответ на удаление профиля
type DeleteProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
}

func (x *DeleteProfileResponse) Reset() {
	*x = DeleteProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_profile_messages_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProfileResponse) ProtoMessage() {}

func (x *DeleteProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_profile_messages_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProfileResponse.ProtoReflect.Descriptor instead.
func (*DeleteProfileResponse) Descriptor() ([]byte, []int) {
	return file_user_profile_messages_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteProfileResponse) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_user_profile_messages_proto protoreflect.FileDescriptor

var file_user_profile_messages_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x6f, 0x6c, 0x6c, 0x79, 0x68,
	0x6f, 0x78, 0x5f, 0x32, 0x31, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x62,
	0x6f, 0x75, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x62, 0x6f, 0x75, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x3a, 0x45, 0x92, 0x41, 0x42, 0x0a, 0x40, 0x2a, 0x0b,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x31, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x20, 0x2d, 0x20, 0xd0, 0xb8, 0xd0, 0xbd, 0xd1,
	0x84, 0xd0, 0xbe, 0xd1, 0x80, 0xd0, 0xbc, 0xd0, 0xb0, 0xd1, 0x86, 0xd0, 0xb8, 0xd1, 0x8f, 0x20,
	0xd0, 0xbf, 0xd1, 0x80, 0xd0, 0xbe, 0xd1, 0x84, 0xd0, 0xb8, 0xd0, 0xbb, 0xd1, 0x8f, 0x22, 0x7d,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x12, 0x58, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x44, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x68,
	0x6f, 0x6c, 0x6c, 0x79, 0x68, 0x6f, 0x78, 0x5f, 0x32, 0x31, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x2d, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x78, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x62, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x68, 0x6f, 0x6c, 0x6c, 0x79, 0x68, 0x6f, 0x78, 0x5f, 0x32, 0x31, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0xcc, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x09, 0xe0, 0x41, 0x02, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xe0, 0x41, 0x02, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xe0, 0x41, 0x02, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x19,
	0x0a, 0x0a, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x31, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0xae, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a,
	0x0a, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2a, 0x0a,
	0x10, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x77, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5e, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x68, 0x6f, 0x6c, 0x6c, 0x79, 0x68, 0x6f, 0x78, 0x5f, 0x32, 0x31, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x22, 0x30, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x5f, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x6c, 0x6c, 0x79, 0x68, 0x6f, 0x78, 0x2d, 0x32,
	0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x3b, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_profile_messages_proto_rawDescOnce sync.Once
	file_user_profile_messages_proto_rawDescData = file_user_profile_messages_proto_rawDesc
)

func file_user_profile_messages_proto_rawDescGZIP() []byte {
	file_user_profile_messages_proto_rawDescOnce.Do(func() {
		file_user_profile_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_profile_messages_proto_rawDescData)
	})
	return file_user_profile_messages_proto_rawDescData
}

var file_user_profile_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_user_profile_messages_proto_goTypes = []interface{}{
	(*ProfileInfo)(nil),           // 0: github.com.hollyhox_21.discord_project.user_profile_service.ProfileInfo
	(*Profile)(nil),               // 1: github.com.hollyhox_21.discord_project.user_profile_service.Profile
	(*GetProfileRequest)(nil),     // 2: github.com.hollyhox_21.discord_project.user_profile_service.GetProfileRequest
	(*GetProfileResponse)(nil),    // 3: github.com.hollyhox_21.discord_project.user_profile_service.GetProfileResponse
	(*CreateProfileRequest)(nil),  // 4: github.com.hollyhox_21.discord_project.user_profile_service.CreateProfileRequest
	(*CreateProfileResponse)(nil), // 5: github.com.hollyhox_21.discord_project.user_profile_service.CreateProfileResponse
	(*UpdateProfileRequest)(nil),  // 6: github.com.hollyhox_21.discord_project.user_profile_service.UpdateProfileRequest
	(*UpdateProfileResponse)(nil), // 7: github.com.hollyhox_21.discord_project.user_profile_service.UpdateProfileResponse
	(*DeleteProfileRequest)(nil),  // 8: github.com.hollyhox_21.discord_project.user_profile_service.DeleteProfileRequest
	(*DeleteProfileResponse)(nil), // 9: github.com.hollyhox_21.discord_project.user_profile_service.DeleteProfileResponse
}
var file_user_profile_messages_proto_depIdxs = []int32{
	1, // 0: github.com.hollyhox_21.discord_project.user_profile_service.Profile.info:type_name -> github.com.hollyhox_21.discord_project.user_profile_service.Profile
	0, // 1: github.com.hollyhox_21.discord_project.user_profile_service.GetProfileResponse.profile:type_name -> github.com.hollyhox_21.discord_project.user_profile_service.ProfileInfo
	1, // 2: github.com.hollyhox_21.discord_project.user_profile_service.UpdateProfileResponse.profile:type_name -> github.com.hollyhox_21.discord_project.user_profile_service.Profile
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_user_profile_messages_proto_init() }
func file_user_profile_messages_proto_init() {
	if File_user_profile_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_profile_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileInfo); i {
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
		file_user_profile_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_user_profile_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileRequest); i {
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
		file_user_profile_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileResponse); i {
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
		file_user_profile_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProfileRequest); i {
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
		file_user_profile_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProfileResponse); i {
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
		file_user_profile_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProfileRequest); i {
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
		file_user_profile_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProfileResponse); i {
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
		file_user_profile_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProfileRequest); i {
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
		file_user_profile_messages_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProfileResponse); i {
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
			RawDescriptor: file_user_profile_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_profile_messages_proto_goTypes,
		DependencyIndexes: file_user_profile_messages_proto_depIdxs,
		MessageInfos:      file_user_profile_messages_proto_msgTypes,
	}.Build()
	File_user_profile_messages_proto = out.File
	file_user_profile_messages_proto_rawDesc = nil
	file_user_profile_messages_proto_goTypes = nil
	file_user_profile_messages_proto_depIdxs = nil
}