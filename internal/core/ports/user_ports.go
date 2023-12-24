package ports

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/domain/payload"

	fiber "github.com/gofiber/fiber/v2"
)

type IUserService interface {
	GetUserInfo(userId string) (*payload.UserInfo, error)
	Login(email string, password string) error
	Register(payload *database.User) error
}

type IUserRepository interface {
	GetUserInfo(userId string) (*payload.UserInfo, error)
	Login(email string, password string) error
	Register(payload *database.User) error
}

type IUserHandlers interface {
	GetUserInfo(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type IServer interface {
	Initialize()
}
