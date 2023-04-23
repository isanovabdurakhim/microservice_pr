// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: store_service.proto

package organization_service

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

// StoreServiceClient is the client API for StoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreServiceClient interface {
	Create(ctx context.Context, in *CreateStore, opts ...grpc.CallOption) (*Store, error)
	GetByID(ctx context.Context, in *StorePrimaryKey, opts ...grpc.CallOption) (*Store, error)
	GetList(ctx context.Context, in *GetListStoreRequest, opts ...grpc.CallOption) (*GetListStoreResponse, error)
	Update(ctx context.Context, in *UpdateStore, opts ...grpc.CallOption) (*Store, error)
	Delete(ctx context.Context, in *StorePrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error)
}

type storeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreServiceClient(cc grpc.ClientConnInterface) StoreServiceClient {
	return &storeServiceClient{cc}
}

func (c *storeServiceClient) Create(ctx context.Context, in *CreateStore, opts ...grpc.CallOption) (*Store, error) {
	out := new(Store)
	err := c.cc.Invoke(ctx, "/organization_service.StoreService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) GetByID(ctx context.Context, in *StorePrimaryKey, opts ...grpc.CallOption) (*Store, error) {
	out := new(Store)
	err := c.cc.Invoke(ctx, "/organization_service.StoreService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) GetList(ctx context.Context, in *GetListStoreRequest, opts ...grpc.CallOption) (*GetListStoreResponse, error) {
	out := new(GetListStoreResponse)
	err := c.cc.Invoke(ctx, "/organization_service.StoreService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) Update(ctx context.Context, in *UpdateStore, opts ...grpc.CallOption) (*Store, error) {
	out := new(Store)
	err := c.cc.Invoke(ctx, "/organization_service.StoreService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) Delete(ctx context.Context, in *StorePrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/organization_service.StoreService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServiceServer is the server API for StoreService service.
// All implementations must embed UnimplementedStoreServiceServer
// for forward compatibility
type StoreServiceServer interface {
	Create(context.Context, *CreateStore) (*Store, error)
	GetByID(context.Context, *StorePrimaryKey) (*Store, error)
	GetList(context.Context, *GetListStoreRequest) (*GetListStoreResponse, error)
	Update(context.Context, *UpdateStore) (*Store, error)
	Delete(context.Context, *StorePrimaryKey) (*empty.Empty, error)
	mustEmbedUnimplementedStoreServiceServer()
}

// UnimplementedStoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStoreServiceServer struct {
}

func (UnimplementedStoreServiceServer) Create(context.Context, *CreateStore) (*Store, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStoreServiceServer) GetByID(context.Context, *StorePrimaryKey) (*Store, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedStoreServiceServer) GetList(context.Context, *GetListStoreRequest) (*GetListStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedStoreServiceServer) Update(context.Context, *UpdateStore) (*Store, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStoreServiceServer) Delete(context.Context, *StorePrimaryKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStoreServiceServer) mustEmbedUnimplementedStoreServiceServer() {}

// UnsafeStoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreServiceServer will
// result in compilation errors.
type UnsafeStoreServiceServer interface {
	mustEmbedUnimplementedStoreServiceServer()
}

func RegisterStoreServiceServer(s grpc.ServiceRegistrar, srv StoreServiceServer) {
	s.RegisterService(&StoreService_ServiceDesc, srv)
}

func _StoreService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.StoreService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).Create(ctx, req.(*CreateStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.StoreService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).GetByID(ctx, req.(*StorePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.StoreService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).GetList(ctx, req.(*GetListStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.StoreService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).Update(ctx, req.(*UpdateStore))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorePrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.StoreService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).Delete(ctx, req.(*StorePrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// StoreService_ServiceDesc is the grpc.ServiceDesc for StoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "organization_service.StoreService",
	HandlerType: (*StoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StoreService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _StoreService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _StoreService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StoreService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StoreService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "store_service.proto",
}