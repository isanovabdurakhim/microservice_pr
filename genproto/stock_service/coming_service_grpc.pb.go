// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: coming_service.proto

package stock_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ComingServiceClient is the client API for ComingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComingServiceClient interface {
	Create(ctx context.Context, in *CreateComing, opts ...grpc.CallOption) (*Coming, error)
	GetById(ctx context.Context, in *ComingPrimaryKey, opts ...grpc.CallOption) (*Coming, error)
	GetList(ctx context.Context, in *GetListComingRequest, opts ...grpc.CallOption) (*GetListComingResponse, error)
	Update(ctx context.Context, in *UpdateComing, opts ...grpc.CallOption) (*Coming, error)
	Delete(ctx context.Context, in *ComingPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error)
}

type comingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComingServiceClient(cc grpc.ClientConnInterface) ComingServiceClient {
	return &comingServiceClient{cc}
}

func (c *comingServiceClient) Create(ctx context.Context, in *CreateComing, opts ...grpc.CallOption) (*Coming, error) {
	out := new(Coming)
	err := c.cc.Invoke(ctx, "/stock_service.ComingService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comingServiceClient) GetById(ctx context.Context, in *ComingPrimaryKey, opts ...grpc.CallOption) (*Coming, error) {
	out := new(Coming)
	err := c.cc.Invoke(ctx, "/stock_service.ComingService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comingServiceClient) GetList(ctx context.Context, in *GetListComingRequest, opts ...grpc.CallOption) (*GetListComingResponse, error) {
	out := new(GetListComingResponse)
	err := c.cc.Invoke(ctx, "/stock_service.ComingService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comingServiceClient) Update(ctx context.Context, in *UpdateComing, opts ...grpc.CallOption) (*Coming, error) {
	out := new(Coming)
	err := c.cc.Invoke(ctx, "/stock_service.ComingService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comingServiceClient) Delete(ctx context.Context, in *ComingPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/stock_service.ComingService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComingServiceServer is the server API for ComingService service.
// All implementations must embed UnimplementedComingServiceServer
// for forward compatibility
type ComingServiceServer interface {
	Create(context.Context, *CreateComing) (*Coming, error)
	GetById(context.Context, *ComingPrimaryKey) (*Coming, error)
	GetList(context.Context, *GetListComingRequest) (*GetListComingResponse, error)
	Update(context.Context, *UpdateComing) (*Coming, error)
	Delete(context.Context, *ComingPrimaryKey) (*empty.Empty, error)
	mustEmbedUnimplementedComingServiceServer()
}

// UnimplementedComingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedComingServiceServer struct {
}

func (UnimplementedComingServiceServer) Create(context.Context, *CreateComing) (*Coming, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedComingServiceServer) GetById(context.Context, *ComingPrimaryKey) (*Coming, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedComingServiceServer) GetList(context.Context, *GetListComingRequest) (*GetListComingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedComingServiceServer) Update(context.Context, *UpdateComing) (*Coming, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedComingServiceServer) Delete(context.Context, *ComingPrimaryKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedComingServiceServer) mustEmbedUnimplementedComingServiceServer() {}

// UnsafeComingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComingServiceServer will
// result in compilation errors.
type UnsafeComingServiceServer interface {
	mustEmbedUnimplementedComingServiceServer()
}

func RegisterComingServiceServer(s grpc.ServiceRegistrar, srv ComingServiceServer) {
	s.RegisterService(&ComingService_ServiceDesc, srv)
}

func _ComingService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateComing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComingServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock_service.ComingService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComingServiceServer).Create(ctx, req.(*CreateComing))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComingService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComingPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComingServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock_service.ComingService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComingServiceServer).GetById(ctx, req.(*ComingPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComingService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListComingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComingServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock_service.ComingService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComingServiceServer).GetList(ctx, req.(*GetListComingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComingService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateComing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComingServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock_service.ComingService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComingServiceServer).Update(ctx, req.(*UpdateComing))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComingService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComingPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComingServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stock_service.ComingService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComingServiceServer).Delete(ctx, req.(*ComingPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// ComingService_ServiceDesc is the grpc.ServiceDesc for ComingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stock_service.ComingService",
	HandlerType: (*ComingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ComingService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _ComingService_GetById_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _ComingService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ComingService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ComingService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coming_service.proto",
}