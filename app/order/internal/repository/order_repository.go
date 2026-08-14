package repository

import (
	"context"
	"gin-micro-shop/app/order/internal/model"
	"gorm.io/gorm"
	"time"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (or *OrderRepository) GetOrderById(ctx context.Context, id int64) (order model.Order, err error) {
	err = or.db.Where("id = ?", id).First(&order).Error
	return
}

func (or *OrderRepository) CreateOrder(ctx context.Context, order *model.Order) (bool, error) {
	order.CreateTime = time.Now()
	tx := or.db.Create(order)
	return tx.Error == nil, tx.Error
}
