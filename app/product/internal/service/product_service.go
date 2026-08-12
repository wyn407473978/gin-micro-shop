package service

import (
	"context"
	"gin-micro-shop/app/product/internal/entity/req"
	"gin-micro-shop/app/product/internal/entity/resp"
	"gin-micro-shop/app/product/internal/model"
	"gin-micro-shop/app/product/internal/repository"
)

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (s *ProductService) CreateProduct(ctx context.Context, createProductReq *request.CreateProductReq) (bool, error) {
	var product = &model.Product{Name: createProductReq.Name, Price: createProductReq.Price, Stock: createProductReq.Stock}
	return s.productRepository.CreateProduct(ctx, product)
}

func (s *ProductService) GetProductById(ctx context.Context, productId int64) (*response.GetProductByIdResp, error) {
	product, err := s.productRepository.GetProduct(ctx, int(productId))
	return &response.GetProductByIdResp{ID: product.ID, Name: product.Name, Price: product.Price, Stock: product.Stock}, err
}

func (s *ProductService) ReduceStock(ctx context.Context, id, count int64) (bool, error) {
	return s.productRepository.ReduceStock(ctx, id, count)
}
