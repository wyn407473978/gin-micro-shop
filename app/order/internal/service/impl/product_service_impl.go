package impl

import (
	"context"
	productv1 "gin-micro-shop/api/proto/product/v1"
	"log"
)

type ProductServiceImpl struct {
	grpcProductService productv1.ProductGrpcServiceClient
}

func NewProductServiceImpl(grpcProductService productv1.ProductGrpcServiceClient) *ProductServiceImpl {
	return &ProductServiceImpl{grpcProductService: grpcProductService}
}

func (p *ProductServiceImpl) StockProduct(ctx context.Context, productId, count int64) (bool, error) {
	stockResponse, err := p.grpcProductService.ReduceStock(ctx, &productv1.ReduceStockRequest{Id: productId, Count: count})
	if err != nil {
		log.Printf("Failed to reduce stock: %v", err)
		return false, err
	}
	return stockResponse.Success, err
}
