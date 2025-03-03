package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"www.genji.xin/backend/CareZero/authServer/auth"
	"www.genji.xin/backend/CareZero/cartServer/cart"
	"www.genji.xin/backend/CareZero/cartServer/internal/svc"
	"www.genji.xin/backend/CareZero/model"
)

type CreatedCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatedCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatedCartLogic {
	return &CreatedCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatedCartLogic) CreatedCart(in *cart.CreatedCartReq) (*cart.CreatedCartResp, error) {
	// todo: add your logic here and delete this line
	// 校验身份
	verifyResp, err := l.svcCtx.AuthRpc.VerifyTokenByRPC(l.ctx, &auth.VerifyTokenReq{Token: in.GetAccessToken()})
	if err != nil {
		l.Logger.Errorf("CreatedCart|VerifyToken err:%v", err)
		return nil, err
	}
	if !verifyResp.Res {
		return &cart.CreatedCartResp{
			Msg: "登录信息失效，请登录后再试",
		}, err
	}

	// 从Token中获取UserId
	userId := verifyResp.Claims.Fields["user_id"].GetNumberValue()
	// 获取User信息
	usr := &model.User{Model: gorm.Model{ID: uint(userId)}}
	sqlRes := l.svcCtx.DB.First(&usr)
	if sqlRes.Error != nil {
		l.Logger.Errorf(sqlRes.Error.Error())
		return &cart.CreatedCartResp{
			Msg: "用户信息获取失败，请联系管理员",
		}, sqlRes.Error
	}
	var cartEx model.Cart
	if err := l.svcCtx.DB.Where("user_id = ?", uint(userId)).First(&cartEx).Error; err == nil {
		return &cart.CreatedCartResp{Msg: "已有购物车"}, nil
	}
	// 创建购物车
	cartEx = model.Cart{
		UserID: int64(userId),
	}
	if err := l.svcCtx.DB.Create(&cartEx).Error; err != nil {
		return &cart.CreatedCartResp{Msg: "创建失败"}, err
	}
	return &cart.CreatedCartResp{Msg: "创建成功"}, nil
}
