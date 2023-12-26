package ports

import (
	"backend/internal/core/domain/database"
	fiber "github.com/gofiber/fiber/v2"
)

type ICardService interface {
	GetCardById(cardId string) (*database.Card, error)
	ListCard(userId string) ([]*database.Card, error)
	CreateCard(payload *database.Card) error
	EditCard(newCard *database.Card) (*database.Card, error)
	DeleteCard(cardId string) error
}

type ICardRepository interface {
	GetCardById(cardId string) (*database.Card, error)
	ListCard(userId string) ([]*database.Card, error)
	CreateCard(payload *database.Card) error
	EditCard(newCard *database.Card) (*database.Card, error)
	DeleteCard(cardId string) error
}

type ICardHandlers interface {
	GetCardById(c *fiber.Ctx) error
	ListCard(c *fiber.Ctx) error
	CreateCard(c *fiber.Ctx) error
	EditCard(c *fiber.Ctx) error
	DeleteCard(c *fiber.Ctx) error
}
