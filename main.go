package main

import (
	"unnantended/database"
	"unnantended/deck"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getDeck(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"success":  true,
		"database": database.Get(),
	})
}

func createDeck(c *fiber.Ctx) error {
	d := deck.NewDeck()
	database.Insert(d)

	return c.JSON(&fiber.Map{
		"success": true,
		"message": "Deck successfully created",
		"data":    d,
	})
}

func Setup() *fiber.App {
	database.Connect()

	app := fiber.New()
	app.Use(cors.New())
	setupRoutes(app)

	return app
}

func setupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Get("/deck", getDeck)
	v1.Put("/deck", createDeck)
}

func main() {
	app := Setup()
	app.Listen(":3000")
}
