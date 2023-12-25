package ports

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/domain/payload"

	"github.com/gofiber/fiber/v2"
)

type IUserService interface {
	GetUserInfo(userId string) (*payload.User, error)
	Login(email string, password string) error
	Register(payload *database.User) error
	CreateUserName(userId string, userName string) error
}

type IUserRepository interface {
	GetUserInfo(userId string) (*payload.User, error)
	Login(email string, password string) error
	Register(payload *database.User) error
	CreateUserName(userId string, userName string) error
}

type IUserHandlers interface {
	GetUserInfo(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	UpsertUserName(c *fiber.Ctx) error
}

type IServer interface {
	Initialize()
}
