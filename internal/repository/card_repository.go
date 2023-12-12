package repository

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/ports"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CardRepository struct {
	client *gorm.DB
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.ICardRepository = (*CardRepository)(nil)

func NewCardRepository(db *gorm.DB) *CardRepository {
	return &CardRepository{
		db,
	}
}

func (r *CardRepository) ListCard(id uuid.UUID) ([]database.Card, error) {
	//Here your code for fetching cards by id
	return nil, nil
}

func (r *CardRepository) CreateCard(payload *database.Card) (*database.Card, error) {
	return nil, nil
}
