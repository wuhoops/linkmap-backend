package handler

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/domain/response"
	"backend/internal/core/ports"

	fiber "github.com/gofiber/fiber/v2"
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
	return c.JSON(response.New("Get user info successfully", user))
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var email string
	var password string
	err := h.userService.Login(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	user := database.User{}
	if err := c.BodyParser(&user); err != nil {
		return &response.Error{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	if user.Email == "" || user.Password == "" {
		return &response.Error{
			Message: "Unable to parse body",
		}
	}

	err := h.userService.Register(&user)
	if err != nil {
		return &response.Error{
			Message: "Unable to register user",
			Err:     err,
		}
	}

	return c.JSON(response.New("User registered successfully", user))
}
