package handler

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/domain/payload"
	"backend/internal/core/domain/response"
	"backend/internal/core/ports"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService ports.IUserService
}

var _ ports.IUserHandlers = (*UserHandler)(nil)

func NewUserHandlers(userService ports.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUserInfo(c *fiber.Ctx) error {
	userId := c.Query("user_id")
	user, err := h.userService.GetUserInfo(userId)
	if err != nil {
		return c.JSON(response.NewError("Unable to get user info", err))
	}
	userMap := payload.UserInfo{User: *user}
	return c.JSON(response.New("Get user info successfully", userMap))
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var email string
	var password string
	err := h.userService.Login(email, password)
	if err != nil {
		return err
	}
	return c.JSON(response.New("Login successfully"))
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	user := database.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.JSON(response.NewError("Unable to parse body"))
	}
	if user.Email == "" {
		return c.JSON(response.NewError("Email is not provided"))
	}
	if user.Password == "" {
		return c.JSON(response.NewError("Password is not provided"))
	}

	err := h.userService.Register(&user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return c.JSON(response.NewError("Email already exist"))
		}
		return c.JSON(response.NewError("Unable to register user"))
	}
	userMap := payload.UserInfo{User: payload.User{UserId: user.UserId, Email: user.Email}}

	return c.JSON(response.New("User registered successfully", userMap))
}
