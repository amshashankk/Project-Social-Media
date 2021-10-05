package routes

import (
	"github.com/amshashankk/GoAuth/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	app.Post("/register", controllers.Register)
}
