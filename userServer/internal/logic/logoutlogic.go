package logic

import (
	"context"
	"fmt"
	"www.genji.xin/backend/CareZero/authServer/authservice"
	"www.genji.xin/backend/CareZero/model"
	"www.genji.xin/backend/CareZero/utils"

	"www.genji.xin/backend/CareZero/userServer/internal/svc"
	"www.genji.xin/backend/CareZero/userServer/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *user.LogoutReq) (*user.LogoutResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.AuthRpc.VerifyTokenByRPC(l.ctx, &authservice.VerifyTokenReq{Token: in.Token})
	if err != nil {
		return &user.LogoutResp{Res: "登出失败，系统内部错误"}, err
	}
	if !rpcResp.Res {
		return &user.LogoutResp{Res: "登出失败，用户信息失效"}, err
	}

	// 把token放入黑名单
	err = l.svcCtx.Rds.Set(fmt.Sprintf("%s%s", model.BlackList_AccessToken, in.Token), "1")
	if err != nil {
		return &user.LogoutResp{Res: "登出失败，系统内部错误"}, err
	}
	err = l.svcCtx.Rds.Expire(fmt.Sprintf("%s%s", model.BlackList_AccessToken, in.Token), utils.AccessExpire*60*60)
	if err != nil {
		l.Logger.Errorf("redis expire err:%v", err, fmt.Sprintf("\n 黑名单中未设置成功过期时间的AccessToken是 : %s%s", model.BlackList_AccessToken, in.Token))
	}

	refreshToken := rpcResp.Claims.Fields["refresh"].GetStringValue()
	fmt.Println("refreshToken:", refreshToken)
	_, err = l.svcCtx.Rds.Del(refreshToken)
	if err != nil {
		l.Logger.Errorf("删除refreshToken失败：%v", err, fmt.Sprintf("\n 没删除的refreshToken是 : %s", refreshToken))
		fmt.Printf(fmt.Sprintf("\n 没删除的refreshToken是 : %s", refreshToken))
	}

	return &user.LogoutResp{
		Res: "登出成功",
	}, nil
}
