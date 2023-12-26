package handler

import (
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

	_, err := h.socialService.ListSocial(req.UserId)
	if err != nil {
		return c.Status(400).JSON(response.NewError(err.Error()))
	}
	//userRes := payload.User{
	//	UserId: user.UserId,
	//	Email:  user.Email,
	//}
	res := map[string]interface{}{
		"user": "null",
	}
	return c.JSON(response.New("Get user info successfully", res))
}
