package handler

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/domain/response"
	"backend/internal/core/ports"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	var card database.Card
	if err := c.BodyParser(&card); err != nil {
		return &response.Error{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	emptyUUID, _ := uuid.Parse("")
	if card.OwnerID == emptyUUID || card.Topic == "" || card.Link == "" {
		return &response.Error{
			Message: "Missing fields",
		}
	}

	err := h.cardService.CreateCard(&card)
	if err != nil {
		return &response.Error{
			Message: "Unable to create card",
			Err:     err,
		}
	}

	return c.JSON(response.New("Card created successfully", card))}

func (h *CardHandler) ListCard(c *fiber.Ctx) error {
	return nil
}
