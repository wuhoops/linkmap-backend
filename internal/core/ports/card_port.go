package ports

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/domain/payload"

	fiber "github.com/gofiber/fiber/v2"
)

type ICardService interface {
	CardInfo(cardId string) (*payload.Card, error)
	ListCard(userId string) (*payload.CardList, error)
	CreateCard(payload *database.Card) error
	EditCard(newCard *database.Card) (*database.Card, error)
}

type ICardRepository interface {
	CardInfo(cardId string) (*payload.Card, error)
	ListCard(userId string) ([]payload.Card, error)
	CreateCard(payload *database.Card) error
	EditCard(newCard *database.Card) (*database.Card, error)
}

type ICardHandlers interface {
	CardInfo(c *fiber.Ctx) error
	ListCard(c *fiber.Ctx) error
	CreateCard(c *fiber.Ctx) error
	EditCard(c *fiber.Ctx) error
}
