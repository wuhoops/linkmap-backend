package repository

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/domain/payload"
	"backend/internal/core/ports"
	"errors"

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

func (r *CardRepository) CreateCard(payload *database.Card) error {
	result := r.client.Model(database.Card{}).Create(&payload)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CardRepository) ListCard(id string) (*payload.CardList, error) {
	user := database.User{}
	result1 := r.client.Model(database.User{}).First(&user, "user_id = ?", id)
	if result1.Error != nil {
		if errors.Is(result1.Error, gorm.ErrRecordNotFound) {
			return nil, result1.Error
		}
		return nil, result1.Error
	}

	cardList := []payload.Card{}
	result2 := r.client.Model(database.Card{}).Where("owner_id = ?", id).Take(&cardList)
	if result2.Error != nil {
		return nil, result2.Error
	}

	cardMap := payload.CardList{Card: cardList}
	return &cardMap, nil
}
