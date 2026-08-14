package service

import (
	"context"
	"gin-micro-shop/app/order/internal/model"
	"gin-micro-shop/app/order/internal/repository"
	"gin-micro-shop/app/order/internal/service/dto"
	"time"
)

type OrderService struct {
	orderRepository *repository.OrderRepository
}

func NewOrderService(orderRepository *repository.OrderRepository) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (os *OrderService) GetOrderById(ctx context.Context, id int64) (orderDtoResponse *dto.OrderDetailResponse, err error) {
	order, err := os.orderRepository.GetOrderById(ctx, id)
	return &dto.OrderDetailResponse{
		ID:        order.ID,
		UserId:    order.UserId,
		ProductId: order.ProductId,
		Count:     order.Count,
		Price:     order.Price,
	}, err
}

func (os *OrderService) CreateOrder(ctx context.Context, userId, productId, count int64) (bool, error) {
	return os.orderRepository.CreateOrder(ctx, &model.Order{
		UserId:     userId,
		ProductId:  productId,
		Count:      count,
		Price:      100.00,
		CreateTime: time.Now(),
	})
}
