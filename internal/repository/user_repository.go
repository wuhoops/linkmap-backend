package repository

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/ports"
	"errors"
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

func (r *UserRepository) GetUserById(userId string) (*database.User, error) {
	var user *database.User
	result := r.client.Model(database.User{}).Where("user_id = ?", userId).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
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

func (r *UserRepository) CreateUserName(userId string, userName string) error {
	result := r.client.Model(database.User{}).Where("user_id = ?", userId).Update("user_name", userName)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *UserRepository) GetUserByUsername(userName string) (*database.User, error) {
	var user *database.User
	result := r.client.Model(database.User{}).Where("user_name = ?", userName).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
