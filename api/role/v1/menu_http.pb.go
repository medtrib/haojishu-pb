// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.2
// - protoc             v3.19.4
// source: role/v1/menu.proto

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

const OperationMenuCreateMenu = "/api.role.v1.Menu/CreateMenu"
const OperationMenuDeleteMenu = "/api.role.v1.Menu/DeleteMenu"
const OperationMenuGetMenu = "/api.role.v1.Menu/GetMenu"
const OperationMenuListMenu = "/api.role.v1.Menu/ListMenu"
const OperationMenuListMenuTree = "/api.role.v1.Menu/ListMenuTree"
const OperationMenuUpdateMenu = "/api.role.v1.Menu/UpdateMenu"

type MenuHTTPServer interface {
	// CreateMenu CreateMenu 创建菜单
	CreateMenu(context.Context, *CreateMenuRequest) (*MenuReply, error)
	// DeleteMenu DeleteMenu 删除菜单
	DeleteMenu(context.Context, *DeleteMenuRequest) (*MenuReply, error)
	// GetMenu GetMenu 查询单挑菜单信息
	GetMenu(context.Context, *GetMenuRequest) (*GetMenuReply, error)
	// ListMenu ListMenu 获取菜单列表
	ListMenu(context.Context, *GetMenuRequest) (*ListMenuReply, error)
	// ListMenuTree ListMenuTree 获取树形列表
	ListMenuTree(context.Context, *GetMenuRequest) (*ListMenuReply, error)
	// UpdateMenu UpdateMenu 编辑菜单
	UpdateMenu(context.Context, *UpdateMenuRequest) (*MenuReply, error)
}

func RegisterMenuHTTPServer(s *http.Server, srv MenuHTTPServer) {
	r := s.Route("/")
	r.POST("/admin/role/v1/CreateMenu", _Menu_CreateMenu0_HTTP_Handler(srv))
	r.POST("/admin/role/v1/UpdateMenu", _Menu_UpdateMenu0_HTTP_Handler(srv))
	r.GET("/admin/role/v1/DeleteMenu", _Menu_DeleteMenu0_HTTP_Handler(srv))
	r.GET("/admin/role/v1/GetMenu", _Menu_GetMenu0_HTTP_Handler(srv))
	r.GET("/admin/role/v1/ListMenu", _Menu_ListMenu0_HTTP_Handler(srv))
	r.GET("/admin/role/v1/ListMenuTree", _Menu_ListMenuTree0_HTTP_Handler(srv))
}

func _Menu_CreateMenu0_HTTP_Handler(srv MenuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateMenuRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuCreateMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateMenu(ctx, req.(*CreateMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MenuReply)
		return ctx.Result(200, reply)
	}
}

func _Menu_UpdateMenu0_HTTP_Handler(srv MenuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateMenuRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuUpdateMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateMenu(ctx, req.(*UpdateMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MenuReply)
		return ctx.Result(200, reply)
	}
}

func _Menu_DeleteMenu0_HTTP_Handler(srv MenuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteMenuRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuDeleteMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteMenu(ctx, req.(*DeleteMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MenuReply)
		return ctx.Result(200, reply)
	}
}

func _Menu_GetMenu0_HTTP_Handler(srv MenuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMenuRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuGetMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMenu(ctx, req.(*GetMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetMenuReply)
		return ctx.Result(200, reply)
	}
}

func _Menu_ListMenu0_HTTP_Handler(srv MenuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMenuRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuListMenu)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMenu(ctx, req.(*GetMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMenuReply)
		return ctx.Result(200, reply)
	}
}

func _Menu_ListMenuTree0_HTTP_Handler(srv MenuHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMenuRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMenuListMenuTree)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMenuTree(ctx, req.(*GetMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMenuReply)
		return ctx.Result(200, reply)
	}
}

type MenuHTTPClient interface {
	CreateMenu(ctx context.Context, req *CreateMenuRequest, opts ...http.CallOption) (rsp *MenuReply, err error)
	DeleteMenu(ctx context.Context, req *DeleteMenuRequest, opts ...http.CallOption) (rsp *MenuReply, err error)
	GetMenu(ctx context.Context, req *GetMenuRequest, opts ...http.CallOption) (rsp *GetMenuReply, err error)
	ListMenu(ctx context.Context, req *GetMenuRequest, opts ...http.CallOption) (rsp *ListMenuReply, err error)
	ListMenuTree(ctx context.Context, req *GetMenuRequest, opts ...http.CallOption) (rsp *ListMenuReply, err error)
	UpdateMenu(ctx context.Context, req *UpdateMenuRequest, opts ...http.CallOption) (rsp *MenuReply, err error)
}

type MenuHTTPClientImpl struct {
	cc *http.Client
}

func NewMenuHTTPClient(client *http.Client) MenuHTTPClient {
	return &MenuHTTPClientImpl{client}
}

func (c *MenuHTTPClientImpl) CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...http.CallOption) (*MenuReply, error) {
	var out MenuReply
	pattern := "/admin/role/v1/CreateMenu"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMenuCreateMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuHTTPClientImpl) DeleteMenu(ctx context.Context, in *DeleteMenuRequest, opts ...http.CallOption) (*MenuReply, error) {
	var out MenuReply
	pattern := "/admin/role/v1/DeleteMenu"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenuDeleteMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuHTTPClientImpl) GetMenu(ctx context.Context, in *GetMenuRequest, opts ...http.CallOption) (*GetMenuReply, error) {
	var out GetMenuReply
	pattern := "/admin/role/v1/GetMenu"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenuGetMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuHTTPClientImpl) ListMenu(ctx context.Context, in *GetMenuRequest, opts ...http.CallOption) (*ListMenuReply, error) {
	var out ListMenuReply
	pattern := "/admin/role/v1/ListMenu"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenuListMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuHTTPClientImpl) ListMenuTree(ctx context.Context, in *GetMenuRequest, opts ...http.CallOption) (*ListMenuReply, error) {
	var out ListMenuReply
	pattern := "/admin/role/v1/ListMenuTree"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMenuListMenuTree))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuHTTPClientImpl) UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...http.CallOption) (*MenuReply, error) {
	var out MenuReply
	pattern := "/admin/role/v1/UpdateMenu"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationMenuUpdateMenu))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
