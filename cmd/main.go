package main

import (
	"context"
	"backend/internal/core/services"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/util/config"
	"log/slog"
	"os"
)

func main() {
	// Init database
	ctx := context.Background()
	db, err := repository.NewDB(ctx)

	//User
	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handler.NewUserHandlers(userService)

	//Card
	cardRepository := repository.NewCardRepository(db)
	cardService := services.NewCardService(cardRepository, userRepository)
	cardHandler := handler.NewCardHandlers(cardService)

	// Init router
	router, err := handler.NewRouter(
		*userHandler,
		*cardHandler,
	)
	err = router.Serve(config.C.Address)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
