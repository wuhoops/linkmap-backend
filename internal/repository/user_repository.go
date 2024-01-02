package repository

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/ports"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	client *gorm.DB
	redis  *redis.Client
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.IUserRepository = (*UserRepository)(nil)

func NewUserRepository(db *gorm.DB, redis *redis.Client) *UserRepository {
	return &UserRepository{
		db, redis,
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

func (r *UserRepository) Login(payload *database.User) error {
	result := r.client.Model(database.User{}).Where("email = ?", payload.Email).First(&payload)
	if result.Error != nil {
		return result.Error
	}
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

func (r *UserRepository) SetRefreshToken(key string, refreshToken string, expiration time.Duration) error {
	ctx := context.Background()
	err := r.redis.Set(ctx, key, refreshToken, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetRefreshToken(username string) (string, error) {
	ctx := context.Background()
	val, err := r.redis.Get(ctx, username).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
