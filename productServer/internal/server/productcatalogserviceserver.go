// Code generated by goctl. DO NOT EDIT.
// Source: productServer.proto

package server

import (
	"context"

	"www.genji.xin/backend/CareZero/productServer/internal/logic"
	"www.genji.xin/backend/CareZero/productServer/internal/svc"
	"www.genji.xin/backend/CareZero/productServer/product"
)

type ProductCatalogServiceServer struct {
	svcCtx *svc.ServiceContext
	product.UnimplementedProductCatalogServiceServer
}

func NewProductCatalogServiceServer(svcCtx *svc.ServiceContext) *ProductCatalogServiceServer {
	return &ProductCatalogServiceServer{
		svcCtx: svcCtx,
	}
}

// 通过商品分类获取商品列表
func (s *ProductCatalogServiceServer) ListProducts(ctx context.Context, in *product.ListProductsReq) (*product.ListProductsResp, error) {
	l := logic.NewListProductsLogic(ctx, s.svcCtx)
	return l.ListProducts(in)
}

// 通过商品Id获取商品信息
func (s *ProductCatalogServiceServer) GetProduct(ctx context.Context, in *product.GetProductReq) (*product.GetProductResp, error) {
	l := logic.NewGetProductLogic(ctx, s.svcCtx)
	return l.GetProduct(in)
}

// 通过商品名称获取商品列表
func (s *ProductCatalogServiceServer) SearchProducts(ctx context.Context, in *product.SearchProductsReq) (*product.SearchProductsResp, error) {
	l := logic.NewSearchProductsLogic(ctx, s.svcCtx)
	return l.SearchProducts(in)
}

// 创建商品
func (s *ProductCatalogServiceServer) CreateProduct(ctx context.Context, in *product.CreateProductReq) (*product.CreateProductResp, error) {
	l := logic.NewCreateProductLogic(ctx, s.svcCtx)
	return l.CreateProduct(in)
}

// 更新商品
func (s *ProductCatalogServiceServer) UpdateProduct(ctx context.Context, in *product.UpdateProductReq) (*product.UpdateProductResp, error) {
	l := logic.NewUpdateProductLogic(ctx, s.svcCtx)
	return l.UpdateProduct(in)
}

// 删除商品
func (s *ProductCatalogServiceServer) DelProduct(ctx context.Context, in *product.DelProductReq) (*product.DelProductResp, error) {
	l := logic.NewDelProductLogic(ctx, s.svcCtx)
	return l.DelProduct(in)
}
