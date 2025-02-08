package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"www.genji.xin/backend/CareZero/authServer/auth"
	"www.genji.xin/backend/CareZero/model"
	"www.genji.xin/backend/CareZero/productServer/internal/svc"
	"www.genji.xin/backend/CareZero/productServer/product"
	"www.genji.xin/backend/CareZero/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetProduct 通过商品Id获取商品信息
func (l *GetProductLogic) GetProduct(in *product.GetProductReq) (*product.GetProductResp, error) {
	// todo: add your logic here and delete this line
	// 鉴权
	verifyResp, err := l.svcCtx.AuthRpc.VerifyTokenByRPC(l.ctx, &auth.VerifyTokenReq{Token: in.GetAccessToken()})
	if err != nil {
		l.Logger.Errorf("GetProduct——鉴权失败")
		return nil, err
	}
	if !verifyResp.Res {
		return &product.GetProductResp{
			Product: nil,
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
		return &product.GetProductResp{
			Product: nil,
			Msg:     "用户信息获取失败，请联系管理员",
		}, sqlRes.Error
	}
	// 获取user的角色
	userRoles, err := l.svcCtx.Auth.GetUsersForRole(usr.Username)
	if err != nil {
		return &product.GetProductResp{
			Product: nil,
			Msg:     "获取用户角色信息失败，请联系管理员",
		}, err
	}
	// resProduct 从数据库中查询到的商品信息
	resProduct := &model.Product{}
	result := l.svcCtx.DB.First(resProduct, in.Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &product.GetProductResp{
				Product: nil,
				Msg:     "商品不存在",
			}, nil
		} else {
			l.Logger.Errorf("GetProduct error:%v", result.Error)
			return &product.GetProductResp{
				Product: nil,
				Msg:     "商品查询出错，请联系管理员",
			}, result.Error
		}
	}
	// productRes 需要返回给用户的商品信息
	productRes := &product.Product{}
	productRes.Reset()
	productRes.Id = uint64(resProduct.ID)
	productRes.Name = resProduct.Name
	productRes.Description = resProduct.Description
	productRes.Picture = resProduct.Picture
	productRes.Sku = resProduct.SKU
	productRes.Stock = int64(resProduct.Stock)
	productRes.IsActive = resProduct.IsActive
	productRes.Price = resProduct.Price
	productRes.Tag = resProduct.Tags

	// 判断user是否是Admin
	isAdmin := utils.IsAdmin(userRoles)
	if isAdmin {
		createUser := &model.User{Model: gorm.Model{ID: resProduct.CreatedByID}}
		err = l.svcCtx.DB.First(&createUser).Error
		if err != nil {
			l.Logger.Errorf("获取商品创建人失败")
		}
		productRes.CreateUserName = createUser.Username
		updateUser := &model.User{Model: gorm.Model{ID: resProduct.UpdatedByID}}
		err = l.svcCtx.DB.First(&updateUser).Error
		if err != nil {
			l.Logger.Errorf("获取商品跟新人失败")
		}
		productRes.UpdateUserName = updateUser.Username
		productRes.CreateTime = resProduct.CreatedAt.String()
		productRes.UpdateTime = resProduct.UpdatedAt.String()
	}

	return &product.GetProductResp{
		Product: productRes,
		Msg:     "查询商品成功",
	}, nil
}
