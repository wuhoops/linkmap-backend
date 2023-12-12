package database

import "github.com/google/uuid"

type User struct {
	UserId   uuid.UUID `json:"user_id" gorm:"primaryKey;not null"`
	Email    string    `json:"email" gorm:"not null"`
	Password string    `json:"password" gorm:"not null"`
	Cards    []Card    `json:"cards" gorm:"foreignKey:OwnerID;references:UserId"`
}
