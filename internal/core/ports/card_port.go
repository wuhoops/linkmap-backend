package ports

import (
	"backend/internal/core/domain/database"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ICardService interface {
	ListCard(userId uuid.UUID) ([]database.Card, error)
	CreateCard(payload *database.Card) (*database.Card, error)
}

type ICardRepository interface {
	ListCard(userId uuid.UUID) ([]database.Card, error)
	CreateCard(payload *database.Card) (*database.Card, error)
}

type ICardHandlers interface {
	ListCard(c *fiber.Ctx) error
	CreateCard(c *fiber.Ctx) error
}
