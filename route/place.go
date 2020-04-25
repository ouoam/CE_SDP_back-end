package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func PlaceRoute(route *fiber.Group) {
	route.Get("/:id", func(c *fiber.Ctx) {
		place := new(model.Place)
		controller.GetID(c, place)
	})

	route.Post("/", controller.CheckLogin, func(c *fiber.Ctx) {
		place := new(model.Place)
		controller.Post(c, place)
	})

	route.Put("/:id", controller.CheckLogin, func(c *fiber.Ctx) {
		place := new(model.Place)
		controller.PutID(c, place)
	})

	route.Get("/", func(c *fiber.Ctx) {
		place := new(model.Place)
		controller.List(c, place)
	})
}
