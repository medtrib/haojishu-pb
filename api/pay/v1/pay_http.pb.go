// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             v3.19.4
// source: pay/v1/pay.proto

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

const OperationPayCreatePay = "/api.pay.v1.Pay/CreatePay"

type PayHTTPServer interface {
	CreatePay(context.Context, *CreatePayRequest) (*CreatePayReply, error)
}

func RegisterPayHTTPServer(s *http.Server, srv PayHTTPServer) {
	r := s.Route("/")
	r.POST("/pay/v1/CreatePay", _Pay_CreatePay0_HTTP_Handler(srv))
}

func _Pay_CreatePay0_HTTP_Handler(srv PayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePayRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPayCreatePay)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePay(ctx, req.(*CreatePayRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreatePayReply)
		return ctx.Result(200, reply)
	}
}

type PayHTTPClient interface {
	CreatePay(ctx context.Context, req *CreatePayRequest, opts ...http.CallOption) (rsp *CreatePayReply, err error)
}

type PayHTTPClientImpl struct {
	cc *http.Client
}

func NewPayHTTPClient(client *http.Client) PayHTTPClient {
	return &PayHTTPClientImpl{client}
}

func (c *PayHTTPClientImpl) CreatePay(ctx context.Context, in *CreatePayRequest, opts ...http.CallOption) (*CreatePayReply, error) {
	var out CreatePayReply
	pattern := "/pay/v1/CreatePay"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPayCreatePay))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
