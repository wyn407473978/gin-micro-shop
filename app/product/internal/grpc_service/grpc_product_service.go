package grpc_service

import (
	"context"
	productv1 "gin-micro-shop/api/proto/product/v1"
	request "gin-micro-shop/app/product/internal/entity/req"
	"gin-micro-shop/app/product/internal/service"
)

type ProductGrpcService struct {
	productv1.UnsafeProductGrpcServiceServer
	productService *service.ProductService
}

func (s *ProductGrpcService) CreateProduct(ctx context.Context, req *productv1.CreateProductRequest) (*productv1.CreateProductResponse, error) {
	productServiceReq := &request.CreateProductReq{Name: req.Name, Price: req.Price, Stock: req.Stock}
	productServiceRes, _ := s.productService.CreateProduct(ctx, productServiceReq)
	return &productv1.CreateProductResponse{Success: productServiceRes}, nil
}
func (s *ProductGrpcService) GetProduct(ctx context.Context, req *productv1.GetProductRequest) (*productv1.GetProductResponse, error) {
	getProductByIdResp, err := s.productService.GetProductById(ctx, req.Id)
	return &productv1.GetProductResponse{Id: getProductByIdResp.ID, Name: getProductByIdResp.Name, Price: float32(getProductByIdResp.Price), Stock: getProductByIdResp.Stock}, err
}
func (s *ProductGrpcService) ReduceStock(ctx context.Context, req *productv1.ReduceStockRequest) (*productv1.ReduceStockResponse, error) {
	stock, err := s.productService.ReduceStock(ctx, req.Id, req.Count)
	return &productv1.ReduceStockResponse{Success: stock}, err
}
