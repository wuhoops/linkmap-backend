package handler

import (
	"backend/internal/core/domain/database"
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
	var card database.Card
	if err := c.BodyParser(&card); err != nil {
		return &response.Error{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	if card.OwnerID == "" || card.Topic == "" || card.Link == "" {
		return c.JSON(response.NewError("Missing required field."))

	}

	err := h.cardService.CreateCard(&card)
	if err != nil {
		return c.JSON(response.NewError("Unable to create card.", err))
	}

	return c.JSON(response.New("Card created successfully", card))
}

func (h *CardHandler) ListCard(c *fiber.Ctx) error {
	var userId string
	userId = c.Query("user_id")

	cards, err := h.cardService.ListCard(userId)
	if err != nil {
		return c.JSON(response.NewError("Unable to list card", err))
	}

	return c.JSON(response.New("List card successfully", cards))
}
