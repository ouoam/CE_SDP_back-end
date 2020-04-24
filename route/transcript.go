package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func TranscriptRoute(route *fiber.Group) {
	route.Get("/", func(c *fiber.Ctx) {
		c.Send("test test")
	})

	route.Get("/:id", func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		controller.GetID(c, transcript)
	})

	route.Post("/", func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		controller.Post(c, transcript)
	})

	route.Put("/:id", func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		controller.PutID(c, transcript)
	})

	route.Get("/", func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		controller.List(c, transcript)
	})
}
