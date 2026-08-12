package model

type Product struct {
	ID         int64   `json:"id" gorm:"column:id;primary_key;auto_increment"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Stock      int64   `json:"stock"`
	UpdateTime int64   `json:"update_time"`
}

func (Product) TableName() string {
	return "micro_product"
}
