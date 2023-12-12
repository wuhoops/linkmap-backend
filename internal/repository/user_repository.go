package repository

import (
	"backend/internal/core/domain/database"
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

func (r *UserRepository) Login(email string, password string) error {
	//Here your code for login in mongo database
	return nil
}

func (r *UserRepository) Register(payload *database.User) (*database.User, error) {
	result := UserModel.Create(&payload)
	if result.Error != nil {
		return nil, result.Error
	}
	return payload, nil
}
