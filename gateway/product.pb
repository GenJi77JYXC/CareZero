
�
productServer.protoproduct"P
DelProductReq
	productId (R	productId!
access_token (	RaccessToken"4
DelProductResp
res (Rres
msg (	Rmsg"a
UpdateProductReq*
product (2.product.ProductRproduct!
access_token (	RaccessToken"7
UpdateProductResp
msg (	Rmsg
res (Rres"j
CreateProductReq3
product (2.product.ProductBasicInfoRproduct!
access_token (	RaccessToken"7
CreateProductResp
msg (	Rmsg
res (Rres"�
ProductBasicInfo
name (	Rname 
description (	Rdescription
picture (	Rpicture
sku (	Rsku
stock (Rstock
isActive (RisActive
price (Rprice
tag (	Rtag&
createUserName	 (	RcreateUserName&
updateUserName
 (	RupdateUserName

createTime (	R
createTime

updateTime (	R
updateTime
category (	Rcategory"�
ListProductsReq
page (Rpage
pageSize (RpageSize"
categoryName (	RcategoryName!
access_token (	RaccessToken"�
Product
id (Rid
name (	Rname 
description (	Rdescription
picture (	Rpicture
sku (	Rsku
stock (Rstock
isActive (RisActive
price (Rprice
tag	 (	Rtag&
createUserName
 (	RcreateUserName&
updateUserName (	RupdateUserName

createTime (	R
createTime

updateTime (	R
updateTime
category (	Rcategory"h
ListProductsResp,
products (2.product.ProductRproducts
total (Rtotal
msg (	Rmsg"B
GetProductReq
id (Rid!
access_token (	RaccessToken"N
GetProductResp*
product (2.product.ProductRproduct
msg (	Rmsg"L
SearchProductsReq
query (	Rquery!
access_token (	RaccessToken"R
SearchProductsResp*
results (2.product.ProductRresults
msg (	Rmsg2�
ProductCatalogServiceE
ListProducts.product.ListProductsReq.product.ListProductsResp" ?

GetProduct.product.GetProductReq.product.GetProductResp" K
SearchProducts.product.SearchProductsReq.product.SearchProductsResp" H
CreateProduct.product.CreateProductReq.product.CreateProductResp" H
UpdateProduct.product.UpdateProductReq.product.UpdateProductResp" ?

DelProduct.product.DelProductReq.product.DelProductResp" BZ	./productbproto3