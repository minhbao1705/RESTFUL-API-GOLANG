package main

import (
	"user-management-api/internal/app"
	"user-management-api/internal/config"
)

func main() {
	// Initialize configuration
	cfg := config.NewConfig()

	// Initialize application
	application := app.NewApplication(cfg)

	// // Initialize repository
	// userRepo := repository.NewInMemoryUserRepository()

	// // Initialize service
	// userService := service.NewUserService(userRepo)

	// // Initialize handler
	// userHandler := handler.NewUserHandler(userService)

	// // Initialize routes
	// userRoutes := routes.NewUserRoutes(userHandler)

	// r := gin.Default()

	// routes.RegisterRoutes(r, userRoutes)

	// start server
	if err := application.Run(); err != nil {
		panic(err)
	}
}
