package ports

import (
	"backend/internal/core/domain/database"
	"github.com/gofiber/fiber/v2"
	"time"
)

type IUserService interface {
	GetUserById(userId string) (*database.User, error)
	Login(payload *database.User) error
	Register(payload *database.User) error
	CreateUserName(userId string, userName string) error
	GetUserByUsername(userName string) (*database.User, error)
	SetRefreshToken(username string, refreshToken string, expiration time.Duration) error
	GenerateToken(username string, expiration time.Time) (string, error)
}

type IUserRepository interface {
	GetUserById(userId string) (*database.User, error)
	Login(payload *database.User) error
	Register(payload *database.User) error
	CreateUserName(userId string, userName string) error
	GetUserByUsername(userName string) (*database.User, error)
	SetRefreshToken(key string, refreshToken string, expiration time.Duration) error
	GetRefreshToken(key string) (string, error)
}

type IUserHandlers interface {
	GetUserById(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	UpsertUserName(c *fiber.Ctx) error
	GetUserByUsername(c *fiber.Ctx) error
}

type IServer interface {
	Initialize()
}
