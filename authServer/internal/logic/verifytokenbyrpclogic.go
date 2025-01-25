package logic

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/structpb"
	"www.genji.xin/backend/CareZero/model"
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

	// 查看token是否注销
	get, err := l.svcCtx.Rds.Get(fmt.Sprintf("%s%s", model.BlackList_AccessToken, in.Token))
	if err != nil {
		return nil, err
	}
	if get != "" {
		return &auth.VerifyResp{
			Res: false,
			Msg: "token失效",
		}, nil
	}

	claims, err := utils.VerifyToken(in.Token)
	if err != nil {
		return nil, err
	}

	return &auth.VerifyResp{
		Res: true,
		Claims: &structpb.Struct{Fields: map[string]*structpb.Value{
			"user_id":     {Kind: &structpb.Value_NumberValue{NumberValue: float64(claims.UserId)}},
			"refresh":     {Kind: &structpb.Value_StringValue{StringValue: claims.RefreshToken}},
			"refresh_exp": {Kind: &structpb.Value_NumberValue{NumberValue: float64(claims.RefreshExpires)}},
		}},
	}, nil
}
