// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: user/v1/user.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// CreateUserRequest 创建用户
type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 邮箱
	Email string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	// 密码
	Passwd string `protobuf:"bytes,2,opt,name=Passwd,proto3" json:"Passwd,omitempty"`
	// 电话
	Phone string `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`
	// IP
	LastIp string `protobuf:"bytes,5,opt,name=LastIp,proto3" json:"LastIp,omitempty"`
	// 客户端头部
	AcceptLanguage string `protobuf:"bytes,6,opt,name=AcceptLanguage,proto3" json:"AcceptLanguage,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetPasswd() string {
	if x != nil {
		return x.Passwd
	}
	return ""
}

func (x *CreateUserRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateUserRequest) GetLastIp() string {
	if x != nil {
		return x.LastIp
	}
	return ""
}

func (x *CreateUserRequest) GetAcceptLanguage() string {
	if x != nil {
		return x.AcceptLanguage
	}
	return ""
}

// CreateUserReply 创建用户返回
type CreateUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 邮箱
	Email string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	// 所属国家
	Country string `protobuf:"bytes,3,opt,name=Country,proto3" json:"Country,omitempty"`
	// 电话
	Phone string `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`
	// token
	Token string `protobuf:"bytes,5,opt,name=Token,proto3" json:"Token,omitempty"`
	// 过期时间
	ExpiresAt string `protobuf:"bytes,6,opt,name=ExpiresAt,proto3" json:"ExpiresAt,omitempty"`
}

func (x *CreateUserReply) Reset() {
	*x = CreateUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserReply) ProtoMessage() {}

func (x *CreateUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserReply.ProtoReflect.Descriptor instead.
func (*CreateUserReply) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserReply) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserReply) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CreateUserReply) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateUserReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateUserReply) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

// UserLoginRequest 用户登录
type UserLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 邮箱
	Email string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	// 密码
	Passwd string `protobuf:"bytes,2,opt,name=Passwd,proto3" json:"Passwd,omitempty"`
	// IP
	LastIp string `protobuf:"bytes,5,opt,name=LastIp,proto3" json:"LastIp,omitempty"`
	// 客户端头部
	AcceptLanguage string `protobuf:"bytes,6,opt,name=AcceptLanguage,proto3" json:"AcceptLanguage,omitempty"`
}

func (x *UserLoginRequest) Reset() {
	*x = UserLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginRequest) ProtoMessage() {}

func (x *UserLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginRequest.ProtoReflect.Descriptor instead.
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserLoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserLoginRequest) GetPasswd() string {
	if x != nil {
		return x.Passwd
	}
	return ""
}

func (x *UserLoginRequest) GetLastIp() string {
	if x != nil {
		return x.LastIp
	}
	return ""
}

func (x *UserLoginRequest) GetAcceptLanguage() string {
	if x != nil {
		return x.AcceptLanguage
	}
	return ""
}

// UserLoginReply 用户登录回复
type UserLoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 邮箱
	Email string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	// 所属国家
	Country string `protobuf:"bytes,3,opt,name=Country,proto3" json:"Country,omitempty"`
	// 电话
	Phone string `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`
	// Token
	Token string `protobuf:"bytes,5,opt,name=Token,proto3" json:"Token,omitempty"`
	// 过期时间
	ExpiresAt string `protobuf:"bytes,6,opt,name=ExpiresAt,proto3" json:"ExpiresAt,omitempty"`
}

func (x *UserLoginReply) Reset() {
	*x = UserLoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginReply) ProtoMessage() {}

func (x *UserLoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginReply.ProtoReflect.Descriptor instead.
func (*UserLoginReply) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserLoginReply) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserLoginReply) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *UserLoginReply) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserLoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserLoginReply) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

// UserFindOrModifyPasswdRequest 用户找回或修改密码
type UserFindOrModifyPasswdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	// IP
	LastIp string `protobuf:"bytes,2,opt,name=LastIp,proto3" json:"LastIp,omitempty"`
	// 客户端头部
	AcceptLanguage string `protobuf:"bytes,3,opt,name=AcceptLanguage,proto3" json:"AcceptLanguage,omitempty"`
	// 验证码ID
	CaptchaId string `protobuf:"bytes,4,opt,name=CaptchaId,proto3" json:"CaptchaId,omitempty"`
	// 回答
	Answer string `protobuf:"bytes,5,opt,name=Answer,proto3" json:"Answer,omitempty"`
}

func (x *UserFindOrModifyPasswdRequest) Reset() {
	*x = UserFindOrModifyPasswdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFindOrModifyPasswdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFindOrModifyPasswdRequest) ProtoMessage() {}

func (x *UserFindOrModifyPasswdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFindOrModifyPasswdRequest.ProtoReflect.Descriptor instead.
func (*UserFindOrModifyPasswdRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserFindOrModifyPasswdRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserFindOrModifyPasswdRequest) GetLastIp() string {
	if x != nil {
		return x.LastIp
	}
	return ""
}

func (x *UserFindOrModifyPasswdRequest) GetAcceptLanguage() string {
	if x != nil {
		return x.AcceptLanguage
	}
	return ""
}

func (x *UserFindOrModifyPasswdRequest) GetCaptchaId() string {
	if x != nil {
		return x.CaptchaId
	}
	return ""
}

func (x *UserFindOrModifyPasswdRequest) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

// UserFindOrModifyPasswdReply 用户找回或修改密码返回
type UserFindOrModifyPasswdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool   `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Url    string `protobuf:"bytes,2,opt,name=Url,proto3" json:"Url,omitempty"`
	UrlStr string `protobuf:"bytes,3,opt,name=urlStr,proto3" json:"urlStr,omitempty"`
}

