package logic

import (
	"context"
	"time"
	"www.genji.xin/backend/CareZero/authServer/util"
	"www.genji.xin/backend/CareZero/model"

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
	//password := in.Password

	token, err := util.GenerateToken(1, model.CustomerRole)
	if err != nil {
		return nil, err
	}
	tokenExpire := time.Now().Add(time.Hour).Unix()

	return &auth.DeliveryResp{
		Token:       token,
		TokenExpire: tokenExpire,
	}, nil
}
