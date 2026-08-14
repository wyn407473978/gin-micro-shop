package impl

import (
	"context"
	"fmt"
	orderv1 "gin-micro-shop/api/proto/order/v1"
	"gin-micro-shop/app/order/internal/model"
	"gin-micro-shop/app/order/internal/repository"
	"gin-micro-shop/app/order/internal/service/interface"
	"time"
)

type OrderGrpcService struct {
	orderv1.UnimplementedOrderGrpcServiceServer
	orderRepository *repository.OrderRepository
	productService  _interface.ProductService
	userService     _interface.UserService
}

func NewOrderGrpcService(
	orderRepository *repository.OrderRepository,
	productService _interface.ProductService,
	userService _interface.UserService) *OrderGrpcService {
	return &OrderGrpcService{
		orderRepository: orderRepository,
		productService:  productService,
		userService:     userService,
	}
}
func (os *OrderGrpcService) GetOrder(ctx context.Context, request *orderv1.GetOrderRequest) (*orderv1.GetOrderResponse, error) {
	order, err := os.orderRepository.GetOrderById(ctx, request.OrderId)
	if err != nil {
		return nil, err
	}
	return &orderv1.GetOrderResponse{
		Id:        order.ID,
		UserId:    order.UserId,
		ProductId: order.ProductId,
		Count:     order.Count,
		Price:     order.Price,
	}, err
}

func (os *OrderGrpcService) CreateOrder(ctx context.Context, request *orderv1.CreateOrderRequest) (response *orderv1.CreateOrderResponse, err error) {
	//TODO: 确认用户是否存在扣减用户余额（没有余额）
	userDto, err := os.userService.GetUserById(ctx, request.UserId)
	if err != nil {
		return nil, err
	}
	fmt.Println(userDto)

	//TODO： 确认商品是否存在，扣减商品库存
	product, err := os.productService.StockProduct(ctx, request.ProductId, request.Count)
	if !product {
		return nil, err
	}
	fmt.Println(product)
	order, err := os.orderRepository.CreateOrder(ctx, &model.Order{
		UserId:     request.UserId,
		ProductId:  request.ProductId,
		Count:      request.Count,
		Price:      100.00,
		CreateTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(order)
	return &orderv1.CreateOrderResponse{
		Success: true,
	}, nil
}
