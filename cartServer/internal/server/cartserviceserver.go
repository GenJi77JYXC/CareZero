// Code generated by goctl. DO NOT EDIT.
// Source: cartServer.proto

package server

import (
	"context"

	"www.genji.xin/backend/CareZero/cartServer/cart"
	"www.genji.xin/backend/CareZero/cartServer/internal/logic"
	"www.genji.xin/backend/CareZero/cartServer/internal/svc"
)

type CartServiceServer struct {
	svcCtx *svc.ServiceContext
	cart.UnimplementedCartServiceServer
}

func NewCartServiceServer(svcCtx *svc.ServiceContext) *CartServiceServer {
	return &CartServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *CartServiceServer) CreatedCart(ctx context.Context, in *cart.CreatedCartReq) (*cart.CreatedCartResp, error) {
	l := logic.NewCreatedCartLogic(ctx, s.svcCtx)
	return l.CreatedCart(in)
}

func (s *CartServiceServer) ClearCart(ctx context.Context, in *cart.ClearCartReq) (*cart.ClearCartResp, error) {
	l := logic.NewClearCartLogic(ctx, s.svcCtx)
	return l.ClearCart(in)
}

func (s *CartServiceServer) GetCartInfo(ctx context.Context, in *cart.GetCartInfoReq) (*cart.GetCartInfoResp, error) {
	l := logic.NewGetCartInfoLogic(ctx, s.svcCtx)
	return l.GetCartInfo(in)
}
