package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func ReviewRoute(route *fiber.Group) {
	route.Get("/:id", parseIntParams("id"), func(c *fiber.Ctx) {
		review := new(model.ReviewWithUser)
		review.Tour.SetValid(c.Locals("params_id").(int64))
		controller.Get(c, review, c.Locals("user_id").(int64))
	})

	route.Post("/:id", parseIntParams("id"), func(c *fiber.Ctx) {
		review := new(model.Review)
		review.Tour.SetValid(c.Locals("params_id").(int64))
		review.User.SetValid(c.Locals("user_id").(int64))
		controller.New(c, review)
	})

	route.Put("/:id", parseIntParams("id"), func(c *fiber.Ctx) {
		review := new(model.Review)
		review.Tour.SetValid(c.Locals("params_id").(int64))
		review.User.SetValid(c.Locals("user_id").(int64))
		controller.Update(c, review)
	})

	route.Delete("/:id", parseIntParams("id"), func(c *fiber.Ctx) {
		review := new(model.Review)
		review.Tour.SetValid(c.Locals("params_id").(int64))
		review.User.SetValid(c.Locals("user_id").(int64))
		controller.Delete(c, review)
	})

	route.Get("/", func(c *fiber.Ctx) {
		review := new(model.ReviewWithUser)
		controller.List(c, review, c.Locals("user_id").(int64))
	})
}
