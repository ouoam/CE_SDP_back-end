package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func MessageRoute(route *fiber.Group) {
	route.Get("/:id", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		message := new(model.MessageWithMe)
		controller.List(c, message, c.Locals("user_id").(int64), c.Locals("params_id").(int64))
	})

	route.Post("/:id", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		message := new(model.Message)
		message.From.SetValid(c.Locals("user_id").(int64))
		message.To.SetValid(c.Locals("params_id").(int64))
		controller.Post(c, message)
	})

	route.Get("/", controller.CheckLogin, func(c *fiber.Ctx) {
		message := new(model.MessageListMe)
		controller.List(c, message, c.Locals("user_id").(int64))
	})
}
