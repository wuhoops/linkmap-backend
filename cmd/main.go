package main

import (
	"context"
	"goHexagonalBlog/internal/core/services"
	"goHexagonalBlog/internal/handler"
	"goHexagonalBlog/internal/repository"
	"goHexagonalBlog/internal/util/config"
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

	// Init router
	router, err := handler.NewRouter(
		*userHandler,
	)
	err = router.Serve(config.C.Address)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
