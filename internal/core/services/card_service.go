package services

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/ports"

	"github.com/google/uuid"
)

type CardService struct {
	cardRepository ports.ICardRepository
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.ICardService = (*CardService)(nil)

func NewCardService(repository ports.ICardRepository) *CardService {
	return &CardService{
		cardRepository: repository,
	}
}

func (r *CardService) ListCard(userId uuid.UUID) ([]database.Card, error) {
	return nil, nil
}

func (r *CardService) CreateCard(payload *database.Card) (*database.Card, error) {
	return nil, nil
}
