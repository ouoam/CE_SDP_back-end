package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func MemberRoute(route *fiber.Group) {
	route.Get("/:id", controller.CheckLogin, func(c *fiber.Ctx) {
		member := new(model.Member)
		member.ID.SetValid(c.Locals("params_id").(int64))
		controller.GetID(c, member)
	})

	route.Post("/", func(c *fiber.Ctx) {
		member := new(model.Member)
		controller.Post(c, member)
	})

	route.Put("/:id", controller.CheckLogin, func(c *fiber.Ctx) {
		// todo check update own
		member := new(model.Member)
		member.ID.SetValid(c.Locals("params_id").(int64))
		controller.PutID(c, member)
	})

	route.Get("/", controller.CheckLogin, func(c *fiber.Ctx) {
		member := new(model.Member)
		controller.List(c, member)
	})
}
