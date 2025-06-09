package routers

import (
	"church_user/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	v1 := app.Group("/api/v1")
	{
		v1.Post("/users", userHandler.CreateUser)
		v1.Get("/users/:id", userHandler.GetUserByID)
		v1.Get("/users", userHandler.GetAllUsers)
		v1.Put("/users/:id", userHandler.UpdateUser)
		v1.Delete("/users/:id", userHandler.DeleteUser)
	}
}
