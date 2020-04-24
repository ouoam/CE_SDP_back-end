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

	route.Post("/", func(c *fiber.Ctx) {
		tour := new(model.Tour)
		controller.Post(c, tour)
	})

	route.Put("/:id", func(c *fiber.Ctx) {
		tour := new(model.Tour)
		controller.PutID(c, tour)
	})

	route.Get("/", func(c *fiber.Ctx) {
		tour := new(model.Tour)
		controller.List(c, tour)
	})
}
