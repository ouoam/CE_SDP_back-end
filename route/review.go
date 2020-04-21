package route

import (
	"../model"
	"github.com/gofiber/fiber"
)

func ReviewRoute(route *fiber.Group) {
	route.Get("/", func(c *fiber.Ctx) {
		c.Send("test test")
	})

	route.Get("/:id", func(c *fiber.Ctx) {
		review := new(model.Review)
		GetID(c, review)
	})

	route.Post("/", func(c *fiber.Ctx) {
		review := new(model.Review)
		Post(c, review)
	})

	route.Put("/:id", func(c *fiber.Ctx) {
		review := new(model.Review)
		PutID(c, review)
	})
}
