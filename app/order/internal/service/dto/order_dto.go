package dto

type OrderDetailResponse struct {
	ID        int64   `json:"id" gorm:"column:id;primary_key;auto_increment"`
	UserId    int64   `json:"user_id"`
	ProductId int64   `json:"product_id"`
	Count     int64   `json:"count"`
	Price     float64 `json:"price"`
}
