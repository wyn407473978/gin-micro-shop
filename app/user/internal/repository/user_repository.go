package repository

import (
	"context"
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
