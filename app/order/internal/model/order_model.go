package model

import "time"

type Order struct {
	ID         int64     `json:"id"`
	UserId     int64     `json:"user_id"`
	ProductId  int64     `json:"product_id"`
	Count      int       `json:"count"`
	Price      string    `json:"price"`
	CreateTime time.Time `json:"create_time"`
}

func (Order) TableName() string {
	return "micro_order"
}
