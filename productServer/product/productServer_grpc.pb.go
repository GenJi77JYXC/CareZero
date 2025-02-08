// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.26.1
// source: productServer.proto

package product

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ProductCatalogService_ListProducts_FullMethodName   = "/product.ProductCatalogService/ListProducts"
	ProductCatalogService_GetProduct_FullMethodName     = "/product.ProductCatalogService/GetProduct"
	ProductCatalogService_SearchProducts_FullMethodName = "/product.ProductCatalogService/SearchProducts"
	ProductCatalogService_CreateProduct_FullMethodName  = "/product.ProductCatalogService/CreateProduct"
	ProductCatalogService_UpdateProduct_FullMethodName  = "/product.ProductCatalogService/UpdateProduct"
	ProductCatalogService_DelProduct_FullMethodName     = "/product.ProductCatalogService/DelProduct"
)

// ProductCatalogServiceClient is the client API for ProductCatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductCatalogServiceClient interface {
	// 通过商品分类获取商品列表
	ListProducts(ctx context.Context, in *ListProductsReq, opts ...grpc.CallOption) (*ListProductsResp, error)
	// 通过商品Id获取商品信息
	GetProduct(ctx context.Context, in *GetProductReq, opts ...grpc.CallOption) (*GetProductResp, error)
	// 通过商品名称获取商品列表
	SearchProducts(ctx context.Context, in *SearchProductsReq, opts ...grpc.CallOption) (*SearchProductsResp, error)
	// 创建商品
	CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResp, error)
	// 更新商品
	UpdateProduct(ctx context.Context, in *UpdateProductReq, opts ...grpc.CallOption) (*UpdateProductResp, error)
	// 删除商品
	DelProduct(ctx context.Context, in *DelProductReq, opts ...grpc.CallOption) (*DelProductResp, error)
}

type productCatalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductCatalogServiceClient(cc grpc.ClientConnInterface) ProductCatalogServiceClient {
	return &productCatalogServiceClient{cc}
}

func (c *productCatalogServiceClient) ListProducts(ctx context.Context, in *ListProductsReq, opts ...grpc.CallOption) (*ListProductsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProductsResp)
	err := c.cc.Invoke(ctx, ProductCatalogService_ListProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) GetProduct(ctx context.Context, in *GetProductReq, opts ...grpc.CallOption) (*GetProductResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProductResp)
	err := c.cc.Invoke(ctx, ProductCatalogService_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) SearchProducts(ctx context.Context, in *SearchProductsReq, opts ...grpc.CallOption) (*SearchProductsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchProductsResp)
	err := c.cc.Invoke(ctx, ProductCatalogService_SearchProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProductResp)
	err := c.cc.Invoke(ctx, ProductCatalogService_CreateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) UpdateProduct(ctx context.Context, in *UpdateProductReq, opts ...grpc.CallOption) (*UpdateProductResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProductResp)
	err := c.cc.Invoke(ctx, ProductCatalogService_UpdateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) DelProduct(ctx context.Context, in *DelProductReq, opts ...grpc.CallOption) (*DelProductResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelProductResp)
	err := c.cc.Invoke(ctx, ProductCatalogService_DelProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCatalogServiceServer is the server API for ProductCatalogService service.
// All implementations must embed UnimplementedProductCatalogServiceServer
// for forward compatibility
type ProductCatalogServiceServer interface {
	// 通过商品分类获取商品列表
	ListProducts(context.Context, *ListProductsReq) (*ListProductsResp, error)
	// 通过商品Id获取商品信息
	GetProduct(context.Context, *GetProductReq) (*GetProductResp, error)
	// 通过商品名称获取商品列表
	SearchProducts(context.Context, *SearchProductsReq) (*SearchProductsResp, error)
	// 创建商品
	CreateProduct(context.Context, *CreateProductReq) (*CreateProductResp, error)
	// 更新商品
	UpdateProduct(context.Context, *UpdateProductReq) (*UpdateProductResp, error)
	// 删除商品
	DelProduct(context.Context, *DelProductReq) (*DelProductResp, error)
	mustEmbedUnimplementedProductCatalogServiceServer()
}

// UnimplementedProductCatalogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductCatalogServiceServer struct {
}

func (UnimplementedProductCatalogServiceServer) ListProducts(context.Context, *ListProductsReq) (*ListProductsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedProductCatalogServiceServer) GetProduct(context.Context, *GetProductReq) (*GetProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductCatalogServiceServer) SearchProducts(context.Context, *SearchProductsReq) (*SearchProductsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProducts not implemented")
}
func (UnimplementedProductCatalogServiceServer) CreateProduct(context.Context, *CreateProductReq) (*CreateProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductCatalogServiceServer) UpdateProduct(context.Context, *UpdateProductReq) (*UpdateProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProductCatalogServiceServer) DelProduct(context.Context, *DelProductReq) (*DelProductResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelProduct not implemented")
}
func (UnimplementedProductCatalogServiceServer) mustEmbedUnimplementedProductCatalogServiceServer() {}

// UnsafeProductCatalogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductCatalogServiceServer will
// result in compilation errors.
type UnsafeProductCatalogServiceServer interface {
	mustEmbedUnimplementedProductCatalogServiceServer()
}

func RegisterProductCatalogServiceServer(s grpc.ServiceRegistrar, srv ProductCatalogServiceServer) {
	s.RegisterService(&ProductCatalogService_ServiceDesc, srv)
}

func _ProductCatalogService_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCatalogService_ListProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).ListProducts(ctx, req.(*ListProductsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCatalogService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).GetProduct(ctx, req.(*GetProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_SearchProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProductsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).SearchProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCatalogService_SearchProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).SearchProducts(ctx, req.(*SearchProductsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCatalogService_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).CreateProduct(ctx, req.(*CreateProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCatalogService_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).UpdateProduct(ctx, req.(*UpdateProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_DelProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).DelProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductCatalogService_DelProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).DelProduct(ctx, req.(*DelProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductCatalogService_ServiceDesc is the grpc.ServiceDesc for ProductCatalogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductCatalogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductCatalogService",
	HandlerType: (*ProductCatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProducts",
			Handler:    _ProductCatalogService_ListProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductCatalogService_GetProduct_Handler,
		},
		{
			MethodName: "SearchProducts",
			Handler:    _ProductCatalogService_SearchProducts_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _ProductCatalogService_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductCatalogService_UpdateProduct_Handler,
		},
		{
			MethodName: "DelProduct",
			Handler:    _ProductCatalogService_DelProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "productServer.proto",
}
