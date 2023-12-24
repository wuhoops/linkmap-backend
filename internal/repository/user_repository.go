package repository

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/domain/payload"
	"backend/internal/core/ports"

	"gorm.io/gorm"
)

type UserRepository struct {
	client *gorm.DB
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.IUserRepository = (*UserRepository)(nil)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) GetUserInfo(userId string) (*payload.UserInfo, error) {
	var user payload.UserInfo
	result := r.client.Model(database.User{}).Where("user_id = ?", userId).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Login(email string, password string) error {
	return nil
}

func (r *UserRepository) Register(payload *database.User) error {
	result := r.client.Model(database.User{}).Create(&payload)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
