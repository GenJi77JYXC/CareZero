package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"www.genji.xin/backend/CareZero/authServer/authservice"
	"www.genji.xin/backend/CareZero/model"

	"www.genji.xin/backend/CareZero/userServer/internal/svc"
	"www.genji.xin/backend/CareZero/userServer/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSelfInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSelfInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSelfInfoLogic {
	return &UpdateSelfInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSelfInfoLogic) UpdateSelfInfo(in *user.UpdateSelfInfoReq) (*user.UpdateSelfInfoRes, error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.AuthRpc.VerifyTokenByRPC(l.ctx, &authservice.VerifyTokenReq{
		Token: in.Token,
	})
	if err != nil {
		return nil, err
	}
	if !res.Res {
		return nil, errors.New(model.ErrorMsg[model.TokenInvalid])
	}
	userId := res.Claims.Fields["user_id"].GetNumberValue()

	u := &model.User{}
	result := l.svcCtx.DB.Where("id = ?", uint(userId)).First(u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &user.UpdateSelfInfoRes{Res: "更新失败，未找到该用户"}, errors.New(model.ErrorMsg[model.UserUnExists])
	}

	u.Username = in.Username
	u.Email = in.Email
	u.Phone = in.Phone
	err = l.svcCtx.DB.Save(u).Error
	if err != nil {
		return &user.UpdateSelfInfoRes{Res: "更新失败， 保存数据失败"}, err
	}

	return &user.UpdateSelfInfoRes{
		Res: "更新成功",
	}, nil
}
