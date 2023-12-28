package repository

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/ports"
	"errors"
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

func (r *SocialRepository) ListSocial(userId string) ([]*database.Social, error) {
	var social []*database.Social
	result := r.client.Model(database.Social{}).Where("owner_id = ?", userId).Order("topic").Find(&social)
	if result.Error != nil {
		return nil, result.Error
	}
	return social, nil
}

func (r *SocialRepository) AddSocial(social *database.Social) error {
	result := r.client.Model(database.Social{}).Create(social)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *SocialRepository) UpdateSocial(social *database.Social) error {
	result := r.client.Model(database.Social{}).Where("social_id = ?", social.SocialId).Updates(social)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no social with ID " + social.SocialId + " found")
	}
	return nil
}

func (r *SocialRepository) DeleteSocial(socialId string) error {
	result := r.client.Model(database.Social{}).Where("social_id = ?", socialId).Delete(&database.Social{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no social with ID " + socialId + " found")
	}
	return nil
}
