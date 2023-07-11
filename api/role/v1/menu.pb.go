// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: role/v1/menu.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MenuReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *MenuReply) Reset() {
	*x = MenuReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_v1_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuReply) ProtoMessage() {}

func (x *MenuReply) ProtoReflect() protoreflect.Message {
	mi := &file_role_v1_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuReply.ProtoReflect.Descriptor instead.
func (*MenuReply) Descriptor() ([]byte, []int) {
	return file_role_v1_menu_proto_rawDescGZIP(), []int{0}
}

func (x *MenuReply) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type CreateMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuName   string `protobuf:"bytes,2,opt,name=menu_name,json=menuName,proto3" json:"menu_name,omitempty"`
	Icon       string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Path       string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Ancestors  string `protobuf:"bytes,5,opt,name=ancestors,proto3" json:"ancestors,omitempty"`
	MenuType   string `protobuf:"bytes,6,opt,name=menu_type,json=menuType,proto3" json:"menu_type,omitempty"`
	Action     string `protobuf:"bytes,7,opt,name=action,proto3" json:"action,omitempty"`
	Permission string `protobuf:"bytes,8,opt,name=permission,proto3" json:"permission,omitempty"`
	ParentId   uint64 `protobuf:"varint,9,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Breadcrumb string `protobuf:"bytes,10,opt,name=breadcrumb,proto3" json:"breadcrumb,omitempty"`
	Component  string `protobuf:"bytes,11,opt,name=component,proto3" json:"component,omitempty"`
	Sort       uint64 `protobuf:"varint,12,opt,name=sort,proto3" json:"sort,omitempty"`
	Visible    uint64 `protobuf:"varint,13,opt,name=visible,proto3" json:"visible,omitempty"`
	IsFrame    uint64 `protobuf:"varint,14,opt,name=is_frame,json=isFrame,proto3" json:"is_frame,omitempty"`
}

func (x *CreateMenuRequest) Reset() {
	*x = CreateMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_v1_menu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMenuRequest) ProtoMessage() {}

func (x *CreateMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_v1_menu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMenuRequest.ProtoReflect.Descriptor instead.
func (*CreateMenuRequest) Descriptor() ([]byte, []int) {
	return file_role_v1_menu_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMenuRequest) GetMenuName() string {
	if x != nil {
		return x.MenuName
	}
	return ""
}

func (x *CreateMenuRequest) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *CreateMenuRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CreateMenuRequest) GetAncestors() string {
	if x != nil {
		return x.Ancestors
	}
	return ""
}

func (x *CreateMenuRequest) GetMenuType() string {
	if x != nil {
		return x.MenuType
	}
	return ""
}

func (x *CreateMenuRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *CreateMenuRequest) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *CreateMenuRequest) GetParentId() uint64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *CreateMenuRequest) GetBreadcrumb() string {
	if x != nil {
		return x.Breadcrumb
	}
	return ""
}

func (x *CreateMenuRequest) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *CreateMenuRequest) GetSort() uint64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *CreateMenuRequest) GetVisible() uint64 {
	if x != nil {
		return x.Visible
	}
	return 0
}

func (x *CreateMenuRequest) GetIsFrame() uint64 {
	if x != nil {
		return x.IsFrame
	}
	return 0
}

type UpdateMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuId     uint64 `protobuf:"varint,1,opt,name=menu_id,json=menuId,proto3" json:"menu_id,omitempty"`
	MenuName   string `protobuf:"bytes,2,opt,name=menu_name,json=menuName,proto3" json:"menu_name,omitempty"`
	Icon       string `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Path       string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Ancestors  string `protobuf:"bytes,5,opt,name=ancestors,proto3" json:"ancestors,omitempty"`
	MenuType   string `protobuf:"bytes,6,opt,name=menu_type,json=menuType,proto3" json:"menu_type,omitempty"`
	Action     string `protobuf:"bytes,7,opt,name=action,proto3" json:"action,omitempty"`
	Permission string `protobuf:"bytes,8,opt,name=permission,proto3" json:"permission,omitempty"`
	ParentId   uint64 `protobuf:"varint,9,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Breadcrumb string `protobuf:"bytes,10,opt,name=breadcrumb,proto3" json:"breadcrumb,omitempty"`
	Component  string `protobuf:"bytes,11,opt,name=component,proto3" json:"component,omitempty"`
	Sort       uint64 `protobuf:"varint,12,opt,name=sort,proto3" json:"sort,omitempty"`
	Visible    uint64 `protobuf:"varint,13,opt,name=visible,proto3" json:"visible,omitempty"`
	IsFrame    uint64 `protobuf:"varint,14,opt,name=is_frame,json=isFrame,proto3" json:"is_frame,omitempty"`
}

func (x *UpdateMenuRequest) Reset() {
	*x = UpdateMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_v1_menu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMenuRequest) ProtoMessage() {}

func (x *UpdateMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_v1_menu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMenuRequest.ProtoReflect.Descriptor instead.
func (*UpdateMenuRequest) Descriptor() ([]byte, []int) {
	return file_role_v1_menu_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateMenuRequest) GetMenuId() uint64 {
	if x != nil {
		return x.MenuId
	}
	return 0
}

func (x *UpdateMenuRequest) GetMenuName() string {
	if x != nil {
		return x.MenuName
	}
	return ""
}

func (x *UpdateMenuRequest) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *UpdateMenuRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UpdateMenuRequest) GetAncestors() string {
	if x != nil {
		return x.Ancestors
	}
	return ""
}

func (x *UpdateMenuRequest) GetMenuType() string {
	if x != nil {
		return x.MenuType
	}
	return ""
}

func (x *UpdateMenuRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *UpdateMenuRequest) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *UpdateMenuRequest) GetParentId() uint64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *UpdateMenuRequest) GetBreadcrumb() string {
	if x != nil {
		return x.Breadcrumb
	}
	return ""
}

func (x *UpdateMenuRequest) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *UpdateMenuRequest) GetSort() uint64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *UpdateMenuRequest) GetVisible() uint64 {
	if x != nil {
		return x.Visible
	}
	return 0
}

func (x *UpdateMenuRequest) GetIsFrame() uint64 {
	if x != nil {
		return x.IsFrame
	}
	return 0
}

type DeleteMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids string `protobuf:"bytes,1,opt,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteMenuRequest) Reset() {
	*x = DeleteMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_v1_menu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMenuRequest) ProtoMessage() {}

