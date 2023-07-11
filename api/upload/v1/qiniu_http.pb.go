// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             v3.19.4
// source: upload/v1/qiniu.proto

package v1

import (
	context "context"

	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationQiNiuUpToken = "/api.upload.v1.QiNiu/UpToken"
const OperationQiNiuUploadFileBase = "/api.upload.v1.QiNiu/UploadFileBase"

type QiNiuHTTPServer interface {
	UpToken(context.Context, *UpTokenRequest) (*UpTokenRequestReply, error)
	UploadFileBase(context.Context, *UploadFileBaseRequest) (*UploadFileBaseReply, error)
}

func RegisterQiNiuHTTPServer(s *http.Server, srv QiNiuHTTPServer) {
	r := s.Route("/")
	r.GET("/api/v1/UpToken", _QiNiu_UpToken0_HTTP_Handler(srv))
	r.POST("/api/v1/UploadFileBase", _QiNiu_UploadFileBase0_HTTP_Handler(srv))
}

func _QiNiu_UpToken0_HTTP_Handler(srv QiNiuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpTokenRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQiNiuUpToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpToken(ctx, req.(*UpTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpTokenRequestReply)
		return ctx.Result(200, reply)
	}
}

func _QiNiu_UploadFileBase0_HTTP_Handler(srv QiNiuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UploadFileBaseRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQiNiuUploadFileBase)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UploadFileBase(ctx, req.(*UploadFileBaseRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UploadFileBaseReply)
		return ctx.Result(200, reply)
	}
}

type QiNiuHTTPClient interface {
	UpToken(ctx context.Context, req *UpTokenRequest, opts ...http.CallOption) (rsp *UpTokenRequestReply, err error)
	UploadFileBase(ctx context.Context, req *UploadFileBaseRequest, opts ...http.CallOption) (rsp *UploadFileBaseReply, err error)
}

type QiNiuHTTPClientImpl struct {
	cc *http.Client
}

func NewQiNiuHTTPClient(client *http.Client) QiNiuHTTPClient {
	return &QiNiuHTTPClientImpl{client}
}

func (c *QiNiuHTTPClientImpl) UpToken(ctx context.Context, in *UpTokenRequest, opts ...http.CallOption) (*UpTokenRequestReply, error) {
	var out UpTokenRequestReply
	pattern := "/api/v1/UpToken"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQiNiuUpToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *QiNiuHTTPClientImpl) UploadFileBase(ctx context.Context, in *UploadFileBaseRequest, opts ...http.CallOption) (*UploadFileBaseReply, error) {
	var out UploadFileBaseReply
	pattern := "/api/v1/UploadFileBase"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationQiNiuUploadFileBase))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
