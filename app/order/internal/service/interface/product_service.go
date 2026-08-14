package _interface

import "context"

type ProductService interface {
	StockProduct(ctx context.Context, productId, count int64) (bool, error)
}
