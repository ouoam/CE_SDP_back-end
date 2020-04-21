package route

import (
	"../model"
	"github.com/gofiber/fiber"
)

func TourRoute(route *fiber.Group) {
	route.Get("/", func(c *fiber.Ctx) {
		c.Send("test test")
	})

	route.Get("/:id", func(c *fiber.Ctx) {
		tour := new(model.Tour)
		GetID(c, tour)
	})

	route.Post("/", func(c *fiber.Ctx) {
		tour := new(model.Tour)
		Post(c, tour)
	})

	route.Put("/:id", func(c *fiber.Ctx) {
		tour := new(model.Tour)
		PutID(c, tour)
	})
}
