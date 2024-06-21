// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: branch.proto

package user_service

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

// BranchServiceClient is the client API for BranchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BranchServiceClient interface {
	Create(ctx context.Context, in *CreateBranch, opts ...grpc.CallOption) (*Branch, error)
	GetByID(ctx context.Context, in *BranchPrimaryKey, opts ...grpc.CallOption) (*Branch, error)
	GetList(ctx context.Context, in *GetListBranchRequest, opts ...grpc.CallOption) (*GetListBranchResponse, error)
	Update(ctx context.Context, in *UpdateBranch, opts ...grpc.CallOption) (*Branch, error)
	Delete(ctx context.Context, in *BranchPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error)
}

type branchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBranchServiceClient(cc grpc.ClientConnInterface) BranchServiceClient {
	return &branchServiceClient{cc}
}

func (c *branchServiceClient) Create(ctx context.Context, in *CreateBranch, opts ...grpc.CallOption) (*Branch, error) {
	out := new(Branch)
	err := c.cc.Invoke(ctx, "/user_service.BranchService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) GetByID(ctx context.Context, in *BranchPrimaryKey, opts ...grpc.CallOption) (*Branch, error) {
	out := new(Branch)
	err := c.cc.Invoke(ctx, "/user_service.BranchService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) GetList(ctx context.Context, in *GetListBranchRequest, opts ...grpc.CallOption) (*GetListBranchResponse, error) {
	out := new(GetListBranchResponse)
	err := c.cc.Invoke(ctx, "/user_service.BranchService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) Update(ctx context.Context, in *UpdateBranch, opts ...grpc.CallOption) (*Branch, error) {
	out := new(Branch)
	err := c.cc.Invoke(ctx, "/user_service.BranchService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchServiceClient) Delete(ctx context.Context, in *BranchPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/user_service.BranchService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BranchServiceServer is the server API for BranchService service.
// All implementations must embed UnimplementedBranchServiceServer
// for forward compatibility
type BranchServiceServer interface {
	Create(context.Context, *CreateBranch) (*Branch, error)
	GetByID(context.Context, *BranchPrimaryKey) (*Branch, error)
	GetList(context.Context, *GetListBranchRequest) (*GetListBranchResponse, error)
	Update(context.Context, *UpdateBranch) (*Branch, error)
	Delete(context.Context, *BranchPrimaryKey) (*empty.Empty, error)
	mustEmbedUnimplementedBranchServiceServer()
}

// UnimplementedBranchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBranchServiceServer struct {
}

func (UnimplementedBranchServiceServer) Create(context.Context, *CreateBranch) (*Branch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBranchServiceServer) GetByID(context.Context, *BranchPrimaryKey) (*Branch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedBranchServiceServer) GetList(context.Context, *GetListBranchRequest) (*GetListBranchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedBranchServiceServer) Update(context.Context, *UpdateBranch) (*Branch, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBranchServiceServer) Delete(context.Context, *BranchPrimaryKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBranchServiceServer) mustEmbedUnimplementedBranchServiceServer() {}

// UnsafeBranchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BranchServiceServer will
// result in compilation errors.
type UnsafeBranchServiceServer interface {
	mustEmbedUnimplementedBranchServiceServer()
}

func RegisterBranchServiceServer(s grpc.ServiceRegistrar, srv BranchServiceServer) {
	s.RegisterService(&BranchService_ServiceDesc, srv)
}

func _BranchService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBranch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.BranchService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).Create(ctx, req.(*CreateBranch))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BranchPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.BranchService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).GetByID(ctx, req.(*BranchPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.BranchService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).GetList(ctx, req.(*GetListBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBranch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.BranchService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).Update(ctx, req.(*UpdateBranch))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BranchPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.BranchService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchServiceServer).Delete(ctx, req.(*BranchPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// BranchService_ServiceDesc is the grpc.ServiceDesc for BranchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BranchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_service.BranchService",
	HandlerType: (*BranchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BranchService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _BranchService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _BranchService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BranchService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BranchService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "branch.proto",
}
