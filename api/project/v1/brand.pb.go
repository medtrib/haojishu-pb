// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: project/v1/brand.proto

package v1

import (
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

// 创建品牌
type CreateBrandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectBrandName          string `protobuf:"bytes,1,opt,name=projectBrandName,proto3" json:"projectBrandName,omitempty"`
	ProjectBrandFirstLetter   string `protobuf:"bytes,2,opt,name=projectBrandFirstLetter,proto3" json:"projectBrandFirstLetter,omitempty"`
	ProjectBrandSort          uint32 `protobuf:"varint,3,opt,name=projectBrandSort,proto3" json:"projectBrandSort,omitempty"`
	ProjectBrandFactoryStatus uint32 `protobuf:"varint,4,opt,name=projectBrandFactoryStatus,proto3" json:"projectBrandFactoryStatus,omitempty"`
	ProjectBrandShowStatus    uint32 `protobuf:"varint,5,opt,name=projectBrandShowStatus,proto3" json:"projectBrandShowStatus,omitempty"`
	ProjectBrandLogo          string `protobuf:"bytes,6,opt,name=projectBrandLogo,proto3" json:"projectBrandLogo,omitempty"`
	ProjectBrandBigPic        string `protobuf:"bytes,7,opt,name=projectBrandBigPic,proto3" json:"projectBrandBigPic,omitempty"`
	ProjectBrandBrandInfo     string `protobuf:"bytes,8,opt,name=projectBrandBrandInfo,proto3" json:"projectBrandBrandInfo,omitempty"`
}

func (x *CreateBrandRequest) Reset() {
	*x = CreateBrandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_brand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBrandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBrandRequest) ProtoMessage() {}

func (x *CreateBrandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_brand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBrandRequest.ProtoReflect.Descriptor instead.
func (*CreateBrandRequest) Descriptor() ([]byte, []int) {
	return file_project_v1_brand_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBrandRequest) GetProjectBrandName() string {
	if x != nil {
		return x.ProjectBrandName
	}
	return ""
}

func (x *CreateBrandRequest) GetProjectBrandFirstLetter() string {
	if x != nil {
		return x.ProjectBrandFirstLetter
	}
	return ""
}

func (x *CreateBrandRequest) GetProjectBrandSort() uint32 {
	if x != nil {
		return x.ProjectBrandSort
	}
	return 0
}

func (x *CreateBrandRequest) GetProjectBrandFactoryStatus() uint32 {
	if x != nil {
		return x.ProjectBrandFactoryStatus
	}
	return 0
}

func (x *CreateBrandRequest) GetProjectBrandShowStatus() uint32 {
	if x != nil {
		return x.ProjectBrandShowStatus
	}
	return 0
}

func (x *CreateBrandRequest) GetProjectBrandLogo() string {
	if x != nil {
		return x.ProjectBrandLogo
	}
	return ""
}

func (x *CreateBrandRequest) GetProjectBrandBigPic() string {
	if x != nil {
		return x.ProjectBrandBigPic
	}
	return ""
}

func (x *CreateBrandRequest) GetProjectBrandBrandInfo() string {
	if x != nil {
		return x.ProjectBrandBrandInfo
	}
	return ""
}

type CreateBrandReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateBrandReply) Reset() {
	*x = CreateBrandReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_brand_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBrandReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBrandReply) ProtoMessage() {}

func (x *CreateBrandReply) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_brand_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBrandReply.ProtoReflect.Descriptor instead.
func (*CreateBrandReply) Descriptor() ([]byte, []int) {
	return file_project_v1_brand_proto_rawDescGZIP(), []int{1}
}

type UpdateBrandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBrandRequest) Reset() {
	*x = UpdateBrandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_brand_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBrandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBrandRequest) ProtoMessage() {}

func (x *UpdateBrandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_brand_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBrandRequest.ProtoReflect.Descriptor instead.
func (*UpdateBrandRequest) Descriptor() ([]byte, []int) {
	return file_project_v1_brand_proto_rawDescGZIP(), []int{2}
}

type UpdateBrandReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBrandReply) Reset() {
	*x = UpdateBrandReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_brand_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBrandReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBrandReply) ProtoMessage() {}

