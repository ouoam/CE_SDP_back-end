package route

import (
	"../model"
	"github.com/gofiber/fiber"
)

func MemberRoute(route *fiber.Group) {
	route.Get("/", func(c *fiber.Ctx) {
		c.Send("test test")
	})

	route.Get("/:id", func(c *fiber.Ctx) {
		member := new(model.Member)
		GetID(c, member)
	})

	route.Post("/", func(c *fiber.Ctx) {
		member := new(model.Member)
		Post(c, member)
	})

	route.Put("/:id", func(c *fiber.Ctx) {
		member := new(model.Member)
		PutID(c, member)
	})
}
