// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: project/v1/brand.proto

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
	Brand_CreateBrand_FullMethodName = "/api.project.v1.Brand/CreateBrand"
	Brand_UpdateBrand_FullMethodName = "/api.project.v1.Brand/UpdateBrand"
	Brand_DeleteBrand_FullMethodName = "/api.project.v1.Brand/DeleteBrand"
	Brand_GetBrand_FullMethodName    = "/api.project.v1.Brand/GetBrand"
	Brand_ListBrand_FullMethodName   = "/api.project.v1.Brand/ListBrand"
)

// BrandClient is the client API for Brand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrandClient interface {
	CreateBrand(ctx context.Context, in *CreateBrandRequest, opts ...grpc.CallOption) (*CreateBrandReply, error)
	UpdateBrand(ctx context.Context, in *UpdateBrandRequest, opts ...grpc.CallOption) (*UpdateBrandReply, error)
	DeleteBrand(ctx context.Context, in *DeleteBrandRequest, opts ...grpc.CallOption) (*DeleteBrandReply, error)
	GetBrand(ctx context.Context, in *GetBrandRequest, opts ...grpc.CallOption) (*GetBrandReply, error)
	ListBrand(ctx context.Context, in *ListBrandRequest, opts ...grpc.CallOption) (*ListBrandReply, error)
}

type brandClient struct {
	cc grpc.ClientConnInterface
}

func NewBrandClient(cc grpc.ClientConnInterface) BrandClient {
	return &brandClient{cc}
}

func (c *brandClient) CreateBrand(ctx context.Context, in *CreateBrandRequest, opts ...grpc.CallOption) (*CreateBrandReply, error) {
	out := new(CreateBrandReply)
	err := c.cc.Invoke(ctx, Brand_CreateBrand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandClient) UpdateBrand(ctx context.Context, in *UpdateBrandRequest, opts ...grpc.CallOption) (*UpdateBrandReply, error) {
	out := new(UpdateBrandReply)
	err := c.cc.Invoke(ctx, Brand_UpdateBrand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandClient) DeleteBrand(ctx context.Context, in *DeleteBrandRequest, opts ...grpc.CallOption) (*DeleteBrandReply, error) {
	out := new(DeleteBrandReply)
	err := c.cc.Invoke(ctx, Brand_DeleteBrand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandClient) GetBrand(ctx context.Context, in *GetBrandRequest, opts ...grpc.CallOption) (*GetBrandReply, error) {
	out := new(GetBrandReply)
	err := c.cc.Invoke(ctx, Brand_GetBrand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandClient) ListBrand(ctx context.Context, in *ListBrandRequest, opts ...grpc.CallOption) (*ListBrandReply, error) {
	out := new(ListBrandReply)
	err := c.cc.Invoke(ctx, Brand_ListBrand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrandServer is the server API for Brand service.
// All implementations must embed UnimplementedBrandServer
// for forward compatibility
type BrandServer interface {
	CreateBrand(context.Context, *CreateBrandRequest) (*CreateBrandReply, error)
	UpdateBrand(context.Context, *UpdateBrandRequest) (*UpdateBrandReply, error)
	DeleteBrand(context.Context, *DeleteBrandRequest) (*DeleteBrandReply, error)
	GetBrand(context.Context, *GetBrandRequest) (*GetBrandReply, error)
	ListBrand(context.Context, *ListBrandRequest) (*ListBrandReply, error)
	mustEmbedUnimplementedBrandServer()
}

// UnimplementedBrandServer must be embedded to have forward compatible implementations.
type UnimplementedBrandServer struct {
}

func (UnimplementedBrandServer) CreateBrand(context.Context, *CreateBrandRequest) (*CreateBrandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBrand not implemented")
}
func (UnimplementedBrandServer) UpdateBrand(context.Context, *UpdateBrandRequest) (*UpdateBrandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBrand not implemented")
}
func (UnimplementedBrandServer) DeleteBrand(context.Context, *DeleteBrandRequest) (*DeleteBrandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBrand not implemented")
}
func (UnimplementedBrandServer) GetBrand(context.Context, *GetBrandRequest) (*GetBrandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrand not implemented")
}
func (UnimplementedBrandServer) ListBrand(context.Context, *ListBrandRequest) (*ListBrandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBrand not implemented")
}
func (UnimplementedBrandServer) mustEmbedUnimplementedBrandServer() {}

// UnsafeBrandServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrandServer will
// result in compilation errors.
type UnsafeBrandServer interface {
	mustEmbedUnimplementedBrandServer()
}

func RegisterBrandServer(s grpc.ServiceRegistrar, srv BrandServer) {
	s.RegisterService(&Brand_ServiceDesc, srv)
}

func _Brand_CreateBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServer).CreateBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Brand_CreateBrand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServer).CreateBrand(ctx, req.(*CreateBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brand_UpdateBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServer).UpdateBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Brand_UpdateBrand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServer).UpdateBrand(ctx, req.(*UpdateBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brand_DeleteBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServer).DeleteBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Brand_DeleteBrand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServer).DeleteBrand(ctx, req.(*DeleteBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brand_GetBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServer).GetBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Brand_GetBrand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServer).GetBrand(ctx, req.(*GetBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brand_ListBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrandServer).ListBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Brand_ListBrand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrandServer).ListBrand(ctx, req.(*ListBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Brand_ServiceDesc is the grpc.ServiceDesc for Brand service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Brand_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.project.v1.Brand",
	HandlerType: (*BrandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBrand",
			Handler:    _Brand_CreateBrand_Handler,
		},
		{
			MethodName: "UpdateBrand",
			Handler:    _Brand_UpdateBrand_Handler,
		},
		{
			MethodName: "DeleteBrand",
			Handler:    _Brand_DeleteBrand_Handler,
		},
		{
			MethodName: "GetBrand",
			Handler:    _Brand_GetBrand_Handler,
		},
		{
			MethodName: "ListBrand",
			Handler:    _Brand_ListBrand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "project/v1/brand.proto",
}
