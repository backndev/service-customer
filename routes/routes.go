package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"service-customer/controllers"
)

func SetupRoutes(app *fiber.App) {
	home := app.Group("/", logger.New())
	home.Get("/", controllers.Home)

	api := app.Group("/api", logger.New())

	// Auth
	//app.Use(middlewares.IsAuthenticated)
	auth := api.Group("/login")
	auth.Post("/", controllers.Login)

	// User
	users := api.Group("/users")
	users.Get("/", controllers.GetAllUser)
	user := api.Group("/user")
	user.Get("/:id", controllers.GetUser)
	user.Post("/", controllers.CreateUser)
}
