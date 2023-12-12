package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
	*fiber.App
}

func NewRouter(
	userHandler UserHandler,
	cardHandler CardHandler,
) (*Router, error) {

	router := fiber.New(fiber.Config{
		Prefork:       false,
		StrictRouting: true,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  5 * time.Second,
		AppName:       "GO Hexagonal Practice API",
	})
	router.Use(Cors)

	router.All("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Success": true,
			"Message": "Go Hexagonal Practice API",
		})
	})
	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.Post("/register", userHandler.Register)
			user.Post("/loggin", userHandler.Login)
		}
	}

	return &Router{
		router,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenPort string) error {
	return r.Listen(listenPort)
}
