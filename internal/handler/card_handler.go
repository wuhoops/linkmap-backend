package handler

import (
	"backend/internal/core/domain/response"
	"backend/internal/core/ports"

	fiber "github.com/gofiber/fiber/v2"
)

type CardHandler struct {
	cardService ports.ICardService
}

var _ ports.ICardHandlers = (*CardHandler)(nil)

func NewCardHandlers(cardService ports.ICardService) *CardHandler {
	return &CardHandler{
		cardService: cardService,
	}
}

func (h *CardHandler) CreateCard(c *fiber.Ctx) error {
	return nil
}

func (h *CardHandler) ListCard(c *fiber.Ctx) error {

	return c.JSON(response.New("User registered successfully"))
}
