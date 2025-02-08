package logic

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"www.genji.xin/backend/CareZero/authServer/auth"
	"www.genji.xin/backend/CareZero/model"
	"www.genji.xin/backend/CareZero/utils"

	"www.genji.xin/backend/CareZero/productServer/internal/svc"
	"www.genji.xin/backend/CareZero/productServer/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProductsLogic {
	return &SearchProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchProductsLogic) SearchProducts(in *product.SearchProductsReq) (*product.SearchProductsResp, error) {
	// todo: add your logic here and delete this line
	// 鉴权
	verifyResp, err := l.svcCtx.AuthRpc.VerifyTokenByRPC(l.ctx, &auth.VerifyTokenReq{Token: in.GetAccessToken()})
	if err != nil {
		l.Logger.Errorf("GetProduct——鉴权失败")
		return nil, err
	}
	if !verifyResp.Res {
		return &product.SearchProductsResp{
			Results: nil,
			Msg:     "登录信息失效，请登录后再试",
		}, nil
	}
	// 从Token中获取UserId
	userId := verifyResp.Claims.Fields["user_id"].GetNumberValue()
	// 获取User信息
	usr := &model.User{Model: gorm.Model{ID: uint(userId)}}
	sqlRes := l.svcCtx.DB.First(&usr)
	if sqlRes.Error != nil {
		l.Logger.Errorf(sqlRes.Error.Error())
		return &product.SearchProductsResp{
			Results: nil,
			Msg:     "用户信息获取失败，请联系管理员",
		}, nil
	}
	// 获取user的角色
	userRoles, err := l.svcCtx.Auth.GetUsersForRole(usr.Username)
	if err != nil {
		return &product.SearchProductsResp{
			Results: nil,
			Msg:     "用户角色信息获取失败，请联系管理员",
		}, nil
	}
	// products 从数据库中查询到的商品信息
	var products []*model.Product
	err = l.svcCtx.DB.Where("name LIKE ?", fmt.Sprintf("%%%s%%", in.Query)).Find(products).Error
	if err != nil {
		l.Logger.Errorf("通过名字模糊查询商品失败")
		return nil, err
	}
	productsRes := make([]*product.Product, 0, len(products))
	for _, val := range products {
		productResp := &product.Product{}
		productResp.Reset()
		productResp.Id = uint64(val.ID)
		productResp.Name = val.Name
		productResp.Description = val.Description
		productResp.Picture = val.Picture
		productResp.Sku = val.SKU
		productResp.Stock = int64(val.Stock)
		productResp.IsActive = val.IsActive
		productResp.Price = val.Price
		productResp.Tag = val.Tags
		// 判断user是否是Admin
		isAdmin := utils.IsAdmin(userRoles)
		if isAdmin {
			createUser := &model.User{Model: gorm.Model{ID: val.CreatedByID}}
			err = l.svcCtx.DB.First(&createUser).Error
			if err != nil {
				l.Logger.Errorf("获取商品创建人失败")
			}
			productResp.CreateUserName = createUser.Username
			updateUser := &model.User{Model: gorm.Model{ID: val.UpdatedByID}}
			err = l.svcCtx.DB.First(&updateUser).Error
			if err != nil {
				l.Logger.Errorf("获取商品跟新人失败")
			}
			productResp.UpdateUserName = updateUser.Username
			productResp.CreateTime = val.CreatedAt.String()
			productResp.UpdateTime = val.UpdatedAt.String()
		}
		productsRes = append(productsRes, productResp)
	}
	return &product.SearchProductsResp{
		Results: productsRes,
		Msg:     "查询商品成功",
	}, nil
}
