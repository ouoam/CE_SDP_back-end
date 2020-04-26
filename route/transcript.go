package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func TranscriptRoute(route *fiber.Group) {
	route.Get("/:id", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		transcript.Tour.SetValid(c.Locals("params_id").(int64))
		transcript.User.SetValid(c.Locals("user_id").(int64))
		controller.GetID(c, transcript)
	})

	route.Post("/", controller.CheckLogin, func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		controller.Post(c, transcript)
	})

	route.Put("/:id", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		transcript.Tour.SetValid(c.Locals("params_id").(int64))
		transcript.User.SetValid(c.Locals("user_id").(int64))
		controller.PutID(c, transcript)
	})

	route.Get("/", controller.CheckLogin, func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		transcript.User.SetValid(c.Locals("user_id").(int64))
		controller.List(c, transcript)
	})
}