func (x *DeleteMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_v1_menu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMenuRequest.ProtoReflect.Descriptor instead.
func (*DeleteMenuRequest) Descriptor() ([]byte, []int) {
	return file_role_v1_menu_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteMenuRequest) GetIds() string {
	if x != nil {
		return x.Ids
	}
	return ""
}

type GetMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMenuRequest) Reset() {
	*x = GetMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_v1_menu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMenuRequest) ProtoMessage() {}

func (x *GetMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_role_v1_menu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMenuRequest.ProtoReflect.Descriptor instead.
func (*GetMenuRequest) Descriptor() ([]byte, []int) {
	return file_role_v1_menu_proto_rawDescGZIP(), []int{4}
}

func (x *GetMenuRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetMenuReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuId     uint64          `protobuf:"varint,1,opt,name=menu_id,json=menuId,proto3" json:"menu_id,omitempty"`
	MenuName   string          `protobuf:"bytes,2,opt,name=menu_name,json=menuName,proto3" json:"menu_name,omitempty"`
	Icon       string          `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Path       string          `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Ancestors  string          `protobuf:"bytes,5,opt,name=ancestors,proto3" json:"ancestors,omitempty"`
	MenuType   string          `protobuf:"bytes,6,opt,name=menu_type,json=menuType,proto3" json:"menu_type,omitempty"`
	Action     string          `protobuf:"bytes,7,opt,name=action,proto3" json:"action,omitempty"`
	Permission string          `protobuf:"bytes,8,opt,name=permission,proto3" json:"permission,omitempty"`
	ParentId   uint64          `protobuf:"varint,9,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Breadcrumb string          `protobuf:"bytes,10,opt,name=breadcrumb,proto3" json:"breadcrumb,omitempty"`
	Component  string          `protobuf:"bytes,11,opt,name=component,proto3" json:"component,omitempty"`
	Sort       uint64          `protobuf:"varint,12,opt,name=sort,proto3" json:"sort,omitempty"`
	Visible    uint64          `protobuf:"varint,13,opt,name=visible,proto3" json:"visible,omitempty"`
	IsFrame    uint64          `protobuf:"varint,14,opt,name=is_frame,json=isFrame,proto3" json:"is_frame,omitempty"`
	Children   []*GetMenuReply `protobuf:"bytes,15,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *GetMenuReply) Reset() {
	*x = GetMenuReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_v1_menu_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMenuReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMenuReply) ProtoMessage() {}

func (x *GetMenuReply) ProtoReflect() protoreflect.Message {
	mi := &file_role_v1_menu_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMenuReply.ProtoReflect.Descriptor instead.
func (*GetMenuReply) Descriptor() ([]byte, []int) {
	return file_role_v1_menu_proto_rawDescGZIP(), []int{5}
}

func (x *GetMenuReply) GetMenuId() uint64 {
	if x != nil {
		return x.MenuId
	}
	return 0
}

func (x *GetMenuReply) GetMenuName() string {
	if x != nil {
		return x.MenuName
	}
	return ""
}

func (x *GetMenuReply) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *GetMenuReply) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GetMenuReply) GetAncestors() string {
	if x != nil {
		return x.Ancestors
	}
	return ""
}

