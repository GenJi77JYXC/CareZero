package logic

import (
	"context"
	"errors"
	"strconv"
	"www.genji.xin/backend/CareZero/model"

	"www.genji.xin/backend/CareZero/userServer/internal/svc"
	"www.genji.xin/backend/CareZero/userServer/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line
	var u model.User
	result := l.svcCtx.DB.Where("email = ?", in.Email).First(&u)
	if result.RowsAffected != 0 {
		return nil, errors.New(strconv.Itoa(model.UserExists))
	}
	return &user.RegisterResp{}, nil
}
