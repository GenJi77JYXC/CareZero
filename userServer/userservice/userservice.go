// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package userservice

import (
	"context"

	"www.genji.xin/backend/CareZero/userServer/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetUserInfoReq    = user.GetUserInfoReq
	GetUserInfoResp   = user.GetUserInfoResp
	LoginReq          = user.LoginReq
	LoginResp         = user.LoginResp
	LogoutReq         = user.LogoutReq
	LogoutResp        = user.LogoutResp
	RegisterReq       = user.RegisterReq
	RegisterResp      = user.RegisterResp
	UpdateSelfInfoReq = user.UpdateSelfInfoReq
	UpdateSelfInfoRes = user.UpdateSelfInfoRes

	UserService interface {
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
		UpdateSelfInfo(ctx context.Context, in *UpdateSelfInfoReq, opts ...grpc.CallOption) (*UpdateSelfInfoRes, error)
		Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutResp, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUserService) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUserService) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUserService) UpdateSelfInfo(ctx context.Context, in *UpdateSelfInfoReq, opts ...grpc.CallOption) (*UpdateSelfInfoRes, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UpdateSelfInfo(ctx, in, opts...)
}

func (m *defaultUserService) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutResp, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.Logout(ctx, in, opts...)
}
