package logic

import (
	"context"

	"www.genji.xin/backend/CareZero/productServer/internal/svc"
	"www.genji.xin/backend/CareZero/productServer/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductLogic) GetProduct(in *product.GetProductReq) (*product.GetProductResp, error) {
	// todo: add your logic here and delete this line

	return &product.GetProductResp{}, nil
}
