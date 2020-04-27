package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func PlaceRoute(route *fiber.Group) {
	route.Get("/:id", parseIntParams("id"), func(c *fiber.Ctx) {
		place := new(model.Place)
		place.ID.SetValid(c.Locals("params_id").(int64))
		controller.GetID(c, place)
	})

	route.Post("/", controller.CheckLogin, func(c *fiber.Ctx) {
		place := new(model.Place)
		controller.Post(c, place)
	})

	route.Put("/:id", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		place := new(model.Place)
		place.ID.SetValid(c.Locals("params_id").(int64))
		controller.PutID(c, place)
	})

	route.Get("/", func(c *fiber.Ctx) {
		place := new(model.Place)
		controller.List(c, place)
	})
}
