// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: api/pay/v1/pay.proto

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

type CreatePayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 订单前缀
	OrderPrefix string `protobuf:"bytes,1,opt,name=orderPrefix,proto3" json:"orderPrefix,omitempty"`
	// 用户标识
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// 用户所属国家
	Country string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	// 订单金额
	Amount float64 `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	// 货币
	Currency string `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	// 返回Url
	ReturnUrl string `protobuf:"bytes,6,opt,name=returnUrl,proto3" json:"returnUrl,omitempty"`
	// 支付参数
	Pay *PayOption `protobuf:"bytes,7,opt,name=pay,proto3" json:"pay,omitempty"`
}

func (x *CreatePayRequest) Reset() {
	*x = CreatePayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pay_v1_pay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePayRequest) ProtoMessage() {}

func (x *CreatePayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_pay_v1_pay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePayRequest.ProtoReflect.Descriptor instead.
func (*CreatePayRequest) Descriptor() ([]byte, []int) {
	return file_api_pay_v1_pay_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePayRequest) GetOrderPrefix() string {
	if x != nil {
		return x.OrderPrefix
	}
	return ""
}

func (x *CreatePayRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreatePayRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CreatePayRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreatePayRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *CreatePayRequest) GetReturnUrl() string {
	if x != nil {
		return x.ReturnUrl
	}
	return ""
}

func (x *CreatePayRequest) GetPay() *PayOption {
	if x != nil {
		return x.Pay
	}
	return nil
}

// 支付参数
type PayOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 支付类型
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// 支付参数
	Parameter string `protobuf:"bytes,2,opt,name=parameter,proto3" json:"parameter,omitempty"`
	// 商品参数
	Shop []*Shop `protobuf:"bytes,3,rep,name=shop,proto3" json:"shop,omitempty"`
	// 简介
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *PayOption) Reset() {
	*x = PayOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pay_v1_pay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayOption) ProtoMessage() {}

func (x *PayOption) ProtoReflect() protoreflect.Message {
	mi := &file_api_pay_v1_pay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayOption.ProtoReflect.Descriptor instead.
func (*PayOption) Descriptor() ([]byte, []int) {
	return file_api_pay_v1_pay_proto_rawDescGZIP(), []int{1}
}

func (x *PayOption) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PayOption) GetParameter() string {
	if x != nil {
		return x.Parameter
	}
	return ""
}

func (x *PayOption) GetShop() []*Shop {
	if x != nil {
		return x.Shop
	}
	return nil
}

func (x *PayOption) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// 商品参数
type Shop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price    string `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
	Num      string `protobuf:"bytes,3,opt,name=num,proto3" json:"num,omitempty"`
	Currency string `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *Shop) Reset() {
	*x = Shop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pay_v1_pay_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shop) ProtoMessage() {}

func (x *Shop) ProtoReflect() protoreflect.Message {
	mi := &file_api_pay_v1_pay_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shop.ProtoReflect.Descriptor instead.
func (*Shop) Descriptor() ([]byte, []int) {
	return file_api_pay_v1_pay_proto_rawDescGZIP(), []int{2}
}

func (x *Shop) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Shop) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Shop) GetNum() string {
	if x != nil {
		return x.Num
	}
	return ""
}

func (x *Shop) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type CreatePayReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Url     string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreatePayReply) Reset() {
	*x = CreatePayReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pay_v1_pay_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePayReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePayReply) ProtoMessage() {}

func (x *CreatePayReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_pay_v1_pay_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePayReply.ProtoReflect.Descriptor instead.
func (*CreatePayReply) Descriptor() ([]byte, []int) {
	return file_api_pay_v1_pay_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePayReply) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreatePayReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetPayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPayRequest) Reset() {
	*x = GetPayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pay_v1_pay_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPayRequest) ProtoMessage() {}

func (x *GetPayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_pay_v1_pay_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPayRequest.ProtoReflect.Descriptor instead.
func (*GetPayRequest) Descriptor() ([]byte, []int) {
	return file_api_pay_v1_pay_proto_rawDescGZIP(), []int{4}
}

type GetPayReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPayReply) Reset() {
	*x = GetPayReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pay_v1_pay_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPayReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPayReply) ProtoMessage() {}

