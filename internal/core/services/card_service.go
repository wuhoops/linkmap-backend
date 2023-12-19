package services

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/ports"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CardService struct {
	cardRepository ports.ICardRepository
	userRepository ports.IUserRepository
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.ICardService = (*CardService)(nil)

func NewCardService(cardRepo ports.ICardRepository, userRepo ports.IUserRepository) *CardService {
	return &CardService{
		cardRepository: cardRepo,
		userRepository: userRepo,
	}
}

func (s *CardService) CreateCard(payload *database.Card) error {
	payload.CardId = uuid.New()
	err := s.cardRepository.CreateCard(payload)
	if err != nil {
		return err
	}

	user, err := s.userRepository.GetUserInfo(payload.OwnerID)
	if err != nil {
		return err
	}
	logrus.Info(user)
	payload.Owner = *user
	return nil
}

func (s *CardService) ListCard(userId uuid.UUID) ([]database.Card, error) {
	return nil, nil
}
