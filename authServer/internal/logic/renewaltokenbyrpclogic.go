package logic

import (
	"context"
	"www.genji.xin/backend/CareZero/utils"

	"www.genji.xin/backend/CareZero/authServer/auth"
	"www.genji.xin/backend/CareZero/authServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RenewalTokenByRPCLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRenewalTokenByRPCLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RenewalTokenByRPCLogic {
	return &RenewalTokenByRPCLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RenewalTokenByRPCLogic) RenewalTokenByRPC(in *auth.RenewalTokenReq) (*auth.RenewalTokenResp, error) {
	// todo: add your logic here and delete this line
	accessToken, err := utils.RefreshAccessToken(l.svcCtx.Rds, in.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &auth.RenewalTokenResp{
		AccessToken: accessToken,
	}, nil
}
