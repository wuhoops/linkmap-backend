package repository

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/ports"
	"gorm.io/gorm"
)

type SocialRepository struct {
	client *gorm.DB
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.ISocialRepository = (*SocialRepository)(nil)

func NewSocialRepository(db *gorm.DB) *SocialRepository {
	return &SocialRepository{
		db,
	}
}

func (r *SocialRepository) ListSocial(userId string) (*database.Social, error) {
	var user *database.Social
	result := r.client.Model(database.User{}).Where("user_id = ?", userId).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
