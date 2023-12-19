package ports

import (
	"backend/internal/core/domain/database"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type IUserService interface {
	GetUserInfo(userId uuid.UUID) (*database.User, error)
	Login(email string, password string) error
	Register(payload *database.User) error
}

type IUserRepository interface {
	GetUserInfo(userId uuid.UUID) (*database.User, error)
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