func (x *UpdateBrandReply) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_brand_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBrandReply.ProtoReflect.Descriptor instead.
func (*UpdateBrandReply) Descriptor() ([]byte, []int) {
	return file_project_v1_brand_proto_rawDescGZIP(), []int{3}
}

type DeleteBrandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBrandRequest) Reset() {
	*x = DeleteBrandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_brand_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBrandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBrandRequest) ProtoMessage() {}

func (x *DeleteBrandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_brand_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBrandRequest.ProtoReflect.Descriptor instead.
func (*DeleteBrandRequest) Descriptor() ([]byte, []int) {
	return file_project_v1_brand_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteBrandRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteBrandReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBrandReply) Reset() {
	*x = DeleteBrandReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_brand_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBrandReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBrandReply) ProtoMessage() {}

func (x *DeleteBrandReply) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_brand_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBrandReply.ProtoReflect.Descriptor instead.
func (*DeleteBrandReply) Descriptor() ([]byte, []int) {
	return file_project_v1_brand_proto_rawDescGZIP(), []int{5}
}

type GetBrandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBrandRequest) Reset() {
	*x = GetBrandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_brand_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBrandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBrandRequest) ProtoMessage() {}

func (x *GetBrandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_brand_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBrandRequest.ProtoReflect.Descriptor instead.
func (*GetBrandRequest) Descriptor() ([]byte, []int) {
	return file_project_v1_brand_proto_rawDescGZIP(), []int{6}
}

func (x *GetBrandRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBrandReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBrandReply) Reset() {
	*x = GetBrandReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_brand_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBrandReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBrandReply) ProtoMessage() {}

func (x *GetBrandReply) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_brand_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBrandReply.ProtoReflect.Descriptor instead.
func (*GetBrandReply) Descriptor() ([]byte, []int) {
	return file_project_v1_brand_proto_rawDescGZIP(), []int{7}
}

type ListBrandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBrandRequest) Reset() {
	*x = ListBrandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_brand_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBrandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBrandRequest) ProtoMessage() {}

func (x *ListBrandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_brand_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBrandRequest.ProtoReflect.Descriptor instead.
func (*ListBrandRequest) Descriptor() ([]byte, []int) {
	return file_project_v1_brand_proto_rawDescGZIP(), []int{8}
}

type ListBrandReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBrandReply) Reset() {
	*x = ListBrandReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_v1_brand_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBrandReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBrandReply) ProtoMessage() {}

func (x *ListBrandReply) ProtoReflect() protoreflect.Message {
	mi := &file_project_v1_brand_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBrandReply.ProtoReflect.Descriptor instead.
func (*ListBrandReply) Descriptor() ([]byte, []int) {
	return file_project_v1_brand_proto_rawDescGZIP(), []int{9}
}

var File_project_v1_brand_proto protoreflect.FileDescriptor

var file_project_v1_brand_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x03, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x10, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x65,
	0x74, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x65, 0x74,
	0x74, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x53, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x6f, 0x72, 0x74, 0x12,
	0x3c, 0x0a, 0x19, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x19, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x36, 0x0a,
	0x16, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x6f,
	0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x6f, 0x77, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4c, 0x6f, 0x67,
	0x6f, 0x12, 0x2e, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e,
	0x64, 0x42, 0x69, 0x67, 0x50, 0x69, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x42, 0x69, 0x67, 0x50, 0x69,
	0x63, 0x12, 0x34, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e,
	0x64, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x15, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xef, 0x04, 0x0a, 0x05, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x12, 0x7d, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22,
	0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x12, 0x7d, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e,
	0x64, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a,
	0x01, 0x2a, 0x22, 0x1d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e,
	0x64, 0x12, 0x7f, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x76,
	0x31, 0x2f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x73, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x27,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x72, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x76,
	0x31, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x42, 0x2f, 0x0a, 0x0e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x1b, 0x75, 0x6e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_v1_brand_proto_rawDescOnce sync.Once
	file_project_v1_brand_proto_rawDescData = file_project_v1_brand_proto_rawDesc
)

