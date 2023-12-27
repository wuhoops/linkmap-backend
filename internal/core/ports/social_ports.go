package ports

import (
	"backend/internal/core/domain/database"
	"github.com/gofiber/fiber/v2"
)

type ISocialService interface {
	ListSocial(userId string) ([]*database.Social, error)
	AddSocial(social *database.Social) error
	UpdateSocial(social *database.Social) error
	DeleteSocial(socialId string) error
}

type ISocialRepository interface {
	ListSocial(userId string) ([]*database.Social, error)
	AddSocial(social *database.Social) error
	UpdateSocial(social *database.Social) error
	DeleteSocial(socialId string) error
}

type ISocialHandlers interface {
	ListSocial(c *fiber.Ctx) error
	AddSocial(c *fiber.Ctx) error
	UpdateSocial(c *fiber.Ctx) error
	DeleteSocial(c *fiber.Ctx) error
}
