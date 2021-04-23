package main

import (
	"strconv"
	"strings"
	"unicode"
	"unnantended/database"
	"unnantended/deck"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getDeck(c *fiber.Ctx) error {
	deckId := c.Params("deckId")

	if deckId != "" {
		deck, err := database.GetByDeckId(deckId)
		if err != nil {
			c.SendString("deck not found")
			return c.SendStatus(404)
		}

		if deck.Remaining == 0 {
			c.SendString("deck has no remaining cards")
			return c.SendStatus(412)
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"data":    deck,
		})
	}

	return c.JSON(&fiber.Map{
		"success":  true,
		"database": database.Get(),
	})
}

var sanitizeParams = func(c rune) bool {
	return !unicode.IsLetter(c) && !unicode.IsNumber(c)
}

func createDeck(c *fiber.Ctx) error {
	var d map[string]deck.Deck
	var err error

	params := c.Query("cards")
	cards := strings.FieldsFunc(params, sanitizeParams)

	if len(params) < 0 || params == "" {
		cards = []string{}
		d, err = deck.NewDeck(cards, false)
	} else {
		d, err = deck.NewDeck(cards, false)
		c.JSON(&fiber.Map{
			"error": err,
			"cards": cards,
		})
	}
	if err != nil {
		c.SendString("error when creating deck")
		return c.SendStatus(400)
	}

	database.Insert(d)

	return c.JSON(&fiber.Map{
		"success": true,
		"message": "Deck successfully created",
		"data":    d,
	})
}

func drawCard(c *fiber.Ctx) error {
	deckId := c.Params("deckId")
	countParam := c.Params("count")

	d, err := database.GetByDeckId(deckId)
	count, _ := strconv.Atoi(countParam)

	if err != nil {
		c.SendString("deck not found")
		return c.SendStatus(404)
	}

	if count > d.Remaining {
		c.SendString("you cannot draw a card from the deck")
		return c.SendStatus(422)
	}

	cards, d := deck.Draw(count, d)

	deckBuild := map[string]deck.Deck{
		deckId: {
			Remaining: d.Remaining,
			Shuffled:  d.Shuffled,
			Cards:     d.Cards,
		},
	}

	database.Insert(deckBuild)

	return c.JSON(&fiber.Map{
		"success": true,
		"message": "cards are successfully drawn",
		"data":    cards,
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
	v1.Get("/deck/:deckId", getDeck)
	v1.Put("/deck/:deckId/:count", drawCard)
	v1.Put("/deck", createDeck)
	v1.Put("/deck?cards=:cards", createDeck)
}

func main() {
	app := Setup()
	app.Listen(":3000")
}
