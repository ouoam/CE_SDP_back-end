package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func TranscriptRoute(route *fiber.Group) {
	route.Get("/:id", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		transcript := new(model.TranscriptWithUser)
		transcript.Tour.SetValid(c.Locals("params_id").(int64))
		controller.Get(c, transcript, c.Locals("user_id").(int64))
	})

	route.Post("/:id", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		transcript.Tour.SetValid(c.Locals("params_id").(int64))
		transcript.User.SetValid(c.Locals("user_id").(int64))
		controller.New(c, transcript)
	})

	route.Put("/:id", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		transcript.Tour.SetValid(c.Locals("params_id").(int64))
		transcript.User.SetValid(c.Locals("user_id").(int64))
		controller.Update(c, transcript)
	})

	route.Get("/", controller.CheckLogin, func(c *fiber.Ctx) {
		transcript := new(model.TranscriptWithUser)
		controller.List(c, transcript, c.Locals("user_id").(int64))
	})
}
