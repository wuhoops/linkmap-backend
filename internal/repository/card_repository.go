package repository

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/ports"
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

func (r *CardRepository) GetCardById(cardId string) (*database.Card, error) {
	var card *database.Card
	result := r.client.Model(database.Card{}).First(&card, "card_id = ?", cardId)
	if result.Error != nil {
		return nil, result.Error
	}
	return card, nil
}

func (r *CardRepository) ListCard(id string) ([]*database.Card, error) {
	var cardList []*database.Card
	result := r.client.Model(database.Card{}).Where("owner_id = ?", id).Find(&cardList)
	if result.Error != nil {
		return nil, result.Error
	}
	return cardList, nil
}

func (r *CardRepository) EditCard(newCard *database.Card) (*database.Card, error) {
	result := r.client.Model(database.Card{}).Where("card_id = ?", newCard.CardId).Updates(newCard)
	if result.Error != nil {
		return nil, result.Error
	}
	return newCard, nil
}

func (r *CardRepository) DeleteCard(cardId string) error {
	result := r.client.Model(database.Card{}).Where("card_id = ?", cardId).Delete(&database.Card{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
