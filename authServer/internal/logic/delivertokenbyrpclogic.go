package logic

import (
	"context"
	"www.genji.xin/backend/CareZero/utils"

	"www.genji.xin/backend/CareZero/authServer/auth"
	"www.genji.xin/backend/CareZero/authServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeliverTokenByRPCLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeliverTokenByRPCLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeliverTokenByRPCLogic {
	return &DeliverTokenByRPCLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeliverTokenByRPCLogic) DeliverTokenByRPC(in *auth.DeliverTokenReq) (*auth.DeliveryResp, error) {
	// todo: add your logic here and delete this line
	//phone := in.Phone
	//password := in.Password )

	token, err := utils.GenerateToken(uint(in.Id), l.svcCtx.Rds)
	if err != nil {
		return nil, err
	}

	return &auth.DeliveryResp{
		AccessToken: token,
	}, nil
}
