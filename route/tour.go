package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func TourRoute(route *fiber.Group) {
	route.Get("/:id", func(c *fiber.Ctx) {
		tour := new(model.Tour)
		controller.GetID(c, tour)
	})

	route.Post("/", controller.CheckLogin, func(c *fiber.Ctx) {
		tour := new(model.Tour)
		controller.Post(c, tour)
	})

	route.Put("/:id", controller.CheckLogin, func(c *fiber.Ctx) {
		tour := new(model.Tour)
		controller.PutID(c, tour)
	})

	route.Get("/", func(c *fiber.Ctx) {
		tour := new(model.Tour)
		controller.List(c, tour)
	})
}
