package routes

import (
	// "go-fiber-api/internal/handlers"
	controllers "go-fiber-api/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/users", controllers.GetUsers)
	api.Post("/users", controllers.CreateUser)
	api.Post("/users-new", controllers.CreateUserNew)
	api.Get("/user/:id", controllers.GetUserByID)
	api.Get("/user/name/:name", controllers.GetUserByName)
	api.Put("/user/:id", controllers.UpdateUser)
	api.Delete("/user/:id", controllers.DeleteUser)
	// api.Get("/users", handlers.GetUsers)
	// api.Get("/users/:id", handlers.GetUser)
}
