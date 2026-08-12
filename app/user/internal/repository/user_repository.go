package repository

import (
	"context"
	"gin-micro-shop/app/gateway/response"
	"gin-micro-shop/app/user/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) FindUserByID(ctx context.Context, id int) (user model.User, err error) {
	err = ur.DB.Where("id = ?", id).First(&user).Error
	return
}

func (ur *UserRepository) PageUserByName(ctx context.Context, req response.UserPageReq) (users []model.User, err error) {
	var user []model.User
	var tx *gorm.DB
	if req.UserName != "" {
		tx = ur.DB.Where("username = ?", req.UserName).Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Find(&user)
	} else {
		tx = ur.DB.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Find(&user)
	}
	if tx.Error != nil {
		return
	}
	if len(user) == 0 {
		return
	}
	return user, nil

}

func (ur *UserRepository) CreateUser(ctx context.Context, user model.User) (bool, error) {
	result := ur.DB.Create(&user)
	return result.Error == nil, result.Error
}
