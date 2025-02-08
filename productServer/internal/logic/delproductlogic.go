package logic

import (
	"context"
	"gorm.io/gorm"
	"www.genji.xin/backend/CareZero/authServer/auth"
	"www.genji.xin/backend/CareZero/model"
	"www.genji.xin/backend/CareZero/utils"

	"www.genji.xin/backend/CareZero/productServer/internal/svc"
	"www.genji.xin/backend/CareZero/productServer/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelProductLogic {
	return &DelProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DelProduct 删除商品
func (l *DelProductLogic) DelProduct(in *product.DelProductReq) (*product.DelProductResp, error) {
	// 鉴权
	verifyResp, err := l.svcCtx.AuthRpc.VerifyTokenByRPC(l.ctx, &auth.VerifyTokenReq{Token: in.GetAccessToken()})
	if err != nil {
		l.Logger.Errorf("CreateProduct——鉴权失败")
		return nil, err
	}
	if !verifyResp.Res {
		return &product.DelProductResp{
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
		return &product.DelProductResp{
			Msg: "用户信息获取失败，请联系管理员",
		}, sqlRes.Error
	}
	// 获取user的角色
	userRole := l.svcCtx.Auth.GetRolesForUserInDomain(usr.Username, "domain1")

	if !utils.IsAdmin(userRole) {
		return &product.DelProductResp{
			Msg: "用户权限不足",
			Res: false,
		}, err
	}
	err = l.svcCtx.DB.Delete(&model.Product{}, in.ProductId).Error
	if err != nil {
		l.Logger.Errorf("删除商品失败", err)
		return &product.DelProductResp{Msg: "删除商品失败"}, err
	}
	return &product.DelProductResp{
		Res: true,
		Msg: "删除商品成功",
	}, nil
}
