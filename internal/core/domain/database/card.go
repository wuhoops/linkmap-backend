package database

import "github.com/google/uuid"

type Card struct {
	CardId      uuid.UUID `json:"card_id" gorm:"primaryKey;not null"`
	Topic       string    `json:"card_number" gorm:"not null"`
	Description string    `json:"description"`
	Link        string    `json:"link" gorm:"not null"`
	OwnerID     int       `json:"owner_id" gorm:"not null"`
	Owner       User      `json:"owner" gorm:"foreignKey:OwnerID;references:UserId"`
}
