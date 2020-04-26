package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func TourRoute(route *fiber.Group) {
	route.Get("/:id", func(c *fiber.Ctx) {
		tour := new(model.TourDetail)
		tour.ID.SetValid(c.Locals("params_id").(int64))
		controller.GetID(c, tour)
	})

	route.Post("/", controller.CheckLogin, func(c *fiber.Ctx) {
		tour := new(model.Tour)
		controller.Post(c, tour)
	})

	route.Put("/:id", controller.CheckLogin, func(c *fiber.Ctx) {
		// todo update own
		tour := new(model.Tour)
		tour.ID.SetValid(c.Locals("params_id").(int64))
		controller.PutID(c, tour)
	})

	route.Get("/", func(c *fiber.Ctx) {
		tour := new(model.TourDetail)
		controller.List(c, tour)
	})
}
