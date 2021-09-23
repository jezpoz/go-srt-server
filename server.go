package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type StreamerStatus struct {
	Status string `json:"status"`
}

func main() {
	var streamers = map[string]StreamerStatus{}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title":     "Hello world",
			"Streamers": streamers,
		})
	})

	app.Get("/api/status", func(c *fiber.Ctx) error {
		return json.NewEncoder(c.Response().BodyWriter()).Encode(streamers)
	})

	app.Post("/api/status/:streamerId", func(c *fiber.Ctx) error {
		streamerId := c.Params("streamerId")
		if streamerId == "" {
			return c.Status(400).SendString("Missing streamer id")
		}
		status := new(StreamerStatus)
		if err := c.BodyParser(status); err != nil {
			return c.Status(400).SendString("Missing streamer status")
		}
		streamers[streamerId] = *status
		return c.SendString("Updated status")
	})

	app.Listen(":3000")
}
