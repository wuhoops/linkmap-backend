package repository

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/ports"

	"github.com/google/uuid"
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

func (r *UserRepository) GetUserInfo(userId uuid.UUID) (*database.User, error) {
	var user database.User
	result := UserModel.Find(&user, userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Login(email string, password string) error {
	return nil
}

func (r *UserRepository) Register(payload *database.User) error {
	result := UserModel.Create(&payload)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
