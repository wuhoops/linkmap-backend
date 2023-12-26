package services

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/ports"
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

func (s *SocialService) ListSocial(userId string) (*database.Social, error) {
	user, err := s.socialRepository.ListSocial(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
