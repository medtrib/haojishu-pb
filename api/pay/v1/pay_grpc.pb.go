// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: pay/v1/pay.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Pay_CreatePay_FullMethodName = "/api.pay.v1.Pay/CreatePay"
	Pay_GetPay_FullMethodName    = "/api.pay.v1.Pay/GetPay"
)

// PayClient is the client API for Pay service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PayClient interface {
	CreatePay(ctx context.Context, in *CreatePayRequest, opts ...grpc.CallOption) (*CreatePayReply, error)
	GetPay(ctx context.Context, in *GetPayRequest, opts ...grpc.CallOption) (*GetPayReply, error)
}

type payClient struct {
	cc grpc.ClientConnInterface
}

func NewPayClient(cc grpc.ClientConnInterface) PayClient {
	return &payClient{cc}
}

func (c *payClient) CreatePay(ctx context.Context, in *CreatePayRequest, opts ...grpc.CallOption) (*CreatePayReply, error) {
	out := new(CreatePayReply)
	err := c.cc.Invoke(ctx, Pay_CreatePay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payClient) GetPay(ctx context.Context, in *GetPayRequest, opts ...grpc.CallOption) (*GetPayReply, error) {
	out := new(GetPayReply)
	err := c.cc.Invoke(ctx, Pay_GetPay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayServer is the server API for Pay service.
// All implementations must embed UnimplementedPayServer
// for forward compatibility
type PayServer interface {
	CreatePay(context.Context, *CreatePayRequest) (*CreatePayReply, error)
	GetPay(context.Context, *GetPayRequest) (*GetPayReply, error)
	mustEmbedUnimplementedPayServer()
}

// UnimplementedPayServer must be embedded to have forward compatible implementations.
type UnimplementedPayServer struct {
}

func (UnimplementedPayServer) CreatePay(context.Context, *CreatePayRequest) (*CreatePayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePay not implemented")
}
func (UnimplementedPayServer) GetPay(context.Context, *GetPayRequest) (*GetPayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPay not implemented")
}
func (UnimplementedPayServer) mustEmbedUnimplementedPayServer() {}

// UnsafePayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayServer will
// result in compilation errors.
type UnsafePayServer interface {
	mustEmbedUnimplementedPayServer()
}

func RegisterPayServer(s grpc.ServiceRegistrar, srv PayServer) {
	s.RegisterService(&Pay_ServiceDesc, srv)
}

func _Pay_CreatePay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).CreatePay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pay_CreatePay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).CreatePay(ctx, req.(*CreatePayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pay_GetPay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServer).GetPay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pay_GetPay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServer).GetPay(ctx, req.(*GetPayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Pay_ServiceDesc is the grpc.ServiceDesc for Pay service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pay_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.pay.v1.Pay",
	HandlerType: (*PayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePay",
			Handler:    _Pay_CreatePay_Handler,
		},
		{
			MethodName: "GetPay",
			Handler:    _Pay_GetPay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pay/v1/pay.proto",
}