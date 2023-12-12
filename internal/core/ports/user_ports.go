package ports

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/domain/payload"

	fiber "github.com/gofiber/fiber/v2"
)

type IUserService interface {
	Login(email string, password string) error
	Register(user *payload.NewUser) (*database.User, error)
}

type IUserRepository interface {
	Login(email string, password string) error
	Register(user *database.User) (*database.User, error)
}

type IUserHandlers interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type IServer interface {
	Initialize()
}
