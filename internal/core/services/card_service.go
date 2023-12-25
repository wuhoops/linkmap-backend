package services

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/domain/payload"
	"backend/internal/core/ports"
	"github.com/google/uuid"
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

func (s *CardService) CardInfo(cardId string) (*payload.Card, error) {
	card, err := s.cardRepository.CardInfo(cardId)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (s *CardService) CreateCard(payload *database.Card) error {
	payload.CardId = uuid.New().String()
	err := s.cardRepository.CreateCard(payload)
	if err != nil {
		return err
	}
	return nil
}

func (s *CardService) ListCard(userId string) (*payload.CardList, error) {
	if _, err := s.userRepository.GetUserInfo(userId); err != nil {
		return nil, err
	}

	cards, err := s.cardRepository.ListCard(userId)
	if err != nil {
		return nil, err
	}
	cardMap := payload.CardList{Card: cards}
	return &cardMap, nil
}

func (s *CardService) EditCard(newCard *payload.Card) error {
	if _, err := s.cardRepository.CardInfo(newCard.CardId); err != nil {
		return err
	}

	err := s.cardRepository.EditCard(newCard)
	if err != nil {
		return err
	}
	return nil
}
