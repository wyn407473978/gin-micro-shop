package repository

import (
	"context"
	"gin-micro-shop/app/product/internal/model"
	"gorm.io/gorm"
	"time"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) CreateProduct(ctx context.Context, product *model.Product) (bool, error) {
	product.UpdateTime = time.Now()
	tx := pr.db.Create(product)
	return tx.Error == nil, tx.Error

}
func (pr *ProductRepository) GetProduct(ctx context.Context, id int) (product *model.Product, err error) {
	err = pr.db.Where("id = ?", id).First(&product).Error
	return product, err
}

func (pr *ProductRepository) ReduceStock(ctx context.Context, id, count int64) (bool, error) {
	tx := pr.db.Where("id = ?", id).Update("stock", gorm.Expr("stock - ?", count))
	return tx.Error == nil, tx.Error
}
