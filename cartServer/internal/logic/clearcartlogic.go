package logic

import (
	"context"
	"gorm.io/gorm"
	"www.genji.xin/backend/CareZero/authServer/auth"
	"www.genji.xin/backend/CareZero/model"

	"www.genji.xin/backend/CareZero/cartServer/cart"
	"www.genji.xin/backend/CareZero/cartServer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClearCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearCartLogic {
	return &ClearCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClearCartLogic) ClearCart(in *cart.ClearCartReq) (*cart.ClearCartResp, error) {
	// todo: add your logic here and delete this line
	// 校验身份
	verifyResp, err := l.svcCtx.AuthRpc.VerifyTokenByRPC(l.ctx, &auth.VerifyTokenReq{Token: in.GetAccessToken()})
	if err != nil {
		l.Logger.Errorf("CreatedCart|VerifyToken err:%v", err)
		return nil, err
	}
	if !verifyResp.Res {
		return &cart.ClearCartResp{
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
		return &cart.ClearCartResp{
			Msg: "用户信息获取失败，请联系管理员",
		}, sqlRes.Error
	}
	// 校验Cart的User对应关系对不对
	var cartEx model.Cart
	err = l.svcCtx.DB.Where("id = ?", in.CartID).First(&cartEx).Error
	if err != nil {
		return &cart.ClearCartResp{Msg: "未查找到该购物车"}, err
	}
	if uint(cartEx.UserID) != usr.ID {
		return &cart.ClearCartResp{Msg: "该购物车不属于该用户，无操作权限"}, err
	}

	if err := l.svcCtx.DB.Where("cart_id = ?", in.CartID).Delete(&model.CartItem{}).Error; err != nil {
		return &cart.ClearCartResp{Msg: "清除购物车失败"}, err
	}

	return &cart.ClearCartResp{Msg: "清除购物车成功"}, nil
}
