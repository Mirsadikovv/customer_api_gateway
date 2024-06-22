// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: customer-auth.proto

package auth_service

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

// CustomerAuthClient is the client API for CustomerAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerAuthClient interface {
	LoginByPassword(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GmailCheck(ctx context.Context, in *GmailCheckRequest, opts ...grpc.CallOption) (*GmailCheckResponse, error)
	RegisterByMail(ctx context.Context, in *GmailCheckRequest, opts ...grpc.CallOption) (*Empty, error)
	RegisterByMailConfirm(ctx context.Context, in *RConfirm, opts ...grpc.CallOption) (*RegGmailResp, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Empty, error)
	LoginByGmail(ctx context.Context, in *GmailCheckRequest, opts ...grpc.CallOption) (*Empty, error)
	LoginByGmailComfirm(ctx context.Context, in *LoginByGmailRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	ResetPassword(ctx context.Context, in *GmailCheckRequest, opts ...grpc.CallOption) (*Empty, error)
	ResetPasswordConfirm(ctx context.Context, in *CustomerPasswordConfirm, opts ...grpc.CallOption) (*Empty, error)
}

type customerAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerAuthClient(cc grpc.ClientConnInterface) CustomerAuthClient {
	return &customerAuthClient{cc}
}

func (c *customerAuthClient) LoginByPassword(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth_service.CustomerAuth/LoginByPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerAuthClient) GmailCheck(ctx context.Context, in *GmailCheckRequest, opts ...grpc.CallOption) (*GmailCheckResponse, error) {
	out := new(GmailCheckResponse)
	err := c.cc.Invoke(ctx, "/auth_service.CustomerAuth/GmailCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerAuthClient) RegisterByMail(ctx context.Context, in *GmailCheckRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/auth_service.CustomerAuth/RegisterByMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerAuthClient) RegisterByMailConfirm(ctx context.Context, in *RConfirm, opts ...grpc.CallOption) (*RegGmailResp, error) {
	out := new(RegGmailResp)
	err := c.cc.Invoke(ctx, "/auth_service.CustomerAuth/RegisterByMailConfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerAuthClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/auth_service.CustomerAuth/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerAuthClient) LoginByGmail(ctx context.Context, in *GmailCheckRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/auth_service.CustomerAuth/LoginByGmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerAuthClient) LoginByGmailComfirm(ctx context.Context, in *LoginByGmailRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth_service.CustomerAuth/LoginByGmailComfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerAuthClient) ResetPassword(ctx context.Context, in *GmailCheckRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/auth_service.CustomerAuth/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerAuthClient) ResetPasswordConfirm(ctx context.Context, in *CustomerPasswordConfirm, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/auth_service.CustomerAuth/ResetPasswordConfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerAuthServer is the server API for CustomerAuth service.
// All implementations must embed UnimplementedCustomerAuthServer
// for forward compatibility
type CustomerAuthServer interface {
	LoginByPassword(context.Context, *LoginRequest) (*LoginResponse, error)
	GmailCheck(context.Context, *GmailCheckRequest) (*GmailCheckResponse, error)
	RegisterByMail(context.Context, *GmailCheckRequest) (*Empty, error)
	RegisterByMailConfirm(context.Context, *RConfirm) (*RegGmailResp, error)
	Create(context.Context, *CreateRequest) (*Empty, error)
	LoginByGmail(context.Context, *GmailCheckRequest) (*Empty, error)
	LoginByGmailComfirm(context.Context, *LoginByGmailRequest) (*LoginResponse, error)
	ResetPassword(context.Context, *GmailCheckRequest) (*Empty, error)
	ResetPasswordConfirm(context.Context, *CustomerPasswordConfirm) (*Empty, error)
	mustEmbedUnimplementedCustomerAuthServer()
}

// UnimplementedCustomerAuthServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerAuthServer struct {
}

func (UnimplementedCustomerAuthServer) LoginByPassword(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByPassword not implemented")
}
func (UnimplementedCustomerAuthServer) GmailCheck(context.Context, *GmailCheckRequest) (*GmailCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GmailCheck not implemented")
}
func (UnimplementedCustomerAuthServer) RegisterByMail(context.Context, *GmailCheckRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterByMail not implemented")
}
func (UnimplementedCustomerAuthServer) RegisterByMailConfirm(context.Context, *RConfirm) (*RegGmailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterByMailConfirm not implemented")
}
func (UnimplementedCustomerAuthServer) Create(context.Context, *CreateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCustomerAuthServer) LoginByGmail(context.Context, *GmailCheckRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByGmail not implemented")
}
func (UnimplementedCustomerAuthServer) LoginByGmailComfirm(context.Context, *LoginByGmailRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByGmailComfirm not implemented")
}
func (UnimplementedCustomerAuthServer) ResetPassword(context.Context, *GmailCheckRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedCustomerAuthServer) ResetPasswordConfirm(context.Context, *CustomerPasswordConfirm) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPasswordConfirm not implemented")
}
func (UnimplementedCustomerAuthServer) mustEmbedUnimplementedCustomerAuthServer() {}

// UnsafeCustomerAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerAuthServer will
// result in compilation errors.
type UnsafeCustomerAuthServer interface {
	mustEmbedUnimplementedCustomerAuthServer()
}

func RegisterCustomerAuthServer(s grpc.ServiceRegistrar, srv CustomerAuthServer) {
	s.RegisterService(&CustomerAuth_ServiceDesc, srv)
}

func _CustomerAuth_LoginByPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServer).LoginByPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.CustomerAuth/LoginByPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServer).LoginByPassword(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerAuth_GmailCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GmailCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServer).GmailCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.CustomerAuth/GmailCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServer).GmailCheck(ctx, req.(*GmailCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerAuth_RegisterByMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GmailCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServer).RegisterByMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.CustomerAuth/RegisterByMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServer).RegisterByMail(ctx, req.(*GmailCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerAuth_RegisterByMailConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RConfirm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServer).RegisterByMailConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.CustomerAuth/RegisterByMailConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServer).RegisterByMailConfirm(ctx, req.(*RConfirm))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerAuth_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.CustomerAuth/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerAuth_LoginByGmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GmailCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServer).LoginByGmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.CustomerAuth/LoginByGmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServer).LoginByGmail(ctx, req.(*GmailCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerAuth_LoginByGmailComfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByGmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServer).LoginByGmailComfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.CustomerAuth/LoginByGmailComfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServer).LoginByGmailComfirm(ctx, req.(*LoginByGmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerAuth_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GmailCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.CustomerAuth/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServer).ResetPassword(ctx, req.(*GmailCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerAuth_ResetPasswordConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerPasswordConfirm)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerAuthServer).ResetPasswordConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_service.CustomerAuth/ResetPasswordConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerAuthServer).ResetPasswordConfirm(ctx, req.(*CustomerPasswordConfirm))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerAuth_ServiceDesc is the grpc.ServiceDesc for CustomerAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_service.CustomerAuth",
	HandlerType: (*CustomerAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginByPassword",
			Handler:    _CustomerAuth_LoginByPassword_Handler,
		},
		{
			MethodName: "GmailCheck",
			Handler:    _CustomerAuth_GmailCheck_Handler,
		},
		{
			MethodName: "RegisterByMail",
			Handler:    _CustomerAuth_RegisterByMail_Handler,
		},
		{
			MethodName: "RegisterByMailConfirm",
			Handler:    _CustomerAuth_RegisterByMailConfirm_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CustomerAuth_Create_Handler,
		},
		{
			MethodName: "LoginByGmail",
			Handler:    _CustomerAuth_LoginByGmail_Handler,
		},
		{
			MethodName: "LoginByGmailComfirm",
			Handler:    _CustomerAuth_LoginByGmailComfirm_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _CustomerAuth_ResetPassword_Handler,
		},
		{
			MethodName: "ResetPasswordConfirm",
			Handler:    _CustomerAuth_ResetPasswordConfirm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer-auth.proto",
}