func (x *GetPayReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_pay_v1_pay_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPayReply.ProtoReflect.Descriptor instead.
func (*GetPayReply) Descriptor() ([]byte, []int) {
	return file_api_pay_v1_pay_proto_rawDescGZIP(), []int{5}
}

var File_api_pay_v1_pay_proto protoreflect.FileDescriptor

var file_api_pay_v1_pay_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdf, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x27, 0x0a, 0x03, 0x70, 0x61, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x70,
	0x61, 0x79, 0x22, 0x85, 0x01, 0x0a, 0x09, 0x50, 0x61, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68,
	0x6f, 0x70, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5e, 0x0a, 0x04, 0x53, 0x68,
	0x6f, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x3c, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xaf, 0x01, 0x0a, 0x03, 0x50, 0x61, 0x79,
	0x12, 0x6a, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x12, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a,
	0x01, 0x2a, 0x22, 0x18, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x12, 0x3c, 0x0a, 0x06,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x3c, 0x0a, 0x0a, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x64, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x68,
	0x61, 0x6f, 0x6a, 0x69, 0x73, 0x68, 0x75, 0x2d, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x61, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_pay_v1_pay_proto_rawDescOnce sync.Once
	file_api_pay_v1_pay_proto_rawDescData = file_api_pay_v1_pay_proto_rawDesc
)

func file_api_pay_v1_pay_proto_rawDescGZIP() []byte {
	file_api_pay_v1_pay_proto_rawDescOnce.Do(func() {
		file_api_pay_v1_pay_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_pay_v1_pay_proto_rawDescData)
	})
	return file_api_pay_v1_pay_proto_rawDescData
}

var file_api_pay_v1_pay_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_pay_v1_pay_proto_goTypes = []interface{}{
	(*CreatePayRequest)(nil), // 0: api.pay.v1.CreatePayRequest
	(*PayOption)(nil),        // 1: api.pay.v1.PayOption
	(*Shop)(nil),             // 2: api.pay.v1.Shop
	(*CreatePayReply)(nil),   // 3: api.pay.v1.CreatePayReply
	(*GetPayRequest)(nil),    // 4: api.pay.v1.GetPayRequest
	(*GetPayReply)(nil),      // 5: api.pay.v1.GetPayReply
}
var file_api_pay_v1_pay_proto_depIdxs = []int32{
	1, // 0: api.pay.v1.CreatePayRequest.pay:type_name -> api.pay.v1.PayOption
	2, // 1: api.pay.v1.PayOption.shop:type_name -> api.pay.v1.Shop
	0, // 2: api.pay.v1.Pay.CreatePay:input_type -> api.pay.v1.CreatePayRequest
	4, // 3: api.pay.v1.Pay.GetPay:input_type -> api.pay.v1.GetPayRequest
	3, // 4: api.pay.v1.Pay.CreatePay:output_type -> api.pay.v1.CreatePayReply
	5, // 5: api.pay.v1.Pay.GetPay:output_type -> api.pay.v1.GetPayReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_pay_v1_pay_proto_init() }
func file_api_pay_v1_pay_proto_init() {
	if File_api_pay_v1_pay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_pay_v1_pay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePayRequest); i {
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
		file_api_pay_v1_pay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayOption); i {
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
		file_api_pay_v1_pay_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shop); i {
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
		file_api_pay_v1_pay_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePayReply); i {
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
		file_api_pay_v1_pay_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPayRequest); i {
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
		file_api_pay_v1_pay_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPayReply); i {
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
			RawDescriptor: file_api_pay_v1_pay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_pay_v1_pay_proto_goTypes,
		DependencyIndexes: file_api_pay_v1_pay_proto_depIdxs,
		MessageInfos:      file_api_pay_v1_pay_proto_msgTypes,
	}.Build()
	File_api_pay_v1_pay_proto = out.File
	file_api_pay_v1_pay_proto_rawDesc = nil
	file_api_pay_v1_pay_proto_goTypes = nil
	file_api_pay_v1_pay_proto_depIdxs = nil
}
