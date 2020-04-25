package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func TranscriptRoute(route *fiber.Group) {
	route.Get("/:id", controller.CheckLogin, func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		controller.GetID(c, transcript)
	})

	route.Post("/", controller.CheckLogin, func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		controller.Post(c, transcript)
	})

	route.Put("/:id", controller.CheckLogin, func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		controller.PutID(c, transcript)
	})

	route.Get("/", controller.CheckLogin, func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		controller.List(c, transcript)
	})
}
