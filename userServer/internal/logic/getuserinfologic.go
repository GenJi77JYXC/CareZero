package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"www.genji.xin/backend/CareZero/model"

	"www.genji.xin/backend/CareZero/userServer/internal/svc"
	"www.genji.xin/backend/CareZero/userServer/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	var u model.User
	result := l.svcCtx.DB.Where("email = ?", in.Email).First(&u)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New(model.ErrorMsg[model.UserUnExists])
		}
		return nil, result.Error
	}

	return &user.GetUserInfoResp{
		Username: u.Username,
		Role:     u.Role,
		Email:    u.Email,
		Phone:    u.Phone,
	}, nil
}
