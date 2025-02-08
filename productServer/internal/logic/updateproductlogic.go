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

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProduct 更新商品
func (l *UpdateProductLogic) UpdateProduct(in *product.UpdateProductReq) (*product.UpdateProductResp, error) {
	// 鉴权
	verifyResp, err := l.svcCtx.AuthRpc.VerifyTokenByRPC(l.ctx, &auth.VerifyTokenReq{Token: in.GetAccessToken()})
	if err != nil {
		l.Logger.Errorf("CreateProduct——鉴权失败")
		return nil, err
	}
	if !verifyResp.Res {
		return &product.UpdateProductResp{
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
		return &product.UpdateProductResp{
			Msg: "用户信息获取失败，请联系管理员",
		}, sqlRes.Error
	}
	// 获取user的角色
	userRole := l.svcCtx.Auth.GetRolesForUserInDomain(usr.Username, "domain1")
	if !utils.IsAdmin(userRole) {
		return &product.UpdateProductResp{
			Msg: "用户权限不足",
			Res: false,
		}, err
	}
	// struct 方法更新 只更新非零值字段
	oldProduct := model.Product{}
	err = l.svcCtx.DB.Where("id = ?", in.Product.Id).First(&oldProduct).Error
	if err != nil {
		l.Logger.Errorf("查询的商品失败 : ", err)
		return nil, err
	}
	oldProduct.Name = in.Product.Name
	oldProduct.Description = in.Product.Description
	oldProduct.Picture = in.Product.Picture
	//oldProduct.SKU = in.Product.Sku
	oldProduct.Stock = int(in.Product.Stock)
	oldProduct.IsActive = in.Product.IsActive
	oldProduct.Price = in.Product.Price
	oldProduct.Tags = in.Product.Tag
	oldProduct.UpdatedByID = usr.ID
	oldProduct.UpdatedBy = *usr
	oldProduct.Category = in.Product.Category
	result := l.svcCtx.DB.Model(&oldProduct).Updates(oldProduct)

	// map 方法更新
	//err = l.svcCtx.DB.Model(&model.Product{}).Where("id = ?", in.Product.Id).Updates(map[string]interface{}{
	//	"Name":        in.Product.Name,
	//	"Description": in.Product.Description,
	//	"Picture":     in.Product.Picture,
	//	"Category":    in.Product.Category,
	//	"Stock":       in.Product.Stock,
	//	//"SKU":         in.Product.Sku,
	//	"IsActive":    in.Product.IsActive,
	//	"UpdatedByID": uint(userId),
	//	"UpdatedBy":   *usr,
	//	"Tags":        in.Product.Tag,
	//}).Error
	if result.Error != nil {
		l.Logger.Errorf("更新商品失败", err)
		return &product.UpdateProductResp{
			Msg: "更新商品失败",
			Res: false,
		}, err
	}

	return &product.UpdateProductResp{
		Res: true,
		Msg: "更新商品成功",
	}, nil
}
