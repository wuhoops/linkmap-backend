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
	userService ports.IUserService
}

var _ ports.ICardHandlers = (*CardHandler)(nil)

func NewCardHandlers(cardService ports.ICardService, userService ports.IUserService) *CardHandler {
	return &CardHandler{
		cardService: cardService,
		userService: userService,
	}
}

// Create card
type creatCardReq struct {
	Topic       string `json:"topic"`
	Description string `json:"description"`
	Link        string `json:"link"`
	OwnerId     string `json:"owner_id"`
}

func (h *CardHandler) CreateCard(c *fiber.Ctx) error {
	var req creatCardReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(response.NewError("Unable to parse body", err.Error()))
	}

	if req.OwnerId == "" || req.Topic == "" || req.Link == "" {
		return c.Status(400).JSON(response.NewError("Missing required field"))
	}

	cardReq := database.Card{
		OwnerId:     req.OwnerId,
		Topic:       req.Topic,
		Description: req.Description,
		Link:        req.Link,
	}

	err := h.cardService.CreateCard(&cardReq)
	if err != nil {
		return c.Status(400).JSON(response.NewError(err.Error()))
	}

	cardRes := payload.Card{
		CardId:      cardReq.CardId,
		OwnerId:     cardReq.OwnerId,
		Topic:       cardReq.Topic,
		Description: cardReq.Description,
		Link:        cardReq.Link,
	}
	res := map[string]interface{}{
		"card": cardRes,
	}
	return c.JSON(response.New("Card created successfully", res))
}

// Get card by id
type getCardReq struct {
	CardId string `json:"card_id"`
}

func (h *CardHandler) GetCardById(c *fiber.Ctx) error {
	var req getCardReq
	req.CardId = c.Query("card_id")
	if req.CardId == "" {
		return c.Status(400).JSON(response.NewError("Unable to parse body"))
	}

	card, err := h.cardService.GetCardById(req.CardId)
	if err != nil {
		return c.Status(400).JSON(response.NewError(err.Error()))
	}

	cardRes := payload.Card{
		CardId:      card.CardId,
		OwnerId:     card.OwnerId,
		Topic:       card.Topic,
		Description: card.Description,
		Link:        card.Link,
	}
	res := map[string]interface{}{
		"card": cardRes,
	}
	return c.JSON(response.New("Get card info successfully", res))
}

// List card
type listCardReq struct {
	Username string `json:"username"`
}

func (h *CardHandler) ListCard(c *fiber.Ctx) error {
	var req listCardReq
	req.Username = c.Query("username")
	if req.Username == "" {
		return c.Status(400).JSON(response.NewError("Unable to parse body"))
	}
	user, err := h.userService.GetUserByUsername(req.Username)
	if err != nil {
		return c.Status(400).JSON(response.NewError(err.Error()))
	}
	cards, err := h.cardService.ListCard(user.UserId)
	if err != nil {
		return c.Status(400).JSON(response.NewError(err.Error()))
	}

	cardList := make([]payload.Card, 0)
	for _, card := range cards {
		cardList = append(cardList, payload.Card{
			CardId:      card.CardId,
			OwnerId:     card.OwnerId,
			Topic:       card.Topic,
			Description: card.Description,
			Link:        card.Link,
		})
	}
	res := map[string]interface{}{
		"card_list": cardList,
	}
	return c.JSON(response.New("List card successfully", res))
}

// Edit card
type editCardReq struct {
	CardId      string `json:"card_id"`
	Topic       string `json:"topic"`
	Description string `json:"description"`
	Link        string `json:"link"`
	OwnerId     string `json:"owner_id"`
}

func (h *CardHandler) EditCard(c *fiber.Ctx) error {
	var req editCardReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(response.NewError("Unable to parse body", err.Error()))
	}
	cardReq := database.Card{
		CardId:      req.CardId,
		OwnerId:     req.OwnerId,
		Topic:       req.Topic,
		Description: req.Description,
		Link:        req.Link,
	}
	card, err := h.cardService.EditCard(&cardReq)
	if err != nil {
		return c.Status(400).JSON(response.NewError("Unable to edit card", err.Error()))
	}
	cardRes := payload.Card{
		CardId:      card.CardId,
		OwnerId:     card.OwnerId,
		Topic:       card.Topic,
		Description: card.Description,
		Link:        card.Link,
	}
	res := map[string]interface{}{
		"card": cardRes,
	}
	return c.JSON(response.New("Edit card successfully", res))
}

// Delete card
type deleteCardReq struct {
	CardId string `json:"card_id"`
}

func (h *CardHandler) DeleteCard(c *fiber.Ctx) error {
	var res deleteCardReq
	res.CardId = c.Query("card_id")
	if res.CardId == "" {
		return c.Status(400).JSON(response.NewError("Unable to parse body"))
	}

	err := h.cardService.DeleteCard(res.CardId)
	if err != nil {
		return c.Status(400).JSON(response.NewError(err.Error()))
	}

	return c.JSON(response.New("Delete card successfully"))
}
