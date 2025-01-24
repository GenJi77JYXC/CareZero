// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.26.1
// source: authorization.proto

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AuthService_DeliverTokenByRPC_FullMethodName = "/auth.AuthService/DeliverTokenByRPC"
	AuthService_VerifyTokenByRPC_FullMethodName  = "/auth.AuthService/VerifyTokenByRPC"
	AuthService_RenewalTokenByRPC_FullMethodName = "/auth.AuthService/RenewalTokenByRPC"
	AuthService_Ping_FullMethodName              = "/auth.AuthService/Ping"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	DeliverTokenByRPC(ctx context.Context, in *DeliverTokenReq, opts ...grpc.CallOption) (*DeliveryResp, error)
	VerifyTokenByRPC(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyResp, error)
	RenewalTokenByRPC(ctx context.Context, in *RenewalTokenReq, opts ...grpc.CallOption) (*RenewalTokenResp, error)
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) DeliverTokenByRPC(ctx context.Context, in *DeliverTokenReq, opts ...grpc.CallOption) (*DeliveryResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeliveryResp)
	err := c.cc.Invoke(ctx, AuthService_DeliverTokenByRPC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyTokenByRPC(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyResp)
	err := c.cc.Invoke(ctx, AuthService_VerifyTokenByRPC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RenewalTokenByRPC(ctx context.Context, in *RenewalTokenReq, opts ...grpc.CallOption) (*RenewalTokenResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RenewalTokenResp)
	err := c.cc.Invoke(ctx, AuthService_RenewalTokenByRPC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, AuthService_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	DeliverTokenByRPC(context.Context, *DeliverTokenReq) (*DeliveryResp, error)
	VerifyTokenByRPC(context.Context, *VerifyTokenReq) (*VerifyResp, error)
	RenewalTokenByRPC(context.Context, *RenewalTokenReq) (*RenewalTokenResp, error)
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) DeliverTokenByRPC(context.Context, *DeliverTokenReq) (*DeliveryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeliverTokenByRPC not implemented")
}
func (UnimplementedAuthServiceServer) VerifyTokenByRPC(context.Context, *VerifyTokenReq) (*VerifyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyTokenByRPC not implemented")
}
func (UnimplementedAuthServiceServer) RenewalTokenByRPC(context.Context, *RenewalTokenReq) (*RenewalTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewalTokenByRPC not implemented")
}
func (UnimplementedAuthServiceServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_DeliverTokenByRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliverTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeliverTokenByRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_DeliverTokenByRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeliverTokenByRPC(ctx, req.(*DeliverTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyTokenByRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyTokenByRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_VerifyTokenByRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyTokenByRPC(ctx, req.(*VerifyTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RenewalTokenByRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewalTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RenewalTokenByRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RenewalTokenByRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RenewalTokenByRPC(ctx, req.(*RenewalTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeliverTokenByRPC",
			Handler:    _AuthService_DeliverTokenByRPC_Handler,
		},
		{
			MethodName: "VerifyTokenByRPC",
			Handler:    _AuthService_VerifyTokenByRPC_Handler,
		},
		{
			MethodName: "RenewalTokenByRPC",
			Handler:    _AuthService_RenewalTokenByRPC_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _AuthService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authorization.proto",
}
