package handler

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/domain/payload"
	"backend/internal/core/domain/response"
	"backend/internal/core/ports"
	"github.com/gofiber/fiber/v2"
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
		return c.Status(400).JSON(response.NewError("Unable to parse body", err.Error()))
	}

	if card.OwnerID == "" || card.Topic == "" || card.Link == "" {
		return c.JSON(response.NewError("Missing required field"))
	}

	err := h.cardService.CreateCard(&card)
	if err != nil {
		return c.Status(400).JSON(response.NewError("Unable to create card.", err))
	}

	return c.JSON(response.New("Card created successfully", card))
}

func (h *CardHandler) CardInfo(c *fiber.Ctx) error {
	var cardId string
	cardId = c.Query("card_id")

	card, err := h.cardService.CardInfo(cardId)
	if err != nil {
		return c.Status(400).JSON(response.NewError("Unable to get card info", err.Error()))
	}

	return c.JSON(response.New("Get card info successfully", card))
}

func (h *CardHandler) ListCard(c *fiber.Ctx) error {
	var userId string
	userId = c.Query("user_id")

	cards, err := h.cardService.ListCard(userId)
	if err != nil {
		return c.Status(400).JSON(response.NewError("Unable to list card", err.Error()))
	}

	return c.JSON(response.New("List card successfully", cards))
}

func (h *CardHandler) EditCard(c *fiber.Ctx) error {
	card := payload.Card{}
	if err := c.BodyParser(&card); err != nil {
		return c.Status(400).JSON(response.NewError("Unable to parse body", err.Error()))
	}
	err := h.cardService.EditCard(&card)
	if err != nil {
		return c.Status(400).JSON(response.NewError("Unable to edit card", err.Error()))
	}
	cardMap := payload.CardEdit{Card: card}
	return c.JSON(response.New("Edit card successfully", cardMap))
}

func (h *CardHandler) DeleteCard(c *fiber.Ctx) error {
	var cardId string
	cardId = c.Query("card_id")

	err := h.cardService.DeleteCard(cardId)
	if err != nil {
		return c.Status(400).JSON(response.NewError("Unable to delete card", err.Error()))
	}

	return c.JSON(response.New("Delete card successfully", nil))
}
