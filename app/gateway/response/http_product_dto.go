package response

type CreateProductReq struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int64   `json:"stock"`
}

type ReduceStockReq struct {
	ProductId int64 `json:"productId"`
	Count     int64 `json:"count"`
}
