package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func ReviewRoute(route *fiber.Group) {
	route.Get("/", func(c *fiber.Ctx) {
		c.Send("test test")
	})

	route.Get("/:id", func(c *fiber.Ctx) {
		review := new(model.Review)
		controller.GetID(c, review)
	})

	route.Post("/", func(c *fiber.Ctx) {
		review := new(model.Review)
		controller.Post(c, review)
	})

	route.Put("/:id", func(c *fiber.Ctx) {
		review := new(model.Review)
		controller.PutID(c, review)
	})
}
