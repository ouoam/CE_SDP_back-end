package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func MemberRoute(route *fiber.Group) {
	route.Get("/:id", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		member := new(model.Member)
		member.ID.SetValid(c.Locals("params_id").(int64))
		controller.Get(c, member)
	})

	route.Post("/", func(c *fiber.Ctx) {
		member := new(model.Member)
		controller.New(c, member)
	})

	route.Put("/", controller.CheckLogin, func(c *fiber.Ctx) {
		// todo check update own
		member := new(model.Member)
		member.ID.SetValid(c.Locals("user_id").(int64))
		controller.Update(c, member)
	})

	route.Get("/", controller.CheckLogin, func(c *fiber.Ctx) {
		member := new(model.Member)
		controller.List(c, member)
	})
}
