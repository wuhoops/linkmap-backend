package ports

import (
	"backend/internal/core/domain/database"
	"github.com/gofiber/fiber/v2"
)

type IUserService interface {
	GetUserById(userId string) (*database.User, error)
	Login(email string, password string) error
	Register(payload *database.User) error
	CreateUserName(userId string, userName string) error
}

type IUserRepository interface {
	GetUserById(userId string) (*database.User, error)
	Login(email string, password string) error
	Register(payload *database.User) error
	CreateUserName(userId string, userName string) error
}

type IUserHandlers interface {
	GetUserById(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	UpsertUserName(c *fiber.Ctx) error
}

type IServer interface {
	Initialize()
}
