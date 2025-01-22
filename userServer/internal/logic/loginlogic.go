package logic

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"www.genji.xin/backend/CareZero/model"

	"www.genji.xin/backend/CareZero/userServer/internal/svc"
	"www.genji.xin/backend/CareZero/userServer/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// todo: add your logic here and delete this line

	u := model.User{}
	result := l.svcCtx.DB.Where("email = ?", in.Email).First(&u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &user.LoginResp{}, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return &user.LoginResp{}, result.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(model.Salt), []byte(in.Password))
	if err != nil {
		return &user.LoginResp{}, errors.New(strconv.Itoa(model.PasswordError))
	}
	return &user.LoginResp{UserId: int32(u.ID)}, nil
}
