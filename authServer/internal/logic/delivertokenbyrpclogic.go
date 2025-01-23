package logic

import (
	"context"
	"gorm.io/gorm"
	"www.genji.xin/backend/CareZero/model"
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
	//password := in.Password
	user := &model.User{
		Model: gorm.Model{ID: uint(in.UserId)},
	}

	l.svcCtx.DB.First(user)

	token, err := utils.GenerateToken(user.ID, user.Role, l.svcCtx.Rds)
	if err != nil {
		return nil, err
	}

	return &auth.DeliveryResp{
		AccessToken: token,
	}, nil
}
