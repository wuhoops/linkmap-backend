package services

import (
	"backend/internal/core/domain/database"
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

func (s *CardService) GetCardById(cardId string) (*database.Card, error) {
	card, err := s.cardRepository.GetCardById(cardId)
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

func (s *CardService) ListCard(userId string) ([]*database.Card, error) {
	if _, err := s.userRepository.GetUserById(userId); err != nil {
		return nil, err
	}

	cards, err := s.cardRepository.ListCard(userId)
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (s *CardService) EditCard(newCard *database.Card) (*database.Card, error) {
	if _, err := s.cardRepository.GetCardById(newCard.CardId); err != nil {
		return nil, err
	}

	card, err := s.cardRepository.EditCard(newCard)
	if err != nil {
		return nil, err
	}

	return card, nil
}

func (s *CardService) DeleteCard(cardId string) error {
	if _, err := s.cardRepository.GetCardById(cardId); err != nil {
		return err
	}

	err := s.cardRepository.DeleteCard(cardId)
	if err != nil {
		return err
	}
	return nil
}
