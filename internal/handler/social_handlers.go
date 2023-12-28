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
	userService   ports.IUserService
}

var _ ports.ISocialHandlers = (*SocialHandler)(nil)

func NewSocialHandlers(socialService ports.ISocialService, userService ports.IUserService) *SocialHandler {
	return &SocialHandler{
		socialService: socialService,
		userService:   userService,
	}
}

// List social
type listSocialReq struct {
	UserId string `json:"user_id"`
}

func (h *SocialHandler) ListSocial(c *fiber.Ctx) error {
	var req listSocialReq
	req.UserId = c.Query("user_id")
	if req.UserId == "" {
		return c.Status(400).JSON(response.NewError("Unable to parse body"))
	}
	user, err := h.userService.GetUserByUsername(req.UserId)
	if err != nil {
		return c.Status(400).JSON(response.NewError(err.Error()))
	}
	if req.UserId == "" {
		return c.Status(400).JSON(response.NewError("Unable to parse body"))
	}

	social, err := h.socialService.ListSocial(user.UserId)
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
		return c.Status(400).JSON(response.NewError(err.Error()))
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

// Update social
type updateSocialReq struct {
	SocialId string         `json:"social_id"`
	OwnerId  string         `json:"owner_id"`
	Topic    database.Topic `json:"topic"`
	Link     string         `json:"link"`
}

func (h *SocialHandler) UpdateSocial(c *fiber.Ctx) error {
	var req updateSocialReq
	err := c.BodyParser(&req)
	socialReq := database.Social{
		SocialId: req.SocialId,
		OwnerId:  req.OwnerId,
		Topic:    req.Topic,
		Link:     req.Link,
	}

	err = h.socialService.UpdateSocial(&socialReq)
	if err != nil {
		return c.Status(400).JSON(response.NewError(err.Error()))
	}

	socialRes := payload.Social{
		SocialId: socialReq.SocialId,
		OwnerId:  socialReq.OwnerId,
		Topic:    socialReq.Topic,
		Link:     socialReq.Link,
	}
	res := map[string]interface{}{
		"social": socialRes,
	}
	return c.JSON(response.New("Update social successfully", res))
}

// Delete social
type deleteSocialReq struct {
	SocialId string `json:"social_id"`
}

func (h *SocialHandler) DeleteSocial(c *fiber.Ctx) error {
	var req deleteSocialReq
	req.SocialId = c.Query("social_id")
	if req.SocialId == "" {
		return c.Status(400).JSON(response.NewError("Unable to parse body"))
	}

	err := h.socialService.DeleteSocial(req.SocialId)
	if err != nil {
		return c.Status(400).JSON(response.NewError(err.Error()))
	}

	return c.JSON(response.New("Delete social successfully"))
}
