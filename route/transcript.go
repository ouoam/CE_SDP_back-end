package route

import (
	"../model"
	"github.com/gofiber/fiber"
)

func TranscriptRoute(route *fiber.Group) {
	route.Get("/", func(c *fiber.Ctx) {
		c.Send("test test")
	})

	route.Get("/:id", func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		GetID(c, transcript)
	})

	route.Post("/", func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		Post(c, transcript)
	})

	route.Put("/:id", func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		PutID(c, transcript)
	})
}
