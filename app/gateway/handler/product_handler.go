package handler

import (
	"fmt"
	productv1 "gin-micro-shop/api/proto/product/v1"
	httpproductdto "gin-micro-shop/app/gateway/response"
	"gin-micro-shop/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ProductHandler struct {
	productServiceClient productv1.ProductGrpcServiceClient
}

func NewProductHandler(productServiceClient productv1.ProductGrpcServiceClient) *ProductHandler {
	return &ProductHandler{productServiceClient: productServiceClient}
}

func (h *ProductHandler) ReduceStock(gin *gin.Context) {
	var reduccStockReq httpproductdto.ReduceStockReq
	if err := gin.ShouldBindJSON(&reduccStockReq); err != nil {
		return
	}
	stockResponse, err := h.productServiceClient.ReduceStock(gin.Request.Context(), &productv1.ReduceStockRequest{
		Id:    reduccStockReq.ProductId,
		Count: reduccStockReq.Count,
	})
	if err != nil {
		response.ErrorWithCode(gin, 500, "requestId", "远程调用失败")
		return
	}
	response.SuccessWithData(gin, stockResponse.Success, "requestId")

}

func (h *ProductHandler) GetProduct(gin *gin.Context) {
	id, b := gin.GetQuery("id")
	fmt.Println(id)
	productId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("err:", err)
		response.ErrorWithCode(gin, 500, "requestId", "参数错误")
		return
	}
	if !b {
		response.ErrorWithCode(gin, 400, "requestId", "参数错误")
		return
	}
	productResponse, err := h.productServiceClient.GetProduct(gin.Request.Context(), &productv1.GetProductRequest{
		Id: int64(productId),
	})
	if err != nil {
		fmt.Println("err:", err)
		response.ErrorWithCode(gin, 500, "requestId", "远程调用失败")
		return
	}
	//实际上这里需要在封装一层然后在返回给前端
	response.SuccessWithData(gin, productResponse, "requestId")
}

func (h *ProductHandler) CreateProduct(gin *gin.Context) {
	var createProductReq httpproductdto.CreateProductReq
	if err := gin.ShouldBindJSON(&createProductReq); err != nil {
		return
	}
	fmt.Println(createProductReq)
	createProductResponse, err := h.productServiceClient.CreateProduct(gin.Request.Context(), &productv1.CreateProductRequest{
		Name:  createProductReq.Name,
		Stock: createProductReq.Stock,
		Price: createProductReq.Price,
	})
	if err != nil {
		response.ErrorWithCode(gin, 500, "requestId", "远程调用失败")
		return
	}
	response.SuccessWithData(gin, createProductResponse.Success, "requestId")
}
