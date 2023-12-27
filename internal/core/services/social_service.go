package services

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/ports"
	"github.com/google/uuid"
)

type SocialService struct {
	socialRepository ports.ISocialRepository
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.ISocialService = (*SocialService)(nil)

func NewSocialService(repository ports.ISocialRepository) *SocialService {
	return &SocialService{
		socialRepository: repository,
	}
}

func (s *SocialService) ListSocial(userId string) ([]*database.Social, error) {
	social, err := s.socialRepository.ListSocial(userId)
	if err != nil {
		return nil, err
	}
	return social, nil
}

func (s *SocialService) AddSocial(social *database.Social) error {
	social.SocialId = uuid.New().String()
	err := s.socialRepository.AddSocial(social)
	if err != nil {
		return err
	}
	return nil
}
