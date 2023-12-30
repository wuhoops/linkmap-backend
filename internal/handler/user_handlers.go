package handler

import (
	"backend/internal/core/domain/database"
	"backend/internal/core/domain/payload"
	"backend/internal/core/domain/response"
	"backend/internal/core/ports"
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

// Get user by id
type getUserReq struct {
	UserId string `json:"user_id"`
}

func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	var req getUserReq
	req.UserId = c.Query("user_id")
	if req.UserId == "" {
		return c.Status(400).JSON(response.NewError("Unable to parse body"))
	}

	user, err := h.userService.GetUserById(req.UserId)
	if err != nil {
		return c.Status(400).JSON(response.NewError(err.Error()))
	}
	userRes := payload.User{
		UserId:   user.UserId,
		Email:    user.Email,
		Username: user.UserName,
	}
	res := map[string]interface{}{
		"user": userRes,
	}
	return c.JSON(response.New("Get user info successfully", res))
}

// Login
type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req loginReq
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(response.NewError("Unable to parse body"))
	}

	userReq := database.User{
		Email:    req.Email,
		Password: req.Password,
	}
	token, err := h.userService.Login(&userReq)
	if err != nil {
		return c.Status(400).JSON(response.NewError(err.Error()))
	}
	tokenString := *token
	userRes := payload.User{
		UserId:   userReq.UserId,
		Email:    userReq.Email,
		Username: userReq.UserName,
	}
	res := map[string]interface{}{
		"user":  userRes,
		"token": tokenString,
	}
	return c.JSON(response.New("Login successfully", res))
}

// Register
type registerReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req registerReq
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(response.NewError("Unable to parse body"))
	}
	if req.Email == "" {
		return c.JSON(response.NewError("Email is not provided"))
	}
	if req.Password == "" {
		return c.JSON(response.NewError("Password is not provided"))
	}
	if req.Username == "" {
		return c.JSON(response.NewError("Username is not provided"))
	}

	userReq := database.User{
		Email:    req.Email,
		Password: req.Password,
		UserName: req.Username,
	}
	err := h.userService.Register(&userReq)
	if err != nil {
		return c.Status(400).JSON(response.NewError(err.Error()))
	}
	userRes := payload.User{
		UserId:   userReq.UserId,
		Email:    userReq.Email,
		Username: userReq.UserName,
	}
	res := map[string]interface{}{
		"user": userRes,
	}
	return c.JSON(response.New("User registered successfully", res))
}

// Upsert username
type upsertUsernameReq struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
}

func (h *UserHandler) UpsertUserName(c *fiber.Ctx) error {
	var req upsertUsernameReq
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(response.NewError("Unable to parse body"))
	}
	err := h.userService.CreateUserName(req.UserId, req.Username)
	if err != nil {
		return c.Status(400).JSON(response.NewError(err.Error()))
	}
	return c.JSON(response.New("Create username successfully"))
}

// Get user by username
type getUserByUsernameReq struct {
	Username string `json:"username"`
}

func (h *UserHandler) GetUserByUsername(c *fiber.Ctx) error {
	var req getUserByUsernameReq
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(response.NewError("Unable to parse body"))
	}
	user, err := h.userService.GetUserByUsername(req.Username)
	if err != nil {
		return c.Status(400).JSON(response.NewError(err.Error()))
	}
	userRes := payload.User{
		UserId:   user.UserId,
		Email:    user.Email,
		Username: user.UserName,
	}
	res := map[string]interface{}{
		"user": userRes,
	}
	return c.JSON(response.New("Get user info successfully", res))
}
