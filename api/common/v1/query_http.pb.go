// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             v3.19.4
// source: common/v1/query.proto

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

const OperationQueryGetCountryIso = "/api.captcha.v1.Query/GetCountryIso"

type QueryHTTPServer interface {
	// GetCountryIso 获取国家IOS
	GetCountryIso(context.Context, *GetCountryIsoRequest) (*GetCountryIsoReply, error)
}

func RegisterQueryHTTPServer(s *http.Server, srv QueryHTTPServer) {
	r := s.Route("/")
	r.GET("/common/query/v1/GetCountryIso", _Query_GetCountryIso0_HTTP_Handler(srv))
}

func _Query_GetCountryIso0_HTTP_Handler(srv QueryHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCountryIsoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationQueryGetCountryIso)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCountryIso(ctx, req.(*GetCountryIsoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCountryIsoReply)
		return ctx.Result(200, reply)
	}
}

type QueryHTTPClient interface {
	GetCountryIso(ctx context.Context, req *GetCountryIsoRequest, opts ...http.CallOption) (rsp *GetCountryIsoReply, err error)
}

type QueryHTTPClientImpl struct {
	cc *http.Client
}

func NewQueryHTTPClient(client *http.Client) QueryHTTPClient {
	return &QueryHTTPClientImpl{client}
}

func (c *QueryHTTPClientImpl) GetCountryIso(ctx context.Context, in *GetCountryIsoRequest, opts ...http.CallOption) (*GetCountryIsoReply, error) {
	var out GetCountryIsoReply
	pattern := "/common/query/v1/GetCountryIso"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationQueryGetCountryIso))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
