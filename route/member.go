package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func MemberRoute(route *fiber.Group) {
	route.Get("/:id", func(c *fiber.Ctx) {
		member := new(model.Member)
		controller.GetID(c, member)
	})

	route.Post("/", func(c *fiber.Ctx) {
		member := new(model.Member)
		controller.Post(c, member)
	})

	route.Put("/:id", func(c *fiber.Ctx) {
		member := new(model.Member)
		controller.PutID(c, member)
	})

	route.Get("/", func(c *fiber.Ctx) {
		member := new(model.Member)
		controller.List(c, member)
	})
}
