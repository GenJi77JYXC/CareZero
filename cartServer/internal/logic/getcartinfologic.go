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

type GetCartInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartInfoLogic {
	return &GetCartInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCartInfoLogic) GetCartInfo(in *cart.GetCartInfoReq) (*cart.GetCartInfoResp, error) {
	// todo: add your logic here and delete this line
	// 校验身份
	verifyResp, err := l.svcCtx.AuthRpc.VerifyTokenByRPC(l.ctx, &auth.VerifyTokenReq{Token: in.GetAccessToken()})
	if err != nil {
		l.Logger.Errorf("CreatedCart|VerifyToken err:%v", err)
		return nil, err
	}
	if !verifyResp.Res {
		return &cart.GetCartInfoResp{
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
		return &cart.GetCartInfoResp{
			Msg: "用户信息获取失败，请联系管理员",
		}, sqlRes.Error
	}
	var cartEx model.Cart
	err = l.svcCtx.DB.Model(&model.Cart{}).First(&cartEx).Error
	if err != nil {
		l.Logger.Errorf("GetCartInfo err:%v", err)
		return &cart.GetCartInfoResp{
			Msg: "未查找到该购物车",
		}, err
	}

	var resp = &cart.GetCartInfoResp{
		Msg:       "查找成功",
		CartID:    in.CartID,
		UserID:    int64(usr.ID),
		CreatedAt: cartEx.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: cartEx.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	var cartItems []model.CartItem
	err = l.svcCtx.DB.Where("cart_id = ?", in.CartID).Find(&cartItems).Error
	if err != nil {
		l.Logger.Errorf("GetCartInfo|FindCartItems err:%v", err)
		return &cart.GetCartInfoResp{
			Msg: "查询购物车信息失败",
		}, err
	}

	respCartItems := make([]*cart.CartItem, 0, len(cartItems))
	for i := 0; i < len(cartItems); i++ {
		items := &cart.CartItem{
			CartItemID: cartItems[i].CartID,
			ProductID:  cartItems[i].ProductID,
			Quantity:   cartItems[i].Quantity,
		}
		respCartItems = append(respCartItems, items)
	}
	resp.CartItems = respCartItems

	return resp, nil
}
