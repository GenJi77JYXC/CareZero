package logic

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"www.genji.xin/backend/CareZero/authServer/auth"
	"www.genji.xin/backend/CareZero/model"
	"www.genji.xin/backend/CareZero/utils"

	"www.genji.xin/backend/CareZero/productServer/internal/svc"
	"www.genji.xin/backend/CareZero/productServer/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateProduct 创建商品
func (l *CreateProductLogic) CreateProduct(in *product.CreateProductReq) (*product.CreateProductResp, error) {
	// todo: add your logic here and delete this line
	// 鉴权
	verifyResp, err := l.svcCtx.AuthRpc.VerifyTokenByRPC(l.ctx, &auth.VerifyTokenReq{Token: in.GetAccessToken()})
	if err != nil {
		l.Logger.Errorf("CreateProduct——鉴权失败")
		return nil, err
	}
	if !verifyResp.Res {
		return &product.CreateProductResp{
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
		return &product.CreateProductResp{
			Msg: "用户信息获取失败，请联系管理员",
		}, sqlRes.Error
	}
	fmt.Printf("%+v\n", *usr)
	// 获取user的角色
	userRole := l.svcCtx.Auth.GetRolesForUserInDomain(usr.Username, "domain1")
	//fmt.Println(userRole, "=====================================", usr)

	if !utils.IsAdmin(userRole) {
		return &product.CreateProductResp{
			Msg: "用户权限不足",
			Res: false,
		}, err
	}

	fmt.Println(in.Product)

	newProduct := &model.Product{
		Name:        in.Product.Name,
		Description: in.Product.Description,
		Picture:     in.Product.Picture,
		SKU:         uuid.New().String(),
		Stock:       int(in.Product.Stock),
		IsActive:    in.Product.IsActive,
		Price:       in.Product.Price,
		Tags:        in.Product.Tag,
		CreatedByID: uint(userId),
		UpdatedByID: uint(userId),
		CreatedBy:   *usr,
		UpdatedBy:   *usr,
		Category:    in.Product.Category,
	}
	fmt.Printf("%+v\n", newProduct)
	result := l.svcCtx.DB.Create(newProduct)
	if result.Error != nil {
		l.Logger.Errorf("创建商品失败", err)
		return &product.CreateProductResp{
			Msg: "创建商品失败",
			Res: false,
		}, err
	}

	return &product.CreateProductResp{
		Res: true,
		Msg: "创建商品成功",
	}, nil
}
