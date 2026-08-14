package model

import "time"

type Order struct {
	ID         int64     `json:"id" gorm:"column:id;primary_key;auto_increment"`
	UserId     int64     `json:"user_id"`
	ProductId  int64     `json:"product_id"`
	Count      int64     `json:"count"`
	Price      float64   `json:"price"`
	CreateTime time.Time `json:"create_time"`
}

func (Order) TableName() string {
	return "micro_order"
}
