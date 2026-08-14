package dto

type UserDto struct {
	ID       int64  `json:"id"`
	UserName string `json:"name"`
	Age      string `json:"age"`
	Sex      int64  `json:"sex"`
}
