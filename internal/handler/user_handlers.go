package handler

import (
	"goHexagonalBlog/internal/core/domain/payload"
	"goHexagonalBlog/internal/core/domain/response"
	"goHexagonalBlog/internal/core/ports"

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

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var email string
	var password string
	//Extract the body and get the email and password
	err := h.userService.Login(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var body payload.NewUser
	if err := c.BodyParser(&body); err != nil {
		return &response.Error{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	err := h.userService.Register(body)
	if err != nil {
		return &response.Error{
			Message: "Unable to register user",
			Err:     err,
		}
	}
	return c.JSON(fiber.Map{"Success": true})
}
