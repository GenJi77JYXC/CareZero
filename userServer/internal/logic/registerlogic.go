package logic

import (
	"context"
	"errors"
	"strconv"
	"www.genji.xin/backend/CareZero/utils"

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
	if utils.VerifyEmailFormat(in.Email) {
		return nil, errors.New(strconv.Itoa(model.EmailError))
	}

	var u model.User
	result := l.svcCtx.DB.Where("email = ?", in.Email).First(&u)
	if result.RowsAffected != 0 {
		return nil, errors.New(strconv.Itoa(model.UserExists))
	}
	u.Email = in.Email
	password, err := utils.Encrypt(u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = password
	l.svcCtx.DB.Create(&u)
	return &user.RegisterResp{}, nil
}
