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
		AppName:       "Go Hexagonal LinkMap API",
	})
	router.Use(Cors)

	router.All("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Success": true,
			"Message": "Go Hexagonal LinkMap API",
		})
	})
	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.Get("/info", userHandler.GetUserInfo)
			user.Post("/register", userHandler.Register)
			user.Post("/login", userHandler.Login)
		}

		card := api.Group("/card")
		{
			card.Get("/list", cardHandler.ListCard)
			card.Post("/create", cardHandler.CreateCard)
			card.Patch("/update", cardHandler.EditCard)
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
