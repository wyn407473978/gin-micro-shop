package model

import "time"

type User struct {
	ID         int       `json:"id" gorm:"column:id;primary_key;auto_increment"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Age        int       `json:"age"`
	Sex        int       `json:"sex"`
	CreateTime time.Time `json:"create_time"`
}

func (User) TableName() string {
	return "micro_user"
}
