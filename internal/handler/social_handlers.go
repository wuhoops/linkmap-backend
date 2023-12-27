package handler

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/domain/payload"
	"backend/internal/core/domain/response"
	"backend/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type SocialHandler struct {
	socialService ports.ISocialService
}

var _ ports.ISocialHandlers = (*SocialHandler)(nil)

func NewSocialHandlers(socialService ports.ISocialService) *SocialHandler {
	return &SocialHandler{
		socialService: socialService,
	}
}

// List social
type listSocialReq struct {
	UserId string `json:"user_id"`
}

func (h *SocialHandler) ListSocial(c *fiber.Ctx) error {
	var req getUserReq
	req.UserId = c.Query("user_id")
	if req.UserId == "" {
		return c.Status(400).JSON(response.NewError("Unable to parse body"))
	}

	social, err := h.socialService.ListSocial(req.UserId)
	if err != nil {
		return c.Status(400).JSON(response.NewError(err.Error()))
	}

	res := make(map[string]interface{})
	for _, social := range social {
		res[string(social.Topic)] = payload.Social{
			SocialId: social.SocialId,
			OwnerId:  social.OwnerId,
			Topic:    social.Topic,
			Link:     social.Link,
		}
	}
	return c.JSON(response.New("Get user info successfully", res))
}

// Add social
type addSocialReq struct {
	OwnerId string         `json:"owner_id"`
	Topic   database.Topic `json:"topic"`
	Link    string         `json:"link"`
}

func (h *SocialHandler) AddSocial(c *fiber.Ctx) error {
	var req addSocialReq
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(400).JSON(response.NewError("Unable to parse body"))
	}
	if string(req.Topic) != string(database.Instagram) {
		return c.Status(400).JSON(response.NewError("Invalid social type"))
	}
	socialReq := database.Social{
		OwnerId: req.OwnerId,
		Topic:   req.Topic,
		Link:    req.Link,
	}

	err = h.socialService.AddSocial(&socialReq)
	if err != nil {
		return c.Status(400).JSON(response.NewError(err.Error()))
	}

	return c.JSON(response.New("Add social successfully"))
}
