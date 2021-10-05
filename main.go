package main

import (
	"github.com/amshashankk/GoAuth/database"
	"github.com/amshashankk/GoAuth/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	database.ConnectMongo("Shashank")

	routes.SetUp(app)

	app.Listen(":8080")
}