func (x *UserFindOrModifyPasswdReply) Reset() {
	*x = UserFindOrModifyPasswdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFindOrModifyPasswdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFindOrModifyPasswdReply) ProtoMessage() {}

func (x *UserFindOrModifyPasswdReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFindOrModifyPasswdReply.ProtoReflect.Descriptor instead.
func (*UserFindOrModifyPasswdReply) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserFindOrModifyPasswdReply) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *UserFindOrModifyPasswdReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UserFindOrModifyPasswdReply) GetUrlStr() string {
	if x != nil {
		return x.UrlStr
	}
	return ""
}

// UserModifyPasswdRequest 修改密码
type UserModifyPasswdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 邮箱
	Email string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	// 找回密码返回KEY
	Code string `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	// IP
	LastIp string `protobuf:"bytes,3,opt,name=LastIp,proto3" json:"LastIp,omitempty"`
	// 密码
	Passwd string `protobuf:"bytes,4,opt,name=Passwd,proto3" json:"Passwd,omitempty"`
	// 客户端头部
	AcceptLanguage string `protobuf:"bytes,5,opt,name=AcceptLanguage,proto3" json:"AcceptLanguage,omitempty"`
}

func (x *UserModifyPasswdRequest) Reset() {
	*x = UserModifyPasswdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserModifyPasswdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserModifyPasswdRequest) ProtoMessage() {}

func (x *UserModifyPasswdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserModifyPasswdRequest.ProtoReflect.Descriptor instead.
func (*UserModifyPasswdRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *UserModifyPasswdRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserModifyPasswdRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UserModifyPasswdRequest) GetLastIp() string {
	if x != nil {
		return x.LastIp
	}
	return ""
}

func (x *UserModifyPasswdRequest) GetPasswd() string {
	if x != nil {
		return x.Passwd
	}
	return ""
}

func (x *UserModifyPasswdRequest) GetAcceptLanguage() string {
	if x != nil {
		return x.AcceptLanguage
	}
	return ""
}

// UserModifyPasswdReply 修改密码返回
type UserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *UserReply) Reset() {
	*x = UserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReply) ProtoMessage() {}

func (x *UserReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReply.ProtoReflect.Descriptor instead.
func (*UserReply) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserReply) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type RefreshTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RefreshTokenRequest) Reset() {
	*x = RefreshTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRequest) ProtoMessage() {}

func (x *RefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{8}
}

type RefreshTokenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// token
	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	// 过期时间
	ExpiresAt string `protobuf:"bytes,2,opt,name=ExpiresAt,proto3" json:"ExpiresAt,omitempty"`
}

func (x *RefreshTokenReply) Reset() {
	*x = RefreshTokenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenReply) ProtoMessage() {}

func (x *RefreshTokenReply) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenReply.ProtoReflect.Descriptor instead.
func (*RefreshTokenReply) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{9}
}

func (x *RefreshTokenReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RefreshTokenReply) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

var File_user_v1_user_proto protoreflect.FileDescriptor

var file_user_v1_user_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a,
	0x06, 0x50, 0x61, 0x73, 0x73, 0x77, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x72, 0x04, 0x10, 0x08, 0x18, 0x14, 0x52, 0x06, 0x50, 0x61, 0x73, 0x73, 0x77, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x4c, 0x61, 0x73, 0x74, 0x49, 0x70,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x70, 0x01, 0x52,
	0x06, 0x4c, 0x61, 0x73, 0x74, 0x49, 0x70, 0x12, 0x30, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xd0, 0x01, 0x00, 0x52, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22, 0xa7, 0x01, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x06, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06,
	0x72, 0x04, 0x10, 0x08, 0x18, 0x14, 0x52, 0x06, 0x50, 0x61, 0x73, 0x73, 0x77, 0x64, 0x12, 0x1f,
	0x0a, 0x06, 0x4c, 0x61, 0x73, 0x74, 0x49, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x70, 0x01, 0x52, 0x06, 0x4c, 0x61, 0x73, 0x74, 0x49, 0x70, 0x12,
	0x30, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xd0, 0x01,
	0x00, 0x52, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x22, 0x8a, 0x01, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22, 0xd2,
	0x01, 0x0a, 0x1d, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x72, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x06, 0x4c, 0x61, 0x73, 0x74, 0x49, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x70, 0x01, 0x52,
	0x06, 0x4c, 0x61, 0x73, 0x74, 0x49, 0x70, 0x12, 0x30, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xd0, 0x01, 0x00, 0x52, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x09, 0x43, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x72, 0x03, 0xd0, 0x01, 0x00, 0x52, 0x09, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xd0, 0x01, 0x00, 0x52, 0x06, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x22, 0x5f, 0x0a, 0x1b, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x4f,
	0x72, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x72, 0x6c, 0x53, 0x74, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x72,
	0x6c, 0x53, 0x74, 0x72, 0x22, 0xc2, 0x01, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x4c, 0x61, 0x73, 0x74, 0x49, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x70, 0x01, 0x52, 0x06, 0x4c, 0x61,
	0x73, 0x74, 0x49, 0x70, 0x12, 0x21, 0x0a, 0x06, 0x50, 0x61, 0x73, 0x73, 0x77, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x08, 0x18, 0x14, 0x52,
	0x06, 0x50, 0x61, 0x73, 0x73, 0x77, 0x64, 0x12, 0x30, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xd0, 0x01, 0x00, 0x52, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x23, 0x0a, 0x09, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x15,
	0x0a, 0x13, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x47, 0x0a, 0x11, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x32, 0xea,
	0x04, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x68, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x67, 0x12, 0x67, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x9b, 0x01, 0x0a, 0x16, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x72, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x64, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x72, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x72, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x72, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x64, 0x12, 0x7b, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x64, 0x12, 0x24, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x64, 0x12, 0x74, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12,
	0x1a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x3e, 0x0a, 0x0b, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x64, 0x74, 0x72, 0x69, 0x62,
	0x2f, 0x68, 0x61, 0x6f, 0x6a, 0x69, 0x73, 0x68, 0x75, 0x2d, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_user_v1_user_proto_rawDescOnce sync.Once
	file_user_v1_user_proto_rawDescData = file_user_v1_user_proto_rawDesc
)

func file_user_v1_user_proto_rawDescGZIP() []byte {
	file_user_v1_user_proto_rawDescOnce.Do(func() {
		file_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_v1_user_proto_rawDescData)
	})
	return file_user_v1_user_proto_rawDescData
}

var file_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_user_v1_user_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),             // 0: api.user.v1.CreateUserRequest
	(*CreateUserReply)(nil),               // 1: api.user.v1.CreateUserReply
	(*UserLoginRequest)(nil),              // 2: api.user.v1.UserLoginRequest
	(*UserLoginReply)(nil),                // 3: api.user.v1.UserLoginReply
	(*UserFindOrModifyPasswdRequest)(nil), // 4: api.user.v1.UserFindOrModifyPasswdRequest
	(*UserFindOrModifyPasswdReply)(nil),   // 5: api.user.v1.UserFindOrModifyPasswdReply
	(*UserModifyPasswdRequest)(nil),       // 6: api.user.v1.UserModifyPasswdRequest
	(*UserReply)(nil),                     // 7: api.user.v1.UserReply
	(*RefreshTokenRequest)(nil),           // 8: api.user.v1.RefreshTokenRequest
	(*RefreshTokenReply)(nil),             // 9: api.user.v1.RefreshTokenReply
}
var file_user_v1_user_proto_depIdxs = []int32{
	0, // 0: api.user.v1.User.CreateUser:input_type -> api.user.v1.CreateUserRequest
	2, // 1: api.user.v1.User.UserLogin:input_type -> api.user.v1.UserLoginRequest
	4, // 2: api.user.v1.User.UserFindOrModifyPasswd:input_type -> api.user.v1.UserFindOrModifyPasswdRequest
	6, // 3: api.user.v1.User.UserModifyPasswd:input_type -> api.user.v1.UserModifyPasswdRequest
	8, // 4: api.user.v1.User.RefreshToken:input_type -> api.user.v1.RefreshTokenRequest
	1, // 5: api.user.v1.User.CreateUser:output_type -> api.user.v1.CreateUserReply
	3, // 6: api.user.v1.User.UserLogin:output_type -> api.user.v1.UserLoginReply
	5, // 7: api.user.v1.User.UserFindOrModifyPasswd:output_type -> api.user.v1.UserFindOrModifyPasswdReply
	7, // 8: api.user.v1.User.UserModifyPasswd:output_type -> api.user.v1.UserReply
	9, // 9: api.user.v1.User.RefreshToken:output_type -> api.user.v1.RefreshTokenReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_v1_user_proto_init() }
func file_user_v1_user_proto_init() {
	if File_user_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_user_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserReply); i {
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
		file_user_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginRequest); i {
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
		file_user_v1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginReply); i {
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
		file_user_v1_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFindOrModifyPasswdRequest); i {
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
		file_user_v1_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFindOrModifyPasswdReply); i {
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
		file_user_v1_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserModifyPasswdRequest); i {
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
		file_user_v1_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReply); i {
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
		file_user_v1_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokenRequest); i {
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
		file_user_v1_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokenReply); i {
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
			RawDescriptor: file_user_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_v1_user_proto_goTypes,
		DependencyIndexes: file_user_v1_user_proto_depIdxs,
		MessageInfos:      file_user_v1_user_proto_msgTypes,
	}.Build()
	File_user_v1_user_proto = out.File
	file_user_v1_user_proto_rawDesc = nil
	file_user_v1_user_proto_goTypes = nil
	file_user_v1_user_proto_depIdxs = nil
}