func file_project_v1_brand_proto_rawDescGZIP() []byte {
	file_project_v1_brand_proto_rawDescOnce.Do(func() {
		file_project_v1_brand_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_v1_brand_proto_rawDescData)
	})
	return file_project_v1_brand_proto_rawDescData
}

var file_project_v1_brand_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_project_v1_brand_proto_goTypes = []interface{}{
	(*CreateBrandRequest)(nil), // 0: api.project.v1.CreateBrandRequest
	(*CreateBrandReply)(nil),   // 1: api.project.v1.CreateBrandReply
	(*UpdateBrandRequest)(nil), // 2: api.project.v1.UpdateBrandRequest
	(*UpdateBrandReply)(nil),   // 3: api.project.v1.UpdateBrandReply
	(*DeleteBrandRequest)(nil), // 4: api.project.v1.DeleteBrandRequest
	(*DeleteBrandReply)(nil),   // 5: api.project.v1.DeleteBrandReply
	(*GetBrandRequest)(nil),    // 6: api.project.v1.GetBrandRequest
	(*GetBrandReply)(nil),      // 7: api.project.v1.GetBrandReply
	(*ListBrandRequest)(nil),   // 8: api.project.v1.ListBrandRequest
	(*ListBrandReply)(nil),     // 9: api.project.v1.ListBrandReply
}
var file_project_v1_brand_proto_depIdxs = []int32{
	0, // 0: api.project.v1.Brand.CreateBrand:input_type -> api.project.v1.CreateBrandRequest
	2, // 1: api.project.v1.Brand.UpdateBrand:input_type -> api.project.v1.UpdateBrandRequest
	4, // 2: api.project.v1.Brand.DeleteBrand:input_type -> api.project.v1.DeleteBrandRequest
	6, // 3: api.project.v1.Brand.GetBrand:input_type -> api.project.v1.GetBrandRequest
	8, // 4: api.project.v1.Brand.ListBrand:input_type -> api.project.v1.ListBrandRequest
	1, // 5: api.project.v1.Brand.CreateBrand:output_type -> api.project.v1.CreateBrandReply
	3, // 6: api.project.v1.Brand.UpdateBrand:output_type -> api.project.v1.UpdateBrandReply
	5, // 7: api.project.v1.Brand.DeleteBrand:output_type -> api.project.v1.DeleteBrandReply
	7, // 8: api.project.v1.Brand.GetBrand:output_type -> api.project.v1.GetBrandReply
	9, // 9: api.project.v1.Brand.ListBrand:output_type -> api.project.v1.ListBrandReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_v1_brand_proto_init() }
func file_project_v1_brand_proto_init() {
	if File_project_v1_brand_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_v1_brand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBrandRequest); i {
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
		file_project_v1_brand_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBrandReply); i {
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
		file_project_v1_brand_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBrandRequest); i {
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
		file_project_v1_brand_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBrandReply); i {
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
		file_project_v1_brand_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBrandRequest); i {
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
		file_project_v1_brand_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBrandReply); i {
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
		file_project_v1_brand_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBrandRequest); i {
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
		file_project_v1_brand_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBrandReply); i {
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
		file_project_v1_brand_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBrandRequest); i {
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
		file_project_v1_brand_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBrandReply); i {
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
			RawDescriptor: file_project_v1_brand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_project_v1_brand_proto_goTypes,
		DependencyIndexes: file_project_v1_brand_proto_depIdxs,
		MessageInfos:      file_project_v1_brand_proto_msgTypes,
	}.Build()
	File_project_v1_brand_proto = out.File
	file_project_v1_brand_proto_rawDesc = nil
	file_project_v1_brand_proto_goTypes = nil
	file_project_v1_brand_proto_depIdxs = nil
}
