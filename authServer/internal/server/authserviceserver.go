// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5
// Source: authorization.proto

package server

import (
	"context"

	"www.genji.xin/backend/CareZero/authServer/auth"
	"www.genji.xin/backend/CareZero/authServer/internal/logic"
	"www.genji.xin/backend/CareZero/authServer/internal/svc"
)

type AuthServiceServer struct {
	svcCtx *svc.ServiceContext
	auth.UnimplementedAuthServiceServer
}

func NewAuthServiceServer(svcCtx *svc.ServiceContext) *AuthServiceServer {
	return &AuthServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *AuthServiceServer) DeliverTokenByRPC(ctx context.Context, in *auth.DeliverTokenReq) (*auth.DeliveryResp, error) {
	l := logic.NewDeliverTokenByRPCLogic(ctx, s.svcCtx)
	return l.DeliverTokenByRPC(in)
}

func (s *AuthServiceServer) VerifyTokenByRPC(ctx context.Context, in *auth.VerifyTokenReq) (*auth.VerifyResp, error) {
	l := logic.NewVerifyTokenByRPCLogic(ctx, s.svcCtx)
	return l.VerifyTokenByRPC(in)
}

func (s *AuthServiceServer) Ping(ctx context.Context, in *auth.Request) (*auth.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