func (x *GetMenuReply) GetMenuType() string {
	if x != nil {
		return x.MenuType
	}
	return ""
}

func (x *GetMenuReply) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *GetMenuReply) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *GetMenuReply) GetParentId() uint64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *GetMenuReply) GetBreadcrumb() string {
	if x != nil {
		return x.Breadcrumb
	}
	return ""
}

func (x *GetMenuReply) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *GetMenuReply) GetSort() uint64 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *GetMenuReply) GetVisible() uint64 {
	if x != nil {
		return x.Visible
	}
	return 0
}

func (x *GetMenuReply) GetIsFrame() uint64 {
	if x != nil {
		return x.IsFrame
	}
	return 0
}

func (x *GetMenuReply) GetChildren() []*GetMenuReply {
	if x != nil {
		return x.Children
	}
	return nil
}

type ListMenuReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*GetMenuReply `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListMenuReply) Reset() {
	*x = ListMenuReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_v1_menu_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMenuReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMenuReply) ProtoMessage() {}

func (x *ListMenuReply) ProtoReflect() protoreflect.Message {
	mi := &file_role_v1_menu_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMenuReply.ProtoReflect.Descriptor instead.
func (*ListMenuReply) Descriptor() ([]byte, []int) {
	return file_role_v1_menu_proto_rawDescGZIP(), []int{6}
}

func (x *ListMenuReply) GetList() []*GetMenuReply {
	if x != nil {
		return x.List
	}
	return nil
}

var File_role_v1_menu_proto protoreflect.FileDescriptor

var file_role_v1_menu_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x23, 0x0a, 0x09, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0xef, 0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65,
	0x6e, 0x75, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x65, 0x6e, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x65, 0x6e, 0x75, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x62, 0x72, 0x65, 0x61, 0x64, 0x63, 0x72, 0x75, 0x6d, 0x62, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x72, 0x65, 0x61, 0x64, 0x63, 0x72, 0x75, 0x6d, 0x62, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x73, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x69,
	0x73, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x22, 0x88, 0x03, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6d,
	0x65, 0x6e, 0x75, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6e, 0x75, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6e, 0x75,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6e,
	0x75, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x72,
	0x65, 0x61, 0x64, 0x63, 0x72, 0x75, 0x6d, 0x62, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x62, 0x72, 0x65, 0x61, 0x64, 0x63, 0x72, 0x75, 0x6d, 0x62, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x69, 0x73, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x22, 0x25, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0xba, 0x03, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x6d,
	0x65, 0x6e, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6d, 0x65,
	0x6e, 0x75, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6e, 0x75, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6e, 0x75, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6e, 0x75,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x72, 0x65,
	0x61, 0x64, 0x63, 0x72, 0x75, 0x6d, 0x62, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62,
	0x72, 0x65, 0x61, 0x64, 0x63, 0x72, 0x75, 0x6d, 0x62, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x69,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x69, 0x73, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x12, 0x35, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x0f, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x08, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0x3e, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x6e, 0x75, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xfe, 0x04, 0x0a, 0x04, 0x4d, 0x65, 0x6e, 0x75,
	0x12, 0x6a, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x1e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e,
	0x75, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01,
	0x2a, 0x22, 0x19, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x6a, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x67, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e,
	0x75, 0x12, 0x61, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x1b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x6e, 0x75, 0x12, 0x64, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75,
	0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x12, 0x17, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x6c, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x54, 0x72, 0x65, 0x65, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x65, 0x6e, 0x75, 0x54, 0x72, 0x65, 0x65, 0x42, 0x29, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x18, 0x75, 0x6e, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_v1_menu_proto_rawDescOnce sync.Once
	file_role_v1_menu_proto_rawDescData = file_role_v1_menu_proto_rawDesc
)

func file_role_v1_menu_proto_rawDescGZIP() []byte {
	file_role_v1_menu_proto_rawDescOnce.Do(func() {
		file_role_v1_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_v1_menu_proto_rawDescData)
	})
	return file_role_v1_menu_proto_rawDescData
}

var file_role_v1_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_role_v1_menu_proto_goTypes = []interface{}{
	(*MenuReply)(nil),         // 0: api.role.v1.MenuReply
	(*CreateMenuRequest)(nil), // 1: api.role.v1.CreateMenuRequest
	(*UpdateMenuRequest)(nil), // 2: api.role.v1.UpdateMenuRequest
	(*DeleteMenuRequest)(nil), // 3: api.role.v1.DeleteMenuRequest
	(*GetMenuRequest)(nil),    // 4: api.role.v1.GetMenuRequest
	(*GetMenuReply)(nil),      // 5: api.role.v1.GetMenuReply
	(*ListMenuReply)(nil),     // 6: api.role.v1.ListMenuReply
}
var file_role_v1_menu_proto_depIdxs = []int32{
	5, // 0: api.role.v1.GetMenuReply.children:type_name -> api.role.v1.GetMenuReply
	5, // 1: api.role.v1.ListMenuReply.list:type_name -> api.role.v1.GetMenuReply
	1, // 2: api.role.v1.Menu.CreateMenu:input_type -> api.role.v1.CreateMenuRequest
	2, // 3: api.role.v1.Menu.UpdateMenu:input_type -> api.role.v1.UpdateMenuRequest
	3, // 4: api.role.v1.Menu.DeleteMenu:input_type -> api.role.v1.DeleteMenuRequest
	4, // 5: api.role.v1.Menu.GetMenu:input_type -> api.role.v1.GetMenuRequest
	4, // 6: api.role.v1.Menu.ListMenu:input_type -> api.role.v1.GetMenuRequest
	4, // 7: api.role.v1.Menu.ListMenuTree:input_type -> api.role.v1.GetMenuRequest
	0, // 8: api.role.v1.Menu.CreateMenu:output_type -> api.role.v1.MenuReply
	0, // 9: api.role.v1.Menu.UpdateMenu:output_type -> api.role.v1.MenuReply
	0, // 10: api.role.v1.Menu.DeleteMenu:output_type -> api.role.v1.MenuReply
	5, // 11: api.role.v1.Menu.GetMenu:output_type -> api.role.v1.GetMenuReply
	6, // 12: api.role.v1.Menu.ListMenu:output_type -> api.role.v1.ListMenuReply
	6, // 13: api.role.v1.Menu.ListMenuTree:output_type -> api.role.v1.ListMenuReply
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_role_v1_menu_proto_init() }
func file_role_v1_menu_proto_init() {
	if File_role_v1_menu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_role_v1_menu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuReply); i {
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
		file_role_v1_menu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMenuRequest); i {
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
		file_role_v1_menu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMenuRequest); i {
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
		file_role_v1_menu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMenuRequest); i {
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
		file_role_v1_menu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMenuRequest); i {
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
		file_role_v1_menu_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMenuReply); i {
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
		file_role_v1_menu_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMenuReply); i {
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
			RawDescriptor: file_role_v1_menu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_role_v1_menu_proto_goTypes,
		DependencyIndexes: file_role_v1_menu_proto_depIdxs,
		MessageInfos:      file_role_v1_menu_proto_msgTypes,
	}.Build()
	File_role_v1_menu_proto = out.File
	file_role_v1_menu_proto_rawDesc = nil
	file_role_v1_menu_proto_goTypes = nil
	file_role_v1_menu_proto_depIdxs = nil
}