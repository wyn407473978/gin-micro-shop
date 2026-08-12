package handler

import productv1 "gin-micro-shop/api/proto/product/v1"

type ProductHandler struct {
	ProductServiceClient productv1.ProductGrpcServiceClient
}
