package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func PlaceRoute(route *fiber.Group) {
	route.Get("/", func(c *fiber.Ctx) {
		c.Send("test test")
	})

	route.Get("/:id", func(c *fiber.Ctx) {
		place := new(model.Place)
		controller.GetID(c, place)
	})

	route.Post("/", func(c *fiber.Ctx) {
		place := new(model.Place)
		controller.Post(c, place)
	})

	route.Put("/:id", func(c *fiber.Ctx) {
		place := new(model.Place)
		controller.PutID(c, place)
	})
}
