package request

type CreateProductReq struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int64   `json:"stock"`
}
