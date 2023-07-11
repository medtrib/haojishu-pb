// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: project/v1/project.proto

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
	Project_CreateProject_FullMethodName = "/api.project.v1.Project/CreateProject"
	Project_UpdateProject_FullMethodName = "/api.project.v1.Project/UpdateProject"
	Project_DeleteProject_FullMethodName = "/api.project.v1.Project/DeleteProject"
	Project_GetProject_FullMethodName    = "/api.project.v1.Project/GetProject"
	Project_ListProject_FullMethodName   = "/api.project.v1.Project/ListProject"
)

// ProjectClient is the client API for Project service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectClient interface {
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectReply, error)
	UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*UpdateProjectReply, error)
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectReply, error)
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectReply, error)
	ListProject(ctx context.Context, in *ListProjectRequest, opts ...grpc.CallOption) (*ListProjectReply, error)
}

type projectClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectClient(cc grpc.ClientConnInterface) ProjectClient {
	return &projectClient{cc}
}

func (c *projectClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectReply, error) {
	out := new(CreateProjectReply)
	err := c.cc.Invoke(ctx, Project_CreateProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*UpdateProjectReply, error) {
	out := new(UpdateProjectReply)
	err := c.cc.Invoke(ctx, Project_UpdateProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectReply, error) {
	out := new(DeleteProjectReply)
	err := c.cc.Invoke(ctx, Project_DeleteProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectReply, error) {
	out := new(GetProjectReply)
	err := c.cc.Invoke(ctx, Project_GetProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectClient) ListProject(ctx context.Context, in *ListProjectRequest, opts ...grpc.CallOption) (*ListProjectReply, error) {
	out := new(ListProjectReply)
	err := c.cc.Invoke(ctx, Project_ListProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServer is the server API for Project service.
// All implementations must embed UnimplementedProjectServer
// for forward compatibility
type ProjectServer interface {
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectReply, error)
	UpdateProject(context.Context, *UpdateProjectRequest) (*UpdateProjectReply, error)
	DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectReply, error)
	GetProject(context.Context, *GetProjectRequest) (*GetProjectReply, error)
	ListProject(context.Context, *ListProjectRequest) (*ListProjectReply, error)
	mustEmbedUnimplementedProjectServer()
}

// UnimplementedProjectServer must be embedded to have forward compatible implementations.
type UnimplementedProjectServer struct {
}

func (UnimplementedProjectServer) CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedProjectServer) UpdateProject(context.Context, *UpdateProjectRequest) (*UpdateProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProject not implemented")
}
func (UnimplementedProjectServer) DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedProjectServer) GetProject(context.Context, *GetProjectRequest) (*GetProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedProjectServer) ListProject(context.Context, *ListProjectRequest) (*ListProjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProject not implemented")
}
func (UnimplementedProjectServer) mustEmbedUnimplementedProjectServer() {}

// UnsafeProjectServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServer will
// result in compilation errors.
type UnsafeProjectServer interface {
	mustEmbedUnimplementedProjectServer()
}

func RegisterProjectServer(s grpc.ServiceRegistrar, srv ProjectServer) {
	s.RegisterService(&Project_ServiceDesc, srv)
}

func _Project_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Project_CreateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_UpdateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).UpdateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Project_UpdateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).UpdateProject(ctx, req.(*UpdateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Project_DeleteProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).DeleteProject(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Project_GetProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Project_ListProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServer).ListProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Project_ListProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServer).ListProject(ctx, req.(*ListProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Project_ServiceDesc is the grpc.ServiceDesc for Project service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Project_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.project.v1.Project",
	HandlerType: (*ProjectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _Project_CreateProject_Handler,
		},
		{
			MethodName: "UpdateProject",
			Handler:    _Project_UpdateProject_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _Project_DeleteProject_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _Project_GetProject_Handler,
		},
		{
			MethodName: "ListProject",
			Handler:    _Project_ListProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "project/v1/project.proto",
}
