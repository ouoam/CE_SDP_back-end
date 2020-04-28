package route

import (
	"../controller"
	"../model"
	"github.com/gofiber/fiber"
)

func TourRoute(route *fiber.Group) {
	route.Get("/:id", parseIntParams("id"), func(c *fiber.Ctx) {
		tour := new(model.TourDetail)
		tour.ID.SetValid(c.Locals("params_id").(int64))
		controller.GetID(c, tour)
	})

	route.Get("/:id/reviews", parseIntParams("id"), func(c *fiber.Ctx) {
		review := new(model.ReviewWithTour)
		controller.List(c, review, c.Locals("params_id").(int64))
	})

	route.Get("/:id/transcripts", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		transcript := new(model.TranscriptWithTour)
		controller.List(c, transcript, c.Locals("params_id").(int64))
	})

	route.Get("/:id/transcripts/:user", controller.CheckLogin, parseIntParams("id", "user"), func(c *fiber.Ctx) {
		transcript := new(model.TranscriptWithTour)
		transcript.User.SetValid(c.Locals("params_user").(int64))
		controller.GetID(c, transcript, c.Locals("params_id").(int64))
	})

	route.Post("/:id/transcripts/:user", controller.CheckLogin, parseIntParams("id", "user"), func(c *fiber.Ctx) {
		transcript := new(model.Transcript)
		transcript.Tour.SetValid(c.Locals("params_id").(int64))
		transcript.User.SetValid(c.Locals("params_user").(int64))
		controller.PutID(c, transcript)
	})

	route.Post("/", controller.CheckLogin, func(c *fiber.Ctx) {
		tour := new(model.Tour)
		tour.Owner.SetValid(c.Locals("user_id").(int64))
		controller.Post(c, tour)
	})

	route.Put("/:id", controller.CheckLogin, parseIntParams("id"), func(c *fiber.Ctx) {
		// todo update own
		tour := new(model.Tour)
		tour.ID.SetValid(c.Locals("params_id").(int64))
		tour.Owner.SetValid(c.Locals("user_id").(int64))
		controller.PutID(c, tour)
	})

	route.Get("/", func(c *fiber.Ctx) {
		tour := new(model.TourDetail)
		controller.List(c, tour)
	})
}
