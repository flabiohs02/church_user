package main

import (
	"log"

	"church_user/config"
	"church_user/handler"
	"church_user/repository"
	"church_user/routers"
	"church_user/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get database connection string from environment
	db := config.InitDb()
	// Dependency Injection
	userRepo := repository.NewGormUserRepository(db)
	userService := usecase.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Set up Fiber app
	app := fiber.New()

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", // You can restrict this to specific origins, e.g., "https://yourfrontend.com"
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: false,
	}))

	// Setup routes
	routers.SetupUserRoutes(app, userHandler)

	log.Println("Server listening on :8585")
	if err := app.Listen(":8585"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
