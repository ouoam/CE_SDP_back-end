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
		controller.Get(c, place)
	})

	route.Post("/", controller.CheckLogin, func(c *fiber.Ctx) {
		place := new(model.Place)
		controller.New(c, place)
	})

	route.Put("/:id", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		place := new(model.Place)
		place.ID.SetValid(c.Locals("params_id").(int64))
		controller.Update(c, place)
	})

	route.Get("/", func(c *fiber.Ctx) {
		place := new(model.PlaceSearch)
		controller.List(c, place)
	})
}
