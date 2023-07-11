// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: captcha/v1/captcha.proto

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
	Captcha_GetCaptcha_FullMethodName    = "/api.captcha.v1.Captcha/GetCaptcha"
	Captcha_VerifyCaptcha_FullMethodName = "/api.captcha.v1.Captcha/VerifyCaptcha"
)

// CaptchaClient is the client API for Captcha service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaptchaClient interface {
	// 获取验证码
	GetCaptcha(ctx context.Context, in *GetCaptchaRequest, opts ...grpc.CallOption) (*GetCaptchaReply, error)
	// 验证验证码
	VerifyCaptcha(ctx context.Context, in *VerifyCaptchaRequest, opts ...grpc.CallOption) (*VerifyCaptchaReply, error)
}

type captchaClient struct {
	cc grpc.ClientConnInterface
}

func NewCaptchaClient(cc grpc.ClientConnInterface) CaptchaClient {
	return &captchaClient{cc}
}

func (c *captchaClient) GetCaptcha(ctx context.Context, in *GetCaptchaRequest, opts ...grpc.CallOption) (*GetCaptchaReply, error) {
	out := new(GetCaptchaReply)
	err := c.cc.Invoke(ctx, Captcha_GetCaptcha_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captchaClient) VerifyCaptcha(ctx context.Context, in *VerifyCaptchaRequest, opts ...grpc.CallOption) (*VerifyCaptchaReply, error) {
	out := new(VerifyCaptchaReply)
	err := c.cc.Invoke(ctx, Captcha_VerifyCaptcha_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaptchaServer is the server API for Captcha service.
// All implementations must embed UnimplementedCaptchaServer
// for forward compatibility
type CaptchaServer interface {
	// 获取验证码
	GetCaptcha(context.Context, *GetCaptchaRequest) (*GetCaptchaReply, error)
	// 验证验证码
	VerifyCaptcha(context.Context, *VerifyCaptchaRequest) (*VerifyCaptchaReply, error)
	mustEmbedUnimplementedCaptchaServer()
}

// UnimplementedCaptchaServer must be embedded to have forward compatible implementations.
type UnimplementedCaptchaServer struct {
}

func (UnimplementedCaptchaServer) GetCaptcha(context.Context, *GetCaptchaRequest) (*GetCaptchaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCaptcha not implemented")
}
func (UnimplementedCaptchaServer) VerifyCaptcha(context.Context, *VerifyCaptchaRequest) (*VerifyCaptchaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCaptcha not implemented")
}
func (UnimplementedCaptchaServer) mustEmbedUnimplementedCaptchaServer() {}

// UnsafeCaptchaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaptchaServer will
// result in compilation errors.
type UnsafeCaptchaServer interface {
	mustEmbedUnimplementedCaptchaServer()
}

func RegisterCaptchaServer(s grpc.ServiceRegistrar, srv CaptchaServer) {
	s.RegisterService(&Captcha_ServiceDesc, srv)
}

func _Captcha_GetCaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptchaServer).GetCaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Captcha_GetCaptcha_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptchaServer).GetCaptcha(ctx, req.(*GetCaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Captcha_VerifyCaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptchaServer).VerifyCaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Captcha_VerifyCaptcha_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptchaServer).VerifyCaptcha(ctx, req.(*VerifyCaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Captcha_ServiceDesc is the grpc.ServiceDesc for Captcha service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Captcha_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.captcha.v1.Captcha",
	HandlerType: (*CaptchaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCaptcha",
			Handler:    _Captcha_GetCaptcha_Handler,
		},
		{
			MethodName: "VerifyCaptcha",
			Handler:    _Captcha_VerifyCaptcha_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "captcha/v1/captcha.proto",
}
