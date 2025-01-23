package logic

import (
	"context"
	"www.genji.xin/backend/CareZero/utils"

	"www.genji.xin/backend/CareZero/authServer/auth"
	"www.genji.xin/backend/CareZero/authServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyTokenByRPCLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyTokenByRPCLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenByRPCLogic {
	return &VerifyTokenByRPCLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyTokenByRPCLogic) VerifyTokenByRPC(in *auth.VerifyTokenReq) (*auth.VerifyResp, error) {
	// todo: add your logic here and delete this line
	_, err := utils.VerifyToken(in.Token)
	if err != nil {
		return nil, err
	}

	return &auth.VerifyResp{
		Res: true,
	}, nil
}
