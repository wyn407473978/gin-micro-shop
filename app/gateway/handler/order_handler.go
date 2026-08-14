package handler

import (
	"fmt"
	orderv1 "gin-micro-shop/api/proto/order/v1"
	"gin-micro-shop/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type OrderHandler struct {
	orderGrpcServiceClient orderv1.OrderGrpcServiceClient
}

func NewOrderHandler(orderGrpcServiceClient orderv1.OrderGrpcServiceClient) *OrderHandler {
	return &OrderHandler{orderGrpcServiceClient: orderGrpcServiceClient}
}

func (os *OrderHandler) GetOrderById(ctx *gin.Context) {
	id := ctx.Query("id")
	idInt64, _ := strconv.ParseInt(id, 10, 64)
	orderResponse, err := os.orderGrpcServiceClient.GetOrder(ctx.Request.Context(), &orderv1.GetOrderRequest{
		OrderId: idInt64,
	})
	if err != nil {
		fmt.Println(err)
		response.ErrorWithCode(ctx, 500, "requestId", "远程调用失败")
		return
	}
	response.SuccessWithData(ctx, orderResponse, "requestId")
	return
}
func (os *OrderHandler) CreateOrder(ctx *gin.Context) {
	orderResponse, err := os.orderGrpcServiceClient.CreateOrder(ctx.Request.Context(), &orderv1.CreateOrderRequest{
		UserId:    1,
		ProductId: 1,
		Count:     1,
	})
	if err != nil {
		response.ErrorWithCode(ctx, 500, "requestId", "远程调用失败")
		return
	}
	response.SuccessWithData(ctx, orderResponse.Success, "requestId")
	return
}
