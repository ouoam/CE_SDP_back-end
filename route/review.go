package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func ReviewRoute(route *fiber.Group) {
	route.Get("/:id", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		review := new(model.Review)
		review.Tour.SetValid(c.Locals("params_id").(int64))
		review.User.SetValid(c.Locals("user_id").(int64))
		controller.GetID(c, review)
	})

	route.Post("/", controller.CheckLogin, func(c *fiber.Ctx) {
		review := new(model.Review)
		controller.Post(c, review)
	})

	route.Put("/:id", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		review := new(model.Review)
		review.Tour.SetValid(c.Locals("params_id").(int64))
		review.User.SetValid(c.Locals("user_id").(int64))
		controller.PutID(c, review)
	})

	route.Get("/", controller.CheckLogin, func(c *fiber.Ctx) {
		review := new(model.Review)
		review.User.SetValid(c.Locals("user_id").(int64))
		controller.List(c, review)
	})
}
