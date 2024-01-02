package main

import (
	"backend/internal/core/services"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/util/config"
	"log/slog"
	"os"
)

func main() {
	// Init database
	db, err := repository.NewDB()
	redis, err := repository.InitRedis()

	//User
	userRepository := repository.NewUserRepository(db, redis)
	userService := services.NewUserService(userRepository)
	userHandler := handler.NewUserHandlers(userService)

	//Card
	cardRepository := repository.NewCardRepository(db)
	cardService := services.NewCardService(cardRepository, userRepository)
	cardHandler := handler.NewCardHandlers(cardService, userService)

	//Social
	socialRepository := repository.NewSocialRepository(db)
	socialService := services.NewSocialService(socialRepository)
	socialHandler := handler.NewSocialHandlers(socialService, userService)

	// Init router
	router, err := handler.NewRouter(
		*userHandler,
		*cardHandler,
		*socialHandler,
	)
	err = router.Serve(config.C.Address)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
