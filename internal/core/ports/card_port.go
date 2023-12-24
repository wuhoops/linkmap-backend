package ports

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/domain/payload"

	fiber "github.com/gofiber/fiber/v2"
)

type ICardService interface {
	ListCard(userId string) (*payload.CardList, error)
	CreateCard(payload *database.Card) error
}

type ICardRepository interface {
	ListCard(userId string) (*payload.CardList, error)
	CreateCard(payload *database.Card) error
}

type ICardHandlers interface {
	ListCard(c *fiber.Ctx) error
	CreateCard(c *fiber.Ctx) error
}
